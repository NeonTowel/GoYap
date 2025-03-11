package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

// SendToAzureOpenAI sends a message to Azure OpenAI and returns the response
func SendToAzureOpenAI(message Message, context []Message) (Response, error) {
	// Prepare request payload
	payload := Request{
		Messages: append(context, message),
		Context:  map[string]interface{}{}, // Add any context properties here
	}

	// Convert payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	// Get Azure OpenAI endpoint and API key from environment variables
	endpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	apiKey := os.Getenv("AZURE_OPENAI_API_KEY")

	// Create HTTP request
	req, err := http.NewRequest("POST", endpoint+"/chat", bytes.NewBuffer(data))
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send request to Azure OpenAI
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	// Parse response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, err
	}

	return response, nil
}
