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
	// ASParamIP gives a chance to search AS by the IP address.
	// Operators: eq.
	ASParamIP = "ip"

	// ASParamASOrg gives a chance to search by by the associated organization name.
	// Operators: eq, contains, starts.
	ASParamASOrg = "as_org"

	// ASParamASN gives a chance to search by its number.
	// Operators: eq.
	ASParamASN = "asn"

	// ASParamDomain gives a chance to search AS by the domain name.
	// Operators: eq.
	ASParamDomain = "domain"
)

// Options for Certificate search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/ssltls-certificates#certificate_search
const (
	// CertificateParamIssuedForDomain gives a chance to search by the domain the certificate is issued for.
	// Operators: eq.
	CertificateParamIssuedForDomain = "issued_for_domain"

	// CertificateParamIssuedForIP gives a chance to search by the IP address the certificate is issued for.
	// Operators: eq.
	CertificateParamIssuedForIP = "issued_for_ip"

	// CertificateParamIssuerCountry gives a chance to search by the country ISO code the issuer organization is
	// registered in. Operators: eq,  exists, not_exists.
	CertificateParamIssuerCountry = "issuer_country"

	// CertificateParamIssuerOrg gives a chance to search by the issuer organization name.
	// Operators: eq, contains, exists, not_exists.
	CertificateParamIssuerOrg = "issuer_org"

	// CertificateParamIssuerCommonName gives a chance to search by the issuer common name.
	// Example: “DigiCert SHA2 Secure Server CA”. Operators: eq, contains, exists, not_exists.
	CertificateParamIssuerCommonName = "issuer_common_name"

	// CertificateParamIssuerEmail gives a chance to search by the issuer email.
	// Operators: eq, contains, exists, not_exists.
	CertificateParamIssuerEmail = "issuer_email"

	// CertificateParamSubjectCountry gives a chance to search by the domain the country ISO code the subject
	// organization is registered in. Operators: eq,  exists, not_exists.
	CertificateParamSubjectCountry = "subject_country"

	// CertificateParamSubjectOrg gives a chance to search by the domain the subject organization name.
	// Operators: eq, contains, exists, not_exists.
	CertificateParamSubjectOrg = "subject_org"

	// CertificateParamSubjectCommonName gives a chance to search by the domain the subject common name.
	// Example: ssl902285.cloudflaressl.com. Operators: eq, contains, exists, not_exists.
	CertificateParamSubjectCommonName = "subject_common_name"

	// CertificateParamSubjectEmail gives a chance to search by the domain the subject email.
	// Operators: eq, contains, exists, not_exists.
	CertificateParamSubjectEmail = "subject_email"

	// CertificateParamFingerprintMD5 gives a chance to search by the MD5 fingerprint. Enter a hexadecimal
	// string without spaces in lowercase. Example:  500d1a99a670675a2e0b147b31ea17ce. Operators: eq.
	CertificateParamFingerprintMD5 = "fingerprint_md5"

	// CertificateParamFingerprintSHA1 gives a chance to search by the domain the SHA1 fingerprint. Enter a
	// hexadecimal string without spaces in lowercase. Example: f817a45f860581ff9911ce6405d3e194e90a0f3f.
	// Operators: eq.
	CertificateParamFingerprintSHA1 = "fingerprint_sha1"

	// CertificateParamFingerprintSHA256 gives a chance to search by the domain the SHA256 fingerprint. Enter a
	// hexadecimal string without spaces in lowercase.
	// Example: f239e183459c727c7b2694937f04e5925ffaf969a10cc31ce848fdd0ccf6e40f. Operators: eq.
	CertificateParamFingerprintSHA256 = "fingerprint_sha256"

	// CertificateParamValidityEnd gives a chance to search by the domain the certificate validity end date.
	// Format: YYYY-MM-DD. Operators: eq, gte, lte.
	CertificateParamValidityEnd = "validity_end"

	// CertificateParamValidityStart gives a chance to search by the domain the certificate validity start date.
	// Format: YYYY-MM-DD. Operators: eq, gte, lte.
	CertificateParamValidityStart = "validity_start"

	// CertificateParamIsTrusted gives a chance to search certificates that have been verified and validated by Spyse.
	// Operators: eq.
	CertificateParamIsTrusted = "is_trusted"
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

// Options for Email search params:
//
// ALl search parameters see at https://spyse-dev.readme.io/reference/emailss#email_search
const (
	// EmailParamEmail gives a chance to search by the email address.
	// Operators: eq, not_eq, contains, not_contains, starts, ends.
	EmailParamEmail = "email"
)
