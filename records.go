package gandiapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Record represents a Record object.
type Record struct {
	RrsetType   string   `json:"rrset_type"`
	RrsetTTL    int      `json:"rrset_ttl,omitempty"`
	RrsetName   string   `json:"rrset_name"`
	RrsetHref   string   `json:"rrset_href,omitempty"`
	RrsetValues []string `json:"rrset_values"`
}

// GetRecords returns a slice of Record.
func (c *Client) GetRecords(fqdn string) ([]Record, error) {
	reqURL := fmt.Sprintf("%s/livedns/domains/%s/records", c.defaultBaseURL, fqdn)
	req, err := create_request(reqURL, http.MethodGet, c.APIKey, nil, nil)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	recordResponse := make([]Record, 0)
	err = json.Unmarshal(body, &recordResponse)
	if err != nil {
		return nil, err
	}

	return recordResponse, nil
}

// GetRecordsText returns a text version of the zone.
func (c *Client) GetRecordsText(fqdn string) (string, error) {
	reqURL := fmt.Sprintf("%s/livedns/domains/%s/records", c.defaultBaseURL, fqdn)

	extraHeaders := make(map[string]string)
	extraHeaders["Accept"] = "text/plain"
	req, err := create_request(reqURL, http.MethodGet, c.APIKey, extraHeaders, nil)
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

// CreateRecord adds a new record for a given zone.
func (c *Client) CreateRecord(data Record, fqdn string) error {
	reqURL := fmt.Sprintf("%s/livedns/domains/%s/records", c.defaultBaseURL, fqdn)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	dataSend := bytes.NewReader(dataJSON)
	extraHeaders := make(map[string]string)
	extraHeaders["Content-Type"] = "application/json"
	req, err := create_request(reqURL, http.MethodPost, c.APIKey, extraHeaders, dataSend)
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
	case http.StatusConflict:
		return ErrAlreadyExists
	}

	return nil
}

// GetRecord returns a the single requested record for the given fqdn, if it exists.
func (c *Client) GetRecord(fqdn, recordName, recordType string) (Record, error) {
	reqURL := fmt.Sprintf("%s/livedns/domains/%s/record/%s/%s", c.defaultBaseURL, fqdn, recordName, recordType)
	req, err := create_request(reqURL, http.MethodGet, c.APIKey, nil, nil)

	recordResponse := Record{}

	if err != nil {
		return recordResponse, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return recordResponse, err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return recordResponse, ErrHTTPForbidden
	case http.StatusUnauthorized:
		return recordResponse, ErrNonAPIKey
	case http.StatusNotFound:
		return recordResponse, ErrNotFound
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return recordResponse, err
	}

	err = json.Unmarshal(body, &recordResponse)
	if err != nil {
		return recordResponse, err
	}

	return recordResponse, nil
}

// UpdateRecord updates an existing record with the provided value.
func (c *Client) UpdateRecord(data Record, fqdn string) error {
	reqURL := fmt.Sprintf("%s/livedns/domains/%s/records/%s/%s", c.defaultBaseURL, fqdn, data.RrsetName, data.RrsetType)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	dataSend := bytes.NewReader(dataJSON)
	extraHeaders := make(map[string]string)
	extraHeaders["Content-Type"] = "application/json"
	req, err := create_request(reqURL, http.MethodPut, c.APIKey, extraHeaders, dataSend)
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
	case http.StatusNotFound:
		return ErrNotFound
	}

	return nil
}

// DeleteRecord removes a given record name.
func (c *Client) DeleteRecord(fqdn, recordName, recordType string) error {
	reqURL := fmt.Sprintf("%s/livedns/domains/%s/records/%s/%s", c.defaultBaseURL, fqdn, recordName, recordType)
	req, err := create_request(reqURL, http.MethodDelete, c.APIKey, nil, nil)
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
