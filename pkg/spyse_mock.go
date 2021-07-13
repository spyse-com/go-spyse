package spyse

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type mock struct {
	// TestMux is the HTTP request multiplexer used with the test server.
	TestMux *http.ServeMux

	Client *HTTPClient
	// TestServer is a test HTTP server used to provide mock API responses.
	TestServer *httptest.Server
}

func setup() *mock {
	// Test server
	testMux := http.NewServeMux()
	testServer := httptest.NewServer(testMux)

	apiBaseUrl := testServer.URL
	if !strings.HasSuffix(apiBaseUrl, "/") {
		apiBaseUrl += "/"
	}

	parsedBaseURL, _ := url.Parse(apiBaseUrl)

	c := &HTTPClient{
		client:  http.DefaultClient,
		baseURL: parsedBaseURL,
	}

	return &mock{
		TestMux:    testMux,
		Client:     c,
		TestServer: testServer,
	}
}

// teardown closes the test HTTP server.
func (s *mock) teardown() {
	s.TestServer.Close()
}

func (s *mock) testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func (s *mock) testRequestURL(t *testing.T, r *http.Request, want string) {
	if got := r.URL.String(); !strings.HasPrefix(got, want) {
		t.Errorf("Request URL: %v, want %v", got, want)
	}
}
