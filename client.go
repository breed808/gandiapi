package gandiapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

// Client manages requests Gandi API.
type Client struct {
	apiKey     string
	endpoint   string
	debug      bool
	dryRun     bool
	httpClient *http.Client
}

// NewClient returns a client.
func NewClient(apiKey string, debug, dryRun bool) *Client {
	return &Client{apiKey: apiKey, endpoint: "https://api.gandi.net/v5/", debug: debug, dryRun: dryRun, httpClient: &http.Client{}}
}

func (c *Client) ask(method, path string, params, recipient interface{}) (http.Header, error) {
	resp, err := c.do(method, path, params, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(recipient)

	return resp.Header, nil
}

type StandardResponse struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	UUID    string          `json:"uuid,omitempty"`
	Object  string          `json:"object,omitempty"`
	Cause   string          `json:"cause,omitempty"`
	Status  string          `json:"status,omitempty"`
	Errors  []StandardError `json:"errors,omitempty"`
}

type StandardError struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Client) do(method, path string, p interface{}, extraHeaders [][2]string) (*http.Response, error) {
	var (
		err error
		req *http.Request
	)

	params, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	suffix := ""

	if params != nil && string(params) != "null" {
		req, err = http.NewRequest(method, c.endpoint+path+suffix, bytes.NewReader(params))
	} else {
		req, err = http.NewRequest(method, c.endpoint+path+suffix, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Apikey "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	if c.dryRun {
		req.Header.Add("Dry-Run", "1")
	}

	for _, header := range extraHeaders {
		req.Header.Add(header[0], header[1])
	}

	if c.debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println("=======================================\nREQUEST:")
		fmt.Println(string(dump))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	if c.debug {
		dump, _ := httputil.DumpResponse(resp, true)

		fmt.Println("=======================================\nRESPONSE:")
		fmt.Println(string(dump))
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		defer resp.Body.Close()

		var message StandardResponse

		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&message)
		if message.Message != "" {
			err = fmt.Errorf("%d: %s", resp.StatusCode, message.Message)
		} else if len(message.Errors) > 0 {
			var errors []string

			for _, oneError := range message.Errors {
				errors = append(errors, fmt.Sprintf("%s: %s", oneError.Name, oneError.Description))
			}

			err = fmt.Errorf(strings.Join(errors, ", "))
		} else {
			err = fmt.Errorf("%d", resp.StatusCode)
		}
	}

	return resp, err
}

func (c *Client) get(path string, params, recipient interface{}) (http.Header, error) {
	return c.ask(http.MethodGet, path, params, recipient)
}

func (c *Client) post(path string, params, recipient interface{}) (http.Header, error) {
	return c.ask(http.MethodPost, path, params, recipient)
}

func (c *Client) put(path string, params, recipient interface{}) (http.Header, error) {
	return c.ask(http.MethodPut, path, params, recipient)
}

func (c *Client) patch(path string, params, recipient interface{}) (http.Header, error) {
	return c.ask(http.MethodPatch, path, params, recipient)
}

func (c *Client) delete(path string, params, recipient interface{}) (http.Header, error) {
	return c.ask(http.MethodDelete, path, params, recipient)
}
