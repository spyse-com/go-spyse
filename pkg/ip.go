package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ipDetailsEndpoint      = "ip/"
	ipSearchEndpoint       = "ip/search"
	ipScrollSearchEndpoint = "ip/scroll/search"
	ipSearchCountEndpoint  = "ip/search/count"
)

// IPService handles IPs for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips
type IPService struct {
	Client *Client
}

func NewIPService(c *Client) *IPService {
	return &IPService{
		Client: c,
	}
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

type IPCVE struct {
	ID        string  `json:"id,omitempty"`
	BaseScore float32 `json:"base_score_cvss2,omitempty"`
	Ports     []int   `json:"ports,omitempty"`
	Tech      string  `json:"technology,omitempty"`
}

// LocationData - geo information
type LocationData struct {
	CityName       string   `json:"city_name,omitempty"`
	Country        string   `json:"country,omitempty"`
	CountryISOCode string   `json:"country_iso_code,omitempty"`
	Location       GeoPoint `json:"location,omitempty"`
}

// GeoPoint is a geographic position described via latitude and longitude.
type GeoPoint struct {
	Lat float64 `json:"lat,omitempty"`
	Lon float64 `json:"lon,omitempty"`
}

type Score struct {
	Score *int `json:"score,omitempty"`
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
	Value     string `json:"value,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	GeoInfo   []Info `json:"geo_info,omitempty"`
}

type Info struct {
	Score           Score           `json:"score,omitempty"`
	SeverityDetails SeverityDetails `json:"severity_details,omitempty"`
	CVEList         []IPCVE         `json:"cve_list,omitempty"`
	OSH             *int64          `json:"osh,omitempty"`
	GeoData
}

type GeoData struct {
	IP    string `json:"ip"`
	ASNum *int   `json:"as_num"`
	ASOrg string `json:"as_org"`
	ISP   string `json:"isp"`
	LocationData
}

type SeverityDetails struct {
	High   int `json:"HIGH,omitempty"`
	Medium int `json:"MEDIUM,omitempty"`
	Low    int `json:"LOW,omitempty"`
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

type Trackers struct {
	AdSenseID              string `json:"adsense_id,omitempty"`
	AppleItunesApp         string `json:"apple_itunes_app,omitempty"`
	GooglePlayApp          string `json:"google_play_app,omitempty"`
	GoogleAnalyticsKey     string `json:"google_analytics_key,omitempty"`
	GoogleSiteVerification string `json:"google_site_verification,omitempty"`
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

// Details returns a full representation of the IP for the given IP address.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips#ip_details
func (s *IPService) Details(ctx context.Context, ip string) (*IP, error) {
	req, err := s.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(ipDetailsEndpoint+"%s", ip), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, &IP{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*IP), nil
	}

	return nil, nil
}

// Search returns a paginated list of IPs that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips#ip_search
func (s *IPService) Search(
	ctx context.Context,
	params []map[string]SearchOption,
	limit, offset int,
) ([]IP, error) {
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: params,
			PaginatedRequest: PaginatedRequest{
				Size: limit,
				From: offset,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, ipSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, IP{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []IP

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(IP))
		}
		return items, nil
	}
	return nil, nil
}

// SearchCount returns a count of IPs that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips#ip_search_count
func (s *IPService) SearchCount(ctx context.Context, params []map[string]SearchOption) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: params})
	if err != nil {
		return 0, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, ipSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.Client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}

type IPScrollResponse struct {
	SearchID   string `json:"search_id"`
	TotalItems int64  `json:"total_items"`
	Offset     int    `json:"offset"`
	Items      []IP   `json:"items"`
}

// ScrollSearch returns a list of IPs that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ips#ip_scroll_search
func (s *IPService) ScrollSearch(
	ctx context.Context,
	params []map[string]SearchOption,
	searchID string,
) (*IPScrollResponse, error) {
	scrollRequest := ScrollSearchRequest{SearchParams: params}
	if searchID != "" {
		scrollRequest.SearchID = searchID
	}
	body, err := json.Marshal(scrollRequest)
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, ipScrollSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, IP{})
	if err != nil {
		return nil, NewSpyseError(err)
	}
	response := &IPScrollResponse{
		SearchID:   *resp.Data.SearchID,
		TotalItems: *resp.Data.TotalCount,
		Offset:     *resp.Data.Offset,
	}
	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			response.Items = append(response.Items, i.(IP))
		}
	}
	return response, err
}
