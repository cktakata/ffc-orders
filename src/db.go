package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Client

type user struct {
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}

type Order struct {
	Store    string `bson:"store"`
	Name     string `bson:"name"`
	Date     string `bson:"date"`
	Value    string `bson:"value"`
	PrevHash string `bson:"prevHash"`
	Hash     string `bson:"hash"`
}

func db() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	collectionName := os.Getenv("COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017") // Connect to //MongoDB
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

	// Get a handle for your database
	database := client.Database("ffc_database")

	// Check if the collection exists
	collections, err := database.ListCollectionNames(context.TODO(), map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	}

	collectionExists := false
	for _, name := range collections {
		if name == collectionName {
			collectionExists = true
			break
		}
	}

	if collectionExists {
		fmt.Printf("Collection %s already exists.\n", collectionName)
	} else {
		// Create the collection
		err = database.CreateCollection(context.TODO(), collectionName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Collection %s created.\n", collectionName)
	}

	return client
}
