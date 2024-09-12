package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// Block represents a single block in the blockchain
type Block struct {
	Index       int    // Position of the block in the blockchain
	Timestamp   string // The time when the block was created
	DID         string // The DID being registered
	DIDDocument string // Reference to the DID document (e.g., IPFS hash)
	PrevHash    string // Hash of the previous block
	Hash        string // Hash of the current block
	Nonce       int    // Nonce for proof of work (optional)
}

// CalculateHash calculates the SHA256 hash of a block
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.DID + block.DIDDocument + block.PrevHash + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
