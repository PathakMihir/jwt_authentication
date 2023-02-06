package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserId       string             `json:"user_id"  bson:"user_id"`
	FirstName    string             `json:"first_name" binding:"required" bson:"first_name"`
	LastName     string             `json:"last_name" binding:"required" bson:"last_name"`
	Email        string             `json:"email" binding:"required,email" bson:"email"`
	PhoneNumber  string             `json:"phone_number" binding:"required,max=10" bson:"phone_number"`
	Password     string             `json:"password" binding:"required,min=6" bson:"password"`
	Token        string             `json:"token" bson:"token"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	Created_at   time.Time          `json:"created_at" bson:"created_at"`
	Updated_at   time.Time          `json:"updated_at" bson:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"required"`
}

type PasswordChange struct{
	Email string `json:"email" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ErrorResponse struct {
	Error   ErrorDetail
	Status  string
	Message string
}

type ErrorDetail struct {
	ErrorType    string
	ErrorMessage string
}

// type Permission struct{
// 	PermissionId int
// 	PermissionTag

// }
