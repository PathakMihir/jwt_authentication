package controllers

import (
	"context"
	"errors"
	"jwt_athentication/connections"
	"jwt_athentication/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func UpdateToken(token string, refresh_token string, email string) error {
	log.Println(email)
	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)

	collections := connections.GetCollection(db_connector, "Users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: token}, {Key: "refresh_token", Value: refresh_token}}}}
	var updatedDocument bson.M
	err := collections.FindOneAndUpdate(
		ctx,
		filter,
		update,
		opts,
	).Decode(&updatedDocument)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		log.Println("Error in Updating the token")
		if err == mongo.ErrNoDocuments {
			return errors.New("User not found")
		}
		log.Fatal(err)
	}
	return nil

}

func GetAllUsers() ([]models.User, error) {

	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)

	collections := connections.GetCollection(db_connector, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var results []models.User
	cursor, err := collections.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("User Not Found")
		return results, errors.New("User Not Found")
	}

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println(err)
			return nil, err
		}
		log.Printf("%+v\n", user)
		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return results, nil
}
