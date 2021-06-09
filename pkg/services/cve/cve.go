package cve

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spyse-com/go-spyse/pkg"
	"github.com/spyse-com/go-spyse/pkg/services/as"
	"net/http"
	"time"
)

const (
	cveDetailsEndpoint      = "cve/"
	cveSearchEndpoint       = "cve/search"
	cveScrollSearchEndpoint = "cve/scroll/search"
	cveSearchCountEndpoint  = "cve/search/count"
)

// CVEService handles CVEs for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/cves
type CVEService struct {
	Client *spyse.Client
}

// CVE represents cve record with vulnerability information
type CVE struct {
	ID               string       `json:"id,omitempty"`
	References       References   `json:"references,omitempty"`
	Description      Descriptions `json:"description,omitempty"`
	Impact           CVEImpact    `json:"impact,omitempty"`
	Conditions       []Conditions `json:"conditions,omitempty"`
	ProblemType      ProblemType  `json:"problemtype,omitempty"`
	PublishedDate    string       `json:"publishedDate,omitempty"`
	LastModifiedDate string       `json:"lastModifiedDate,omitempty"`
}

type References struct {
	ReferencesData []ReferencesData `json:"reference_data,omitempty"`
}

type ReferencesData struct {
	URL       string   `json:"url,omitempty"`
	Name      string   `json:"name,omitempty"`
	RefSource string   `json:"refsource,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type Descriptions struct {
	DescriptionData []DescriptionData `json:"description_data,omitempty"`
}

type DescriptionData struct {
	Language string `json:"lang,omitempty"`
	Value    string `json:"value,omitempty"`
}

type CVEBaseMetricCVSSV2 struct {
	Version               string  `json:"version,omitempty"`
	VectorString          string  `json:"vectorString,omitempty"`
	AccessVector          string  `json:"accessVector,omitempty"`
	AccessComplexity      string  `json:"accessComplexity,omitempty"`
	Authentication        string  `json:"authentication,omitempty"`
	ConfidentialityImpact string  `json:"confidentialityImpact,omitempty"`
	IntegrityImpact       string  `json:"integrityImpact,omitempty"`
	AvailabilityImpact    string  `json:"availabilityImpact,omitempty"`
	BaseScore             float32 `json:"baseScore,omitempty"`
}

type CVEBaseMetricV2 struct {
	CVSSV2                  CVEBaseMetricCVSSV2 `json:"cvssV2,omitempty"`
	Severity                string              `json:"severity,omitempty"`
	ExploitabilityScore     float32             `json:"exploitabilityScore,omitempty"`
	ImpactScore             float32             `json:"impactScore,omitempty"`
	AcInSufInfo             *bool               `json:"acInsufInfo,omitempty"`
	ObtainAllPrivilege      *bool               `json:"obtainAllPrivilege,omitempty"`
	ObtainUserPrivilege     *bool               `json:"obtainUserPrivilege,omitempty"`
	ObtainOtherPrivilege    *bool               `json:"obtainOtherPrivilege,omitempty"`
	UserInteractionRequired *bool               `json:"userInteractionRequired,omitempty"`
}

type CVEBaseMetricCVSSV3 struct {
	AttackComplexity      string  `json:"attackComplexity,omitempty"`
	AttackVector          string  `json:"attackVector,omitempty"`
	AvailabilityImpact    string  `json:"availabilityImpact,omitempty"`
	BaseScore             float32 `json:"baseScore,omitempty"`
	BaseSeverity          string  `json:"baseSeverity,omitempty"`
	ConfidentialityImpact string  `json:"confidentialityImpact,omitempty"`
	IntegrityImpact       string  `json:"integrityImpact,omitempty"`
	PrivilegesRequired    string  `json:"privilegesRequired,omitempty"`
	Scope                 string  `json:"scope,omitempty"`
	UserInteraction       string  `json:"userInteraction,omitempty"`
	VectorString          string  `json:"vectorString,omitempty"`
	Version               string  `json:"version,omitempty"`
}

type CVEBaseMetricV3 struct {
	CVSSV3              CVEBaseMetricCVSSV3 `json:"cvssV3,omitempty"`
	ExploitabilityScore float32             `json:"exploitabilityScore,omitempty"`
	ImpactScore         float32             `json:"impactScore,omitempty"`
}

type CVEImpact struct {
	BaseMetricV2 CVEBaseMetricV2 `json:"baseMetricV2,omitempty"`
	BaseMetricV3 CVEBaseMetricV3 `json:"baseMetricV3,omitempty"`
}

type Conditions struct {
	Application                    string `json:"application,omitempty"`
	CPEPrefix                      string `json:"cpe_prefix,omitempty"`
	Hardware                       string `json:"hardware,omitempty"`
	OperationSystem                string `json:"operation_system,omitempty"`
	VersionEndExcluding            string `json:"version_end_excluding,omitempty"`
	VersionEndExcludingRepresent   int64  `json:"version_end_excluding_representation,omitempty"`
	VersionEndIncluding            string `json:"version_end_including,omitempty"`
	VersionEndIncludingRepresent   int64  `json:"version_end_including_representation,omitempty"`
	VersionStartExcluding          string `json:"version_start_excluding,omitempty"`
	VersionStartExcludingRepresent int64  `json:"version_start_excluding_representation,omitempty"`
	VersionStartIncluding          string `json:"version_start_including,omitempty"`
	VersionStartIncludingRepresent int64  `json:"version_start_including_representation,omitempty"`
}

type ProblemType struct {
	ProblemTypeData []ProblemTypeData `json:"problemtype_data,omitempty"`
}

type ProblemTypeData struct {
	Description []DescriptionData `json:"description,omitempty"`
}

type CVETime struct {
	time.Time
}

// Details returns a full representation of the CVE for the given CVE ID.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/cves#cve_details
func (s *CVEService) Details(ctx context.Context, cve string) (*CVE, error) {
	req, err := s.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(cveDetailsEndpoint+"%s", cve), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, &CVE{})
	if err != nil {
		return nil, spyse.NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*CVE), nil
	}

	return nil, nil
}

// Search returns a list of CVEs that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/cves#cve_search
func (s *CVEService) Search(
	ctx context.Context,
	params []map[string]spyse.SearchOption,
	limit, offset int,
) ([]CVE, error) {
	body, err := json.Marshal(
		spyse.SearchRequest{
			SearchParams: params,
			PaginatedRequest: spyse.PaginatedRequest{
				Size: limit,
				From: offset,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, cveSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, CVE{})
	if err != nil {
		return nil, spyse.NewSpyseError(err)
	}

	var items []CVE

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(CVE))
		}
		return items, nil
	}
	return nil, nil
}

// SearchCount returns a count of CVEs that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/cves#cve_search_count
func (s *CVEService) SearchCount(ctx context.Context, params []map[string]spyse.SearchOption) (int64, error) {
	body, err := json.Marshal(spyse.SearchRequest{SearchParams: params})
	if err != nil {
		return 0, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, cveSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.Client.Do(req, &as.TotalCountResponseData{})
	if err != nil {
		return 0, spyse.NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}

type CVEScrollResponse struct {
	SearchID   string `json:"search_id"`
	TotalItems int64  `json:"total_items"`
	Offset     int    `json:"offset"`
	Items      []CVE  `json:"items"`
}

// ScrollSearch returns a list of CVEs that match the specified search params.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/cves#cve_scroll_search
func (s *CVEService) ScrollSearch(
	ctx context.Context,
	params []map[string]spyse.SearchOption,
	searchID string,
) (*CVEScrollResponse, error) {
	scrollRequest := spyse.ScrollSearchRequest{SearchParams: params}
	if searchID != "" {
		scrollRequest.SearchID = searchID
	}
	body, err := json.Marshal(scrollRequest)
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPost, cveScrollSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, CVE{})
	if err != nil {
		return nil, spyse.NewSpyseError(err)
	}
	response := &CVEScrollResponse{
		SearchID:   *resp.Data.SearchID,
		TotalItems: *resp.Data.TotalCount,
		Offset:     *resp.Data.Offset,
	}
	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			response.Items = append(response.Items, i.(CVE))
		}
	}
	return response, err
}
