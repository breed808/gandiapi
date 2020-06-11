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
