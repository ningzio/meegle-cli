package meegle

import "net/http"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(baseURL string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
}
