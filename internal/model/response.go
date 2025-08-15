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