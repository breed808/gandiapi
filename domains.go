package gandiapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ErrNonAPIKey holds errors when API key is not found.
var ErrNonAPIKey = errors.New("Bad authentication attempt due to incorrect API key")

// ErrHTTPForbidden returns HTTP 403.
var ErrHTTPForbidden = errors.New("request forbidden")

// ErrBadRequest returns HTTP 400.
var ErrBadRequest = errors.New("HTTP 400 bad request")

var ErrNotFound = errors.New("HTTP 404 API object not found")

var ErrAlreadyExists = errors.New("HTTP 409 API object already exists")

// DomainDates stores dates for a retrieved FQDN.
type DomainDates struct {
	RegistryCreatedAt   time.Time `json:"registry_created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	AuthinfoExpiresAt   time.Time `json:"authinfo_expires_at,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	DeletesAt           time.Time `json:"deletes_at,omitempty"`
	HoldBeginsAt        time.Time `json:"hold_begins_at,omitempty"`
	HoldEndsAt          time.Time `json:"hold_ends_at,omitempty"`
	PendingDeleteEndsAt time.Time `json:"pending_delete_ends_at,omitempty"`
	RegistryEndsAt      time.Time `json:"registry_ends_at,omitempty"`
	RenewBeginsAt       time.Time `json:"renew_begins_at,omitempty"`
	RestoreEndsAt       time.Time `json:"restore_ends_at,omitempty"`
}

// DomainNameserver stores nameserver information for a retrieved FQDN.
type DomainNameserver struct {
	Current string   `json:"current"`
	Hosts   []string `json:"hosts,omitempty"`
}

// DomainResponse struct stores response from zones endpoint.
type DomainResponse struct {
	Autorenew   bool             `json:"autorenew"`
	Dates       DomainDates      `json:"dates"`
	DomainOwner string           `json:"domain_owner"`
	Fqdn        string           `json:"fqdn"`
	FqdnUnicode string           `json:"fqdn_unicode"`
	Href        string           `json:"href"`
	ID          string           `json:"id"`
	Nameserver  DomainNameserver `json:"nameserver"`
	OrgaOwner   string           `json:"orga_owner"`
	Owner       string           `json:"owner"`
	SharingID   string           `json:"sharing_id,omitempty"`
	Status      []string         `json:"status"`
	Tags        []string         `json:"tags,omitempty"`
	Tld         string           `json:"tld"`
}

// GetDomains retrieves a list of zones.
func (c *Client) GetDomains() ([]DomainResponse, error) {
	reqURL := fmt.Sprintf("%s/domain/domains", c.defaultBaseURL)
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

	domainResponse := make([]DomainResponse, 0)
	err = json.Unmarshal(body, &domainResponse)
	if err != nil {
		return nil, err
	}

	return domainResponse, nil
}
