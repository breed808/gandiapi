package gandigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Snapshot represents information about a given snapshot.
type Snapshot struct {
	DateCreated time.Time `json:"date_created"`
	Automatic   bool      `json:"automatic"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
}

// SnapshotContent represents details of a snapshot.
type SnapshotContent struct {
	ZoneUUID string `json:"zone_uuid"`
	UUID     string `json:"uuid"`
	ZoneData []struct {
		RrsetType   string   `json:"rrset_type"`
		RrsetTTL    int      `json:"rrset_ttl"`
		RrsetName   string   `json:"rrset_name"`
		RrsetValues []string `json:"rrset_values"`
	} `json:"zone_data"`
	DateCreated time.Time `json:"date_created"`
	Automatic   bool      `json:"automatic"`
	Name        string    `json:"name"`
}

// SnapshotCreate reprensets a response when creating a new snapshot.
type SnapshotCreate struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

// GetSnapshots returns a list of snapshots
func (c *Client) GetSnapshots(zoneID string) ([]Snapshot, error) {
	urlRecords := fmt.Sprintf("%s/zones/%s/snapshots", defaultBaseURL, zoneID)
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

	snapshotResponse := make([]Snapshot, 0)
	err = json.NewDecoder(resp.Body).Decode(&snapshotResponse)
	if err != nil {
		return nil, err
	}

	return snapshotResponse, nil
}

// GetSnapshotDetails returns details of a snapshot
func (c *Client) GetSnapshotDetails(zoneID, snapshotID string) (*SnapshotContent, error) {
	urlRecords := fmt.Sprintf("%s/zones/%s/snapshots/%s", defaultBaseURL, zoneID, snapshotID)
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

	snapshotResponse := new(SnapshotContent)
	err = json.NewDecoder(resp.Body).Decode(snapshotResponse)
	if err != nil {
		return nil, err
	}

	return snapshotResponse, nil
}

// CreateSnapshot creates a new snapshopt for the given zone.
func (c *Client) CreateSnapshot(zoneID string) (*SnapshotCreate, error) {
	urlRecords := fmt.Sprintf("%s/zones/%s/snapshots", defaultBaseURL, zoneID)
	u, err := url.Parse(urlRecords)
	if err != nil {
		return nil, err
	}

	req := http.Request{
		URL:    u,
		Header: make(http.Header),
		Method: http.MethodPost,
	}

	req.Header.Add("X-Api-Key", c.APIKey)
	req.Header.Add("Content-Type", "application/json")

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

	snapshotResponse := new(SnapshotCreate)
	err = json.NewDecoder(resp.Body).Decode(snapshotResponse)
	if err != nil {
		return nil, err
	}

	return snapshotResponse, nil
}
