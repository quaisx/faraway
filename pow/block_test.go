package pow

import (
	"encoding/json"
	"testing"
)

func TestBlock(t *testing.T) {
	block := &Block{}
	payload, err := json.Marshal(block)
	if err != nil {
		t.Errorf("error marshalling block: %v", err)
	}
	err = json.Unmarshal(payload, &block)
	if err != nil {
		t.Errorf("error unmarshalling block: %v", err)
	}
}
