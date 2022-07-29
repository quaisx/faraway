package pow

import (
	"fmt"
	"strings"
)

const (
	POW_DIFFICULTY = 3
)

func ValidProof(nonce int, timestampHash [32]byte, connection *Connection, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{nonce, timestampHash, connection}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func ProofOfWork(timestampHash [32]byte, connection *Connection, difficulty int) int {
	nonce := 0
	for !ValidProof(nonce, timestampHash, connection, difficulty) {
		nonce += 1
	}
	return nonce
}
