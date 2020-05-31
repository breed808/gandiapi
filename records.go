package gandigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sgmac/gandigo/internal/requests"
)

// Record represents a Record object.
type Record struct {
	RrsetType   string   `json:"rrset_type"`
	RrsetTTL    int      `json:"rrset_ttl"`
	RrsetName   string   `json:"rrset_name"`
	RrsetHref   string   `json:"rrset_href,omitempty"`
	RrsetValues []string `json:"rrset_values"`
}

// GetRecords returns a slice of Record.
func (c *Client) GetRecords(zoneID string) ([]Record, error) {
	reqURL := fmt.Sprintf("%s/zones/%s/records", defaultBaseURL, zoneID)
	req, err := requests.Do(reqURL, http.MethodGet, c.APIKey, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
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
	reqURL := fmt.Sprintf("%s/zones/%s/records", defaultBaseURL, zoneID)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	dataSend := bytes.NewReader(dataJSON)
	extraHeaders := make(map[string]string)
	extraHeaders["Content-Type"] = "application/json"
	req, err := requests.Do(reqURL, http.MethodPost, c.APIKey, extraHeaders, dataSend)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
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
		return err
	}

	return nil
}

// GetRecordsText returns a text version of the zone.
func (c *Client) GetRecordsText(zoneID string) (string, error) {
	reqURL := fmt.Sprintf("%s/zones/%s/records", defaultBaseURL, zoneID)

	extraHeaders := make(map[string]string)
	extraHeaders["Accept"] = "text/plain"
	req, err := requests.Do(reqURL, http.MethodGet, c.APIKey, extraHeaders, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return "", ErrHTTPForbidden
	case http.StatusUnauthorized:
		return "", ErrNonAPIKey
	}
	defer resp.Body.Close()

	textRecordData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	text := string(textRecordData)

	return text, nil
}

// DeleteRecord removes a given record name.
func (c *Client) DeleteRecord(zoneID, recordName string) error {
	reqURL := fmt.Sprintf("%s/zones/%s/records/%s", defaultBaseURL, zoneID, recordName)
	req, err := requests.Do(reqURL, http.MethodDelete, c.APIKey, nil, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
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
	reqURL := fmt.Sprintf("%s/zones/%s/records", defaultBaseURL, zoneID)

	extraHeaders := make(map[string]string)
	extraHeaders["Content-Type"] = "application/json"
	req, err := requests.Do(reqURL, http.MethodDelete, c.APIKey, extraHeaders, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
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

// DeleteRecordType deletes a record that matches name and type.
func (c *Client) DeleteRecordType(zoneID, recordName, recordType string) error {
	reqURL := fmt.Sprintf("%s/zones/%s/records/%s/%s", defaultBaseURL, zoneID, recordName, recordType)

	extraHeaders := make(map[string]string)
	extraHeaders["Content-Type"] = "application/json"
	req, err := requests.Do(reqURL, http.MethodDelete, c.APIKey, extraHeaders, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
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
