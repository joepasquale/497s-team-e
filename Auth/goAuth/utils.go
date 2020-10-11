package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// create a token based on userID
func createToken(userID string) (UserToken, error) {
	var err error
	td := UserToken{}
	empty := UserToken{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessUUID = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 14).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	// fmt.Println("AtExpires: " + fmt.Sprint(td.AtExpires))
	// fmt.Println("RefreshUUID: " + td.AccessUUID)
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["access_uuid"] = td.AccessUUID
	atClaim["userID"] = userID
	atClaim["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	td.AccessToken, err = at.SignedString(accessSecret)
	if err != nil {
		return empty, err
	}

	rtClaim := jwt.MapClaims{}
	rtClaim["refresh_uuid"] = td.RefreshUUID
	rtClaim["userID"] = userID
	rtClaim["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaim)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return empty, err
	}
	return td, nil
}
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization: Bearer the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func tokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func extractTokenMetadata(r *http.Request) (*AccessToken, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, ok := claims["userID"].(string)

		return &AccessToken{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

func fetchToken(auth *AccessToken) (string, error) {
	userID, err := rdb.Get(auth.AccessUUID).Result()
	if err != nil {
		return "", err
	}
	return userID, nil
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
		refreshUUID, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
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
			c.JSON(http.StatusUnauthorized, "unauthorized access")
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

// TODO: implement find user from mongodb, return user id
func findUser(username string, password string) (string, bool) {
	filter := bson.M{"username": username}
	var result User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return "", false
	}
	match := checkPasswordHash(password, result.Password)
	if match {
		return result.UserID, true
	}
	return "", false
}

func userExist(username string) bool {
	filter := bson.M{"username": username}
	var result User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}

// store token into redis
func createAuth(userid string, td UserToken) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := rdb.Set(td.AccessUUID, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := rdb.Set(td.RefreshUUID, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	// errAccessToken := rdb.Set(td.AccessToken, userid, at.Sub(now)).Err()
	// if errAccessToken != nil {
	// 	return errAccess
	// }
	// errRefreshToken := rdb.Set(td.RefreshToken, userid, rt.Sub(now)).Err()
	// if errRefreshToken != nil {
	// 	return errRefresh
	// }

	userID, _ := rdb.Get(td.AccessUUID).Result()
	fmt.Println("access token stored into redis")
	fmt.Println("AccessUUID: " + td.AccessUUID)
	fmt.Println("AccessToken: " + td.AccessToken)
	fmt.Println("userID: " + userID)
	return nil
}

func deleteAuth(UUID string) (int64, error) {
	deleted, err := rdb.Del(UUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func pingRedis() {
	_, err := rdb.Ping().Result() // check redis is connected
	if err != nil {
		panic(err)
	}
}
