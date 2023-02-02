package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserId       string             `json:"user_id"`
	FirstName    string             `json:"first_name" binding:"required"`
	LastName     string             `json:"last_name" binding:"required"`
	Email        string             `json:"email" binding:"required,email"`
	PhoneNumber  string             `json:"phone_number" binding:"required,max=10"`
	Password     string             `json:"password" binding:"required,min=6"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refresh_token"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"required"`
}
