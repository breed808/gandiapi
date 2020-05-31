package requests

import (
	"io"
	"net/http"
)

// Do represents and http request for a given url.
func Do(reqURL, httpMethod, APIKey string, headers map[string]string, data io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, reqURL, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Api-Key", APIKey)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return req, nil
}
