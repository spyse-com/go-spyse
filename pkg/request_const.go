package spyse

const (
	OperatorEqual              = "eq"
	OperatorNotEqual           = "not_eq"
	OperatorContains           = "contains"
	OperatorStartsWith         = "starts"
	OperatorEndsWith           = "ends"
	OperatorExists             = "exists"
	OperatorNotExists          = "not_exists"
	OperatorGreaterThanOrEqual = "gte"
	OperatorLessThanOrEqual    = "lte"
)

// Options for Domain search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/domains#domain_search
const (
	// DomainParamName gives a chance to search by the domain name. Operators: eq, starts, ends, not_eq
	DomainParamName                                     = "name"
	DomainParamAdsenseID                                = "http_extract_tracker_adsense_id"
	DomainParamAlexaRank                                = "alexa_rank"
	DomainParamDNSA                                     = "dns_a"
	DomainParamDNSAAAA                                  = "dns_aaaa"
	DomainParamDNSNS                                    = "dns_ns"
	DomainParamDNSMX                                    = "dns_mx"
	DomainParamDNSTXT                                   = "dns_txt"
	DomainParamDNSCAA                                   = "dns_caa"
	DomainParamDNSCNAME                                 = "dns_cname"
	DomainParamDNSSPFRAW                                = "dns_spf_raw"
	DomainParamDNSSPFVersion                            = "dns_spf_version"
	DomainParamDNSSPFErrorsTarget                       = "dns_spf_errors_target"
	DomainParamDNSSPFModifiersName                      = "dns_spf_modifiers_name"
	DomainParamDNSSPFMechanismsName                     = "dns_spf_mechanisms_name"
	DomainParamDNSSPFModifiersValue                     = "dns_spf_modifiers_value"
	DomainParamDNSSPFMechanismsValue                    = "dns_spf_mechanisms_value"
	DomainParamDNSSPFErrorsDescription                  = "dns_spf_errors_description"
	DomainParamDNSSPFMechanismsQualifier                = "dns_spf_mechanisms_qualifier"
	DomainParamHTTPExtractTitle                         = "http_extract_title"
	DomainParamHTTPExtractEmail                         = "http_extract_email"
	DomainParamHTTPExtractRobots                        = "http_extract_robots"
	DomainParamHTTPExtractStyles                        = "http_extract_styles"
	DomainParamHTTPExtractScripts                       = "http_extract_scripts"
	DomainParamHTTPExtractMetaName                      = "http_extract_meta_name"
	DomainParamHTTPExtractLinkHost                      = "http_extract_link_host"
	DomainParamHTTPExtractMetaValue                     = "http_extract_meta_value"
	DomainParamHTTPExtractFaviconURI                    = "http_extract_favicon_uri"
	DomainParamHTTPExtractStatusCode                    = "http_extract_status_code"
	DomainParamHTTPExtractFaviconSHA256                 = "http_extract_favicon_sha256"
	DomainParamHTTPExtractHeadersName                   = "http_extract_headers_name"
	DomainParamHTTPExtractDescription                   = "http_extract_description"
	DomainParamHTTPExtractHeadersValue                  = "http_extract_headers_value"
	DomainParamHTTPExtractLinkURL                       = "http_extract_link_url"
	DomainParamHTTPExtractFinalRedirectURL              = "http_extract_final_redirect_url"
	DomainParamCVEID                                    = "cve_id"
	DomainParamCVESeverity                              = "cve_severity"
	DomainParamTechnologyName                           = "technology_name"
	DomainParamTechnologyVersion                        = "technology_version"
	DomainParamTechnologyCPE                            = "technology_cpe"
	DomainParamCertificateSHA256                        = "certificate_sha256"
	DomainParamCertificateVersion                       = "certificate_version"
	DomainParamWhoisRegistrarWhoisServer                = "whois_registrar_whois_server"
	DomainParamWhoisRegistrantOrg                       = "whois_registrant_org"
	DomainParamWhoisRegistrarName                       = "whois_registrar_name"
	DomainParamWhoisRegistrantName                      = "whois_registrant_name"
	DomainParamWhoisRegistrarEmail                      = "whois_registrar_email"
	DomainParamWhoisRegistrantPhone                     = "whois_registrant_phone"
	DomainParamWhoisRegistrantEmail                     = "whois_registrant_email"
	DomainParamWithoutSuffix                            = "without_suffix"
	DomainParamHTTPExtractTrackerGoogleAnalyticsKey     = "http_extract_tracker_google_analytics_key"
	DomainParamHTTPExtractTrackerGooglePlayApp          = "http_extract_tracker_google_play_app"
	DomainParamHTTPExtractTrackerAppleItunesApp         = "http_extract_tracker_apple_itunes_app"
	DomainParamHTTPExtractTrackerGoogleSiteVerification = "http_extract_tracker_google_site_verification"
	DomainParamCertificateIssuerOrg                     = "certificate_issuer_org"
	DomainParamCertificateIssuerCname                   = "certificate_issuer_cname"
	DomainParamCertificateIssuerEmail                   = "certificate_issuer_email"
	DomainParamCertificateIssuerOrganizationalUnit      = "certificate_issuer_organizational_unit"
	DomainParamCertificateIssuerCountry                 = "certificate_issuer_country"
	DomainParamCertificateIssuerState                   = "certificate_issuer_state"
	DomainParamCertificateIssuerLocality                = "certificate_issuer_locality"
	DomainParamCertificateSubjectOrg                    = "certificate_subject_org"
	DomainParamCertificateSubjectCNAME                  = "certificate_subject_cname"
	DomainParamCertificateSubjectEmail                  = "certificate_subject_email"
	DomainParamCertificateSubjectOrganizationalUnit     = "certificate_subject_organizational_unit"
	DomainParamCertificateSubjectCountry                = "certificate_subject_country"
	DomainParamCertificateSubjectState                  = "certificate_subject_state"
	DomainParamCertificateSubjectLocality               = "certificate_subject_locality"
	DomainParamCertificateSubjectSerialNumber           = "certificate_subject_serial_number"
	DomainParamCertificateValidityEnd                   = "certificate_validity_end"
	DomainParamGeoCountryISOCode                        = "geo_country_iso_code"
	DomainParamGeoCountry                               = "geo_country"
	DomainParamGeoCity                                  = "geo_city"
	DomainParamASNum                                    = "as_num"
	DomainParamOrganizationIndustry                     = "organization_industry"
	DomainParamOrganizationEmail                        = "organization_email"
	DomainParamOrganizationName                         = "organization_name"
	DomainParamOrganizationLegalName                    = "organization_legal_name"
	DomainParamISP                                      = "isp"
	DomainParamASOrg                                    = "as_org"
	DomainParamIsPTR                                    = "is_ptr"
	DomainParamIsMX                                     = "is_mx"
	DomainParamIsNS                                     = "is_ns"
	DomainParamIsSubDomainParam                         = "is_subdomain"
	DomainParamIsCNAME                                  = "is_cname"
)

