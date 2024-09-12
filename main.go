package main

import (
	"DID-Webauthn/blockchain"
	"fmt"
)

func main() {
	// Initialize the blockchain with the genesis block
	blockchain.Blockchain = append(blockchain.Blockchain, blockchain.GenesisBlock())

	fmt.Println("Blockchain initialized with genesis block.")

	// Add a new DID block to the blockchain
	blockchain.AddBlock("did:web:example.com", "QmXoY...")

	// Print the blockchain to verify the new block
	for _, block := range blockchain.Blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("DID: %s\n", block.DID)
		fmt.Printf("DID Document: %s\n", block.DIDDocument)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}
