package gandiapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetDomains(t *testing.T) {
	responseData := []byte(`
[
  {
    "status": [
      "clientTransferProhibited"
    ],
    "dates": {
      "created_at": "2019-02-13T11:04:18Z",
      "registry_created_at": "2019-02-13T10:04:18Z",
      "registry_ends_at": "2021-02-13T10:04:18Z",
      "updated_at": "2019-02-25T16:20:49Z"
    },
    "tags": [],
    "fqdn": "example.net",
    "id": "ba1167be-2f76-11e9-9dfb-00163ec4cb00",
    "autorenew": false,
    "tld": "net",
    "owner": "alice_doe",
    "orga_owner": "alice_doe",
    "domain_owner": "Alice Doe",
    "nameserver": {
      "current": "livedns"
    },
    "href": "https://api.test/v5/domain/domains/example.net",
    "fqdn_unicode": "example.net"
  },
  {
    "status": [],
    "dates": {
      "created_at": "2019-01-15T14:19:59Z",
      "registry_created_at": "2019-01-15T13:19:58Z",
      "registry_ends_at": "2020-01-15T13:19:58Z",
      "updated_at": "2019-01-15T13:30:42Z"
    },
    "tags": [],
    "fqdn": "example.com",
    "id": "42927d64-18c8-11e9-b9b5-00163ec4cb00",
    "autorenew": false,
    "tld": "fr",
    "owner": "alice_doe",
    "orga_owner": "alice_doe",
    "domain_owner": "Alice Doe",
    "nameserver": {
      "current": "livedns"
    },
    "href": "https://api.test/v5/domain/domains/example.com",
    "fqdn_unicode": "example.com"
  }
]
`)

	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/domain/domains", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(responseData))
	})

	domains, err := client.GetDomains()
	if err != nil {
		t.Errorf("got error with GetDomains %v", err)
	}

	var expected []DomainResponse
	err = json.Unmarshal(responseData, &expected)
	if err != nil {
		t.Errorf("got error with Unmarshal %v %v", err, domains)
	}

	if !reflect.DeepEqual(domains, expected) {
		t.Errorf("client.GetDomains got %v expected %v\n", domains, expected)
	}

}
