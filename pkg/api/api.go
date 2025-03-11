package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ChatHandler handles chat messages and sends them to Azure OpenAI
func ChatHandler(c echo.Context) error {
	// Parse incoming message
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid message format"})
	}

	// Retrieve past messages from context
	pastMessages := getPastMessages(c)

	// Send message to Azure OpenAI
	response, err := SendToAzureOpenAI(message, pastMessages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to communicate with Azure OpenAI"})
	}

	return c.JSON(http.StatusOK, response)
}

// getPastMessages retrieves up to 10 past messages from the context
func getPastMessages(c echo.Context) []Message {
	// Logic to retrieve past messages
	return []Message{}
}
