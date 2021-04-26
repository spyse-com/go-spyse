package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DomainDetailsEndpoint     = "domain/"
	DomainSearchEndpoint      = "domain/search"
	DomainSearchCountEndpoint = "domain/search/count"
)

// DomainService handles Domains for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains
type DomainService struct {
	client *Client
}

// Domain represents the web site domain model
type Domain struct {
	AlexaInfo       AlexaInfo         `json:"alexa,omitempty"`
	CertSummary     DomainCertSummary `json:"cert_summary,omitempty"`
	DNSRecords      DNSRecords        `json:"dns_records,omitempty"`
	HostsEnrichment []GeoData         `json:"hosts_enrichment"`
	Extract         DomainExtractData `json:"http_extract,omitempty"`
	IsCNAME         *bool             `json:"is_CNAME,omitempty"`
	IsMX            *bool             `json:"is_MX,omitempty"`
	IsNS            *bool             `json:"is_NS,omitempty"`
	IsPTR           *bool             `json:"is_PTR,omitempty"`
	IsSubdomain     *bool             `json:"is_subdomain,omitempty"`
	Name            string            `json:"name,omitempty"`
	NameWithoutTLD  string            `json:"name_without_suffix,omitempty"`
	UpdatedAt       string            `json:"updated_at,omitempty"`
	WHOISParsed     WHOISParsedData   `json:"whois_parsed,omitempty"`
	WhoisUpdatedAt  string            `json:"whois_updated_at,omitempty"`
	ScreenshotURL   string            `json:"screenshot_url,omitempty"`
	Score           DomainScore       `json:"security_score,omitempty"`
	CVEList         []CVEInfo         `json:"cve_lists,omitempty"`
	Organization    *DomainOrg        `json:"organization,omitempty"`
	Technologies    []Technology      `json:"technology"`
	Trackers        Trackers          `json:"trackers"`
}

type Trackers struct {
	AdSenseID              string `json:"adsense_id,omitempty"`
	AppleItunesApp         string `json:"apple_itunes_app,omitempty"`
	GooglePlayApp          string `json:"google_play_app,omitempty"`
	GoogleAnalyticsKey     string `json:"google_analytics_key,omitempty"`
	GoogleSiteVerification string `json:"google_site_verification,omitempty"`
}

type Technology struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AlexaInfo struct {
	Rank      *int   `json:"rank,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type DomainCertSummary struct {
	FingerprintSHA256 string              `json:"fingerprint_sha256,omitempty"`
	IssuerDN          DomainCertIssuerDN  `json:"issuer_dn,omitempty"`
	SubjectDN         DomainCertSubjectDN `json:"subject_dn,omitempty"`
	TLSVersion        string              `json:"tls_version,omitempty"`
	ValidityEnd       string              `json:"validity_end,omitempty"`
}

type DomainCertIssuerDN struct {
	C            string `json:"C,omitempty"`
	CN           string `json:"CN,omitempty"`
	L            string `json:"L,omitempty"`
	O            string `json:"O,omitempty"`
	OU           string `json:"OU,omitempty"`
	ST           string `json:"ST,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	Raw          string `json:"raw,omitempty"`
}

type DomainCertSubjectDN struct {
	C                           string `json:"C,omitempty"`
	CN                          string `json:"CN,omitempty"`
	L                           string `json:"L,omitempty"`
	O                           string `json:"O,omitempty"`
	OU                          string `json:"OU,omitempty"`
	ST                          string `json:"ST,omitempty"`
	BusinessCategory            string `json:"businessCategory,omitempty"`
	EmailAddress                string `json:"emailAddress,omitempty"`
	JurisdictionCountry         string `json:"jurisdictionCountry,omitempty"`
	JurisdictionStateOrProvince string `json:"jurisdictionStateOrProvince,omitempty"`
	PostalCode                  string `json:"postalCode,omitempty"`
	Raw                         string `json:"raw,omitempty"`
	SerialNumber                string `json:"serialNumber,omitempty"`
	Street                      string `json:"street,omitempty"`
}

