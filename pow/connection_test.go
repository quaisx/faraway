package pow

import (
	"encoding/json"
	"testing"
)

func TestConnection(t *testing.T) {
	conn := &Connection{}
	payload, err := json.Marshal(conn)
	if err != nil {
		t.Errorf("error marshalling payload: %v", err)
	}
	err = json.Unmarshal(payload, conn)
	if err != nil {
		t.Errorf("error unmarshalling payload: %v", err)
	}
}
