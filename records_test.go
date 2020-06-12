package gandigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetRecords(t *testing.T) {
	responseData := []byte(`
 [
  {
    "rrset_type": "A",
    "rrset_ttl": 1800,
    "rrset_name": "@",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/40/A",
    "rrset_values": [
      "45.79.62.185"
    ]
  },
  {
    "rrset_type": "MX",
    "rrset_ttl": 10800,
    "rrset_name": "@",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/40/MX",
    "rrset_values": [
      "10 spool.mail.gandi.net.",
      "50 fb.mail.gandi.net."
    ]
  },
  {
    "rrset_type": "TXT",
    "rrset_ttl": 10800,
    "rrset_name": "@",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/40/TXT",
    "rrset_values": [
      "\"v=spf1 include:_mailcust.gandi.net ?all\""
    ]
  },
  {
    "rrset_type": "CNAME",
    "rrset_ttl": 10800,
    "rrset_name": "blog",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/blog/CNAME",
    "rrset_values": [
      "blogs.vip.gandi.net."
    ]
  },
  {
    "rrset_type": "A",
    "rrset_ttl": 1800,
    "rrset_name": "raneto",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/raneto/A",
    "rrset_values": [
      "45.79.62.185"
    ]
  },
  {
    "rrset_type": "CNAME",
    "rrset_ttl": 10800,
    "rrset_name": "webmail",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/webmail/CNAME",
    "rrset_values": [
      "webmail.gandi.net."
    ]
  },
  {
    "rrset_type": "CNAME",
    "rrset_ttl": 10800,
    "rrset_name": "www",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/www/CNAME",
    "rrset_values": [
      "webredir.vip.gandi.net."
    ]
  },
  {
    "rrset_type": "AAAA",
    "rrset_ttl": 1800,
    "rrset_name": "x23",
    "rrset_href": "https://dns.api.gandi.net/api/v5/zones/12345678/records/x23/AAAA",
    "rrset_values": [
      "2a02:8108:13b:bb00:211:32ff:fe15:d122"
    ]
  }
]
 `)
	mockZoneID := "12345678"

	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/zones/12345678/records", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(responseData))
	})

	records, err := client.GetRecords(mockZoneID)
	if err != nil {
		t.Errorf("got error with GetZones %v", err)
	}

	var expected []Record
	err = json.Unmarshal(responseData, &expected)
	if err != nil {
		t.Errorf("got error with Unmarshal %v %v", err, records)
	}

	if !reflect.DeepEqual(records, expected) {
		t.Errorf("client.GetRecords() got %v expected %v\n", records, expected)
	}
}

func TestCreateRecord(t *testing.T) {
	mockZoneID := "12345678"
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/zones/12345678/records", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string([]byte(`{"message": "DNS Record Created"}`)))
	})

	data := Record{
		RrsetType:   "A",
		RrsetTTL:    300,
		RrsetName:   "amazing-cli",
		RrsetValues: []string{"18.185.88.103"},
	}

	err := client.CreateRecord(data, mockZoneID)
	if err != nil {
		t.Errorf("got error with CreateError %v", err)
	}
}
