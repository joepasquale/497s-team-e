package main

/* Forms are what server suppose to recieve or send*/

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

// TokenRequestForm data structure for client requests
type TokenRequestForm struct {
	RefID string `json:"reference_id"`
}

// ResponseForm server response format
type ResponseForm struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	RequestStatus string `json:"status"`
	ErrorMsg      string `json:"error"`
}

// TokenPairForm data structure for handling delete requests
type TokenPairForm struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshTokenForm --
type RefreshTokenForm struct {
	RefreshToken string `json:"refresh_token"`
}

// AccessTokenForm structure for validate a token
type AccessTokenForm struct {
	AccessToken string `json:"access_token"`
}

// SingleToken data structure for a simple token
type SingleToken struct {
	Token     string
	UUID      string
	RefID     string
	ExpiresIn int64
}

// TokenPair data structure for a token pair
type TokenPair struct {
	RefID        string
	AccessToken  string
	AccessUUID   string
	AtExpires    int64
	RefreshToken string
	RefreshUUID  string
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
