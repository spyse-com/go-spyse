package spyse

func (s *CVEService) Params() CVEParams {
	return CVEParams{
		ID: CVEParamID{
			Name: "id",
			Operator: CVEIDOperators{
				Equal: OperatorEqual,
			},
		},
		CPE: CVEParamCPE{
			Name: "cpe",
			Operator: CVECPEOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
			},
		},
		ScoreCVSS2: CVEParamScoreCVSS2{
			Name: "score_cvss2",
			Operator: CVEScoreCVSS2Operators{
				Gte: OperatorGreaterThanOrEqual,
				Lte: OperatorLessThanOrEqual,
			},
		},
		ScoreCVSS3: CVEParamScoreCVSS3{
			Name: "score_cvss3",
			Operator: CVEScoreCVSS3Operators{
				Gte: OperatorGreaterThanOrEqual,
				Lte: OperatorLessThanOrEqual,
			},
		},
		SeverityCVSS2: CVEParamSeverityCVSS2{
			Name: "severity_cvss2",
			Operator: CVESeverityCVSS2Operators{
				Equal: OperatorEqual,
			},
		},
		SeverityCVSS3: CVEParamSeverityCVSS3{
			Name: "severity_cvss3",
			Operator: CVESeverityCVSS3Operators{
				Equal: OperatorEqual,
			},
		},
		PublishedAt: CVEParamPublishedAt{
			Name: "published_at",
			Operator: CVEPublishedAtOperators{
				Equal: OperatorEqual,
			},
		},
		ModifiedAt: CVEParamModifiedAt{
			Name: "modified_at",
			Operator: CVEModifiedAtOperators{
				Equal: OperatorEqual,
			},
		},
	}
}

// CVEParams for CVE search:
//
// All search parameters see at https://spyse-dev.readme.io/reference/cves#cve_search
type CVEParams struct {
	// ID search by the CVE ID defined by the MITRE Corporation.
	ID CVEParamID
	// CPE search by the Common Platform Enumeration (CPE) name or prefix. Example:
	// cpe:2.3:o:canonical:ubuntu_linux:12.04.
	CPE CVEParamCPE
	// ScoreCVSS2 search by the CVE score according to the Common Vulnerability Scoring System Version 2 (CVSS2).
	ScoreCVSS2 CVEParamScoreCVSS2
	// ScoreCVSS3 search by the CVE score according to the Common Vulnerability Scoring System Version 3 (CVSS3).
	ScoreCVSS3 CVEParamScoreCVSS3
	// SeverityCVSS2 search by the CVE severity according to CVSSv2. Supported options: high, medium, low.
	SeverityCVSS2 CVEParamSeverityCVSS2
	// SeverityCVSS3 search by the CVE severity according to CVSSv3. Supported options: high, medium, low, critical.
	SeverityCVSS3 CVEParamSeverityCVSS3
	// PublishedAt search by the vulnerability publication date. Format: YYYY-MM-DD.
	PublishedAt CVEParamPublishedAt
	// ModifiedAt search by the vulnerability modification date. Format: YYYY-MM-DD.
	ModifiedAt CVEParamModifiedAt
}

type CVEParamID struct {
	Name     string
	Operator CVEIDOperators
}

type CVEIDOperators struct {
	Equal string
}

type CVEParamCPE struct {
	Name     string
	Operator CVECPEOperators
}

type CVECPEOperators struct {
	Equal      string
	StartsWith string
}

type CVEParamScoreCVSS2 struct {
	Name     string
	Operator CVEScoreCVSS2Operators
}

type CVEScoreCVSS2Operators struct {
	Gte string
	Lte string
}

type CVEParamScoreCVSS3 struct {
	Name     string
	Operator CVEScoreCVSS3Operators
}

type CVEScoreCVSS3Operators struct {
	Gte string
	Lte string
}

type CVEParamSeverityCVSS2 struct {
	Name     string
	Operator CVESeverityCVSS2Operators
}

type CVESeverityCVSS2Operators struct {
	Equal string
}

type CVEParamSeverityCVSS3 struct {
	Name     string
	Operator CVESeverityCVSS3Operators
}

type CVESeverityCVSS3Operators struct {
	Equal string
}

type CVEParamPublishedAt struct {
	Name     string
	Operator CVEPublishedAtOperators
}

type CVEPublishedAtOperators struct {
	Equal string
}

type CVEParamModifiedAt struct {
	Name     string
	Operator CVEModifiedAtOperators
}

type CVEModifiedAtOperators struct {
	Equal string
}
