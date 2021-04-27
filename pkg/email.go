package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	EmailDetailsEndpoint     = "email/"
	EmailSearchEndpoint      = "email/search"
	EmailSearchCountEndpoint = "email/search/count"
)

// EmailService handles Emails for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/emails
type EmailService struct {
	client *Client
}

type Email struct {
	Email     string         `json:"email,omitempty"`
	Sources   []EmailSources `json:"sources,omitempty"`
	UpdatedAt string         `json:"updated_at,omitempty"`
}

type EmailSources struct {
	Target   string `json:"target,omitempty"`
	Type     string `json:"type,omitempty"`
	LastSeen string `json:"last_seen,omitempty"`
}

// Details returns a full representation of the emails for the given emails address.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/emails#email_details
func (s *EmailService) Details(ctx context.Context, email string) (*Email, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(EmailDetailsEndpoint+"%s", email), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &Email{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*Email), nil
	}

	return nil, nil
}

// Search returns a list of Domains that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/emails#email_search
func (s *EmailService) Search(ctx context.Context, filters []map[string]Filter, limit, offset int) ([]*Email, error) {
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: filters,
			PaginatedRequest: PaginatedRequest{
				Limit:  limit,
				Offset: offset,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, EmailSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &Email{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []*Email

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(*Email))
		}

		return items, nil
	}

	return nil, nil
}

// SearchCount returns a count of emails that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/emails#email_search_count
func (s *EmailService) SearchCount(ctx context.Context, filters []map[string]Filter) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: filters})
	if err != nil {
		return 0, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, EmailSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}
