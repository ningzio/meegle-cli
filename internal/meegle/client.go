package meegle

import "net/http"

// Client holds configuration for Meegle API requests.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient creates a Meegle API client targeting the provided base URL.
func NewClient(baseURL string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
}
