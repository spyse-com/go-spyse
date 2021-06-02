package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	AutonomousSystemDetailsEndpoint      = "as/"
	AutonomousSystemSearchEndpoint       = "as/search"
	AutonomousSystemScrollSearchEndpoint = "as/scroll/search"
	AutonomousSystemSearchCountEndpoint  = "as/search/count"
)

// ASService handles Autonomous Systems for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems
type ASService struct {
	client *Client
}

type TotalCountResponseData struct {
	// The precise number of search results that matched the search query.
	TotalCount int64 `json:"total_count"`
}

type AS struct {
	Number       int            `json:"asn,omitempty"`
	Organization string         `json:"as_org,omitempty"`
	IPv4Prefixes []IPV4Prefixes `json:"ipv4_prefixes,omitempty"`
	IPv6Prefixes []IPV6Prefixes `json:"ipv6_prefixes,omitempty"`
}

type IPV4Prefixes struct {
	CIDR string `json:"cidr,omitempty"`
	ISP  string `json:"isp,omitempty"`
}

type IPV6Prefixes struct {
	CIDR *string `json:"cidr,omitempty"`
	ISP  *string `json:"isp,omitempty"`
}

// Details returns a full representation of the Autonomous System for the given AS number.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_details
func (s *ASService) Details(ctx context.Context, asn string) (*AS, error) {
	refURI := fmt.Sprintf(AutonomousSystemDetailsEndpoint+"%s", asn)
	req, err := s.client.NewRequest(ctx, http.MethodGet, refURI, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &AS{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*AS), nil
	}

	return nil, nil
}

// Search returns a list of Autonomous Systems that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_search
func (s *ASService) Search(ctx context.Context, params []map[string]SearchParameter, limit, offset int) ([]AS, error) {
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: params,
			PaginatedRequest: PaginatedRequest{
				Size: limit,
				From: offset,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, AutonomousSystemSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, AS{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []AS

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(AS))
		}

		return items, nil
	}

	return nil, nil
}

// SearchCount returns a count of Autonomous Systems that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_search_count
func (s *ASService) SearchCount(ctx context.Context, params []map[string]SearchParameter) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: params})
	if err != nil {
		return 0, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, AutonomousSystemSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}

type ASScrollResponse struct {
	SearchID   string `json:"search_id"`
	TotalItems int64  `json:"total_items"`
	Offset     int    `json:"offset"`
	Items      []AS   `json:"items"`
}

// ScrollSearch returns a list of autonomous systems that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_scroll_search
func (s *ASService) ScrollSearch(
	ctx context.Context,
	params []map[string]SearchParameter,
	searchID string,
) (*ASScrollResponse, error) {
	scrollRequest := ScrollSearchRequest{SearchParams: params}
	if searchID != "" {
		scrollRequest.SearchID = searchID
	}
	body, err := json.Marshal(scrollRequest)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, AutonomousSystemScrollSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, AS{})
	if err != nil {
		return nil, NewSpyseError(err)
	}
	response := &ASScrollResponse{
		SearchID:   *resp.Data.SearchID,
		TotalItems: *resp.Data.TotalCount,
		Offset:     *resp.Data.Offset,
	}
	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			response.Items = append(response.Items, i.(AS))
		}
	}
	return response, err
}
