package main

import "fmt"

var accessSecret = []byte("secret") // change this to os.Getenv("ACCESS_SECRET")

var refreshSecret = []byte("secret") // change this to os.Getenv("REFRESH_SECRET")

var dbName = "uplink-test"

var collectionName = "users"

var dbUsername = "root"

var dbPassword = "password"

//var mongodbURL = "mongodb+srv://Macbook:acLe9y63lP7ORH3l@cluster0.yfyl0.gcp.mongodb.net/uplink-test?retryWrites=true&w=majority"
var mongodbURL = fmt.Sprintf("mongodb://%s:%s@mongo:27017", dbUsername, dbPassword)

var redisURL = "redis:6379"
