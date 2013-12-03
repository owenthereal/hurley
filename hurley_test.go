package hurley

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func tearDown() {
	server.Close()
}

type TestHandler struct {
}

func (h *TestHandler) PrepareRequest(req *http.Request) error {
	req.Header.Add("Content-Type", "application/json")
	return nil
}

func (h *TestHandler) PrepareResponse(resp *http.Response) error {
	resp.Status = "201"
	return nil
}

func TestClient_Get(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("HTTP method should be GET, but it's %s", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("HTTP header Content-Type should be application/json, but it's %s", r.Header.Get("Content-Type"))
		}

		fmt.Fprint(w, "hi")
	})

	c := New()
	h := &TestHandler{}
	c.Use(h)

	resp, err := c.Get(server.URL + "/foo")
	if err != nil {
		t.Fatalf("error should be nil but it's %s", err)
	}

	if resp.Status != "201" {
		t.Errorf("response status should be 201, but it's %s", resp.Status)
	}
}
