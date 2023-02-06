package repositories

import (
	"context"
	"errors"
	"jwt_athentication/connections"
	"jwt_athentication/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAll() ([]models.User,error)  {
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

func UserGetById(Id primitive.ObjectID) (models.User,error) {

	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)

	collections := connections.GetCollection(db_connector, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	filter := bson.D{{Key: "_id", Value: Id}}
	err := collections.FindOne(ctx,filter).Decode(&user)

	if err!=nil{
		log.Println(err)
		return models.User{}, err
	}
	return user ,nil

}

func UserDeleteById(Id primitive.ObjectID) error {
	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)

	collections := connections.GetCollection(db_connector, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: Id}}
	_,err := collections.DeleteOne(ctx,filter)

	if err!=nil{
		log.Println(err)
		return err
	}
	return nil

}

func UserCreate(userRequest *models.User) error {

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
	
func UserUpdatePassword(email string, new_password string,old_password string) error {

	db_connector := connections.DB_Connect()
	defer connections.CloseClientDB(db_connector)
	
	collections := connections.GetCollection(db_connector, "Users")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter:=bson.D{{Key: "email", Value: email},{Key: "password",Value: old_password}}
	update:= bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: new_password}}}}

	_, err := collections.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}