package main

import "fmt"

var accessSecret = []byte("secret") // change this to os.Getenv("ACCESS_SECRET") for deployment

var refreshSecret = []byte("secret") // change this to os.Getenv("REFRESH_SECRET") for deployment

var dbName = "uplink-test" // change this to "uplink" for deployment

var collectionName = "users"

var dbUsername = "admin" //admin

var dbPassword = "password" //password

//var mongodbURL = "mongodb+srv://Macbook:acLe9y63lP7ORH3l@cluster0.yfyl0.gcp.mongodb.net/uplink-test?retryWrites=true&w=majority"
var mongodbURL = fmt.Sprintf("mongodb://%s:%s@mongos1:27017", dbUsername, dbPassword)

var redisURL = "redis:6379"
