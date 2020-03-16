package gandigo

import (
	"net/http"
	"os"
)

// Client manages requests Gandi API.
type Client struct {
	http *http.Client

	APIKey string
}

// NewClient returns a client.
func NewClient() (*Client, error) {
	APIKey := os.Getenv("GANDI_API_KEY")
	c := &Client{
		http:   &http.Client{},
		APIKey: APIKey,
	}
	return c, nil
}
