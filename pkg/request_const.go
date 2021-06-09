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
	// DomainParamName search by the domain name. Operators: eq, starts, ends, not_eq
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
	// IPParamCIDR search IPs by CIDR. Operators: eq.
	IPParamCIDR = "cidr"

	// IPParamISP search IPs by the internet service provider name. Operators: eq, contains, exists, not_exists.
	IPParamISP = "isp"

	// IPParamPTR search IPs PTR record value. Operators: eq, contains, exists, not_exists.
	IPParamPTR = "ptr"

	// IPParamASOrg search IPs by the autonomous system organization name. Operators: eq, contains, exists, not_exists.
	IPParamASOrg = "as_org"

	// IPParamASNum search IPs by the autonomous system number. Operators: eq, exists, not_exists.
	IPParamASNum = "as_num"

	// IPParamGeoCity search IPs by the city they are located in. Operators: eq, exists, not_exists.
	IPParamGeoCity = "geo_city"

	// IPParamGeoCountry search IPs  by the country they are located in. Operators: eq, exists, not_exists.
	IPParamGeoCountry = "geo_country"

	// IPParamGeoCountryISOCode search IPs by the country ISO code. Operators: eq, exists, not_exists.
	IPParamGeoCountryISOCode = "geo_country_iso_code"

	// IPParamTechnologyCPE search IPs by technology CPE. Example: "cpe:2.3:a:apache:http_server".
	// Operators: eq, contains.
	IPParamTechnologyCPE = "technology_cpe"

	// IPParamPortTechnologyName find IPs hosting the specified technology. Examples: “Pure-FTPd”, “Dovecot pop3d”,
	// "Postfix smtpd". Operators: eq, contains.
	IPParamPortTechnologyName = "port_technology_name"

	// IPParamPortTechnologyVersion find IPs hosting the specified technology version. Examples: “2.4.18”
	// (Apache version), “7.4” (OpenSSH version), "1.3.5rc3" (ProFTPD version).
	// Operators: eq, gte, lte.
	IPParamPortTechnologyVersion = "port_technology_version"

	// IPParamOpenPort the open port number to search for. For ex. by setting the value "22" and operator "eq"
	// you will search for all IP hosts that have port 22 opened.
	// Operators: eq, exists, not_exists.
	IPParamOpenPort = "open_port"

	// IPParamPortCVEID find IPs affected by a specific CVE. Example: CVE-2019-12407.
	// Operators: eq.
	IPParamPortCVEID = "port_cve_id"

	// IPParamPortBanner search IPs by its port banner content. Operators: contains, exists, not_exists.
	IPParamPortBanner = "port_banner"

	// IPParamPortService search IPs by services detected on ports. Examples: "ssh”, “http”, “ftp” etc.
	// Operators: eq, exists, not_exists.
	IPParamPortService = "port_service"

	// IPParamHTTPExtractDescription search IPs that host websites (on any port) with the homepage containing
	// the specified description.
	// Operators: eq, starts, exists, not_exists.
	IPParamHTTPExtractDescription = "http_extract_description"

	// IPParamHTTPExtractTitle search IPs that host websites (on any port) with the homepage containing the specified
	// title. Operators: eq, starts, exists, not_exists.
	IPParamHTTPExtractTitle = "http_extract_title"

	// IPParamHTTPExtractEmail search IPs that host websites (on any port) with the homepage containing the specified
	// email address. Operators: eq, contains, exists, not_exists.
	IPParamHTTPExtractEmail = "http_extract_email"

	// IPParamHTTPExtractRobots search IPs that host websites (on any port) with the homepage containing the
	// specified robots.txt content. Operators: contains, exists, not_exists.
	IPParamHTTPExtractRobots = "http_extract_robots"

	// IPParamHTTPExtractStyles search IPs that host websites (on any port) with the homepage containing <link>
	// tags with specified "href" attribute substring. Operators: contains, exists, not_exists.
	IPParamHTTPExtractStyles = "http_extract_styles"

	// IPParamHTTPExtractScripts search IPs that host websites (on any port) with the homepage containing <script>
	// tags with specified "src" attribute substring. Operators: contains, exists, not_exists.
	IPParamHTTPExtractScripts = "http_extract_scripts"

	// IPParamHTTPExtractTrackerAdsenseID search IPs that host websites (on any port) with the homepage containing the
	// specified Google AdSense identifier. Operators: eq, exists, not_exists..
	IPParamHTTPExtractTrackerAdsenseID = "http_extract_tracker_adsense_id"

	// IPParamHTTPExtractMetaName search IPs that host websites (on any port) with the homepage containing the
	// specified HTML meta tag name. Operators: eq, contains.
	IPParamHTTPExtractMetaName = "http_extract_meta_name"

	// IPParamHTTPExtractMetaValue search IPs that host websites (on any port) with the homepage containing the
	// specified HTML meta tag value. Operators: eq, contains..
	IPParamHTTPExtractMetaValue = "http_extract_meta_value"

	// IPParamHTTPExtractLinkHost search IPs that host websites (on any port) with the homepage containing links
	// with the specified substrings in their URL host. Operators: contains, starts, ends.
	IPParamHTTPExtractLinkHost = "http_extract_link_host"

	// IPParamHTTPExtractStatusCode search IPs that host websites (on any port) returning the specified status code.
	// Operators: eq, gte, lte, exists, not_exists.
	IPParamHTTPExtractStatusCode = "http_extract_status_code"

	// IPParamHTTPExtractFaviconSHA256 search IPs that host websites (on any port) with the homepage containing
	// favicon with the specified SHA256 hash value. Operators: eq, exists, not_exists.
	IPParamHTTPExtractFaviconSHA256 = "http_extract_favicon_sha256"

	// IPParamHTTPExtractHeadersName search IPs that host websites (on any port) returning the specified HTTP
	// header name. Operators: eq, contains.
	IPParamHTTPExtractHeadersName = "http_extract_headers_name"

	// IPParamHTTPExtractHeadersValue search IPs that host websites (on any port) returning the specified HTTP
	// header value. Operators: eq, contains.
	IPParamHTTPExtractHeadersValue = "http_extract_headers_value"

	// IPParamHTTPExtractLinkURL search IPs that host websites (on any port) with the homepage containing links
	// with the specified substrings in their URL.
	// Operators: contains, starts, ends.
	IPParamHTTPExtractLinkURL = "http_extract_link_url"

	// IPParamHTTPExtractFinalRedirectURL find IPs that host websites redirecting to the specified URL.
	// Operators: eq, contains, exists, not_exists.
	IPParamHTTPExtractFinalRedirectURL = "http_extract_final_redirect_url"

	// IPParamHTTPExtractTrackerGoogleAnalyticsKey search IPs that host websites (on any port) with the homepage
	// containing the specified Google Analytics key. Operators: eq, exists, not_exists.
	IPParamHTTPExtractTrackerGoogleAnalyticsKey = "http_extract_tracker_google_analytics_key"

	// IPParamHTTPExtractTrackerGooglePlayApp search IPs that host websites (on any port) with the homepage containing
	// the specified Google Play app identifier. Operators: eq, contains, exists, not_exists.
	IPParamHTTPExtractTrackerGooglePlayApp = "http_extract_tracker_google_play_app"

	// IPParamHTTPExtractTrackerAppleItunesApp search IPs that host websites (on any port) with the homepage containing
	// the specified Apple iTunes app identifier. Operators: eq, contains, exists, not_exists.
	IPParamHTTPExtractTrackerAppleItunesApp = "http_extract_tracker_apple_itunes_app"

	// IPParamHTTPExtractTrackerGoogleSiteVerification Search IPs that host websites (on any port) with the homepage
	// containing the specified Google Site verification identifier.
	// Operators: eq, exists, not_exists.
	IPParamHTTPExtractTrackerGoogleSiteVerification = "http_extract_tracker_google_site_verification"

	// IPParamAbusesReportsNum find IPs by the number of abuses received. Operators: gte, lte.
	IPParamAbusesReportsNum = "abuses_reports_num"

	// IPParamAbusesReportsComment find IPs by the abuses comment. Operators: eq, contains, exists, not_exists.
	IPParamAbusesReportsComment = "abuses_reports_comment"

	// IPParamAbusesConfidenceScore find IPs by abuse confidence score. Operators: gte, lte.
	IPParamAbusesConfidenceScore = "abuses_confidence_score"

	// IPParamAbusesCategoryName find IPs by abuse category name. Operators: eq, contains, exists, not_exists.
	IPParamAbusesCategoryName = "abuses_category_name"

	// IPParamAbusesIsWhitelistStrong find IPs, which contains benign crawlers verified by reverse-forward DNS
	// resolution checks. Operators: eq.
	IPParamAbusesIsWhitelistStrong = "abuses_is_whitelist_strong"

	// IPParamAbusesIsWhitelistWeak find IPs, that are within our CIDR whitelist, which contain net blocks
	// from organizations we believe to be benign or often mistaken for malicious. Operators: eq.
	IPParamAbusesIsWhitelistWeak = "abuses_is_whitelist_weak"

	// IPParamSecurityScore search IPs by Spyse security score. Operators: gte, lte.
	IPParamSecurityScore = "security_score"
)

