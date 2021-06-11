package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	domainDetailsEndpoint      = "domain/"
	domainSearchEndpoint       = "domain/search"
	domainScrollSearchEndpoint = "domain/scroll/search"
	domainSearchCountEndpoint  = "domain/search/count"
)

// DomainService handles Domains for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains
type DomainService struct {
	Client *HTTPClient
}

func NewDomainService(c *HTTPClient) *DomainService {
	return &DomainService{
		Client: c,
	}
}

// Domain represents the web site domain model
type Domain struct {
	AlexaInfo       AlexaInfo       `json:"alexa,omitempty"`
	CertSummary     CertSummary     `json:"cert_summary,omitempty"`
	DNSRecords      DNSRecords      `json:"dns_records,omitempty"`
	HostsEnrichment []GeoData       `json:"hosts_enrichment"`
	Extract         ExtractData     `json:"http_extract,omitempty"`
	IsCNAME         *bool           `json:"is_CNAME,omitempty"`
	IsMX            *bool           `json:"is_MX,omitempty"`
	IsNS            *bool           `json:"is_NS,omitempty"`
	IsPTR           *bool           `json:"is_PTR,omitempty"`
	IsSubdomain     *bool           `json:"is_subdomain,omitempty"`
	Name            string          `json:"name,omitempty"`
	NameWithoutTLD  string          `json:"name_without_suffix,omitempty"`
	UpdatedAt       string          `json:"updated_at,omitempty"`
	WHOISParsed     WHOISParsedData `json:"whois_parsed,omitempty"`
	ScreenshotURL   string          `json:"screenshot_url,omitempty"`
	Score           Score           `json:"security_score,omitempty"`
	CVEList         []CVEInfo       `json:"cve_list,omitempty"`
	Technologies    []Technology    `json:"technologies,omitempty"`
	Trackers        Trackers        `json:"trackers,omitempty"`
	Organizations   []Organization  `json:"organizations,omitempty"`
}

type Organization struct {
	CrunchBase CrunchBase `json:"crunchbase"`
}

type CrunchBase struct {
	Name             string   `json:"name"`
	LegalName        string   `json:"legal_name"`
	HomepageURL      string   `json:"homepage_url"`
	Description      string   `json:"description"`
	ShortDescription string   `json:"short_description"`
	Address          string   `json:"address"`
	Categories       []string `json:"categories"`
	FoundedOn        string   `json:"founded_on"`
	ClosedOn         string   `json:"closed_on"`
	ContactEmail     string   `json:"contact_email"`
	ImageURL         string   `json:"image_url"`
	NumberEmployees  string   `json:"num_employees_enum"`
	OperatingStatus  string   `json:"operating_status"`
	Phone            string   `json:"phone_number"`
	RevenueRange     string   `json:"revenue_range"`
	Status           string   `json:"status"`
	CountryCode      string   `json:"country_code"`
	StatusCode       string   `json:"state_code"`
	Region           string   `json:"region"`
	City             string   `json:"city"`
	PostalCode       string   `json:"postal_code"`
	CrunchBaseURL    string   `json:"cb_url"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}

type AlexaInfo struct {
	Rank      *int   `json:"rank,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type CertSummary struct {
	FingerprintSHA256 string              `json:"fingerprint_sha256,omitempty"`
	Issuer            DomainCertIssuerDN  `json:"issuer,omitempty"`
	IssuerDN          string              `json:"issuer_dn,omitempty"`
	Subject           DomainCertSubjectDN `json:"subject,omitempty"`
	SubjectDN         string              `json:"subject_dn,omitempty"`
	TLSVersion        string              `json:"tls_version,omitempty"`
	ValidityEnd       string              `json:"validity_end,omitempty"`
}

type DomainCertIssuerDN struct {
	C            string `json:"country,omitempty"`             // Country
	CN           string `json:"common_name,omitempty"`         // Common Name
	L            string `json:"locality,omitempty"`            // Locality
	O            string `json:"organization,omitempty"`        // Organization
	OU           string `json:"organizational_unit,omitempty"` // Organization unit
	ST           string `json:"province,omitempty"`            // State
	EmailAddress string `json:"emailAddress,omitempty"`
}

