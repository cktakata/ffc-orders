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
	return Order{
		Store:    store,
		Name:     store,
		Date:     time.DateTime,
		Value:    0,
		PrevHash: "",
		Hash:     "",
	}
}

// func (bc *Blockchain) addBlock(data string) {
// 	prevBlock := bc.Blocks[len(bc.Blocks)-1]
// 	newBlock := createBlock(prevBlock, data)
// 	bc.Blocks = append(bc.Blocks, newBlock)
// }

// func newBlockchain() *Blockchain {
// 	genesisBlock := createGenesisBlock()
// 	genesisBlock.Hash = calculateHash(genesisBlock)
// 	return &Blockchain{[]Block{genesisBlock}}
// }

/*
func (bc *Blockchain) isValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		// Check if the current block's hash is correct
		if currentBlock.Hash != calculateHash(currentBlock) {
			fmt.Printf("Block %d has an invalid hash\n", currentBlock.Index)
			return false
		}

		// Check if the current block's previous hash matches the previous block's hash
		if currentBlock.PrevHash != prevBlock.Hash {
			fmt.Printf("Block %d has an invalid previous hash\n", currentBlock.Index)
			return false
		}
	}
	return true
}
*/
