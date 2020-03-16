package gandigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://dns.api.gandi.net/api/v5/"

// ErrErrNonAPIKey holds errors when API key is not found.
var ErrNonAPIKey = errors.New("not API Key defined")

// ErrHttpForbidden returns HTTP 403
var ErrHttpForbidden = errors.New("request forbidden")

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

// ListZones retrieves a list of zones.
func (c *Client) ListZones() ([]ZoneResponse, error) {
	urlZones := fmt.Sprintf("%s/zones", defaultBaseURL)
	u, err := url.Parse(urlZones)
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

	zoneResponse := make([]ZoneResponse, 0)
	err = json.NewDecoder(resp.Body).Decode(&zoneResponse)
	if err != nil {
		return nil, err
	}

	return zoneResponse, nil
}
