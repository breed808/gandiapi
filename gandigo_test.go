package gandigo

import (
	"log"
	"net/http"
	"net/http/httptest"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	options := OptsClient{APIURL: server.URL}

	client, err := NewClient(&options)
	if err != nil {
		log.Fatalf("got error with NewClient %s", err)
	}
	return client, mux, server.Close
}
