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

func TestRequest_empty_params(t *testing.T) {
	req := McpRequest{
		JsonRpc: "2.0",
		Id:      1,
		Method:  "test_method",
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal struct: %v", err)
	}

	expectedJson := `{"jsonrpc":"2.0","id":1,"method":"test_method"}`

	assert.Equal(t, expectedJson, string(jsonData), "nto equal")
}

func Test_WaterInstance(t *testing.T) {
	w := Water{}
	val := w.doSomething()
	assert.Equal(t, Water{}, val, "not equal")
}

func Test_WaterPointer(t *testing.T) {
	w := new(Water)
	val := w.doSomething()
	assert.Equal(t, Water{}, val, "not equal")
}

func Test_Oven_Make_Struct(t *testing.T) {
	o := Oven{}
	w := o.Make("roast chicken")
	assert.Equal(t, Water{}, w, "not equal")
}

func Test_Oven_Make_Pointer(t *testing.T) {
	app := new(Oven)
	w := app.Make("roast chicken")
	assert.Equal(t, Water{}, w, "not equal")
}

func Test_Appliance_Make_Pointer(t *testing.T) {
	// o := Oven{}
	var o Oven
	var app Appliance = &o

	w := app.Make("roast chicken")
	assert.Equal(t, Water{}, w, "not equal")
}

func Test_Oven(t *testing.T) {
	var o Oven
	assert.Equal(t, Oven{}, o, "not equal")
}
