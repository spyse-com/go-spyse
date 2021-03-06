package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

const (
	bulkSearchDomainEndpoint = "bulk-search/domain"
	bulkSearchIPEndpoint     = "/bulk-search/ip"
)

// BulkService handles Bulk search for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/bulk-search
type BulkService struct {
	Client *Client
}

func NewBulkService(c *Client) *BulkService {
	return &BulkService{
		Client: c,
	}
}

// Domain lookup returns a full representation of the domains for the given domain names.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/bulk-search#bulk_search_domain
func (s *BulkService) Domain(ctx context.Context, domainNames []string) ([]Domain, error) {
	body, err := json.Marshal(DomainBulkSearchRequest{DomainList: domainNames})
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, bulkSearchDomainEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, Domain{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []Domain

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(Domain))
		}

		return items, nil
	}

	return nil, nil
}

// IP lookup returns a full representation of the ips for the given ip addresses.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/bulk-search#bulk_search_ip
func (s *BulkService) IP(ctx context.Context, ipList []string) ([]IP, error) {
	body, err := json.Marshal(IPBulkSearchRequest{IPList: ipList})
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, bulkSearchIPEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, IP{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []IP

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(IP))
		}

		return items, nil
	}

	return nil, nil
}
