package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

type OrderSchema struct {
	Name  string  `bson:"name"`
	Value float64 `bson:"value"`
}

func calculateHash(order Order) string {
	record := order.store + order.name + string(order.date) + strconv.FormatFloat(order.value, 'f', -1, 64) + order.prevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createOrder(prevOrder bson.M, order OrderSchema) Order {
	store := os.Getenv("COLLECTION_NAME")
	newOrder := Order{
		store:    store + "_" + time.Now().Format("20060102"),
		name:     order.Name,
		date:     time.Now().Format("2006-01-02 15:04:05"),
		value:    order.Value,
		prevHash: prevOrder["prevHash"].(string),
		hash:     "",
	}
	newOrder.hash = calculateHash(newOrder)
	return newOrder
}
