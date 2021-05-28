package spyse

import (
	"context"
	"io"
	"net/http"
)

const (
	MaxSearchSize = 100
	MaxTotalItems = 10000
)

type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type PaginatedRequest struct {
	// The limit of rows to receive, value must be an integer in range from 1 to 100
	// required: false
	Size int `json:"limit"`
	// The offset of rows iterator,value must be an integer in range from 0 to 9999
	From int `json:"offset"`
}

type SearchRequest struct {
	SearchParams []map[string]Filter `json:"search_params"`
	PaginatedRequest
}

type ScrollSearchRequest struct {
	SearchParams []map[string]Filter `json:"search_params"`
	SearchID     string              `json:"search_id"`
}

type DomainBulkSearchRequest struct {
	DomainList []string `json:"domain_list" schema:"domain_list"`
	PaginatedRequest
}

type IPBulkSearchRequest struct {
	IPList []string `json:"ip_list" schema:"ip_list"`
	PaginatedRequest
}

func newRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}
