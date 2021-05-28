package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	IPDetailsEndpoint     = "ip/"
	IPSearchEndpoint      = "ip/search"
	IPSearchCountEndpoint = "ip/search/count"
)

// IPService handles IPs for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips
type IPService struct {
	client *Client
}

// IP represents IP record with geo info and DNS PTR record
type IP struct {
	CVEList      []IPCVE         `json:"cve_list,omitempty"`
	IPAddress    string          `json:"ip,omitempty"`
	GEOInfo      LocationData    `json:"geo_info,omitempty"`
	ISPInfo      ISPInfo         `json:"isp_info,omitempty"`
	PtrRecord    PtrRecord       `json:"ptr_record,omitempty"`
	Ports        []Port          `json:"ports,omitempty"`
	Score        Score           `json:"security_score,omitempty"`
	UpdatedAt    string          `json:"updated_at,omitempty"`
	CIDR         string          `json:"cidr,omitempty"`
	Technologies []Technology    `json:"technologies,omitempty"`
	Abuses       ShortAbusesInfo `json:"abuses,omitempty"`
}

type ShortAbusesInfo struct {
	ReportsNum int `json:"reports_num"`
	Score      int `json:"score"`
}

// ISPInfo - info about ISP, AS, organization
type ISPInfo struct {
	ASNum     *int   `json:"as_num,omitempty"`
	ASOrg     string `json:"as_org,omitempty"`
	ISP       string `json:"isp,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// PtrRecord of current IP
type PtrRecord struct {
	Value     string   `json:"value,omitempty"`
	UpdatedAt string   `json:"updated_at,omitempty"`
	GeoInfo   []IPInfo `json:"geo_info,omitempty"`
}

// Port represents port record with CPE, CVE info
type Port struct {
	Banner    string       `json:"banner,omitempty"`
	Extract   PortExtract  `json:"http_extract,omitempty"`
	Port      int          `json:"port,omitempty"`
	Tech      []Technology `json:"technology,omitempty"`
	Service   string       `json:"masscan_service_name,omitempty"`
	UpdatedAt string       `json:"updated_at,omitempty"`
	Trackers  Trackers     `json:"trackers,omitempty"`
}

type PortExtract struct {
	Cookies             []ExtractCookie `json:"cookies,omitempty"`
	Description         string          `json:"description,omitempty"`
	Emails              []string        `json:"emails,omitempty"`
	ExternalRedirectURI URIParts        `json:"final_redirect_url,omitempty"`
	ExtractedAt         string          `json:"extracted_at,omitempty"`
	FaviconSha256       string          `json:"favicon_sha256,omitempty"`
	HTTPHeaders         []HTTPHeaders   `json:"http_headers,omitempty"`
	HTTPStatusCode      *int            `json:"http_status_code,omitempty"`
	Links               []Hyperlink     `json:"links,omitempty"`
	MetaTags            []MetaTag       `json:"meta_tags,omitempty"`
	RobotsTxt           string          `json:"robots_txt,omitempty"`
	Scripts             []string        `json:"scripts,omitempty"`
	Styles              []string        `json:"styles,omitempty"`
	Title               string          `json:"title,omitempty"`
}

type Technology struct {
	Port    int    `json:"port,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type IPTechnologies struct {
	HTTP []Technology `json:"http,omitempty"`
}

// Details returns a full representation of the IP for the given IP address.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips#ip
func (s *IPService) Details(ctx context.Context, ip string) (*IP, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(IPDetailsEndpoint+"%s", ip), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &IP{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*IP), nil
	}

	return nil, nil
}

// Search returns a paginated list of IPs that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips#ip_search
func (s *IPService) Search(ctx context.Context, filters []map[string]Filter, limit, offset int) ([]*IP, error) {
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: filters,
			PaginatedRequest: PaginatedRequest{
				Size: limit,
				From: offset,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, IPSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &IP{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []*IP

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(*IP))
		}
		return items, nil
	}
	return nil, nil
}

// SearchCount returns a count of IPs that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/autonomous-systems#ip_search_count
func (s *IPService) SearchCount(ctx context.Context, filters []map[string]Filter) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: filters})
	if err != nil {
		return 0, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, IPSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}

// SearchAll returns a list of IPs that match the specified filters.
func (s *IPService) SearchAll(ctx context.Context, filters []map[string]Filter) (items []*IP, err error) {
	var from int

	for {
		body, err := json.Marshal(
			SearchRequest{
				SearchParams: filters,
				PaginatedRequest: PaginatedRequest{
					Size: MaxSearchSize,
					From: from,
				},
			},
		)
		if err != nil {
			return nil, err
		}

		req, err := s.client.NewRequest(ctx, http.MethodPost, IPSearchEndpoint, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}

		resp, err := s.client.Do(req, &IP{})
		if err != nil {
			return nil, NewSpyseError(err)
		}

		if len(resp.Data.Items) > 0 {
			for _, i := range resp.Data.Items {
				items = append(items, i.(*IP))
			}
			from += MaxSearchSize
			if from >= MaxTotalItems || len(resp.Data.Items) < MaxSearchSize {
				break
			}
			continue
		}
		break
	}
	return
}