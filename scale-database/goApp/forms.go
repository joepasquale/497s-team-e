package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Forms are what server suppose to recieve or send*/

// User - info required for login
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty" `
	Email    string             `form:"email" binding:"required"`
	Password string             `form:"password" binding:"required"`
}

// Feedback - info required for sending feedback
type Feedback struct {
	Status string      `json:"status"`
	Msgs   []string    `json:"msgs"`
	Data   interface{} `json:"data"`
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
