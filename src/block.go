package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func calculateHash(order Order) string {
	record := order.Store + order.Name + string(order.Date) + strconv.FormatFloat(order.Value, 'f', -1, 64) + order.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createOrder(prevOrder Order, data string) Order {
	newOrder := Order{
		Store:    prevOrder.Store,
		Name:     prevOrder.Name,
		Date:     prevOrder.Date,
		Value:    prevOrder.Value,
		PrevHash: prevOrder.PrevHash,
		Hash:     "",
	}
	newOrder.Hash = calculateHash(newOrder)
	return newOrder
}
