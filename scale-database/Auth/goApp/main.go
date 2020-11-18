package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     redisURL,
	PoolSize: 10,
	DB:       0, // use default DB
})

// Set client options
var clientOptions = options.Client().ApplyURI(mongodbURL)

// Connect to MongoDB
var mongodb, err = mongo.Connect(context.Background(), clientOptions)

// connect to collection
var userCollection = mongodb.Database("uplink-test").Collection("users")

func main() {
	var r = gin.Default()
	// Check the connection
	err = mongodb.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, err)

	r.POST("/login", login)
	r.POST("/register", register)
	r.Run(":8000")
}
