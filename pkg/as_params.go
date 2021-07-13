package spyse

// Params for Autonomous system search:
//
// All search parameters see at https://spyse-dev.readme.io/reference/autonomous-systems#as_search
func (s *ASService) Params() ASParams {
	return ASParams{
		IP: ASParamIP{
			Name: "ip",
			Operator: ASIPOperators{
				Equal: OperatorEqual,
			},
		},
		ASN: ASParamASN{
			Name: "asn",
			Operator: ASASNOperators{
				Equal: OperatorEqual,
			},
		},
		Domain: ASParamDomain{
			Name: "domain",
			Operator: ASDomainOperators{
				Equal: OperatorEqual,
			},
		},
		Organization: ASParamOrg{
			Name: "as_org",
			Operator: ASOrgOperators{
				Equal:      OperatorEqual,
				Contains:   OperatorContains,
				StartsWith: OperatorStartsWith,
			},
		},
	}
}

type ASParams struct {
	// ASParamIP search AS by the IP address.
	IP ASParamIP
	// ASParamASN search by its number.
	ASN ASParamASN
	// ASParamDomain search AS by the domain name.
	Domain ASParamDomain
	// ASParamOrg search by by the associated organization name.
	Organization ASParamOrg
}

type ASParamASN struct {
	Name     string
	Operator ASASNOperators
}

type ASASNOperators struct {
	Equal string
}

type ASParamIP struct {
	Name     string
	Operator ASIPOperators
}

type ASIPOperators struct {
	Equal string
}

type ASParamDomain struct {
	Name     string
	Operator ASDomainOperators
}

type ASDomainOperators struct {
	Equal string
}

type ASParamOrg struct {
	Name     string
	Operator ASOrgOperators
}

type ASOrgOperators struct {
	Equal      string
	Contains   string
	StartsWith string
}