// Options for IP search params:
//
// ALl search parameters see at https://spyse-dev.readme.io/reference/ips#ip_search
const (
	IPParamCIDR                                     = "cidr"
	IPParamISP                                      = "isp"
	IPParamPTR                                      = "ptr"
	IPParamASOrg                                    = "as_org"
	IPParamASNum                                    = "as_num"
	IPParamGeoCity                                  = "geo_city"
	IPParamGeoCountry                               = "geo_country"
	IPParamGeoCountryISOCode                        = "geo_country_iso_code"
	IPParamTechnologyCPE                            = "technology_cpe"
	IPParamPortTechnologyName                       = "port_technology_name"
	IPParamPortTechnologyVersion                    = "port_technology_version"
	IPParamOpenPort                                 = "open_port"
	IPParamPortCVEID                                = "port_cve_id"
	IPParamPortCVESeverity                          = "port_cve_severity"
	IPParamPortBanner                               = "port_banner"
	IPParamPortService                              = "port_service"
	IPParamHTTPExtractDescription                   = "http_extract_description"
	IPParamHTTPExtractTitle                         = "http_extract_title"
	IPParamHTTPExtractEmail                         = "http_extract_email"
	IPParamHTTPExtractRobots                        = "http_extract_robots"
	IPParamHTTPExtractStyles                        = "http_extract_styles"
	IPParamHTTPExtractScripts                       = "http_extract_scripts"
	IPParamHTTPExtractTrackerAdsenseID              = "http_extract_tracker_adsense_id"
	IPParamHTTPExtractMetaName                      = "http_extract_meta_name"
	IPParamHTTPExtractMetaValue                     = "http_extract_meta_value"
	IPParamHTTPExtractLinkHost                      = "http_extract_link_host"
	IPParamHTTPExtractStatusCode                    = "http_extract_status_code"
	IPParamHTTPExtractFaviconURI                    = "http_extract_favicon_uri"
	IPParamHTTPExtractFaviconSHA256                 = "http_extract_favicon_sha256"
	IPParamHTTPExtractHeadersName                   = "http_extract_headers_name"
	IPParamHTTPExtractHeadersValue                  = "http_extract_headers_value"
	IPParamHTTPExtractLinkURL                       = "http_extract_link_url"
	IPParamHTTPExtractFinalRedirectURL              = "http_extract_final_redirect_url"
	IPParamHTTPExtractTrackerGoogleAnalyticsKey     = "http_extract_tracker_google_analytics_key"
	IPParamHTTPExtractTrackerGooglePlayApp          = "http_extract_tracker_google_play_app"
	IPParamHTTPExtractTrackerAppleItunesApp         = "http_extract_tracker_apple_itunes_app"
	IPParamHTTPExtractTrackerGoogleSiteVerification = "http_extract_tracker_google_site_verification"
	IPParamAbusesReportsNum                         = "abuses_reports_num"
	IPParamAbusesReportsComment                     = "abuses_reports_comment"
	IPParamAbusesConfidenceScore                    = "abuses_confidence_score"
	IPParamAbusesCategoryName                       = "abuses_category_name"
	IPParamAbusesCategoryDescription                = "abuses_category_description"
	IPParamAbusesIsWhitelistStrong                  = "abuses_is_whitelist_strong"
	IPParamAbusesIsWhitelistWeak                    = "abuses_is_whitelist_weak"
	IPParamSecurityScore                            = "security_score"
)