// Options for Autonomous system search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/autonomous-systems#as_search
const (
	// ASParamIP search AS by the IP address. Operators: eq.
	ASParamIP = "ip"

	// ASParamASOrg search by by the associated organization name. Operators: eq, contains, starts.
	ASParamASOrg = "as_org"

	// ASParamASN search by its number. Operators: eq.
	ASParamASN = "asn"

	// ASParamDomain search AS by the domain name. Operators: eq.
	ASParamDomain = "domain"
)

// Options for Certificate search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/ssltls-certificates#certificate_search
const (
	// CertificateParamIssuedForDomain search by the domain the certificate is issued for.
	// Operators: eq.
	CertificateParamIssuedForDomain = "issued_for_domain"

	// CertificateParamIssuedForIP search by the IP address the certificate is issued for.
	// Operators: eq.
	CertificateParamIssuedForIP = "issued_for_ip"

	// CertificateParamIssuerCountry search by the country ISO code the issuer organization is registered in.
	// Operators: eq,  exists, not_exists.
	CertificateParamIssuerCountry = "issuer_country"

	// CertificateParamIssuerOrg search by the issuer organization name.
	// Operators: eq, contains, exists, not_exists.
	CertificateParamIssuerOrg = "issuer_org"

	// CertificateParamIssuerCommonName search by the issuer common name. Example: “DigiCert SHA2 Secure Server CA”.
	// Operators: eq, contains, exists, not_exists.
	CertificateParamIssuerCommonName = "issuer_common_name"

	// CertificateParamIssuerEmail search by the issuer email. Operators: eq, contains, exists, not_exists.
	CertificateParamIssuerEmail = "issuer_email"

	// CertificateParamSubjectCountry search by the domain the country ISO code the subject  organization is registered
	// in. Operators: eq,  exists, not_exists.
	CertificateParamSubjectCountry = "subject_country"

	// CertificateParamSubjectOrg search by the domain the subject organization name.
	// Operators: eq, contains, exists, not_exists.
	CertificateParamSubjectOrg = "subject_org"

	// CertificateParamSubjectCommonName search by the domain the subject common name.
	// Example: ssl902285.cloudflaressl.com. Operators: eq, contains, exists, not_exists.
	CertificateParamSubjectCommonName = "subject_common_name"

	// CertificateParamSubjectEmail search by the domain the subject email. Operators: eq, contains, exists, not_exists.
	CertificateParamSubjectEmail = "subject_email"

	// CertificateParamFingerprintMD5 search by the MD5 fingerprint. Enter a hexadecimal string without spaces in
	// lowercase. Example:  500d1a99a670675a2e0b147b31ea17ce. Operators: eq.
	CertificateParamFingerprintMD5 = "fingerprint_md5"

	// CertificateParamFingerprintSHA1 search by the domain the SHA1 fingerprint. Enter a hexadecimal string without
	// spaces in lowercase. Example: f817a45f860581ff9911ce6405d3e194e90a0f3f.
	// Operators: eq.
	CertificateParamFingerprintSHA1 = "fingerprint_sha1"

	// CertificateParamFingerprintSHA256 search by the domain the SHA256 fingerprint. Enter a hexadecimal string
	// without spaces in lowercase. Example: f239e183459c727c7b2694937f04e5925ffaf969a10cc31ce848fdd0ccf6e40f.
	// Operators: eq.
	CertificateParamFingerprintSHA256 = "fingerprint_sha256"

	// CertificateParamValidityEnd search by the domain the certificate validity end date. Format: YYYY-MM-DD.
	// Operators: eq, gte, lte.
	CertificateParamValidityEnd = "validity_end"

	// CertificateParamValidityStart search by the domain the certificate validity start date. Format: YYYY-MM-DD.
	// Operators: eq, gte, lte.
	CertificateParamValidityStart = "validity_start"

	// CertificateParamIsTrusted search certificates that have been verified and validated by Spyse.
	// Operators: eq.
	CertificateParamIsTrusted = "is_trusted"
)

