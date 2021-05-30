package gandiapi

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

// OptsClient allows to define options for the client.
type OptsClient struct {
	APIURL string
	APIKey string
}

// NewClient returns a client.
func NewClient(opts *OptsClient) (*Client, error) {
	if opts == nil {
		opts = &OptsClient{}
		opts.APIURL = "https://api.gandi.net/v5/"
		opts.APIKey = os.Getenv("GANDI_API_KEY")
	}

	c := &Client{
		http:           &http.Client{},
		APIKey:         opts.APIKey,
		defaultBaseURL: opts.APIURL,
	}

	return c, nil
}
