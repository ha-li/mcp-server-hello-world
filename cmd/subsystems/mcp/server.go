package mcp

import "fmt"

type (
	McpRequest struct {
		JsonRpc string      `json:"jsonrpc"`
		Id      interface{} `json:"id"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
	}
	McpResponse struct {
		JsonRpc string      `json:"jsonrpc"`
		Id      interface{} `json:"id"`
		Result  interface{} `json:"result,omitempty"`
		Error   *McpError   `json:"error,omitempty"`
	}

	McpError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Pineapple interface {
		Juice(string, Water) Water
	}

	Water struct{}

	Appliance interface {
		Make(string) Water
	}

	Oven struct {
		temperature int
	}
)

func (w *Water) doSomething() Water {
	return Water{}
}

func (o *Oven) Make(sm string) Water {
	fmt.Printf("Making %s", sm)
	return Water{}
}
