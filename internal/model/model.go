package model

// Model represents the structure of a model's information
type Model struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Temperature    float64  `json:"temperature,omitempty"`
	TopP           float64  `json:"top_p,omitempty"`
	TokenLimit     int      `json:"tokenLimit,omitempty"`
	Parameters     string   `json:"parameters,omitempty"`
	Architecture   string   `json:"architecture,omitempty"`
	HFRepo         string   `json:"hf_repo,omitempty"`
	AboutContent   string   `json:"aboutContent"`
	InfoContent    string   `json:"infoContent"`
	ThumbnailID    string   `json:"thumbnailId"`
	DeployURL      string   `json:"deployUrl,omitempty"`
	Available      bool     `json:"available"`
}