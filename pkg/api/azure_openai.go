package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
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
		log.Errorf("JSON marshalling error: %v", err)
		return Response{}, err
	}

	// Get Azure OpenAI endpoint, API key, and API version from environment variables
	endpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	apiKey := os.Getenv("AZURE_OPENAI_API_KEY")

	// Validate endpoint, API key, and API version
	if endpoint == "" || apiKey == "" {
		log.Error("Azure OpenAI endpoint or API key missing")
		return Response{}, fmt.Errorf("Azure OpenAI endpoint or API key missing")
	}

	// Log the outgoing request data
	log.Infof("Sending request to Azure OpenAI with payload: %s", data)

	// Create HTTP request with api-version as query parameters
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		log.Errorf("Failed to create HTTP request: %v", err)
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey) // Use 'api-key' for header

	// Send request to Azure OpenAI
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("HTTP request to Azure OpenAI failed: %v", err)
		return Response{}, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Warnf("Error closing response body: %v", err)
		}
	}()

	// Log the response status and headers
	log.Infof("Received response from Azure OpenAI with status: %s", resp.Status)

	// Parse response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Errorf("Failed to decode Azure OpenAI response: %v", err)
		return Response{}, err
	}

	// Log all choice messages in response
	for _, choice := range response.Choices {
		log.Infof("Received message from Azure OpenAI: %+v", choice.Message)
	}

	return response, nil
}
