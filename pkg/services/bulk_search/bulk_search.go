package bulk_search

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/spyse-com/go-spyse/pkg"
	"github.com/spyse-com/go-spyse/pkg/services/domain"
	"github.com/spyse-com/go-spyse/pkg/services/ip"
	"net/http"
)

const (
	bulkSearchDomainEndpoint = "bulk-search/domain"
	bulkSearchIPEndpoint     = "/bulk-search/ip"
)

// BulkSearchService handles Bulk search for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/bulk-search
type BulkSearchService struct {
	client *spyse.Client
}

// Domain lookup returns a full representation of the domains for the given domain names.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/bulk-search#bulk_search_domain
func (s *BulkSearchService) Domain(ctx context.Context, domainNames []string) ([]domain.Domain, error) {
	body, err := json.Marshal(spyse.DomainBulkSearchRequest{DomainList: domainNames})
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, bulkSearchDomainEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, domain.Domain{})
	if err != nil {
		return nil, spyse.NewSpyseError(err)
	}

	var items []domain.Domain

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(domain.Domain))
		}

		return items, nil
	}

	return nil, nil
}

// IP lookup returns a full representation of the ips for the given ip addresses.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/bulk-search#bulk_search_ip
func (s *BulkSearchService) IP(ctx context.Context, ipList []string) ([]ip.IP, error) {
	body, err := json.Marshal(spyse.IPBulkSearchRequest{IPList: ipList})
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, bulkSearchIPEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, ip.IP{})
	if err != nil {
		return nil, spyse.NewSpyseError(err)
	}

	var items []ip.IP

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(ip.IP))
		}

		return items, nil
	}

	return nil, nil
}
