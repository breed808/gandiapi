package gandigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetZones(t *testing.T) {
	responseData := []byte(`
[
  {
    "retry": 3600,
    "uuid": "539119b8-9008-4e90-b65c-12eb0fdfc245",
    "zone_href": "https://dns.api.gandi.net/api/v5/zones/539119b8-9008-4e90-b65c-12eb0fdfc245",
    "minimum": 10800,
    "domains_href": "https://dns.api.gandi.net/api/v5/zones/539119b8-9008-4e90-b65c-12eb0fdfc245/domains",
    "refresh": 10800,
    "zone_records_href": "https://dns.api.gandi.net/api/v5/zones/539119b8-9008-4e90-b65c-12eb0fdfc245/records",
    "expire": 604800,
    "sharing_id": "d510c560-35fb-4d69-be5a-00a0848ef0a2",
    "serial": 2222222222,
    "email": "hostmaster.gandi.net.",
    "primary_ns": "ns1.gandi.net",
    "name": "example.net"
  },
  {
    "retry": 3600,
    "uuid": "b45167c7-31fa-4665-9325-90db47c6aee3",
    "zone_href": "https://dns.api.gandi.net/api/v5/zones/b45167c7-31fa-4665-9325-90db47c6aee3",
    "minimum": 10800,
    "domains_href": "https://dns.api.gandi.net/api/v5/zones/b45167c7-31fa-4665-9325-90db47c6aee3/domains",
    "refresh": 10800,
    "zone_records_href": "https://dns.api.gandi.net/api/v5/zones/b45167c7-31fa-4665-9325-90db47c6aee3/records",
    "expire": 604800,
    "sharing_id": "d510c560-35fb-4d69-be5a-00a0848ef0a2",
    "serial": 1111111111,
    "email": "hostmaster.gandi.net.",
    "primary_ns": "ns1.gandi.net",
    "name": "ontheway.com"
  }
]
`)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	options := OptsClient{APIURL: server.URL}

	client, err := NewClient(&options)
	if err != nil {
		t.Errorf("got error with NewClient %v", err)
	}

	mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(responseData))
	})

	zones, err := client.GetZones()
	if err != nil {
		t.Errorf("got error with GetZones %v", err)
	}

	var expected []ZoneResponse
	err = json.Unmarshal(responseData, &expected)
	if err != nil {
		t.Errorf("got error with Unmarshal %v %v", err, zones)
	}

	if !reflect.DeepEqual(zones, expected) {
		t.Errorf("client.Getzones got %v expected %v\n", zones, expected)
	}

}
