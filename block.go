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

type OrderSchema struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func calculateHash(order Order) string {
	record := order.Store + order.Name + string(order.Date) + strconv.FormatFloat(order.Value, 'f', -1, 64) + order.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
