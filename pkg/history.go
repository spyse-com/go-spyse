package spyse

import (
	"context"
	"fmt"
	"net/http"
)

const (
	dnsHistoryEndpoint   = "/history/dns/%s/%s?limit=%d&offset=%d"
	whoisHistoryEndpoint = "/history/domain-whois/%s?limit=%d&offset=%d"
)

// HistoryService handles History for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/history
type HistoryService struct {
	client *Client
}

type DNSSimpleHistorical struct {
	Value     string `json:"value,omitempty"`
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

type WhoisHistory struct {
	Admin      DomainWHOIS          `json:"admin,omitempty"`
	Registrant DomainWHOIS          `json:"registrant,omitempty"`
	Registrar  DomainWHOISRegistrar `json:"registrar,omitempty"`
	Tech       DomainWHOIS          `json:"tech,omitempty"`
	UpdatedAt  string               `json:"updated_at,omitempty"`
	CreatedAt  string               `json:"created_at,omitempty"`
}

// DNS returns the DNS history of the given domain.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/history#dns_history
func (s *HistoryService) DNS(
	ctx context.Context,
	domain, dnsType string,
	limit, offset int,
) ([]DNSSimpleHistorical, error) {
	path := fmt.Sprintf(dnsHistoryEndpoint, dnsType, domain, limit, offset)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, DNSSimpleHistorical{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []DNSSimpleHistorical

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(DNSSimpleHistorical))
		}

		return items, nil
	}

	return nil, nil
}

// Whois returns the WHOIS history of the given domain.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/history#domain_whois_history
func (s *HistoryService) Whois(ctx context.Context, domain string, limit, offset int) ([]WhoisHistory, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(whoisHistoryEndpoint, domain, limit, offset), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, WhoisHistory{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []WhoisHistory

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(WhoisHistory))
		}

		return items, nil
	}

	return nil, nil
}
