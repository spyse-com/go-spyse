package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// ASService handles Autonomous Systems for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems
type ASService struct {
	client *Client
}

// NewSearchService creates a new service for searching in Elasticsearch.
func NewASService(client *Client) *ASService {
	builder := &ASService{
		client: client,
	}
	return builder
}

// ASSearchService handles Autonomous Systems for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems
type ASSearchService struct {
	client *Client

	searchSource *SearchSource
}

// NewSearchService creates a new service for searching in ElASSearchticsearch.
func NewASSearchService(client *Client) *ASSearchService {
	builder := &ASSearchService{
		client:       client,
		searchSource: NewSearchSource(),
	}
	return builder
}

type SearchSource struct {
	// Search filters to retrieve data
	filters []map[string]Filter
	// The offset of rows iterator,value must be an integer in range from 0 to 9999
	from int
	// The limit of rows to receive, value must be an integer in range from 1 to 100
	size int
}

// NewSearchSource initializes a new SearchSource.
func NewSearchSource() *SearchSource {
	return &SearchSource{
		from:    0,
		size:    1,
		filters: []map[string]Filter{},
	}
}

// Filter sets the filter to search source.
func (s *SearchSource) Filter(filter ...Filters) *SearchSource {
	for _, f := range filter {
		s.filters = append(s.filters, f.Get())
	}

	return s
}

// From index to start the search from. Defaults to 0.
func (s *SearchSource) From(from int) *SearchSource {
	s.from = from
	return s
}

// Size is the number of search hits to return. Defaults to 1.
func (s *SearchSource) Size(size int) *SearchSource {
	s.size = size
	return s
}

type ASData struct {
	Data ASItems `json:"data"`
}

type ASItems struct {
	Items []AS `json:"items"`
	PaginatedResponse
}

type PaginatedResponse struct {
	// The total number of records stored in the database
	TotalCount *int64 `json:"total_count,omitempty"`
	// Maximum allowed number of records for viewing
	MaxViewCount *int `json:"max_view_count,omitempty"`
	// The offset of rows iterator
	Offset *int `json:"offset,omitempty"`
	// Received Rows Limit
	Limit *int `json:"limit,omitempty"`
}

type AS struct {
	ASN           *int        `json:"asn,omitempty"`
	ASOrg         *string     `json:"as_org,omitempty"`
	IPv4CIDRList  []IPV4Range `json:"ipv4_cidr,omitempty"`
	IPv6CIDRList  []IPV6Range `json:"ipv6_cidr,omitempty"`
	IPv4CIDRArray []string    `json:"ipv4_cidr_array,omitempty"`
	IPv6CIDRArray []string    `json:"ipv6_cidr_array,omitempty"`
}

type IPV4Range struct {
	CIDR         *string `json:"cidr,omitempty"`
	ISP          *string `json:"isp,omitempty"`
	DomainsCount *int64  `json:"domains_count,omitempty"`
}

type IPV6Range struct {
	CIDR *string `json:"cidr,omitempty"`
	ISP  *string `json:"isp,omitempty"`
}

// Details returns a full representation of the Autonomous System for the given AS number.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as
func (s *ASService) Details(ctx context.Context, asn int) (*ASData, error) {
	refURI := fmt.Sprintf("as?asn=%s", strconv.Itoa(asn))
	req, err := s.client.NewRequest(ctx, "GET", refURI, nil)
	if err != nil {
		return nil, err
	}

	as := new(ASData)
	if err = s.client.Do(req, as); err != nil {
		return nil, NewSpyseError(err)
	}

	return as, nil
}

type Filters interface {
	Get() map[string]Filter
	Set(filter map[string]Filter)
}

type ASSearchFilters struct {
	Filters map[string]Filter
}

func (p *ASSearchFilters) Get() map[string]Filter {
	return p.Filters
}

func (p *ASSearchFilters) Set(filter map[string]Filter) {
	p.Filters = filter
}

type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type SearchRequest struct {
	SearchParams []map[string]Filter `json:"search_params"`
	PaginatedRequest
}

// PaginatedRequest struct for pagination params
type PaginatedRequest struct {
	// The limit of rows to receive, value must be an integer in range from 1 to 100
	// required: false
	Size int `json:"limit"`
	// The offset of rows iterator,value must be an integer in range from 0 to 9999
	From int `json:"offset"`
}

// Do returns a list of Autonomous Systems that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_search
func (s *ASSearchService) Do(ctx context.Context) (*ASData, error) {
	refURI := "as/search"
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: s.searchSource.filters,
			PaginatedRequest: PaginatedRequest{
				Size: s.searchSource.size,
				From: s.searchSource.from,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest(ctx, "POST", refURI, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	as := new(ASData)
	if err = s.client.Do(req, as); err != nil {
		return nil, NewSpyseError(err)
	}
	return as, nil
}

// Filter sets the filter to search request.
func (s *ASSearchService) Filter(filter ...Filters) *ASSearchService {
	s.searchSource = s.searchSource.Filter(filter...)
	return s
}

// From index to start the search from. Defaults to 0.
func (s *ASSearchService) From(from int) *ASSearchService {
	s.searchSource = s.searchSource.From(from)
	return s
}

// Size is the number of search hits to return. Defaults to 1.
func (s *ASSearchService) Size(size int) *ASSearchService {
	s.searchSource = s.searchSource.Size(size)
	return s
}
