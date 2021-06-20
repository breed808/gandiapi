package gandiapi

import (
	"net/http"
	"net/http/httptest"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	client = NewClient("fake_api_key", false, true)
	client.SetEndpoint(server.URL + "/")
	return client, mux, server.Close
}
