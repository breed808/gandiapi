package gandigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetSnapshotsNoEmpty(t *testing.T) {
	responseData := []byte(`
[
  {
    "date_created": "2020-04-05T13:23:46Z",
    "automatic": true,
    "uuid": "uuid1",
    "name": "2020-04-05 #22"
  },
  {
    "date_created": "2020-04-05T13:42:00Z",
    "automatic": false,
    "uuid": "uuid2",
    "name": "2020-04-05 #23"
  },
  {
    "date_created": "2020-04-05T13:47:00Z",
    "automatic": false,
    "uuid": "uuid3",
    "name": "2020-04-05 #24"
  }
 ]
`)

	mockZoneID := "12345678"
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/zones/12345678/snapshots", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(responseData))
	})

	snapshots, err := client.GetSnapshots(mockZoneID)
	if err != nil {
		t.Errorf("got error with GetZones %v", err)
	}

	var expected []Snapshot
	err = json.Unmarshal(responseData, &expected)
	if err != nil {
		t.Errorf("got error with Unmarshal %v %v", err, snapshots)
	}

	if !reflect.DeepEqual(snapshots, expected) {
		t.Errorf("client.GetSnapshots got %v expected %v\n", snapshots, expected)
	}

}

func TestGetSnapshotsEmpty(t *testing.T) {
	responseData := []byte(`
[]
`)

	mockZoneID := "12345678"
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/zones/12345678/snapshots", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(responseData))
	})

	snapshots, err := client.GetSnapshots(mockZoneID)
	if err != nil {
		t.Errorf("got error with GetZones %v", err)
	}

	var expected []Snapshot
	err = json.Unmarshal(responseData, &expected)
	if err != nil {
		t.Errorf("got error with Unmarshal %v %v", err, snapshots)
	}

	if !reflect.DeepEqual(snapshots, expected) {
		t.Errorf("client.GetSnapshots got %v expected %v\n", snapshots, expected)
	}

}

func TestGetSnapshotDetail(t *testing.T) {
	responseData := []byte(`
{
	"zone_uuid": "12345678",
	"uuid": "f7d9bfb2-a322-11ea-b635-00163e890b7b",
	"zone_data": [
	  {
		"rrset_type": "A",
		"rrset_ttl": 1800,
		"rrset_name": "@",
		"rrset_values": [
		  "45.79.62.185"
		]
	  },
	  {
		"rrset_type": "MX",
		"rrset_ttl": 10800,
		"rrset_name": "@",
		"rrset_values": [
		  "10 spool.mail.gandi.net.",
		  "50 fb.mail.gandi.net."
		]
	  },
	  {
		"rrset_type": "TXT",
		"rrset_ttl": 10800,
		"rrset_name": "@",
		"rrset_values": [
		  "\"v=spf1 include:_mailcust.gandi.net ?all\""
		]
	  },
	  {
		"rrset_type": "CNAME",
		"rrset_ttl": 10800,
		"rrset_name": "blog",
		"rrset_values": [
		  "blogs.vip.gandi.net."
		]
	  },
	  {
		"rrset_type": "A",
		"rrset_ttl": 1800,
		"rrset_name": "raneto",
		"rrset_values": [
		  "45.79.62.185"
		]
	  },
	  {
		"rrset_type": "CNAME",
		"rrset_ttl": 10800,
		"rrset_name": "webmail",
		"rrset_values": [
		  "webmail.gandi.net."
		]
	  },
	  {
		"rrset_type": "CNAME",
		"rrset_ttl": 10800,
		"rrset_name": "www",
		"rrset_values": [
		  "webredir.vip.gandi.net."
		]
	  },
	  {
		"rrset_type": "AAAA",
		"rrset_ttl": 1800,
		"rrset_name": "x23",
		"rrset_values": [
		  "2a02:8308:23b:ae00:211:32ff:fe15:d148"
		]
	  }
	],
	"date_created": "2020-05-31T09:41:57Z",
	"automatic": false,
	"name": "2020-05-31 #1"
  }
`)

	mockZoneID := "12345678"
	mockUUID := "f7d9bfb2-a322-11ea-b635-00163e890b7b"

	client, mux, teardown := setup()
	defer teardown()

	endpoint := fmt.Sprintf("/zones/12345678/snapshots/%s", mockUUID)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(responseData))
	})

	snapshotDetail, err := client.GetSnapshotDetails(mockZoneID, mockUUID)
	if err != nil {
		t.Errorf("got error with GetZones %v", err)
	}

	var expected SnapshotContent
	err = json.Unmarshal(responseData, &expected)
	if err != nil {
		t.Errorf("got error with Unmarshal %v %v", err, snapshotDetail)
	}

	if !reflect.DeepEqual(*snapshotDetail, expected) {
		t.Errorf("client.GetSnapshotDetail got %v expected %v\n", snapshotDetail, expected)
	}
}

func TestCreateSnapshot(t *testing.T) {
	responseData := []byte(`
	{
		"message": "Zone Snapshot Created",
		"uuid": "ef07a2ee-acca-11ea-bdd5-00163ec14bc8"
	}
	`)
	mockZoneID := "12345678"
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/zones/12345678/snapshots", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(responseData))
	})

	createResponse, err := client.CreateSnapshot(mockZoneID)
	if err != nil {
		t.Errorf("got error with GetZones %v", err)
	}

	var expected SnapshotCreate
	err = json.Unmarshal(responseData, &expected)
	if err != nil {
		t.Errorf("got error with Unmarshal %v %v", err, createResponse)
	}

	if !reflect.DeepEqual(*createResponse, expected) {
		t.Errorf("client.CreateSnapshot got %v expected %v\n", createResponse, expected)
	}
}
