package controllers

import (
	"context"
	"errors"
	"jwt_athentication/connections"
	"jwt_athentication/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ValidateUser(userRequest *models.User) error {
	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)

	collections := connections.GetCollection(db_connector, "Users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collections.CountDocuments(ctx, bson.D{{Key: "email", Value: userRequest.Email}})
	if err != nil {
		log.Printf("Insertion Error")
		return err
	}
	if count > 0 {

		return errors.New("User Already Exists")
	}
	return nil

}

func AuthenticateUser(loginDetails *models.LoginRequest) (models.User, error) {

	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)

	collections := connections.GetCollection(db_connector, "Users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usr models.User
	err := collections.FindOne(ctx, bson.D{{Key: "email", Value: loginDetails.Email}}).Decode(&usr)

	if err != nil {
		log.Printf("User Not Found")
		return usr, errors.New("User Not Found")
	}

	if usr.Password == loginDetails.Password {
		return usr, nil
	}

	log.Println("Authentication Failed.....")
	return models.User{}, errors.New("Authentication Failed")
}
