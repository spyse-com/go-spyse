package spyse

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type TotalCountResponseData struct {
	// The precise number of search results that matched the search query.
	TotalCount int64 `json:"total_count"`
}

type Response struct {
	Data  *Data  `json:"data"`
	Error *Error `json:"error"`
}

type Data struct {
	Items []interface{} `json:"items"`
	// The total number of records stored in the database
	TotalCount *int64 `json:"total_items,omitempty"`
	// The offset of rows iterator
	Offset *int `json:"offset,omitempty"`
	// Received Rows Limit
	Limit *int `json:"limit,omitempty"`
	// Received search_id in scroll requests
	SearchID *string `json:"search_id,omitempty"`
}

func (res *Response) decodeFromJSON(source []byte, result interface{}) error {
	var err error
	var items []interface{}

	if err = json.Unmarshal(source, res); err != nil {
		return err
	}

	for _, i := range res.Data.Items {
		config := &mapstructure.DecoderConfig{
			TagName: "json",
		}

		r := result
		config.Result = &r
		decoder, _ := mapstructure.NewDecoder(config)
		err = decoder.Decode(i)

		if err != nil {
			return err
		}

		items = append(items, r)
	}

	res.Data.Items = items

	return err
}

func newResponse() *Response {
	resp := new(Response)
	return resp
}
