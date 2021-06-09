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
	// DomainParamName search domains by the domain name. Operators: eq, starts, ends, not_eq.
	DomainParamName = "name"

	// DomainParamAdsenseID search domains with the homepage containing the specified Google AdSense identifier.
	// Operators: eq, exists, not_exists.
	DomainParamAdsenseID = "http_extract_tracker_adsense_id"

	// DomainParamAlexaRank search by Alexa Rank value. Operators: eq, gte, lte, exists, not_exists.
	DomainParamAlexaRank = "alexa_rank"

	// DomainParamDNSA search by a specific IPv4 address, stored in DNS A record, or CIDR.
	// Operators: eq, exists, not_exists.
	DomainParamDNSA = "dns_a"

	// DomainParamDNSAAAA search by a specific IPv6 address stored in DNS AAAA record.
	// Operators: eq, exists, not_exists.
	DomainParamDNSAAAA = "dns_aaaa"

	// DomainParamDNSNS search by name server (NS) address stored in DNS NS record. Example: ns1.google.com.
	// Operators: eq, contains, starts, exists, not_exists.
	DomainParamDNSNS = "dns_ns"

	// DomainParamDNSMX search by name server (MX) address stored in DNS MX record. Example: ns1.google.com.
	//	// Operators: eq, contains, starts, exists, not_exists.
	DomainParamDNSMX = "dns_mx"

	// DomainParamDNSTXT search by content of DNS TXT record. Operators: eq, contains, exists, not_exists.
	DomainParamDNSTXT = "dns_txt"

	// DomainParamDNSCAA search by content of DNS CAA record. Operators: eq, contains, exists, not_exists.
	DomainParamDNSCAA = "dns_caa"

	// DomainParamDNSCNAME search by content of DNS CNAME record. Operators: eq, contains, exists, not_exists.
	DomainParamDNSCNAME = "dns_cname"

	// DomainParamDNSSPFRAW search by not parsed SPF record. Operators: contains, exists, not_exists.
	DomainParamDNSSPFRAW = "dns_spf_raw"

	// DomainParamDNSSPFVersion search by an SPF record version. F.e., 'spf1'. Operators: eq, exists, not_exists.
	DomainParamDNSSPFVersion = "dns_spf_version"

	// DomainParamDNSSPFErrorsTarget search for mechanism or modifier, which includes an error.
	// Operators: contains, exists, not_exists.
	DomainParamDNSSPFErrorsTarget = "dns_spf_errors_target"

	// DomainParamDNSSPFModifiersName search by a modifier name. F.e., 'ip4', 'redirect', 'include'.
	// Operators: contains, exists, not_exists.
	DomainParamDNSSPFModifiersName = "dns_spf_modifiers_name"

	// DomainParamDNSSPFMechanismsName search by the name of the mechanism e.g. 'ip4', 'ip6', 'all'.
	// Operators: eq, exists, not_exists.
	DomainParamDNSSPFMechanismsName = "dns_spf_mechanisms_name"

	// DomainParamDNSSPFModifiersValue search by a modifier value, e.g. '_spf.mailhostbox.com', '_spf.yandex.net.
	// Operators: contains, exists, not_exists.
	DomainParamDNSSPFModifiersValue = "dns_spf_modifiers_value"

	// DomainParamDNSSPFMechanismsValue search by a mechanism value, e.g. '195.201.206.85', '_spf.google.com'.
	// Operators: contains, exists, not_exists.
	DomainParamDNSSPFMechanismsValue = "dns_spf_mechanisms_value"

	// DomainParamDNSSPFErrorsDescription search by the description of a validation error.
	// Operators: contains, exists, not_exists.
	DomainParamDNSSPFErrorsDescription = "dns_spf_errors_description"

	// DomainParamDNSSPFMechanismsQualifier search by a qualifier of the mechanism.
	// Operators: contains, exists, not_exists.
	DomainParamDNSSPFMechanismsQualifier = "dns_spf_mechanisms_qualifier"

	// DomainParamHTTPExtractTitle search by content of the HTML title tag.
	// Operators: eq, contains, starts, exists, not_exists.
	DomainParamHTTPExtractTitle = "http_extract_title"

	// DomainParamHTTPExtractEmail search by emails found on the web-page. Operators: eq, contains, exists, not_exists.
	DomainParamHTTPExtractEmail = "http_extract_email"

	// DomainParamHTTPExtractRobots search by content of the robots.txt file. Operators: contains, exists, not_exists.
	DomainParamHTTPExtractRobots = "http_extract_robots"

	// DomainParamHTTPExtractStyles search by links found in the href attribute of the link HTML tag.
	// Operators: contains, exists, not_exists.
	DomainParamHTTPExtractStyles = "http_extract_styles"

	// DomainParamHTTPExtractScripts search by links found in the src attribute of the script HTML tag.
	// Operators: contains, exists, not_exists.
	DomainParamHTTPExtractScripts = "http_extract_scripts"

	// DomainParamHTTPExtractMetaName search by HTML meta tag name. Operators: eq, contains.
	DomainParamHTTPExtractMetaName = "http_extract_meta_name"

	// DomainParamHTTPExtractLinkHost search by a set of links found on sites. Operators: eq, starts, ends.
	DomainParamHTTPExtractLinkHost = "http_extract_link_host"

	// DomainParamHTTPExtractMetaValue search by HTML meta tag value. Operators: eq, contains.
	DomainParamHTTPExtractMetaValue = "http_extract_meta_value"

	// DomainParamHTTPExtractFaviconURI search by website favicon URI. Operators: eq, contains, exists, not_exists.
	DomainParamHTTPExtractFaviconURI = "http_extract_favicon_uri"

	// DomainParamHTTPExtractStatusCode search by HTTP response status code.
	// Operators: eq, gte, lte, exists, not_exists.
	DomainParamHTTPExtractStatusCode = "http_extract_status_code"

	// DomainParamHTTPExtractFaviconSHA256 search by SHA256 hash of the site favicon.
	// Operators: eq, exists, not_exists.
	DomainParamHTTPExtractFaviconSHA256 = "http_extract_favicon_sha256"

	// DomainParamHTTPExtractHeadersName search by HTTP response header name.
	// Operators: eq, contains.
	DomainParamHTTPExtractHeadersName = "http_extract_headers_name"

	// DomainParamHTTPExtractDescription search by content of the HTML description meta tag.
	// Operators: contains, exists, not_exists.
	DomainParamHTTPExtractDescription = "http_extract_description"

	// DomainParamHTTPExtractHeadersValue search by HTTP response header value.
	// Operators: eq, contains.
	DomainParamHTTPExtractHeadersValue = "http_extract_headers_value"

	// DomainParamHTTPExtractLinkURL search by a full link found on sites. Operators: eq, starts, ends.
	DomainParamHTTPExtractLinkURL = "http_extract_link_url"

	// DomainParamHTTPExtractFinalRedirectURL find hosts that redirect to a specific domain.
	// Operators: eq, contains, exists, not_exists.
	DomainParamHTTPExtractFinalRedirectURL = "http_extract_final_redirect_url"

	// DomainParamCVEID search by Common Vulnerability and Exposures (CVE) ID. Example: CVE-2019-16905.
	// Operators: eq.
	DomainParamCVEID = "cve_id"

	// DomainParamCVESeverity search by CVE severity. Options: "high", "medium", "low". Operators: eq.
	DomainParamCVESeverity = "cve_severity"

	// DomainParamTechnologyName search by technology name. Example: PHP. Operators: eq, contains.
	DomainParamTechnologyName = "technology_name"

	// DomainParamTechnologyVersion search by technology version. Example: 5.6.40. Operators: eq, gte, lte.
	DomainParamTechnologyVersion = "technology_version"

	// DomainParamCertificateSHA256 search by certificate SHA256 fingerprint. Operators: eq, exists, not_exists.
	DomainParamCertificateSHA256 = "certificate_sha256"

	// DomainParamCertificateVersion search by a specific SSL/TLS version. Example: SSLv3.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateVersion = "certificate_version"

	// DomainParamWhoisRegistrarWhoisServer search by WHOIS server domain name. Operators: eq, exists, not_exists.
	DomainParamWhoisRegistrarWhoisServer = "whois_registrar_whois_server"

	// DomainParamWhoisRegistrantOrg search by the name of organization who has registered the domain name.
	// Operators: eq, contains, exists, not_exists.
	DomainParamWhoisRegistrantOrg = "whois_registrant_org"

	// DomainParamWhoisRegistrarName search by registrar name. Operators: eq, contains, exists, not_exists.
	DomainParamWhoisRegistrarName = "whois_registrar_name"

	// DomainParamWhoisRegistrantName search by registrant name. Operators: eq, contains, exists, not_exists.
	DomainParamWhoisRegistrantName = "whois_registrant_name"

	// DomainParamWhoisRegistrarEmail search by registrar email. Operators: eq, contains, exists, not_exists.
	DomainParamWhoisRegistrarEmail = "whois_registrar_email"

	// DomainParamWhoisRegistrantPhone search by registrant phone. Operators: eq, exists, not_exists.
	DomainParamWhoisRegistrantPhone = "whois_registrant_phone"

	// DomainParamWhoisRegistrantEmail search by registrant email. Operators: eq, contains, exists, not_exists.
	DomainParamWhoisRegistrantEmail = "whois_registrant_email"

	// DomainParamWithoutSuffix Search by domain name without public suffix. Example: google. Operators: eq.
	DomainParamWithoutSuffix = "without_suffix"

	// DomainParamHTTPExtractTrackerGoogleAnalyticsKey search domains with the homepage containing the specified
	// Google Analytics key. Operators: eq, exists, not_exists.
	DomainParamHTTPExtractTrackerGoogleAnalyticsKey = "http_extract_tracker_google_analytics_key"

	// DomainParamHTTPExtractTrackerGooglePlayApp search domains with the homepage containing the specified Google
	// Play app identifier. Operators: eq, contains, exists, not_exists.
	DomainParamHTTPExtractTrackerGooglePlayApp = "http_extract_tracker_google_play_app"

	// DomainParamHTTPExtractTrackerAppleItunesApp search domains with the homepage containing the specified Apple
	// iTunes app identifier. Operators: eq, contains, exists, not_exists.
	DomainParamHTTPExtractTrackerAppleItunesApp = "http_extract_tracker_apple_itunes_app"

	// DomainParamHTTPExtractTrackerGoogleSiteVerification search domains with the homepage containing the specified
	// Google Site verification identifier. Operators: eq, exists, not_exists.
	DomainParamHTTPExtractTrackerGoogleSiteVerification = "http_extract_tracker_google_site_verification"

	// DomainParamCertificateIssuerOrg search domains by certificate issuer organization.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateIssuerOrg = "certificate_issuer_org"

	// DomainParamCertificateIssuerCname search domains by certificate issuer CNAME.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateIssuerCname = "certificate_issuer_cname"

	// DomainParamCertificateIssuerEmail search domains by certificate issuer email.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateIssuerEmail = "certificate_issuer_email"

	// DomainParamCertificateIssuerOrganizationalUnit search domains by certificate issuer organizational unit.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateIssuerOrganizationalUnit = "certificate_issuer_organizational_unit"

	// DomainParamCertificateIssuerCountry search domains by certificate issuer country.
	// Operators: eq, exists, not_exists.
	DomainParamCertificateIssuerCountry = "certificate_issuer_country"

	// DomainParamCertificateIssuerState search domains by certificate issuer state.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateIssuerState = "certificate_issuer_state"

	// DomainParamCertificateIssuerLocality search domains by certificate issuer locality.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateIssuerLocality = "certificate_issuer_locality"

	// DomainParamCertificateSubjectOrg search domains by certificate subject organization.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateSubjectOrg = "certificate_subject_org"

	// DomainParamCertificateSubjectCNAME search domains by certificate subject CNAME.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateSubjectCNAME = "certificate_subject_cname"

	// DomainParamCertificateSubjectEmail search domains by certificate subject email.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateSubjectEmail = "certificate_subject_email"

	// DomainParamCertificateSubjectOrganizationalUnit search domains by certificate subject organizational unit.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateSubjectOrganizationalUnit = "certificate_subject_organizational_unit"

	// DomainParamCertificateSubjectCountry search domains by certificate subject country.
	// Operators: eq, exists, not_exists.
	DomainParamCertificateSubjectCountry = "certificate_subject_country"

	// DomainParamCertificateSubjectState search domains by certificate subject state.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateSubjectState = "certificate_subject_state"

	// DomainParamCertificateSubjectLocality search domains by certificate subject locality.
	// Operators: eq, contains, exists, not_exists.
	DomainParamCertificateSubjectLocality = "certificate_subject_locality"

	// DomainParamCertificateSubjectSerialNumber search domains by certificate subject serial number.
	// Operators: eq, exists, not_exists.
	DomainParamCertificateSubjectSerialNumber = "certificate_subject_serial_number"

	// DomainParamCertificateValidityEnd search domains by certificate validity end date.
	// Operators: eq, gte, lte.
	DomainParamCertificateValidityEnd = "certificate_validity_end"

	// DomainParamGeoCountryISOCode find domains by the ISO code of the country they are located in.
	// Operators: eq, exists, not_exists.
	DomainParamGeoCountryISOCode = "geo_country_iso_code"

	// DomainParamGeoCountry find domains by the country they are located in.
	// Operators: eq, exists, not_exists.
	DomainParamGeoCountry = "geo_country"

	// DomainParamGeoCity find domains by the city they are located in. Operators: eq, contains, exists, not_exists.
	DomainParamGeoCity = "geo_city"

	// DomainParamASNum find domains by the autonomous system number. Operators: eq, exists, not_exists.
	DomainParamASNum = "as_num"

	// DomainParamOrganizationIndustry find domains by the organization industry.
	// Operators: eq, contains, exists, not_exists.
	DomainParamOrganizationIndustry = "organization_industry"

	// DomainParamOrganizationEmail find domains by the organization email. Operators: eq, contains, exists, not_exists.
	DomainParamOrganizationEmail = "organization_email"

	// DomainParamOrganizationName find domains by the organization name. Operators: eq, contains, exists, not_exists.
	DomainParamOrganizationName = "organization_name"

	// DomainParamOrganizationLegalName find domains by the organization legal name.
	// Operators: eq, contains, exists, not_exists.
	DomainParamOrganizationLegalName = "organization_legal_name"

	// DomainParamISP find domains by the autonomous system ISP. Operators: eq, contains, exists, not_exists.
	DomainParamISP = "isp"

	// DomainParamASOrg find domains by the autonomous system organization. Operators: eq, contains, exists, not_exists.
	DomainParamASOrg = "as_org"

	// DomainParamIsPTR find domain names that have been found in any DNS PTR record. Operators: eq.
	DomainParamIsPTR = "is_ptr"

	// DomainParamIsMX find domains that are fully qualified domain names of a mail server. Operators: eq.
	DomainParamIsMX = "is_mx"

	// DomainParamIsNS find domains that are fully qualified domain names of a nameserver. Operators: eq.
	DomainParamIsNS = "is_ns"

	// DomainParamIsSubDomainParam find domains that are subdomains. Operators: eq.
	DomainParamIsSubDomainParam = "is_subdomain"

	// DomainParamIsCNAME find domain names that have been found in any CNAME record. Operators: eq.
	DomainParamIsCNAME = "is_cname"
)

// Options for IP search params:
//
// ALl search parameters see at https://spyse-dev.readme.io/reference/ips#ip_search
const (
	// IPParamCIDR search IPs by CIDR. Operators: eq.
	IPParamCIDR = "cidr"

	// IPParamISP search IPs by the internet services provider name. Operators: eq, contains, exists, not_exists.
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
