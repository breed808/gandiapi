package gandiapi

import (
	"io"
	"net/http"
)

// create_request combines request details with an API key and outputs a request.
func create_request(reqURL, httpMethod, APIKey string, headers map[string]string, data io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, reqURL, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Apikey "+APIKey)

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}
