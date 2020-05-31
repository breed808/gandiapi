package gandigo

import (
	"net/http"
	"os"
)

// Client manages requests Gandi API.
type Client struct {
	http *http.Client

	APIKey         string
	defaultBaseURL string
}

// OptsClient allows to define options for the client
type OptsClient struct {
	APIURL string
}

// NewClient returns a client.
func NewClient(opts *OptsClient) (*Client, error) {
	APIKey := os.Getenv("GANDI_API_KEY")

	if opts == nil {
		opts = &OptsClient{}
		opts.APIURL = "https://dns.api.gandi.net/api/v5/"
	}

	c := &Client{
		http:           &http.Client{},
		APIKey:         APIKey,
		defaultBaseURL: opts.APIURL,
	}
	return c, nil
}
