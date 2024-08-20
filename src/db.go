package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Client

type Order struct {
	store    string  `bson:"store"`
	name     string  `bson:"name"`
	date     string  `bson:"date"`
	value    float64 `bson:"value"`
	prevHash string  `bson:"prevHash"`
	hash     string  `bson:"hash"`
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
	fmt.Println("Current date and time:", time.Now())
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
		// Get a handle for your collection
		orderCollection := database.Collection(collectionName)
		order := createGenesisOrder()
		insertResult, err := orderCollection.InsertOne(context.TODO(), order)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Genesis block added. %v\n", insertResult)
	}

	return client
}
