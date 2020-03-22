package gandigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Record represents a Record object
type Record struct {
	RrsetType   string   `json:"rrset_type"`
	RrsetTTL    int      `json:"rrset_ttl"`
	RrsetName   string   `json:"rrset_name"`
	RrsetHref   string   `json:"rrset_href"`
	RrsetValues []string `json:"rrset_values"`
}

// GetRecords returns a slice of Record
func (c *Client) GetRecords(zoneID string) ([]Record, error) {
	urlRecords := fmt.Sprintf("%s/zones/%s/records", defaultBaseURL, zoneID)
	u, err := url.Parse(urlRecords)
	if err != nil {
		return nil, err
	}

	req := http.Request{
		URL:    u,
		Header: make(http.Header),
	}

	req.Header.Add("X-Api-Key", c.APIKey)

	resp, err := c.http.Do(&req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return nil, ErrHttpForbidden
	case http.StatusUnauthorized:
		return nil, ErrNonAPIKey
	}
	defer resp.Body.Close()

	recordResponse := make([]Record, 0)
	err = json.NewDecoder(resp.Body).Decode(&recordResponse)
	if err != nil {
		return nil, err
	}

	return recordResponse, nil
}
