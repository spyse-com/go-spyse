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
	AutonomousSystemDetailsEndpoint = "as/"
	AutonomousSystemSearchEndpoint  = "as/search"
)

// ASService handles Autonomous Systems for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems
type ASService struct {
	client *Client
}

type AS struct {
	Number       int            `json:"num,omitempty"`
	Organization string         `json:"org,omitempty"`
	IPv4Prefixes []IPV4Prefixes `json:"ipv4_prefixes,omitempty"`
	IPv6Prefixes []IPV6Prefixes `json:"ipv6_prefixes,omitempty"`
}

type IPV4Prefixes struct {
	CIDR         string `json:"cidr,omitempty"`
	ISP          string `json:"isp,omitempty"`
	DomainsCount int64  `json:"domains_count,omitempty"`
}

type IPV6Prefixes struct {
	CIDR *string `json:"cidr,omitempty"`
	ISP  *string `json:"isp,omitempty"`
}

// Details returns a full representation of the Autonomous System for the given AS number.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as
func (s *ASService) Details(ctx context.Context, asn int) (*AS, error) {
	// TODO: refactor
	refURI := fmt.Sprintf(AutonomousSystemDetailsEndpoint+"%s", strconv.Itoa(asn))
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

// Search returns a list of Autonomous Systems that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#as_search
func (s *ASService) Search(ctx context.Context, filters []map[string]Filter, limit, offset int) ([]*AS, error) {
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: filters,
			Limit:        limit,
			Offset:       offset,
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
