package connections

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DB_Connect() *mongo.Client{
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))

	if err != nil {
		log.Fatal(err)
	}

	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)
	err=client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to mongoDB")
 	return client


}

func CloseClientDB(client *mongo.Client) {
    if client == nil {
        return
    }

    err := client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    // TODO optional you can log your closed MongoDB client
    log.Println("Connection to MongoDB closed.")
}

func GetCollection(dbconnector *mongo.Client,collection_name string)  *mongo.Collection{
	collection:=dbconnector.Database("JWT").Collection(collection_name)
	return collection
	
}
