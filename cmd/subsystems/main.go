package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"mcp-server-hello-world/cmd/subsystems/mcp"
)

type (
	InitializeParams struct {
		ProtocolVersion string                 `json:"protocolVersion"`
		Capabilities    map[string]interface{} `json:"capabilities"`
		ClientInfo      ClientInfo             `json:"clientInfo"`
	}

	ClientInfo struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	InitializeResult struct {
		ProtocolVersion string                 `json:"protocolVersion"`
		Capabilities    map[string]interface{} `json:"capabilities"`
		ServerInfo      ServerInfo             `json:"serverInfo"`
	}

	/* Capabilities struct {
		Tools map[string]interface{} `json:"tools,omitempty"`
	} */

	ServerInfo struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	Tool struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		InputSchema InputSchema `json:"inputSchema"`
	}

	InputSchema struct {
		Type       string              `json:"type"`
		Properties map[string]Property `json:"properties,omitempty"`
	}

	Property struct {
		Type        string `json:"type"`
		Description string `json:"description,omitempty"`
	}

	ListToolsResult struct {
		Tools []Tool `json:"tools"`
	}

	CallToolParams struct {
		Name      string                 `json:"name"`
		Arguments map[string]interface{} `json:"arguments,omitempty"`
	}

	CallToolResult struct {
		Content []Content `json:"content"`
	}

	Content struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
)

//func Bootstrap() *Tool {
//	return &Tool{Name: "bootstrapper", Description: "kick start it all"}
//}
//
//func kickStarter(t *Tool) *Generic {
//	return &Generic{Tool: *t, Type: "bootstrapper", Name: "boostrap"}
//}

func main() {
	fmt.Fprintln(os.Stderr, "Hello World MCP Server starting...")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Fprintf(os.Stderr, "Received: %s\n", line)

		var request mcp.McpRequest
		if err := json.Unmarshal([]byte(line), &request); err != nil {
			fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
			sendError(request.Id, -32700, "Parse error")
			continue
		}
		fmt.Fprintf(os.Stderr, "Handling method: %s\n", request.Method)
		handleRequest(request)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
	}
}

func handleRequest(req mcp.McpRequest) {
	fmt.Fprintf(os.Stderr, "Processing method: %s with ID: %v\n", req.Method, req.Id)

	switch req.Method {
	case "initialize":
		handleInitialize(req)
	case "tools/list":
		handleListTools(req)
	case "tools/call":
		handleCallTool(req)
	case "notifications/initialized":
		fmt.Fprintln(os.Stderr, "Received initialized notification")
		// No response needed for notifications
		return
	default:
		sendError(req.Id, -32601, fmt.Sprintf("Method not found: %s", req.Method))
	}
}

func handleInitialize(req mcp.McpRequest) {
	// Parse the initialize parameters
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal params: %v\n", err)
		sendError(req.Id, -32602, "Invalid params")
		return
	}

	var params InitializeParams
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal initialize params: %v\n", err)
		sendError(req.Id, -32602, "Invalid params")
		return
	}

	fmt.Fprintf(os.Stderr, "Client: %s v%s, Protocol: %s\n",
		params.ClientInfo.Name, params.ClientInfo.Version, params.ProtocolVersion)

	// Respond with our capabilities
	result := InitializeResult{
		ProtocolVersion: "2025-06-18",
		Capabilities: map[string]interface{}{
			"tools": map[string]interface{}{},
		},
		ServerInfo: ServerInfo{
			Name:    "hello-world-server",
			Version: "1.0.0",
		},
	}

	sendResponse(req.Id, result)
}

func handleListTools(req mcp.McpRequest) {
	tools := []Tool{
		{
			Name:        "hello_world",
			Description: "Returns a hello world message",
			InputSchema: InputSchema{
				Type: "object",
				Properties: map[string]Property{
					"name": {
						Type:        "string",
						Description: "Optional name to greet (defaults to 'world')",
					},
				},
			},
		},
	}

	result := ListToolsResult{
		Tools: tools,
	}

	sendResponse(req.Id, result)
}

func handleCallTool(req mcp.McpRequest) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		sendError(req.Id, -32602, "Invalid params")
		return
	}

	var params CallToolParams
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		sendError(req.Id, -32602, "Invalid params")
		return
	}

	switch params.Name {
	case "hello_world":
		name := "world"
		if nameArg, ok := params.Arguments["name"].(string); ok {
			name = nameArg
		}

		message := fmt.Sprintf("Hello %s!", name)

		result := CallToolResult{
			Content: []Content{
				{
					Type: "text",
					Text: message,
				},
			},
		}

		sendResponse(req.Id, result)
	default:
		sendError(req.Id, -32602, fmt.Sprintf("Unknown tool: %s", params.Name))
	}
}

func sendResponse(id interface{}, result interface{}) {
	response := mcp.McpResponse{
		JsonRpc: "2.0",
		Id:      id,
		Result:  result,
	}

	jsonBytes, _ := json.Marshal(response)
	fmt.Println(string(jsonBytes))
}

func sendError(id interface{}, code int, message string) {
	response := mcp.McpResponse{
		JsonRpc: "2.0",
		Id:      id,
		Error: &mcp.McpError{
			Code:    code,
			Message: message,
		},
	}

	jsonBytes, _ := json.Marshal(response)
	fmt.Println(string(jsonBytes))
}
