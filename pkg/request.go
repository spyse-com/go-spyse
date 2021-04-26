package spyse

import (
	"context"
	"io"
	"net/http"
)

type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type SearchRequest struct {
	SearchParams []map[string]Filter `json:"search_params"`
	// The limit of rows to receive, value must be an integer in range from 1 to 100
	// required: false
	Limit int `json:"limit"`
	// The offset of rows iterator,value must be an integer in range from 0 to 9999
	Offset int `json:"offset"`
}

func newRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}
