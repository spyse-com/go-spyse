package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ASService handles Autonomous Systems for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems
type ASService struct {
	client *Client
}

type ASData struct {
	Data ASItems `json:"data"`
}

type ASItems struct {
	Items []AS `json:"items"`
	PaginatedResponse
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
	req, err := s.client.NewRequest(ctx, http.MethodGet, refURI, nil)
	if err != nil {
		return nil, err
	}

	as := new(ASData)
	if err = s.client.Do(req, as); err != nil {
		return nil, NewSpyseError(err)
	}

	return as, nil
}

type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type SearchRequest struct {
	SearchParams []map[string]Filter `json:"search_params"`
	PaginatedRequest
}

// Do returns a list of Autonomous Systems that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_search
func (s *ASService) Search(ctx context.Context, filters []map[string]Filter, limit, offset int) (*ASData, error) {
	refURI := "as/search"
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: filters,
			PaginatedRequest: PaginatedRequest{
				Size: limit,
				From: offset,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, refURI, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	as := new(ASData)
	if err = s.client.Do(req, as); err != nil {
		return nil, NewSpyseError(err)
	}
	return as, nil
}
