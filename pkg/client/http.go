package client

import (
	"io"
	"net/http"
	"time"
)

// HTTPClient wraps http.Client with additional functionality
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTPClient instance
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// Get performs a GET request with optional headers
func (c *HTTPClient) Get(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

// Post performs a POST request with body and optional headers
func (c *HTTPClient) Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// Add headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}