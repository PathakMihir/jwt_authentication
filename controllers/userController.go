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

	count,err:=collections.CountDocuments(ctx,bson.D{{Key: "email",Value: userRequest.Email}})
	if err != nil {
		log.Printf("Insertion Error")
		return err
	}
	if count>0{
		
		return errors.New("User Already Exists")
	}
	return nil

}

func InsertUser(userRequest *models.User) error {

	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)


	collections := connections.GetCollection(db_connector, "Users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collections.InsertOne(ctx, userRequest)

	if err != nil {
		log.Printf("Insertion Error")
		return err
	}
	log.Println("User Info Successfully Inserted...")
	return nil

}

func AuthenticateUser(loginDetails *models.LoginRequest) error{

	return nil
}

