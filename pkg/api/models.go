package api

// Message represents a chat message
type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"` // Role can be "assistant" or "user"
}

// Request represents the request body for Azure OpenAI
type Request struct {
	Messages     []Message              `json:"messages"`
	Context      map[string]interface{} `json:"context,omitempty"`
	SessionState interface{}            `json:"session_state,omitempty"`
}

// Response represents the response from Azure OpenAI
type Response struct {
	Content string `json:"content"`
}
