package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     redisURL,
	PoolSize: 10,
	DB:       0, // use default DB
})

func init() {
	fmt.Println("Hello World! and hashed as: ")
	fmt.Println(uuid.New([]byte("Hello world")).String())
}

// Set client options
var clientOptions = options.Client().ApplyURI(mongodbURL)

// Connect to MongoDB
var mongodb, err = mongo.Connect(context.Background(), clientOptions)

// connect to collection
var userCollection = mongodb.Database("uplink-test").Collection("users")

func main() {
	var r = gin.Default()
	r.Use(cors.Default())
	// Check the connection
	err = mongodb.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, err)
	r.StaticFile("login.html", "./static/login.html")
	r.StaticFile("jwt.js", "./static/jwt.js")
	r.POST("/login", login)
	r.POST("/logout", tokenAuthMiddleware(), logout)
	r.POST("/register", register)
	r.POST("/refresh_access_token", RefreshAccessToken)
	r.POST("/validate_access_token", ValidateAccessToken)
	r.POST("/validate_refresh_token", ValidateRefreshToken)
	r.POST("/delete_token_pair", DeleteTokens)
	r.POST("/get_token_pair", GetTokenPair)
	r.Run(":5555")
}
