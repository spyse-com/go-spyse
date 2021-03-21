package spyse

type PaginatedResponse struct {
	// The total number of records stored in the database
	TotalCount *int64 `json:"total_count,omitempty"`
	// Maximum allowed number of records for viewing
	MaxViewCount *int `json:"max_view_count,omitempty"`
	// The offset of rows iterator
	Offset *int `json:"offset,omitempty"`
	// Received Rows Limit
	Limit *int `json:"limit,omitempty"`
}
