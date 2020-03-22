package gandigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Record represents a Record object
type Record struct {
	RrsetType   string   `json:"rrset_type"`
	RrsetTTL    int      `json:"rrset_ttl"`
	RrsetName   string   `json:"rrset_name"`
	RrsetHref   string   `json:"rrset_href,omitempty"`
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
		Method: http.MethodGet,
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

// CreateRecord adds a new record for a given zone.
func (c *Client) CreateRecord(data Record, zoneID string) error {
	urlRecords := fmt.Sprintf("%szones/%s/records", defaultBaseURL, zoneID)
	dataJSON, err := json.Marshal(data)
	fmt.Println(string(dataJSON))
	if err != nil {
		return err
	}

	fmt.Println(data)
	dataSend := bytes.NewReader(dataJSON)
	req, err := http.NewRequest(http.MethodPost, urlRecords, dataSend)

	req.Header.Add("X-Api-Key", c.APIKey)
	req.Header.Add("Content-Type", "application/json")
	fmt.Println(req.Header)

	resp, err := c.http.Do(req)
	if err != nil {
		log.Println("Do request")
		return err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return ErrHttpForbidden
	case http.StatusUnauthorized:
		return ErrNonAPIKey
	}
	defer resp.Body.Close()

	createResponse := struct{ Message string }{}

	err = json.NewDecoder(resp.Body).Decode(&createResponse)
	if err != nil {
		log.Println("DECODING")
		return err
	}
	fmt.Println(createResponse)

	return nil
}
