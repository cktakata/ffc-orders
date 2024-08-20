package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"time"
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
	record := order.Store + order.Name + string(order.Date) + strconv.FormatFloat(order.Value, 'f', -1, 64) + order.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createOrder(prevOrder Order, order OrderSchema) Order {
	store := os.Getenv("COLLECTION_NAME")
	newOrder := Order{
		Store:    store + "_" + time.Now().Format("20060102"),
		Name:     order.Name,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Value:    order.Value,
		PrevHash: prevOrder.PrevHash,
		Hash:     "",
	}
	newOrder.Hash = calculateHash(newOrder)
	return newOrder
}
