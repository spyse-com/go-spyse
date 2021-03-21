package spyse

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

type Response struct {
	Data  *Data  `json:"data"`
	Error *Error `json:"error"`
}

type Data struct {
	Items []interface{} `json:"items"`
	// The total number of records stored in the database
	TotalCount *int64 `json:"total_count,omitempty"`
	// Maximum allowed number of records for viewing
	MaxViewCount *int `json:"max_view_count,omitempty"`
	// The offset of rows iterator
	Offset *int `json:"offset,omitempty"`
	// Received Rows Limit
	Limit *int `json:"limit,omitempty"`
}

func (r *Response) CheckForErrors() {

}

func (r *Response) decodeFromJSON(source []byte, result interface{}) error {
	var err error
	var items []interface{}

	if err = json.Unmarshal(source, r); err != nil {
		return err
	}

	for _, i := range r.Data.Items {
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

	r.Data.Items = items

	return err
}

func newResponse() *Response {
	//if nil == r {
	//	return nil,
	//		&IllegalArgumentError{ErrString: "*http.Response cannot be null"}
	//}
	resp := new(Response)
	//resp.httpResponse = r
	////populate first
	//resp.populatePagingInfo()
	//resp.populateRateLimits()
	////read the response
	//buf, err := ioutil.ReadAll(resp.httpResponse.Body)
	//if err != nil {
	//	return nil, err
	//}
	//resp.body = &buf
	////now check for errors
	//err = resp.checkForErrors()
	//if err != nil {
	//	return nil, err
	//}
	return resp
}
