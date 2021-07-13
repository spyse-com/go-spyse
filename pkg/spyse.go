package spyse

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	authorizationHeaderName = "Authorization"
	authorizationType       = "Bearer"
	contentTypeHeaderName   = "Content-Type"
	defaultContentType      = "application/json"
)

// client defines an interface for an http.Client implementation so that alternative
// http Clients can be passed in for making requests
type client interface {
	Do(request *http.Request) (response *http.Response, err error)
}

type HTTPClient struct {
	// The Spyse API's uses accessToken-based authentication, which means that developers must pass their API accessToken.
	accessToken string

	// HTTP httpClient used to communicate with the API.
	client
	// Base URL for API requests.
	baseURL *url.URL
}

// NewClient returns a new Spyse API httpClient.
// If a nil httpClient is provided, http.DefaultClient will be used.
// To use API methods you must provide your API accessToken.
// See https://spyse-dev.readme.io/reference/quick-start
func NewClient(accessToken string, httpClient client) (*HTTPClient, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	parsedBaseURL, err := url.Parse("https://api.spyse.com/v4/data/")
	if err != nil {
		return nil, err
	}
	return &HTTPClient{
		client:      httpClient,
		baseURL:     parsedBaseURL,
		accessToken: accessToken,
	}, nil
}

func newRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// SetBaseURL set base URL for API endpoints.
func (c *HTTPClient) SetBaseURL(baseURL string) error {
	// ensure the baseURL contains a trailing slash so that all paths are preserved in later calls
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	c.baseURL = parsedBaseURL
	return nil
}

// NewRequest creates an API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *HTTPClient) NewRequest(ctx context.Context, method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// Relative URLs should be specified without a preceding slash since baseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := c.baseURL.ResolveReference(rel)

	req, err := newRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set(contentTypeHeaderName, defaultContentType)
	req.Header.Set(authorizationHeaderName, authorizationType+" "+c.accessToken)

	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to result,
// or returned an error if an API error has occurred.
func (c *HTTPClient) Do(req *http.Request, result interface{}) (*Response, error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, getErrorFromResponse(httpResp)
	}

	if result != nil {
		var body []byte
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}

		response := newResponse()
		if err = response.decodeFromJSON(body, result); err != nil {
			return nil, err
		}

		return response, nil

	} // todo: if result == nil ?
	return nil, err
}
