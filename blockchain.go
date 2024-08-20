package main

import (
	"os"
	"time"
)

type Blockchain struct {
	Order []Order
}

func createGenesisOrder() Order {
	store := os.Getenv("COLLECTION_NAME")
	newOrder := Order{
		Store:    store + "_" + time.Now().Format("20060102"),
		Name:     store,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Value:    0,
		PrevHash: "",
		Hash:     "",
	}
	newOrder.Hash = calculateHash(newOrder)
	return newOrder
}
