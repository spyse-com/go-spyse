package spyse

func (s *CertificateService) Params() CertificateParams { // nolint:funlen
	return CertificateParams{
		IssuedForDomain: CertificateParamIssuedForDomain{
			Name: "issued_for_domain",
			Operator: CertificateIssuedForDomainOperators{
				Equal: OperatorEqual,
			},
		},
		IssuedForIP: CertificateParamIssuedForIP{
			Name: "issued_for_ip",
			Operator: CertificateIssuedForIPOperators{
				Equal: OperatorEqual,
			},
		},
		IssuerCountry: CertificateParamIssuerCountry{
			Name: "issuer_country",
			Operator: CertificateIssuerCountryOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		IssuerOrg: CertificateParamIssuerOrg{
			Name: "issuer_org",
			Operator: CertificateIssuerOrgOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		IssuerCommonName: CertificateParamIssuerCommonName{
			Name: "issuer_common_name",
			Operator: CertificateIssuerCommonNameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		IssuerEmail: CertificateParamIssuerEmail{
			Name: "issuer_email",
			Operator: CertificateIssuerEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		SubjectCountry: CertificateParamSubjectCountry{
			Name: "subject_country",
			Operator: CertificateSubjectCountryOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		SubjectOrg: CertificateParamSubjectOrg{
			Name: "subject_org",
			Operator: CertificateSubjectOrgOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		SubjectCommonName: CertificateParamSubjectCommonName{
			Name: "subject_common_name",
			Operator: CertificateSubjectCommonNameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		SubjectEmail: CertificateParamSubjectEmail{
			Name: "subject_email",
			Operator: CertificateSubjectEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		FingerprintMD5: CertificateParamFingerprintMD5{
			Name: "fingerprint_md5",
			Operator: CertificateFingerprintMD5Operators{
				Equal: OperatorEqual,
			},
		},
		FingerprintSHA1: CertificateParamFingerprintSHA1{
			Name: "fingerprint_sha1",
			Operator: CertificateFingerprintSHA1Operators{
				Equal: OperatorEqual,
			},
		},
		FingerprintSHA256: CertificateParamFingerprintSHA256{
			Name: "fingerprint_sha256",
			Operator: CertificateFingerprintSHA256Operators{
				Equal: OperatorEqual,
			},
		},
		ValidityEnd: CertificateParamValidityEnd{
			Name: "validity_end",
			Operator: CertificateValidityEndOperators{
				Equal: OperatorEqual,
				Gte:   OperatorGreaterThanOrEqual,
				Lte:   OperatorLessThanOrEqual,
			},
		},
		ValidityStart: CertificateParamValidityStart{
			Name: "validity_start",
			Operator: CertificateValidityStartOperators{
				Equal: OperatorEqual,
				Gte:   OperatorGreaterThanOrEqual,
				Lte:   OperatorLessThanOrEqual,
			},
		},
		IsTrusted: CertificateParamIsTrusted{
			Name: "is_trusted",
			Operator: CertificateIsTrustedOperators{
				Equal: OperatorEqual,
			},
		},
	}
}

// CertificateParams for Certificate search:
//
// All search parameters see at https://spyse-dev.readme.io/reference/ssltls-certificates#certificate_search
type CertificateParams struct {
	// IssuedForDomain search by the domain the certificate is issued for.
	IssuedForDomain CertificateParamIssuedForDomain
	// IssuedForIP search by the IP address the certificate is issued for.
	IssuedForIP CertificateParamIssuedForIP
	// IssuerCountry search by the country ISO code the issuer organization is registered in.
	IssuerCountry CertificateParamIssuerCountry
	// IssuerOrg search by the issuer organization name.
	IssuerOrg CertificateParamIssuerOrg
	// IssuerCommonName search by the issuer common name. Example: “DigiCert SHA2 Secure Server CA”.
	IssuerCommonName CertificateParamIssuerCommonName
	// IssuerEmail search by the issuer email.
	IssuerEmail CertificateParamIssuerEmail
	// SubjectCountry search by the domain the country ISO code the subject  organization is registered in.
	SubjectCountry CertificateParamSubjectCountry
	// SubjectOrg search by the domain the subject organization name.
	SubjectOrg CertificateParamSubjectOrg
	// SubjectCommonName search by the domain the subject common name. Example: ssl902285.cloudflaressl.com.
	SubjectCommonName CertificateParamSubjectCommonName
	// SubjectEmail search by the domain the subject email.
	SubjectEmail CertificateParamSubjectEmail
	// FingerprintMD5 search by the MD5 fingerprint. Enter a hexadecimal string without spaces in
	// lowercase. Example:  500d1a99a670675a2e0b147b31ea17ce.
	FingerprintMD5 CertificateParamFingerprintMD5
	// FingerprintSHA1 search by the domain the SHA1 fingerprint. Enter a hexadecimal string without
	// spaces in lowercase. Example: f817a45f860581ff9911ce6405d3e194e90a0f3f.
	FingerprintSHA1 CertificateParamFingerprintSHA1
	// FingerprintSHA256 search by the domain the SHA256 fingerprint. Enter a hexadecimal string
	// without spaces in lowercase. Example: f239e183459c727c7b2694937f04e5925ffaf969a10cc31ce848fdd0ccf6e40f.
	FingerprintSHA256 CertificateParamFingerprintSHA256
	// ValidityEnd search by the domain the certificate validity end date. Format: YYYY-MM-DD.
	ValidityEnd CertificateParamValidityEnd
	// ValidityStart search by the domain the certificate validity start date. Format: YYYY-MM-DD.
	ValidityStart CertificateParamValidityStart
	// IsTrusted search certificates that have been verified and validated by Spyse.
	IsTrusted CertificateParamIsTrusted
}

type CertificateParamIssuedForDomain struct {
	Name     string
	Operator CertificateIssuedForDomainOperators
}

type CertificateIssuedForDomainOperators struct {
	Equal string
}

type CertificateParamIssuedForIP struct {
	Name     string
	Operator CertificateIssuedForIPOperators
}

type CertificateIssuedForIPOperators struct {
	Equal string
}

type CertificateParamIssuerCountry struct {
	Name     string
	Operator CertificateIssuerCountryOperators
}

type CertificateIssuerCountryOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type CertificateParamIssuerOrg struct {
	Name     string
	Operator CertificateIssuerOrgOperators
}

type CertificateIssuerOrgOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type CertificateParamIssuerCommonName struct {
	Name     string
	Operator CertificateIssuerCommonNameOperators
}

type CertificateIssuerCommonNameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type CertificateParamIssuerEmail struct {
	Name     string
	Operator CertificateIssuerEmailOperators
}

type CertificateIssuerEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type CertificateParamSubjectCountry struct {
	Name     string
	Operator CertificateSubjectCountryOperators
}

type CertificateSubjectCountryOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type CertificateParamSubjectOrg struct {
	Name     string
	Operator CertificateSubjectOrgOperators
}

type CertificateSubjectOrgOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type CertificateParamSubjectCommonName struct {
	Name     string
	Operator CertificateSubjectCommonNameOperators
}

type CertificateSubjectCommonNameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type CertificateParamSubjectEmail struct {
	Name     string
	Operator CertificateSubjectEmailOperators
}

type CertificateSubjectEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type CertificateParamFingerprintMD5 struct {
	Name     string
	Operator CertificateFingerprintMD5Operators
}

type CertificateFingerprintMD5Operators struct {
	Equal string
}

type CertificateParamFingerprintSHA1 struct {
	Name     string
	Operator CertificateFingerprintSHA1Operators
}

type CertificateFingerprintSHA1Operators struct {
	Equal string
}

type CertificateParamFingerprintSHA256 struct {
	Name     string
	Operator CertificateFingerprintSHA256Operators
}

type CertificateFingerprintSHA256Operators struct {
	Equal string
}

type CertificateParamValidityEnd struct {
	Name     string
	Operator CertificateValidityEndOperators
}

type CertificateValidityEndOperators struct {
	Equal string
	Gte   string
	Lte   string
}

type CertificateParamValidityStart struct {
	Name     string
	Operator CertificateValidityStartOperators
}

type CertificateValidityStartOperators struct {
	Equal string
	Gte   string
	Lte   string
}

type CertificateParamIsTrusted struct {
	Name     string
	Operator CertificateIsTrustedOperators
}

type CertificateIsTrustedOperators struct {
	Equal string
}
