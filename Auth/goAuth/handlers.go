package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func login(c *gin.Context) {
	var u LoginForm
	// accept both json and form data
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid form provided")
		return
	}
	fmt.Println("Request form:")
	fmt.Println(u)
	//compare the user from the request, with the one we defined:
	userID, found := findUser(u.Username, u.Password)
	if found == false {
		c.JSON(http.StatusUnauthorized, "Username or password incorrect")
		return
	}
	ts, err := createToken(userID)
	fmt.Println("Access token created: " + ts.AccessToken)
	if err != nil {
		fmt.Println("Token not created")
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := createAuth(userID, ts)
	if saveErr != nil {
		fmt.Println("Token not saved")
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

// TODO: implement register router
func register(c *gin.Context) {
	var user LoginForm
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if userExist(user.Username) {
		c.JSON(http.StatusUnprocessableEntity, "Username exists, choose another one")
		return
	}

	hashedPw, err := hashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	u := bson.M{
		"userID":   uuid.New([]byte(user.Username)).String(),
		"username": user.Username,
		"password": hashedPw,
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult)
	c.JSON(http.StatusOK, insertResult)
}

func logout(c *gin.Context) {
	au, err := extractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized access")
		return
	}
	deleted, delErr := deleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized access")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func tokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func checkToken(c *gin.Context) {
	user, err := extractTokenMetadata(c.Request)
	if err != nil {
		data := map[string]interface{}{
			"userID":  "",
			"islogin": "false",
			"status":  "invalid token",
		}
		c.JSON(http.StatusUnauthorized, data)
		return
	}

	userID, err := fetchToken(user)
	if err != nil {
		data := map[string]interface{}{
			"userID":  "",
			"islogin": "false",
			"status":  "access token expired",
		}
		c.JSON(http.StatusUnauthorized, data)
		return
	}

	data := map[string]interface{}{
		"userID":  userID,
		"islogin": "true",
		"status":  "ok",
	}

	c.JSON(http.StatusOK, data)

}

func getToken(c *gin.Context) {
	var req TokenRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "request info invalid")
		return
	}

	var simple SingleToken
	simple.RefID = req.RefID
	simple.UUID = uuid.NewV4().String()
	simple.ExpiresIn = time.Now().Add(time.Minute * 15).Unix()
	atClaim := jwt.MapClaims{
		"access_uuid": simple.UUID,
		"exp":         simple.ExpiresIn,
		"refID":       simple.RefID,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	token, err := at.SignedString(accessSecret)
	simple.Token = token
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = storeToken(simple)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]string{
		"access_token": token,
	})
}

// DeleteTokens delete a pair of access and refresh tokens in the db
func DeleteTokens(c *gin.Context) {
	var req TokenPairForm
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "request info invalid")
		return
	}
	if len(req.AccessToken) != 0 {
		token, err := jwt.Parse(req.AccessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Access token has unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(accessSecret), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, "access_uuid invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			accessUUID, ok := claims["access_uuid"].(string)
			if !ok {
				c.JSON(http.StatusBadRequest, "access_uuid invalid")
			}
			deleteAuth(accessUUID)
		}
	}

	if len(req.RefreshToken) != 0 {
		token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Refresh token has unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(refreshSecret), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, "refresh_uuid invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			refreshUUID, ok := claims["refresh_uuid"].(string)
			if !ok {
				c.JSON(http.StatusBadRequest, "ref_uuid invalid")
			}
			deleteAuth(refreshUUID)
		}
	}
}

// GetTokenPair a pair of tokens
func GetTokenPair(c *gin.Context) {
	var req TokenRequestForm
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, "request info invalid")
		return
	}
	var access SingleToken
	access.RefID = req.RefID
	access.UUID = uuid.NewV4().String()
	access.ExpiresIn = time.Now().Add(time.Minute * 15).Unix()
	atClaim := jwt.MapClaims{
		"access_uuid": access.UUID,
		"exp":         access.ExpiresIn,
		"refID":       access.RefID,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	accessToken, err := at.SignedString(accessSecret)
	access.Token = accessToken
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = storeToken(access)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	var refresh SingleToken
	refresh.RefID = req.RefID
	refresh.UUID = uuid.NewV4().String()
	refresh.ExpiresIn = time.Now().Add(time.Hour * 24 * 14).Unix()
	refClaim := jwt.MapClaims{
		"refresh_uuid": refresh.UUID,
		"exp":          refresh.ExpiresIn,
		"refID":        refresh.RefID,
	}
	ref := jwt.NewWithClaims(jwt.SigningMethodHS256, refClaim)
	refreshToken, err := ref.SignedString(refreshSecret)
	refresh.Token = refreshToken
	err = storeToken(refresh)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]string{
		"refresh_token": refreshToken,
		"access_token":  accessToken,
	})
	return
}