// Options for Autonomous system search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/autonomous-systems#as_search
const (
	ASParamIP     = "ip"
	ASParamASOrg  = "as_org"
	ASParamASN    = "asn"
	ASParamDomain = "domain"
)

// Options for Certificate search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/ssltls-certificates#certificate_search
const (
	CertificateParamIssuedForDomain   = "issued_for_domain"
	CertificateParamIssuedForIP       = "issued_for_ip"
	CertificateParamIssuerCountry     = "issuer_country"
	CertificateParamIssuerOrg         = "issuer_org"
	CertificateParamIssuerCommonName  = "issuer_common_name"
	CertificateParamIssuerEmail       = "issuer_email"
	CertificateParamSubjectCountry    = "subject_country"
	CertificateParamSubjectOrg        = "subject_org"
	CertificateParamSubjectCommonName = "subject_common_name"
	CertificateParamSubjectEmail      = "subject_email"
	CertificateParamFingerprintMD5    = "fingerprint_md5"
	CertificateParamFingerprintSHA1   = "fingerprint_sha1"
	CertificateParamFingerprintSHA256 = "fingerprint_sha256"
	CertificateParamValidityEnd       = "validity_end"
	CertificateParamValidityStart     = "validity_start"
	CertificateParamIsTrusted         = "is_trusted"
)

// Options for CVE search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/cves#cve_search
const (
	// CVEParamID gives a chance to search by the CVE ID defined by the MITRE Corporation. Operators: eq.
	CVEParamID = "id"

	// CVEParamCPE gives a chance to search by the Common Platform Enumeration (CPE) name or prefix. Example:
	// cpe:2.3:o:canonical:ubuntu_linux:12.04. Operators: eq, starts.
	CVEParamCPE = "cpe"

	// CVEParamScoreCVSS2 gives a chance to search by the CVE score according to the Common Vulnerability Scoring System
	// Version 2 (CVSS2). Operators: gte, lte.
	CVEParamScoreCVSS2 = "score_cvss2"

	// CVEParamScoreCVSS3 gives a chance to search by the CVE score according to the Common Vulnerability Scoring System
	// Version 3 (CVSS3). Operators: gte, lte.
	CVEParamScoreCVSS3 = "score_cvss3"

	// CVEParamSeverityCVSS2 gives a chance to search by the CVE severity according to CVSSv2.
	// Supported options: high, medium, low. Operators: eq.
	CVEParamSeverityCVSS2 = "severity_cvss2"

	// CVEParamSeverityCVSS3 gives a chance to search by the CVE severity according to CVSSv3.
	// Supported options: high, medium, low, critical. Operators: eq.
	CVEParamSeverityCVSS3 = "severity_cvss3"

	// CVEParamPublishedAt gives a chance to search by the vulnerability publication date. Format: YYYY-MM-DD .
	// Operators: gte, lte.
	CVEParamPublishedAt = "published_at"

	// CVEParamModifiedAt gives a chance to search by the vulnerability modification date. Format: YYYY-MM-DD .
	// Operators: gte, lte.
	CVEParamModifiedAt = "modified_at"
)

// EmailParamEmail option for Email search params:
//
// ALl search parameters see at https://spyse-dev.readme.io/reference/emailss#email_search
const (
	EmailParamEmail = "email"
)