type DNSRecords struct {
	A         []string      `json:"A,omitempty"`
	AAAA      []string      `json:"AAAA,omitempty"`
	CAA       []string      `json:"CAA,omitempty"`
	CNAME     []string      `json:"CNAME,omitempty"`
	MX        []string      `json:"MX,omitempty"`
	NS        []string      `json:"NS,omitempty"`
	SOA       *DNSSOARecord `json:"SOA,omitempty"`
	TXT       []string      `json:"TXT,omitempty"`
	SPF       []*SPF        `json:"SPF,omitempty"`
	UpdatedAt string        `json:"updated_at"`
}

type IPInfo struct {
	Score           Score           `json:"score,omitempty"`
	SeverityDetails SeverityDetails `json:"severity_details,omitempty"`
	CVEList         []IPCVE         `json:"cve_list,omitempty"`
	OSH             *int64          `json:"osh,omitempty"`
	GeoData
}

type Score struct {
	Score *int `json:"score,omitempty"`
}

type SeverityDetails struct {
	High   int `json:"HIGH,omitempty"`
	Medium int `json:"MEDIUM,omitempty"`
	Low    int `json:"LOW,omitempty"`
}

type IPCVE struct {
	Ports          []int  `json:"ports,omitempty"`
	Tech           string `json:"technology,omitempty"`
	TechVersion    string `json:"technology_version,omitempty"`
	TechWebsite    string `json:"technology_website,omitempty"`
	TechFaviconURL string `json:"technology_favicon_url,omitempty"`
	CVEInfo
}

type CVEInfo struct {
	ID          string  `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	BaseScore   float32 `json:"base_score,omitempty"`
	Severity    string  `json:"severity,omitempty"`
	Vector      string  `json:"vector,omitempty"`
}

type GeoData struct {
	IP    string `json:"ip"`
	ASNum *int   `json:"as_num"`
	ASOrg string `json:"as_org"`
	ISP   string `json:"isp"`
	LocationData
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

type DNSSOARecord struct {
	Email   string `json:"email,omitempty"`
	Expire  int64  `json:"expire,omitempty"`
	MinTTL  int64  `json:"min_ttl,omitempty"`
	NS      string `json:"ns,omitempty"`
	Refresh int64  `json:"refresh,omitempty"`
	Retry   int64  `json:"retry,omitempty"`
	Serial  int64  `json:"serial,omitempty"`
}

type SPF struct {
	Mechanisms       []Mechanisms      `json:"mechanisms,omitempty"`
	Modifiers        []Modifiers       `json:"modifiers,omitempty"`
	Raw              string            `json:"raw,omitempty"`
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
	Version          string            `json:"version,omitempty"`
}

type Mechanisms struct {
	Name      string `json:"name,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Modifiers struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ValidationError struct {
	Description string `json:"description,omitempty"`
	Target      string `json:"target,omitempty"`
}

type DomainExtractData struct {
	Cookies             []ExtractCookie `json:"cookies,omitempty"`
	Description         string          `json:"description,omitempty"`
	Emails              []string        `json:"emails,omitempty"`
	ExternalRedirectURI URIParts        `json:"final_redirect_url,omitempty"`
	ExtractedAt         string          `json:"extracted_at,omitempty"`
	FaviconSHA256       string          `json:"favicon_sha256,omitempty"`
	HTTPHeaders         []HTTPHeaders   `json:"http_headers,omitempty"`
	HTTPStatusCode      *int64          `json:"http_status_code,omitempty"`
	Links               []Hyperlink     `json:"links,omitempty"`
	MetaTags            []MetaTag       `json:"meta_tags,omitempty"`
	RobotsTXT           string          `json:"robots_txt,omitempty"`
	Scripts             []string        `json:"scripts,omitempty"`
	Styles              []string        `json:"styles,omitempty"`
	Title               string          `json:"title,omitempty"`
}

