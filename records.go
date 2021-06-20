package gandiapi

import (
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
func (c *Client) GetRecords(fqdn string) (records []Record, err error) {
	_, err = c.get("livedns/domains/"+fqdn+"/records", nil, &records)
	return
}

// GetRecordsText returns a text version of the zone.
func (c *Client) GetRecordsText(fqdn string) (string, error) {
	var resp http.Request

	_, err := c.get("livedns/domains/"+fqdn+"/records", nil, &resp)
	if err != nil {
		return "", err
	}

	textRecordData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	text := string(textRecordData)

	return text, nil
}

// CreateRecord adds a new record for a given zone.
func (c *Client) CreateRecord(data Record, fqdn string) (response StandardResponse, err error) {
	_, err = c.post("livedns/domains/"+fqdn+"/records", data, &response)
	return
}

// GetRecord returns a the single requested record for the given fqdn, if it exists.
func (c *Client) GetRecord(fqdn, recordName, recordType string) (record Record, err error) {
	_, err = c.get("livedns/domains/"+fqdn+"/records/"+recordName+"/"+recordType, nil, &record)
	return
}

// UpdateRecord updates an existing record with the provided value.
func (c *Client) UpdateRecord(data Record, fqdn string) (response StandardResponse, err error) {
	c.put("livedns/domains/"+fqdn+"/records/"+data.RrsetName+"/"+data.RrsetType, data, &response)
	return
}

// DeleteRecord removes a given record name.
func (c *Client) DeleteRecord(fqdn, recordName, recordType string) (err error) {
	c.put("livedns/domains/"+fqdn+"/records/"+recordName+"/"+recordType, nil, nil)
	return
}
