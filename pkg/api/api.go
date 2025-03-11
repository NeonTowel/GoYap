package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ChatHandler handles chat messages and sends them to Azure OpenAI
func ChatHandler(c echo.Context) error {
	// Parse incoming message array
	var request struct {
		Messages []Message `json:"messages"`
	}
	if err := c.Bind(&request); err != nil {
		c.Logger().Errorf("Failed to bind messages: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid message format"})
	}

	// Assume the last message is the current user's input
	currentMessage := request.Messages[len(request.Messages)-1]
	pastMessages := request.Messages[:len(request.Messages)-1]

	// Send messages to Azure OpenAI
	response, err := SendToAzureOpenAI(currentMessage, pastMessages)
	if err != nil {
		c.Logger().Errorf("Failed to communicate with Azure OpenAI: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to communicate with Azure OpenAI"})
	}

	return c.JSON(http.StatusOK, response)
}