type ExtractCookie struct {
	Domain   string `json:"domain,omitempty"`
	Expire   string `json:"expire,omitempty"`
	HTTPOnly *bool  `json:"http_only,omitempty"`
	Key      string `json:"key,omitempty"`
	MaxAge   int64  `json:"max_age,omitempty"`
	Path     string `json:"path,omitempty"`
	Security *bool  `json:"security,omitempty"`
	Value    string `json:"value,omitempty"`
}

type URIParts struct {
	FullURI string `json:"full_uri,omitempty"`
	Host    string `json:"host,omitempty"`
	Path    string `json:"path,omitempty"`
}

type HTTPHeaders struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Hyperlink struct {
	Anchor              string              `json:"anchor,omitempty"`
	HyperlinkAttributes HyperlinkAttributes `json:"attributes,omitempty"`
}

type HyperlinkAttributes struct {
	NoFollow *bool    `json:"nofollow,omitempty"`
	URI      URIParts `json:"uri,omitempty"`
}

type MetaTag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// DomainScore represents scoring information about target
type DomainScore struct {
	Score *int `json:"score,omitempty"`
}

type WHOISParsedData struct {
	Admin      DomainWHOIS          `json:"admin,omitempty"`
	Registrant DomainWHOIS          `json:"registrant,omitempty"`
	Registrar  DomainWHOISRegistrar `json:"registrar,omitempty"`
	Tech       DomainWHOIS          `json:"tech,omitempty"`
	UpdatedAt  string               `json:"updated_at"`
}

type DomainWHOIS struct {
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	Email        string `json:"email,omitempty"`
	Fax          string `json:"fax,omitempty"`
	FaxExt       string `json:"fax_ext,omitempty"`
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PhoneExt     string `json:"phone_ext,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	Province     string `json:"province,omitempty"`
	Street       string `json:"street,omitempty"`
	StreetExt    string `json:"street_ext,omitempty"`
}

type DomainWHOISRegistrar struct {
	CreatedDate    string `json:"created_date,omitempty"`
	DomainDNSSec   string `json:"domain_dnssec,omitempty"`
	DomainID       string `json:"domain_id,omitempty"`
	DomainName     string `json:"domain_name,omitempty"`
	DomainStatus   string `json:"domain_status,omitempty"`
	Emails         string `json:"emails,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	NameServers    string `json:"name_servers,omitempty"`
	ReferralURL    string `json:"referral_url,omitempty"`
	RegistrarID    string `json:"registrar_id,omitempty"`
	RegistrarName  string `json:"registrar_name,omitempty"`
	UpdatedDate    string `json:"updated_date,omitempty"`
	WHOISServer    string `json:"whois_server,omitempty"`
}
type CVELists struct {
	HTTPCVELists []CVEInfo `json:"http_cve_list,omitempty"`
}

type DomainOrg struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

type DNSRecordsData struct {
	Data DomainDNSRecords `json:"data"`
}

type DomainDNSRecords struct {
	DNSRecords
	DNSRecordsInfo
	Count int64 `json:"count,omitempty"`
}

type DNSRecordsInfo struct {
	MXInfo    map[string][]IPInfo `json:"mx_info,omitempty"`
	NSInfo    map[string][]IPInfo `json:"ns_info,omitempty"`
	CNAMEInfo map[string][]IPInfo `json:"cname_info,omitempty"`
}

// Details returns a full representation of the Domain for the given Domain name.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains#domain
func (s *DomainService) Details(ctx context.Context, domainName string) (*Domain, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(DomainDetailsEndpoint+"%s", domainName), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &Domain{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*Domain), nil
	}

	return nil, nil
}

// Search returns a list of Domains that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains#domain_search
func (s *DomainService) Search(ctx context.Context, filters []map[string]Filter, limit, offset int) ([]*Domain, error) {
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: filters,
			Limit:        limit,
			Offset:       offset,
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, DomainSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &Domain{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []*Domain

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(*Domain))
		}

		return items, nil
	}

	return nil, nil
}

// SearchCount returns a count of domains that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ssltls-certificates#domain_search_count
func (s *DomainService) SearchCount(ctx context.Context, filters []map[string]Filter) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: filters})
	if err != nil {
		return 0, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, DomainSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}
