package spyse

// PaginatedRequest struct for pagination params
type PaginatedRequest struct {
	// The limit of rows to receive, value must be an integer in range from 1 to 100
	// required: false
	Size int `json:"limit"`
	// The offset of rows iterator,value must be an integer in range from 0 to 9999
	From int `json:"offset"`
}
