package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain represents the blockchain, which is just a slice of pointers to Block
type BlockChain struct {
	blocks []*Block
}

// Block represents a single block in the blockchain
type Block struct {
	Hash     []byte // The hash of the block's data and previous block's hash
	Data     []byte // The data stored in the block
	PrevHash []byte // The hash of the previous block in the chain
}

// DeriveHash calculates the block's hash by hashing the concatenation of its Data and PrevHash
func (b *Block) DeriveHash() {
	// Concatenate Data and PrevHash into one byte slice
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

	hash := sha256.Sum256(info)

	// Assign the resulting hash to the block's Hash field
	b.Hash = hash[:]
}

// CreateBlock creates a new block with the given data and previous hash, then derives its hash
func CreateBlock(data string, prevHash []byte) *Block {
	// Initialize a new block with empty hash, given data, and previous block's hash
	block := &Block{[]byte{}, []byte(data), prevHash}

	// Calculate and set the block's hash
	block.DeriveHash()

	return block
}

// AddBlock adds a new block to the blockchain with the given data
func (chain *BlockChain) AddBlock(data string) {
	// Get the previous block (the last block in the chain)
	prevBlock := chain.blocks[len(chain.blocks)-1]

	// Create a new block with the given data and the hash of the previous block
	newBlock := CreateBlock(data, prevBlock.Hash)

	// Append the new block to the chain
	chain.blocks = append(chain.blocks, newBlock)
}

// Genesis creates the very first block in the blockchain, known as the "genesis block"
func Genesis() *Block {
	// Create a genesis block with special data and no previous hash (empty slice)
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initializes a new blockchain with the genesis block
func InitBlockChain() *BlockChain {
	// Return a new blockchain with just the genesis block
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	// Initialize the blockchain with the genesis block
	chain := InitBlockChain()

	// Add subsequent blocks to the chain
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// Print the details of each block in the blockchain
	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
