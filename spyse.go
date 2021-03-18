package spyse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

const BaseURL = "https://api.spyse.com/v3/data/"

// httpClient defines an interface for an http.Client implementation so that alternative
// http Clients can be passed in for making requests
type httpClient interface {
	Do(request *http.Request) (response *http.Response, err error)
}

// A Client manages communication with the Spyse API.
type Client struct {
	// HTTP client used to communicate with the API.
	client httpClient

	// The Spyse API's uses token-based authentication, which means that developers must pass their API token.
	token string

	// Base URL for API requests.
	baseURL *url.URL
}

func (c *Client) AS() *ASService {
	return NewASService(c)
}

func (c *Client) ASSearch() *ASSearchService {
	return NewASSearchService(c)
}

// Response represents Spyse API response. It wraps http.Response returned from
// API and provides information about pagination.
type Response struct {
	*http.Response
}

// NewClient returns a new Spyse API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
// To use API methods you must provide your API token.
// See https://spyse-dev.readme.io/reference/quick-start
func NewClient(httpClient httpClient, apiToken string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	parsedBaseURL, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		baseURL: parsedBaseURL,
		token:   apiToken,
	}

	return c, nil
}

// NewRequest creates an API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body io.Reader) (*http.Request, error) {
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to result,
// or returned an error if an API error has occurred.
func (c *Client) Do(req *http.Request, result interface{}) error {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if httpResp.StatusCode != http.StatusOK {
		return getErrorFromResponse(httpResp)
	}

	if result != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer httpResp.Body.Close()
		if err = json.NewDecoder(httpResp.Body).Decode(result); err != nil {
			return errors.New("failed to decode response body")
		}
	}
	return err
}
