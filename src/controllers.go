package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	collectionName := os.Getenv("COLLECTION_NAME")
	w.Header().Set("Content-Type", "application/json")
	// Get a handle for your collection
	orderCollection := Db.Database("ffc_database").Collection(collectionName)

	var results []bson.M
	cursor, err := orderCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	// Iterate through the cursor and decode each document
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(results) // returns a Map containing //mongodb document
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	collectionName := os.Getenv("COLLECTION_NAME")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]                   //get Parameter value as string
	_id, err := primitive.ObjectIDFromHex(params) // convert params to //mongodb Hex ID
	if err != nil {
		fmt.Println(err.Error())
	}
	// Get a handle for your collection
	orderCollection := Db.Database("ffc_database").Collection(collectionName)

	var result bson.M
	err = orderCollection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(&result) // return order
}

func getLatestOrder() bson.M {
	collectionName := os.Getenv("COLLECTION_NAME")
	// Get a handle for your collection
	orderCollection := db().Database("ffc_database").Collection(collectionName)
	// Define an options object to sort by timestamp in descending order
	findOptions := options.FindOne().SetSort(bson.D{{"timestamp", -1}})

	// Find the latest record
	var result bson.M
	err := orderCollection.FindOne(context.TODO(), bson.D{}, findOptions).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func addOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	var order OrderSchema
	err := json.NewDecoder(r.Body).Decode(&order) // storing in order
	if err != nil {
		fmt.Print(err)
	}

	collectionName := os.Getenv("COLLECTION_NAME")
	// Get a handle for your collection
	orderCollection := db().Database("ffc_database").Collection(collectionName)
	latestOrder := getLatestOrder()

	newOrder := createOrder(latestOrder, order)
	insertResult, err := orderCollection.InsertOne(context.TODO(), newOrder)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
}

func chargeBackOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	var order OrderSchema
	err := json.NewDecoder(r.Body).Decode(&order) // storing in order
	if err != nil {
		fmt.Print(err)
	}

	collectionName := os.Getenv("COLLECTION_NAME")
	// Get a handle for your collection
	orderCollection := db().Database("ffc_database").Collection(collectionName)
	latestOrder := getLatestOrder()

	if order.Value == 0 {
		order.Value = latestOrder["value"].(float64) * -1
	}

	newOrder := createOrder(latestOrder, order)
	insertResult, err := orderCollection.InsertOne(context.TODO(), newOrder)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
}

/*
func validateOrders() {
	bc := newBlockchain()

	bc.addBlock("First Block after Genesis")
	bc.addBlock("Second Block after Genesis")

	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}

	fmt.Printf("Blockchain valid: %v\n", bc.isValid())
}
*/
