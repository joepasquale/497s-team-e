package main

import (
	"context"
	"fmt"
	"log"
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

func storeToken(token SingleToken) error {
	at := time.Unix(token.ExpiresIn, 0) //converting Unix to UTC(to Time object)
	now := time.Now()
	err := rdb.Set(token.UUID, token.RefID, at.Sub(now)).Err()
	if err != nil {
		log.Fatal(err)
	}
	return err
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

// Allow CORS access
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Writer.Header().Get("Method") == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Requested-By")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
			c.Next()
		}
	}
}
