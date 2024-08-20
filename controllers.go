package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Message string `json:"message"`
	IsValid bool   `json:"isValid"`
}

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
	err = orderCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(&result) // return order
}

func createOrder(prevOrder bson.M, order OrderSchema) Order {
	store := os.Getenv("COLLECTION_NAME")
	newOrder := Order{
		Store:    store + "_" + time.Now().Format("20060102"),
		Name:     order.Name,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Value:    order.Value,
		PrevHash: prevOrder["hash"].(string),
		Hash:     "",
	}
	newOrder.Hash = calculateHash(newOrder)
	return newOrder
}

func getLatestOrder() bson.M {
	collectionName := os.Getenv("COLLECTION_NAME")
	// Get a handle for your collection
	orderCollection := db().Database("ffc_database").Collection(collectionName)
	// Define an options object to sort by timestamp in descending order
	findOptions := options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}})

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

	if order.Value > 0 {
		order.Value = latestOrder["value"].(float64) * -1
	}

	newOrder := createOrder(latestOrder, order)
	insertResult, err := orderCollection.InsertOne(context.TODO(), newOrder)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document

}

func isValidOrders(w http.ResponseWriter, r *http.Request) {
	collectionName := os.Getenv("COLLECTION_NAME")
	w.Header().Set("Content-Type", "application/json")
	// Get a handle for your collection
	orderCollection := Db.Database("ffc_database").Collection(collectionName)

	// Find documents in the collection
	cursor, err := orderCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// Retrieve all orders into a slice
	var orders []Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		log.Fatal(err)
	}

	// Iterate through the slice using a regular for loop
	orderError := "No error"
	isValid := true
	for i := 1; i < len(orders); i++ {
		currentBlock := orders[i]
		prevBlock := orders[i-1]
		// Check if the current block's hash is correct
		if currentBlock.Hash != calculateHash(currentBlock) {
			orderError = currentBlock.Date + " " + currentBlock.Name + " " + fmt.Sprintf("%f", currentBlock.Value)
			isValid = false
		}

		// Check if the current block's previous hash matches the previous block's hash
		if currentBlock.PrevHash != prevBlock.Hash {
			orderError = currentBlock.Date + " " + currentBlock.Name + " " + fmt.Sprintf("%f", currentBlock.Value)
			isValid = false
		}
	}

	response := Response{
		Message: orderError,
		IsValid: isValid,
	}
	json.NewEncoder(w).Encode(response) // returns a Map containing //mongodb document
}
