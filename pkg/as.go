package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	AutonomousSystemDetailsEndpoint = "as"
	AutonomousSystemSearchEndpoint  = "as/search"
	ASNumberQueryParam              = "asn"
)

// ASService handles Autonomous Systems for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems
type ASService struct {
	client *Client
}

type AS struct {
	ASN           int         `json:"asn,omitempty"`
	ASOrg         string      `json:"as_org,omitempty"`
	IPv4CIDRList  []IPV4Range `json:"ipv4_cidr,omitempty"`
	IPv6CIDRList  []IPV6Range `json:"ipv6_cidr,omitempty"`
	IPv4CIDRArray []string    `json:"ipv4_cidr_array,omitempty"`
	IPv6CIDRArray []string    `json:"ipv6_cidr_array,omitempty"`
}

type IPV4Range struct {
	CIDR         string `json:"cidr,omitempty"`
	ISP          string `json:"isp,omitempty"`
	DomainsCount int64  `json:"domains_count,omitempty"`
}

type IPV6Range struct {
	CIDR *string `json:"cidr,omitempty"`
	ISP  *string `json:"isp,omitempty"`
}

// Details returns a full representation of the Autonomous System for the given AS number.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as
func (s *ASService) Details(ctx context.Context, asn int) (*AS, error) {
	// TODO: refactor
	refURI := fmt.Sprintf(AutonomousSystemDetailsEndpoint+"?"+ASNumberQueryParam+"=%s", strconv.Itoa(asn))
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

// Do returns a list of Autonomous Systems that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_search
func (s *ASService) Search(ctx context.Context, filters []map[string]Filter, limit, offset int) ([]*AS, error) {
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
	req, err := s.client.NewRequest(ctx, http.MethodPost, AutonomousSystemSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &AS{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []*AS

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(*AS))
		}

		return items, nil
	}

	return nil, nil
}

// TODO: Add "SearchAll" method
