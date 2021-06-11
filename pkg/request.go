package spyse

type SearchOption struct {
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
	SearchParams []map[string]SearchOption `json:"search_params"`
	PaginatedRequest
}

type ScrollSearchRequest struct {
	SearchParams []map[string]SearchOption `json:"search_params"`
	SearchID     string                    `json:"search_id"`
}

type DomainBulkSearchRequest struct {
	DomainList []string `json:"domain_list" schema:"domain_list"`
	PaginatedRequest
}

type IPBulkSearchRequest struct {
	IPList []string `json:"ip_list" schema:"ip_list"`
	PaginatedRequest
}

type QueryBuilder struct {
	Query []map[string]SearchOption
}

func (q *QueryBuilder) AppendParam(lst ...QueryParam) {
	for _, v := range lst {
		q.Query = append(q.Query, map[string]SearchOption{
			v.Name: {
				Operator: v.Operator,
				Value:    v.Value,
			},
		})
	}
}

func (q *QueryBuilder) AppendGroupParams(lst ...QueryParam) {
	groupParams := make(map[string]SearchOption, len(lst))
	for _, v := range lst {
		groupParams[v.Name] = SearchOption{
			Value:    v.Value,
			Operator: v.Operator,
		}
	}
	q.Query = append(q.Query, groupParams)
}

func (q *QueryBuilder) Clear() {
	q.Query = []map[string]SearchOption{}
}

type QueryParam struct {
	Name     string
	Operator string
	Value    string
}
