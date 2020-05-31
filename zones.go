package gandigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sgmac/gandigo/internal/requests"
)

// ErrNonAPIKey holds errors when API key is not found.
var ErrNonAPIKey = errors.New("not API Key defined")

// ErrHTTPForbidden returns HTTP 403.
var ErrHTTPForbidden = errors.New("request forbidden")

// ErrBadRequest returns HTTP 400.
var ErrBadRequest = errors.New("HTTP 400 bad request")

// ZoneResponse struct stores response from zones endpoint.
type ZoneResponse struct {
	Retry           int    `json:"retry"`
	UUID            string `json:"uuid"`
	ZoneHref        string `json:"zone_href"`
	Minimum         int    `json:"minimum"`
	DomainsHref     string `json:"domains_href"`
	Refresh         int    `json:"refresh"`
	ZoneRecordsHref string `json:"zone_records_href"`
	Expire          int    `json:"expire"`
	SharingID       string `json:"sharing_id"`
	Serial          int    `json:"serial"`
	Email           string `json:"email"`
	PrimaryNs       string `json:"primary_ns"`
	Name            string `json:"name"`
}

// GetZones retrieves a list of zones.
func (c *Client) GetZones() ([]ZoneResponse, error) {
	reqURL := fmt.Sprintf("%s/zones", c.defaultBaseURL)
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

	zoneResponse := make([]ZoneResponse, 0)
	err = json.NewDecoder(resp.Body).Decode(&zoneResponse)
	if err != nil {
		return nil, err
	}

	return zoneResponse, nil
}
