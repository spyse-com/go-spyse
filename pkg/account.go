package spyse

import (
	"context"
	"net/http"
)

const (
	accountQuotaEndpoint = "account/quota"
)

// AccountService handles Account for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/account-1
type AccountService struct {
	Client *Client
}

func NewAccountService(c *Client) *AccountService {
	return &AccountService{
		Client: c,
	}
}

type Account struct {
	// The starting date of the quota period.
	StartAt string `json:"start_at"`
	// The ending date of the quota period.
	EndAt string `json:"end_at"`
	// The number of remaining API requests during the quota period.
	APIRequestsRemaining int64 `json:"api_requests_remaining"`
	// The number of maximum API requests you can make during the quota period.
	APIRequestsLimit int64 `json:"api_requests_limit"`
	// The number of remaining downloads during the quota period.
	DownloadsLimitRemaining int64 `json:"downloads_limit_remaining"`
	// The number of maximum downloads you can make during the quota period.
	DownloadsLimit int64 `json:"downloads_limit"`
}

// Quota returns details about your account quotas.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/account-1#account
func (s *AccountService) Quota(ctx context.Context) (*Account, error) {
	req, err := s.Client.NewRequest(ctx, http.MethodGet, accountQuotaEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, &Account{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*Account), nil
	}

	return nil, nil
}
