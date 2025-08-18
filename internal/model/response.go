package model

// APIResponse represents the standard API response format
type APIResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// ImageGenerationData represents response data for image generation
type ImageGenerationData struct {
	Model  string `json:"model"`
	JobID  string `json:"jobId"`
	Prompt string `json:"prompt"`
	Pic    string `json:"pic"`
}

// TextGenerationData represents response data for text generation
type TextGenerationData struct {
	Model           string `json:"model"`
	MessageID       string `json:"messageId"`
	AllContent      string `json:"all_content"`
	ThinkingContent string `json:"thinking_content"`
	PureContent     string `json:"pure_content"`
}

// ErrorData represents error response data
type ErrorData struct {
	Message string `json:"msg"`
}

// SessionResponse represents the session response from Akash
type SessionResponse struct {
	SessionToken string `json:"session_token,omitempty"`
}

// OpenAIChatCompletion represents the chat completion response in OpenAI format.
type OpenAIChatCompletion struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice represents a single choice in the chat completion response.
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Message represents a message in the chat completion response.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Usage represents the token usage statistics.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}