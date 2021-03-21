package spyse

import (
	"context"
	"io"
	"net/http"
)

// PaginatedRequest struct for pagination params
type PaginatedRequest struct {
	// The limit of rows to receive, value must be an integer in range from 1 to 100
	// required: false
	Size int `json:"limit"`
	// The offset of rows iterator,value must be an integer in range from 0 to 9999
	From int `json:"offset"`
}

type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type SearchRequest struct {
	SearchParams []map[string]Filter `json:"search_params"`
	PaginatedRequest
}

func newRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}