// Options for CVE search params:
//
// All search parameters see at https://spyse-dev.readme.io/reference/cves#cve_search
const (
	// CVEParamID search by the CVE ID defined by the MITRE Corporation. Operators: eq.
	CVEParamID = "id"

	// CVEParamCPE search by the Common Platform Enumeration (CPE) name or prefix. Example:
	// cpe:2.3:o:canonical:ubuntu_linux:12.04. Operators: eq, starts.
	CVEParamCPE = "cpe"

	// CVEParamScoreCVSS2 search by the CVE score according to the Common Vulnerability Scoring System
	// Version 2 (CVSS2). Operators: gte, lte.
	CVEParamScoreCVSS2 = "score_cvss2"

	// CVEParamScoreCVSS3 search by the CVE score according to the Common Vulnerability Scoring System
	// Version 3 (CVSS3). Operators: gte, lte.
	CVEParamScoreCVSS3 = "score_cvss3"

	// CVEParamSeverityCVSS2 search by the CVE severity according to CVSSv2. Supported options: high, medium, low.
	// Operators: eq.
	CVEParamSeverityCVSS2 = "severity_cvss2"

	// CVEParamSeverityCVSS3 search by the CVE severity according to CVSSv3. Supported options: high, medium, low,
	// critical. Operators: eq.
	CVEParamSeverityCVSS3 = "severity_cvss3"

	// CVEParamPublishedAt search by the vulnerability publication date. Format: YYYY-MM-DD. Operators: gte, lte.
	CVEParamPublishedAt = "published_at"

	// CVEParamModifiedAt search by the vulnerability modification date. Format: YYYY-MM-DD. Operators: gte, lte.
	CVEParamModifiedAt = "modified_at"
)

// Options for Email search params:
//
// ALl search parameters see at https://spyse-dev.readme.io/reference/emailss#email_search
const (
	// EmailParamEmail search by the email address. Operators: eq, not_eq, contains, not_contains, starts, ends.
	EmailParamEmail = "email"
)
