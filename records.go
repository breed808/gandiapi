package gandigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		return nil, ErrHTTPForbidden
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
		return ErrHTTPForbidden
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

// GetRecordsText returns a text version of the zone.
func (c *Client) GetRecordsText(zoneID string) (*string, error) {
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
	req.Header.Add("Accept", "text/plain")

	resp, err := c.http.Do(&req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return nil, ErrHTTPForbidden
	case http.StatusUnauthorized:
		return nil, ErrNonAPIKey
	}
	defer resp.Body.Close()

	textRecordData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	text := string(textRecordData)

	return &text, nil
}

// DeleteRecord removes a given record name
func (c *Client) DeleteRecord(zoneID, recordName string) error {
	urlRecords := fmt.Sprintf("%s/zones/%s/records/%s", defaultBaseURL, zoneID, recordName)
	u, err := url.Parse(urlRecords)
	if err != nil {
		return err
	}

	req := http.Request{
		URL:    u,
		Header: make(http.Header),
		Method: http.MethodDelete,
	}

	req.Header.Add("X-Api-Key", c.APIKey)

	resp, err := c.http.Do(&req)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return ErrHTTPForbidden
	case http.StatusUnauthorized:
		return ErrNonAPIKey
	case http.StatusBadRequest:
		return ErrBadRequest
	}
	defer resp.Body.Close()

	return nil
}

// DeleteRecordsZone removes all the records in a given zone.
// TODO: do testing, I don't want to delete all records of one of my zones.
func (c *Client) DeleteRecordsZone(zoneID string) error {
	urlRecords := fmt.Sprintf("%s/zones/%s/records", defaultBaseURL, zoneID)
	u, err := url.Parse(urlRecords)
	if err != nil {
		return err
	}

	req := http.Request{
		URL:    u,
		Header: make(http.Header),
		Method: http.MethodDelete,
	}

	req.Header.Add("X-Api-Key", c.APIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.http.Do(&req)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return ErrHTTPForbidden
	case http.StatusUnauthorized:
		return ErrNonAPIKey
	case http.StatusBadRequest:
		return ErrBadRequest
	}
	defer resp.Body.Close()

	return nil
}
