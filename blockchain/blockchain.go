package blockchain

import (
	"time"
)

// Blockchain holds the chain of blocks
var Blockchain []Block

// GenesisBlock creates the first block in the blockchain
func GenesisBlock() Block {
	return Block{
		Index:       0,
		Timestamp:   time.Now().String(),
		DID:         "genesis",
		DIDDocument: "genesis-doc",
		PrevHash:    "",
		Hash:        "",
		Nonce:       0,
	}
}

// GenerateBlock creates a new block based on the previous one
func GenerateBlock(oldBlock Block, did string, didDoc string) Block {
	newBlock := Block{}
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.DID = did
	newBlock.DIDDocument = didDoc
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)
	newBlock.Nonce = 0 // Can be used for proof of work (optional)
	return newBlock
}

// IsBlockValid validates a new block to ensure it follows the blockchain rules
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// AddBlock adds a new block to the blockchain
func AddBlock(did string, didDoc string) {
	oldBlock := Blockchain[len(Blockchain)-1]
	newBlock := GenerateBlock(oldBlock, did, didDoc)

	if IsBlockValid(newBlock, oldBlock) {
		Blockchain = append(Blockchain, newBlock)
	}
}
