package mcp

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	req := McpRequest{
		JsonRpc: "2.0",
		Id:      1,
		Method:  "test_method",
		Params: map[string]interface{}{
			"param1": "value1",
			"param2": 123,
		},
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal struct: %v", err)
	}

	expectedJson := `{"jsonrpc":"2.0","id":1,"method":"test_method","params":{"param1":"value1","param2":123}}`

	assert.Equal(t, expectedJson, string(jsonData), "nto equal")
}
