package main

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func login(c *gin.Context) {
	var u User      // request form
	var f Feedback  // feedback to client
	var result User // result from database
	var err error
	if err = c.ShouldBind(&u); err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Invalid login form")
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}
	// TODO: check if the user is in the cache!

	// Find User by email (done)
	filter := bson.M{"email": u.Email}
	err = userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "User not found")
		c.JSON(http.StatusUnauthorized, f)
		return
	}

	// Authenticate user password (done)
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(u.Password))
	if err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Email or password incorrect")
		c.JSON(http.StatusUnauthorized, f)
		return
	}

	// Create token pari
	td, err := createToken(result.ID)
	if err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Token not created")
		f.Msgs = append(f.Msgs, err.Error())
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}

	// save token into redis (done)
	saveErr := createAuth(result.ID, td)
	if saveErr != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Token created but not saved")
		f.Msgs = append(f.Msgs, saveErr.Error())
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}
	f.Status = "success"
	f.Data = tokens
	c.JSON(http.StatusOK, f)
}

// TODO: implement register router
func register(c *gin.Context) {
	var u User      // request form
	var f Feedback  // feedback to client
	var result User // result from database

	if err := c.ShouldBind(&u); err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Invalid register form")
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}
	// TODO: check if the user is in the cache!

	// Find User by email (done)
	filter := bson.M{"email": u.Email}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Email exists")
		c.JSON(http.StatusUnauthorized, f)
		return
	}

	// TODO: check email and password format is valid

	// Hash User password and store it into database
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	insertUser := bson.M{
		"email":    u.Email,
		"password": string(bytes),
	}
	_, err = userCollection.InsertOne(context.TODO(), insertUser)
	if err != nil {
		panic(err)
	}
	f.Status = "success"
	c.JSON(http.StatusOK, f)
}

func createToken(id primitive.ObjectID) (UserToken, error) {
	var err error
	td := UserToken{}
	empty := UserToken{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessUUID = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 14).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["access_uuid"] = td.AccessUUID
	atClaim["_id"] = id
	atClaim["expire"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	td.AccessToken, err = at.SignedString(accessSecret)
	if err != nil {
		return empty, err
	}

	rtClaim := jwt.MapClaims{}
	rtClaim["refresh_uuid"] = td.RefreshUUID
	rtClaim["_id"] = id
	rtClaim["expire"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaim)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return empty, err
	}
	return td, nil
}

// store token into redis
func createAuth(userID primitive.ObjectID, td UserToken) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := rdb.Set(td.AccessUUID, userID.Hex(), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := rdb.Set(td.RefreshUUID, userID.Hex(), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
