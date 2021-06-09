package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	emailDetailsEndpoint      = "email/"
	emailSearchEndpoint       = "email/search"
	emailScrollSearchEndpoint = "email/scroll/search"
	emailSearchCountEndpoint  = "email/search/count"
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
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(emailDetailsEndpoint+"%s", email), nil)
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

// Search returns a list of Emails that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/emails#email_search
func (s *EmailService) Search(
	ctx context.Context, params []map[string]SearchOption, limit, offset int) ([]Email, error,
) {
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

	req, err := s.client.NewRequest(ctx, http.MethodPost, emailSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, Email{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []Email

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(Email))
		}

		return items, nil
	}

	return nil, nil
}

// SearchCount returns a count of emails that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/emails#email_search_count
func (s *EmailService) SearchCount(ctx context.Context, params []map[string]SearchOption) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: params})
	if err != nil {
		return 0, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, emailSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}

type EmailScrollResponse struct {
	SearchID   string  `json:"search_id"`
	TotalItems int64   `json:"total_items"`
	Offset     int     `json:"offset"`
	Items      []Email `json:"items"`
}

// ScrollSearch returns a list of Emails that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/emails#email_scroll_search
func (s *EmailService) ScrollSearch(
	ctx context.Context,
	params []map[string]SearchOption,
	searchID string,
) (*EmailScrollResponse, error) {
	scrollRequest := ScrollSearchRequest{SearchParams: params}
	if searchID != "" {
		scrollRequest.SearchID = searchID
	}
	body, err := json.Marshal(scrollRequest)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, emailScrollSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, Email{})
	if err != nil {
		return nil, NewSpyseError(err)
	}
	response := &EmailScrollResponse{
		SearchID:   *resp.Data.SearchID,
		TotalItems: *resp.Data.TotalCount,
		Offset:     *resp.Data.Offset,
	}
	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			response.Items = append(response.Items, i.(Email))
		}
	}
	return response, err
}
