package requests

import (
	"fmt"
	"net/http"
	"net/url"
)

// Do represents and http request for a given url
func Do(defaultBaseURL, httpMethod, zoneID, APIKey string, headers map[string]string) (http.Request, error) {
	urlRecords := fmt.Sprintf("%s/zones/%s/records", defaultBaseURL, zoneID)
	u, err := url.Parse(urlRecords)
	if err != nil {
		return http.Request{}, err
	}

	req := http.Request{
		URL:    u,
		Header: make(http.Header),
		Method: httpMethod,
	}

	req.Header.Add("X-Api-Key", APIKey)

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}
