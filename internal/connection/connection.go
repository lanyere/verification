package connection

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func Connection() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database("test").Collection("pageData")
}

func InsertOne(pd interface{}) {
	insertResult, err := collection.InsertOne(ctx, pd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// func FindMany(pd interface{}) interface{} {

// 	filter := bson.D{{"Title", "Index Page"}}

// 	err := collection.FindOne(context.TODO(), filter).Decode(&pd)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return pd
// }
