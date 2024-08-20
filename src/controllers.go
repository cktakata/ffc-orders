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

	var result Order
	err = orderCollection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result) // return number of //documents deleted
}

func createProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	var person user
	err := json.NewDecoder(r.Body).Decode(&person) // storing in person   //variable of type user
	if err != nil {
		fmt.Print(err)
	}

	// Get a handle for your collection
	userCollection := db().Database("your_database").Collection("your_collection")

	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
}

func getUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body user
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON //document which is a Map

	// Get a handle for your collection
	userCollection := db().Database("your_database").Collection("your_collection")

	err := userCollection.FindOne(context.TODO(), bson.D{{"name", body.Name}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result) // returns a Map containing //mongodb document
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type updateBody struct {
		Name string `json:"name"` //value that has to be matched
		City string `json:"city"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}
	filter := bson.D{{"name", body.Name}} // converting value to BSON after := options.After         // for returning updated document
	after := options.After                // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"city", body.City}}}}

	// Get a handle for your collection
	userCollection := db().Database("your_database").Collection("your_collection")

	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]                   //get Parameter value as string
	_id, err := primitive.ObjectIDFromHex(params) // convert params to //mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to //specify language-specific rules for string comparison, such as //rules for lettercase

	// Get a handle for your collection
	userCollection := db().Database("your_database").Collection("your_collection")

	res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of //documents deleted
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
