package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func login(c *gin.Context) {
	var u LoginForm
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
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

/*
func findUserByToken(c *gin.Context) {
	tokenString := extractToken(c.Request)
	fmt.Println("Token string: " + tokenString)
	if len(tokenString) == 0 {
		data := map[string]interface{}{
			"userID":  "",
			"islogin": "false",
		}
		c.JSON(http.StatusUnauthorized, data)
		return
	}

	userID, err := rdb.Get(tokenString).Result()
	fmt.Println("User ID found: " + userID)
	if err != nil {
		fmt.Println(err)
		data := map[string]interface{}{
			"userID":  "",
			"islogin": "false",
		}
		c.JSON(http.StatusUnauthorized, data)
		return
	}
	data := map[string]interface{}{
		"userID":  userID,
		"islogin": "true",
	}
	c.JSON(http.StatusOK, data)
	return
}
*/
