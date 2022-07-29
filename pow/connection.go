package pow

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Connection struct {
	ClientAddress string `json:"client_address"`
	ServerAddress string `json:"server_address"`
	LastAttemptAt uint64 `json:"last_attempt_at"`
	Attempts      uint16 `json:"attempts"`
}

func NewConnection(client string, server string, timestamp uint64) *Connection {
	return &Connection{client, server, timestamp, 0}
}

func (c *Connection) String() string {
	output := fmt.Sprintf("%s\n", strings.Repeat(">", 40))
	output += fmt.Sprintf(" %s -> %s\n", c.ClientAddress, c.ServerAddress)
	return output
}

func (c *Connection) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ClientAddress string `json:"client_address"`
		ServerAddress string `json:"server_address"`
		LastAttemptAt uint64 `json:"last_attempt_at"`
		Attempts      uint16 `json:"attempts`
	}{
		ClientAddress: c.ClientAddress,
		ServerAddress: c.ServerAddress,
		LastAttemptAt: c.LastAttemptAt,
		Attempts:      c.Attempts,
	})
}

func (c *Connection) UnmarshalJSON(data []byte) error {
	v := &struct {
		ClientAddress *string `json:"client_address"`
		ServerAddress *string `json:"server_address"`
		LastAttemptAt *uint64 `json:"last_attempt_at"`
		Attempts      *uint16 `json:"attempts"`
	}{
		ClientAddress: &c.ClientAddress,
		ServerAddress: &c.ServerAddress,
		LastAttemptAt: &c.LastAttemptAt,
		Attempts:      &c.Attempts,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}
