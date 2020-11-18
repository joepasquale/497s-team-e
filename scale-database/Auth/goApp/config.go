package main

var accessSecret = []byte("secret") // change this to os.Getenv("ACCESS_SECRET") for deployment

var refreshSecret = []byte("secret") // change this to os.Getenv("REFRESH_SECRET") for deployment

var dbName = "uplink-test" // change this to "uplink" for deployment

var collectionName = "users"

var dbUsername = "root"

var dbPassword = "password"

// var mongodbURL = fmt.Sprintf("mongodb://%s:%s@mongos1:27017/test?w=majority", dbUsername, dbPassword)
var mongodbURL = "mongodb://mongos1:27019"
var redisURL = "redis:6379"
