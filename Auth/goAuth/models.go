package main

// User - info required to authenticate a user when login
// related data should be stored in a postgres database
type User struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginForm - info required for login
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// UserToken data structure for storing token info
type UserToken struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessToken data structure for storing access token info
// Redis store key:AccessUUID value:UserID pairs
type AccessToken struct {
	AccessUUID string
	UserID     string
}

// example user
var u0 = User{
	UserID:   "1",
	Username: "username",
	Password: "password",
}