// Refresh take a refresh token and return a new access token
func Refresh(c *gin.Context) {
	var response ResponseForm
	response.RequestStatus = "fail"
	var ref RefreshTokenForm
	if err := c.ShouldBindJSON(&ref); err != nil {
		response.ErrorMsg = "Invalid refresh token"
		c.JSON(http.StatusBadRequest, response)
	}
	tokenString := ref.RefreshToken
	decoded, ok := parseToken(tokenString, refreshSecret)
	if !ok {
		response.ErrorMsg = "refresh_uuid invalid"
		c.JSON(http.StatusBadRequest, response)
	}
	refreshUUID := decoded["refresh_uuid"].(string)
	refID, err := rdb.Get(refreshUUID).Result()
	if err != nil {
		response.ErrorMsg = "refresh token expired"
		c.JSON(http.StatusBadRequest, response)
	}
	// create and return a new access token
	var access SingleToken
	access.UUID = uuid.NewV4().String()
	access.RefID = refID
	access.ExpiresIn = time.Now().Add(time.Minute * 15).Unix()
	atClaim := jwt.MapClaims{
		"access_uuid": access.UUID,
		"exp":         access.ExpiresIn,
		"refID":       access.RefID,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	accessToken, err := at.SignedString(accessSecret)
	access.Token = accessToken
	err = storeToken(access)
	if err != nil {
		response.ErrorMsg = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.RequestStatus = "success"
	response.AccessToken = accessToken
	c.JSON(http.StatusOK, response)
}

func parseToken(tokenString string, secret []byte) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Refresh token has unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, true
	}
	return nil, false
}

func refreshToken(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecret), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, ok)
			return
		}
		userID, ok := claims["userID"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := deleteAuth(refreshUUID)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "delete refresh token failed")
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := createToken(userID)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := createAuth(userID, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "token expired")
	}
}

// RefreshAccessToken -- Check the coming refresh token is still valid
// Reply new_access token
func RefreshAccessToken(c *gin.Context) {
	var refreshRequest RefreshTokenForm
	var response ResponseForm
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		// form is invalid
		response.RequestStatus = "failed"
		response.ErrorMsg = "invalid request format"
		c.JSON(http.StatusBadRequest, response)
	}
	refreshToken := refreshRequest.RefreshToken
	//verify the token
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecret), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = "invalid refresh token"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if !ok {
		response.RequestStatus = "failed"
		response.ErrorMsg = "failed to parse token"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	refreshUUID := claims["refresh_uuid"].(string)
	refID, err := rdb.Get(refreshUUID).Result() // _:refID
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = "refresh token expired"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	var newAccessToken SingleToken
	newAccessToken.RefID = refID
	newAccessToken.UUID = uuid.NewV4().String()
	newAccessToken.ExpiresIn = time.Now().Add(time.Minute * 15).Unix()
	atClaim := jwt.MapClaims{
		"access_uuid": newAccessToken.UUID,
		"exp":         newAccessToken.ExpiresIn,
		"refID":       newAccessToken.RefID,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	token, err := at.SignedString(accessSecret)
	newAccessToken.Token = token
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = storeToken(newAccessToken)
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// finally
	response.RequestStatus = "success"
	response.ErrorMsg = ""
	response.AccessToken = newAccessToken.Token
	c.JSON(http.StatusOK, response)
}

// ValidateAccessToken check the given token is valid
func ValidateAccessToken(c *gin.Context) {
	var accessRequest AccessTokenForm
	var response ResponseForm
	if err := c.ShouldBindJSON(&accessRequest); err != nil {
		// form is invalid
		response.RequestStatus = "failed"
		response.ErrorMsg = "invalid request format"
		c.JSON(http.StatusBadRequest, response)
	}
	//verify the token
	parsedToken, err := jwt.Parse(accessRequest.AccessToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessSecret), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = "invalid access token"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if !ok {
		response.RequestStatus = "failed"
		response.ErrorMsg = "failed to parse token"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	accessUUID := claims["access_uuid"].(string)
	_, err = rdb.Get(accessUUID).Result() // _:refID
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = "access token expired"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response.RequestStatus = "success"
	c.JSON(http.StatusUnauthorized, response)
	return
}

// ValidateRefreshToken check the given token is valid
func ValidateRefreshToken(c *gin.Context) {
	var refreshRequest RefreshTokenForm
	var response ResponseForm
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		// form is invalid
		response.RequestStatus = "failed"
		response.ErrorMsg = "invalid request format"
		c.JSON(http.StatusBadRequest, response)
	}
	//verify the token
	parsedToken, err := jwt.Parse(refreshRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecret), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = "invalid refresh token"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if !ok {
		response.RequestStatus = "failed"
		response.ErrorMsg = "failed to parse token"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	refreshUUID := claims["refresh_uuid"].(string)
	_, err = rdb.Get(refreshUUID).Result() // _:refID
	if err != nil {
		response.RequestStatus = "failed"
		response.ErrorMsg = "refresh token expired"
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response.RequestStatus = "success"
	c.JSON(http.StatusUnauthorized, response)
	return
}