type DomainCertSubjectDN struct {
	C                           string `json:"country,omitempty"`             // Country
	CN                          string `json:"common_name,omitempty"`         // Common Name
	L                           string `json:"locality,omitempty"`            // Locality
	O                           string `json:"organization,omitempty"`        // Organization
	OU                          string `json:"organizational_unit,omitempty"` // Organization unit
	ST                          string `json:"province,omitempty"`            // State
	BusinessCategory            string `json:"businessCategory,omitempty"`
	EmailAddress                string `json:"emailAddress,omitempty"`
	JurisdictionCountry         string `json:"jurisdictionCountry,omitempty"`
	JurisdictionStateOrProvince string `json:"jurisdictionStateOrProvince,omitempty"`
	PostalCode                  string `json:"postalCode,omitempty"`
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
	UpdatedAt string        `json:"updated_at,omitempty"`
}

type IPInfo struct {
	Score           Score           `json:"score,omitempty"`
	SeverityDetails SeverityDetails `json:"severity_details,omitempty"`
	CVEList         []IPCVE         `json:"cve_list,omitempty"`
	OSH             *int64          `json:"osh,omitempty"`
	GeoData
}

type CVEInfo struct {
	ID        string  `json:"id,omitempty"`
	BaseScore float32 `json:"base_score_cvss2,omitempty"`
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

type ExtractData struct {
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
	URI URIParts `json:"uri,omitempty"`
}

type MetaTag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type WHOISParsedData struct {
	Admin      WHOIS          `json:"admin,omitempty"`
	Registrant WHOIS          `json:"registrant,omitempty"`
	Registrar  WHOISRegistrar `json:"registrar,omitempty"`
	Tech       WHOIS          `json:"tech,omitempty"`
	UpdatedAt  string         `json:"updated_at,omitempty"`
}

type WHOIS struct {
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

type WHOISRegistrar struct {
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

// Details returns a full representation of the Domain for the given Domain name.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains#domain_details
func (s *DomainService) Details(ctx context.Context, domainName string) (*Domain, error) {
	req, err := s.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(domainDetailsEndpoint+"%s", domainName), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, &Domain{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*Domain), nil
	}

	return nil, nil
}

// Search returns a list of Domains that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains#domain_search
func (s *DomainService) Search(
	ctx context.Context,
	params []map[string]SearchOption,
	limit, offset int,
) ([]Domain, error) {
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

	req, err := s.Client.NewRequest(ctx, http.MethodPost, domainSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, Domain{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []Domain

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(Domain))
		}

		return items, nil
	}

	return nil, nil
}

// SearchCount returns a count of domains that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains#domain_search_count
func (s *DomainService) SearchCount(ctx context.Context, params []map[string]SearchOption) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: params})
	if err != nil {
		return 0, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, domainSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.Client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}

type DomainScrollResponse struct {
	SearchID   string   `json:"search_id"`
	TotalItems int64    `json:"total_items"`
	Offset     int      `json:"offset"`
	Items      []Domain `json:"items"`
}

// ScrollSearch returns a list of Domains that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/domains#domain_scroll_search
func (s *DomainService) ScrollSearch(
	ctx context.Context,
	params []map[string]SearchOption,
	searchID string,
) (*DomainScrollResponse, error) {
	scrollRequest := ScrollSearchRequest{SearchParams: params}
	if searchID != "" {
		scrollRequest.SearchID = searchID
	}
	body, err := json.Marshal(scrollRequest)
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, domainScrollSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, Domain{})
	if err != nil {
		return nil, NewSpyseError(err)
	}
	response := &DomainScrollResponse{
		SearchID:   *resp.Data.SearchID,
		TotalItems: *resp.Data.TotalCount,
		Offset:     *resp.Data.Offset,
	}
	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			response.Items = append(response.Items, i.(Domain))
		}
	}
	return response, err
}
