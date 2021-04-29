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
	DefaultBaseURL          = "https://api.spyse.com/v3/data/"
	AuthorizationHeaderName = "Authorization"
	AuthorizationType       = "Bearer"
	ContentTypeHeaderName   = "Content-Type"
	DefaultContentType      = "application/json"
)

// httpClient defines an interface for an http.Client implementation so that alternative
// http Clients can be passed in for making requests
type httpClient interface {
	Do(request *http.Request) (response *http.Response, err error)
}

// A Client manages communication with the Spyse API.
type Client struct {
	// HTTP httpClient used to communicate with the API.
	httpClient httpClient

	// The Spyse API's uses accessToken-based authentication, which means that developers must pass their API accessToken.
	accessToken string

	// Base URL for API requests.
	baseURL *url.URL

	Certificate *CertificateService
	AS          *ASService
	CVE         *CVEService
	Domain      *DomainService
	Email       *EmailService
	IP          *IPService
	BulkSearch  *BulkSearchService
	History     *HistoryService
}

// NewClient returns a new Spyse API httpClient.
// If a nil httpClient is provided, http.DefaultClient will be used.
// To use API methods you must provide your API accessToken.
// See https://spyse-dev.readme.io/reference/quick-start
func NewClient(baseURL, accessToken string, httpClient httpClient) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// ensure the baseURL contains a trailing slash so that all paths are preserved in later calls
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		httpClient:  httpClient,
		baseURL:     parsedBaseURL,
		accessToken: accessToken,
	}

	c.AS = &ASService{client: c}
	c.Certificate = &CertificateService{client: c}
	c.CVE = &CVEService{client: c}
	c.Domain = &DomainService{client: c}
	c.Email = &EmailService{client: c}
	c.IP = &IPService{client: c}
	c.BulkSearch = &BulkSearchService{client: c}
	c.History = &HistoryService{client: c}

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

	req.Header.Set(ContentTypeHeaderName, DefaultContentType)
	req.Header.Set(AuthorizationHeaderName, AuthorizationType+" "+c.accessToken)

	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to result,
// or returned an error if an API error has occurred.
func (c *Client) Do(req *http.Request, result interface{}) (*Response, error) {
	httpResp, err := c.httpClient.Do(req)
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

	}
	return nil, err
}
