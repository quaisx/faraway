package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Block struct {
	nonce         int
	timestampHash [32]byte
	connection    *Connection
}

func (b *Block) Nonce() int {
	return b.nonce
}

func (b *Block) Timestamp() [32]byte {
	return b.timestampHash
}

func (b *Block) Connection() *Connection {
	return b.connection
}

func NewBlock(nonce int, timestampHash [32]byte, connection *Connection) *Block {
	b := &Block{
		nonce:         nonce,
		timestampHash: timestampHash,
		connection:    connection,
	}
	return b
}

func (b *Block) String() string {
	output := fmt.Sprintf("nonce    %d\n", b.nonce)
	output += fmt.Sprintf("hash     %x\n", b.timestampHash)
	output += fmt.Sprintf("%s", b.connection)
	return output
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce         int         `json:"nonce"`
		TimestampHash string      `json:"timestamp_hash"`
		Connection    *Connection `json:"connection"`
	}{
		Nonce: b.nonce,
		// Note: storing previosHash([32]byte) value as hex string. When decoding, must do reverse
		TimestampHash: fmt.Sprintf("%x", b.timestampHash),
		Connection:    b.connection,
	})
}

func (b *Block) UnmarshalJSON(data []byte) error {
	var timestampHash string
	v := &struct {
		Nonce         *int         `json:"nonce"`
		TimestampHash *string      `json:"timestamp_hash"`
		Connection    **Connection `json:"connection"`
	}{
		Nonce:         &b.nonce,
		TimestampHash: &timestampHash,
		Connection:    &b.connection,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	// After unmarshaling, convert previousHash from hex representationt to a string
	ph, _ := hex.DecodeString(*v.TimestampHash)
	// Update previosHash value with its true [32]byte value
	copy(b.timestampHash[:], ph[:32])
	return nil
}
