package spyse

func (s *IPService) Params() IPParams {
	return IPParams{
		CIDR: IPParamCIDR{
			Name: "cidr",
			Operator: IPCIDROperators{
				Equal: OperatorEqual,
			},
		},
		ISP: IPParamISP{
			Name: "isp",
			Operator: IPISPOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		PTR: IPParamPTR{
			Name: "ptr",
			Operator: IPPTROperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		ASOrg: IPParamASOrg{
			Name: "as_org",
			Operator: IPASOrgOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		ASNum: IPParamASNum{
			Name: "as_num",
			Operator: IPASNumOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		GeoCity: IPParamGeoCity{
			Name: "geo_city",
			Operator: IPGeoCityOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		GeoCountry: IPParamGeoCountry{
			Name: "geo_country",
			Operator: IPGeoCountryOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		GeoCountryISOCode: IPParamGeoCountryISOCode{
			Name: "geo_country_iso_code",
			Operator: IPGeoCountryISOCodeOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		TechnologyCPE: IPParamTechnologyCPE{
			Name: "technology_cpe",
			Operator: IPTechnologyCPEOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		PortTechnologyName: IPParamPortTechnologyName{
			Name: "port_technology_name",
			Operator: IPPortTechnologyNameOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		PortTechnologyVersion: IPParamPortTechnologyVersion{
			Name: "port_technology_version",
			Operator: IPPortTechnologyVersionOperators{
				Equal: OperatorEqual,
				Gte:   OperatorGreaterThanOrEqual,
				Lte:   OperatorLessThanOrEqual,
			},
		},
		OpenPort: IPParamOpenPort{
			Name: "open_port",
			Operator: IPOpenPortOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		PortCVEID: IPParamPortCVEID{
			Name: "port_cve_id",
			Operator: IPPortCVEIDOperators{
				Equal: OperatorEqual,
			},
		},
		PortBanner: IPParamPortBanner{
			Name: "port_banner",
			Operator: IPPortBannerOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		PortService: IPParamPortService{
			Name: "port_service",
			Operator: IPPortServiceOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractDescription: IPParamHTTPExtractDescription{
			Name: "http_extract_description",
			Operator: IPHTTPExtractDescriptionOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
				Exists:     OperatorExists,
				NotExists:  OperatorNotExists,
			},
		},
		HTTPExtractTitle: IPParamHTTPExtractTitle{
			Name: "http_extract_title",
			Operator: IPHTTPExtractTitleOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
				Exists:     OperatorExists,
				NotExists:  OperatorNotExists,
			},
		},
		HTTPExtractEmail: IPParamHTTPExtractEmail{
			Name: "http_extract_email",
			Operator: IPHTTPExtractEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractRobots: IPParamHTTPExtractRobots{
			Name: "http_extract_robots",
			Operator: IPHTTPExtractRobotsOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractStyles: IPParamHTTPExtractStyles{
			Name: "http_extract_styles",
			Operator: IPHTTPExtractStylesOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractScripts: IPParamHTTPExtractScripts{
			Name: "http_extract_scripts",
			Operator: IPHTTPExtractScriptsOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerAdsenseID: IPParamHTTPExtractTrackerAdsenseID{
			Name: "http_extract_tracker_adsense_id",
			Operator: IPHTTPExtractTrackerAdsenseIDOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractMetaName: IPParamHTTPExtractMetaName{
			Name: "http_extract_meta_name",
			Operator: IPHTTPExtractMetaNameOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractMetaValue: IPParamHTTPExtractMetaValue{
			Name: "http_extract_meta_value",
			Operator: IPHTTPExtractMetaValueOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractLinkHost: IPParamHTTPExtractLinkHost{
			Name: "http_extract_link_host",
			Operator: IPHTTPExtractLinkHostOperators{
				Contains:   OperatorContains,
				StartsWith: OperatorStartsWith,
				EndsWith:   OperatorEndsWith,
			},
		},
		HTTPExtractStatusCode: IPParamHTTPExtractStatusCode{
			Name: "http_extract_status_code",
			Operator: IPHTTPExtractStatusCodeOperators{
				Equal:     OperatorEqual,
				Gte:       OperatorGreaterThanOrEqual,
				Lte:       OperatorLessThanOrEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractFaviconSHA256: IPParamHTTPExtractFaviconSHA256{
			Name: "http_extract_favicon_sha256",
			Operator: IPHTTPExtractFaviconSHA256Operators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractHeadersName: IPParamHTTPExtractHeadersName{
			Name: "http_extract_headers_name",
			Operator: IPHTTPExtractHeadersNameOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractHeadersValue: IPParamHTTPExtractHeadersValue{
			Name: "http_extract_headers_value",
			Operator: IPHTTPExtractHeadersValueOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractLinkURL: IPParamHTTPExtractLinkURL{
			Name: "http_extract_link_url",
			Operator: IPHTTPExtractLinkURLOperators{
				Contains:   OperatorContains,
				StartsWith: OperatorStartsWith,
				EndsWith:   OperatorEndsWith,
			},
		},
		HTTPExtractFinalRedirectURL: IPParamHTTPExtractFinalRedirectURL{
			Name: "http_extract_final_redirect_url",
			Operator: IPHTTPExtractFinalRedirectURLOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerGoogleAnalyticsKey: IPParamHTTPExtractTrackerGoogleAnalyticsKey{
			Name: "http_extract_tracker_google_analytics_key",
			Operator: IPHTTPExtractTrackerGoogleAnalyticsKeyOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerGooglePlayApp: IPParamHTTPExtractTrackerGooglePlayApp{
			Name: "http_extract_tracker_google_play_app",
			Operator: IPHTTPExtractTrackerGooglePlayAppOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerAppleItunesApp: IPParamHTTPExtractTrackerAppleItunesApp{
			Name: "http_extract_tracker_apple_itunes_app",
			Operator: IPHTTPExtractTrackerAppleItunesAppOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerGoogleSiteVerification: IPParamHTTPExtractTrackerGoogleSiteVerification{
			Name: "http_extract_tracker_google_site_verification",
			Operator: IPHTTPExtractTrackerGoogleSiteVerificationOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		AbusesReportsNum: IPParamAbusesReportsNum{
			Name: "abuses_reports_num",
			Operator: IPAbusesReportsNumOperators{
				Gte: OperatorGreaterThanOrEqual,
				Lte: OperatorLessThanOrEqual,
			},
		},
		AbusesReportsComment: IPParamAbusesReportsComment{
			Name: "abuses_reports_comment",
			Operator: IPAbusesReportsCommentOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		AbusesConfidenceScore: IPParamAbusesConfidenceScore{
			Name: "abuses_confidence_score",
			Operator: IPAbusesConfidenceScoreOperators{
				Gte: OperatorGreaterThanOrEqual,
				Lte: OperatorLessThanOrEqual,
			},
		},
		AbusesCategoryName: IPParamAbusesCategoryName{
			Name: "abuses_category_name",
			Operator: IPAbusesCategoryNameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		AbusesIsWhitelistStrong: IPParamAbusesIsWhitelistStrong{
			Name: "abuses_is_whitelist_strong",
			Operator: IPAbusesIsWhitelistStrongOperators{
				Equal: OperatorEqual,
			},
		},
		AbusesIsWhitelistWeak: IPParamAbusesIsWhitelistWeak{
			Name: "abuses_is_whitelist_weak",
			Operator: IPAbusesIsWhitelistWeakOperators{
				Equal: OperatorEqual,
			},
		},
		SecurityScore: IPParamSecurityScore{
			Name: "security_score",
			Operator: IPSecurityScoreOperators{
				Gte: OperatorGreaterThanOrEqual,
				Lte: OperatorLessThanOrEqual,
			},
		},
	}
}

// IPParams for IPs search:
//
// ALl search parameters see at https://spyse-dev.readme.io/reference/ips#ip_search
type IPParams struct {
	// CIDR search IPs by CIDR.
	CIDR IPParamCIDR
	// ISP search IPs by the internet service provider name.
	ISP IPParamISP
	// PTR search IPs PTR record value.
	PTR IPParamPTR
	// ASOrg search IPs by the autonomous system organization name.
	ASOrg IPParamASOrg
	// ASNum search IPs by the autonomous system number.
	ASNum IPParamASNum
	// GeoCity search IPs by the city they are located in.
	GeoCity IPParamGeoCity
	// GeoCountry search IPs  by the country they are located in.
	GeoCountry IPParamGeoCountry
	// GeoCountryISOCode search IPs by the country ISO code.
	GeoCountryISOCode IPParamGeoCountryISOCode
	// TechnologyCPE search IPs by technology CPE. Example: "cpe:2.3:a:apache:http_server".
	TechnologyCPE IPParamTechnologyCPE
	// PortTechnologyName find IPs hosting the specified technology. Examples: “Pure-FTPd”, “Dovecot pop3d”,
	// "Postfix smtpd".
	PortTechnologyName IPParamPortTechnologyName
	// PortTechnologyVersion find IPs hosting the specified technology version. Examples: “2.4.18”
	// (Apache version), “7.4” (OpenSSH version), "1.3.5rc3" (ProFTPD version).
	PortTechnologyVersion IPParamPortTechnologyVersion
	// OpenPort the open port number to search for. For ex. by setting the value "22" and operator "eq"
	// you will search for all IP hosts that have port 22 opened.
	OpenPort IPParamOpenPort
	// PortCVEID find IPs affected by a specific CVE. Example: CVE-2019-12407.
	PortCVEID IPParamPortCVEID
	// PortBanner search IPs by its port banner content.
	PortBanner IPParamPortBanner
	// PortService search IPs by services detected on ports. Examples: "ssh”, “http”, “ftp” etc.
	PortService IPParamPortService
	// HTTPExtractDescription search IPs that host websites (on any port) with the homepage containing
	// the specified description.
	HTTPExtractDescription IPParamHTTPExtractDescription
	// HTTPExtractTitle search IPs that host websites (on any port) with the homepage containing the specified title.
	HTTPExtractTitle IPParamHTTPExtractTitle
	// HTTPExtractEmail search IPs that host websites (on any port) with the homepage containing the specified
	// email address.
	HTTPExtractEmail IPParamHTTPExtractEmail
	// HTTPExtractRobots search IPs that host websites (on any port) with the homepage containing the
	// specified robots.txt content.
	HTTPExtractRobots IPParamHTTPExtractRobots
	// HTTPExtractStyles search IPs that host websites (on any port) with the homepage containing <link>
	// tags with specified "href" attribute substring.
	HTTPExtractStyles IPParamHTTPExtractStyles
	// HTTPExtractScripts search IPs that host websites (on any port) with the homepage containing <script>
	// tags with specified "src" attribute substring.
	HTTPExtractScripts IPParamHTTPExtractScripts
	// HTTPExtractTrackerAdsenseID search IPs that host websites (on any port) with the homepage containing the
	// specified Google AdSense identifier.
	HTTPExtractTrackerAdsenseID IPParamHTTPExtractTrackerAdsenseID
	// HTTPExtractMetaName search IPs that host websites (on any port) with the homepage containing the
	// specified HTML meta tag name.
	HTTPExtractMetaName IPParamHTTPExtractMetaName
	// HTTPExtractMetaValue search IPs that host websites (on any port) with the homepage containing the
	// specified HTML meta tag value.
	HTTPExtractMetaValue IPParamHTTPExtractMetaValue
	// HTTPExtractLinkHost search IPs that host websites (on any port) with the homepage containing links
	// with the specified substrings in their URL host.
	HTTPExtractLinkHost IPParamHTTPExtractLinkHost
	// HTTPExtractStatusCode search IPs that host websites (on any port) returning the specified status code.
	HTTPExtractStatusCode IPParamHTTPExtractStatusCode
	// HTTPExtractFaviconSHA256 search IPs that host websites (on any port) with the homepage containing
	// favicon with the specified SHA256 hash value.
	HTTPExtractFaviconSHA256 IPParamHTTPExtractFaviconSHA256
	// HTTPExtractHeadersName search IPs that host websites (on any port) returning the specified HTTP
	// header name.
	HTTPExtractHeadersName IPParamHTTPExtractHeadersName
	// HTTPExtractHeadersValue search IPs that host websites (on any port) returning the specified HTTP
	// header value.
	HTTPExtractHeadersValue IPParamHTTPExtractHeadersValue
	// HTTPExtractLinkURL search IPs that host websites (on any port) with the homepage containing links
	// with the specified substrings in their URL.
	HTTPExtractLinkURL IPParamHTTPExtractLinkURL
	// HTTPExtractFinalRedirectURL find IPs that host websites redirecting to the specified URL.
	HTTPExtractFinalRedirectURL IPParamHTTPExtractFinalRedirectURL
	// HTTPExtractTrackerGoogleAnalyticsKey search IPs that host websites (on any port) with the homepage
	// containing the specified Google Analytics key.
	HTTPExtractTrackerGoogleAnalyticsKey IPParamHTTPExtractTrackerGoogleAnalyticsKey
	// HTTPExtractTrackerGooglePlayApp search IPs that host websites (on any port) with the homepage containing
	// the specified Google Play app identifier.
	HTTPExtractTrackerGooglePlayApp IPParamHTTPExtractTrackerGooglePlayApp
	// HTTPExtractTrackerAppleItunesApp search IPs that host websites (on any port) with the homepage containing
	// the specified Apple iTunes app identifier.
	HTTPExtractTrackerAppleItunesApp IPParamHTTPExtractTrackerAppleItunesApp
	// HTTPExtractTrackerGoogleSiteVerification Search IPs that host websites (on any port) with the homepage
	// containing the specified Google Site verification identifier.
	HTTPExtractTrackerGoogleSiteVerification IPParamHTTPExtractTrackerGoogleSiteVerification
	// AbusesReportsNum find IPs by the number of abuses received.
	AbusesReportsNum IPParamAbusesReportsNum
	// AbusesReportsComment find IPs by the abuses comment.
	AbusesReportsComment IPParamAbusesReportsComment
	// AbusesConfidenceScore find IPs by abuse confidence score.
	AbusesConfidenceScore IPParamAbusesConfidenceScore
	// AbusesCategoryName find IPs by abuse category name.
	AbusesCategoryName IPParamAbusesCategoryName
	// AbusesIsWhitelistStrong find IPs, which contains benign crawlers verified by reverse-forward DNS
	// resolution checks.
	AbusesIsWhitelistStrong IPParamAbusesIsWhitelistStrong
	// AbusesIsWhitelistWeak find IPs, that are within our CIDR whitelist, which contain net blocks
	// from organizations we believe to be benign or often mistaken for malicious.
	AbusesIsWhitelistWeak IPParamAbusesIsWhitelistWeak
	// SecurityScore search IPs by Spyse security score.
	SecurityScore IPParamSecurityScore
}

type IPParamCIDR struct {
	Name     string
	Operator IPCIDROperators
}

type IPCIDROperators struct {
	Equal string
}

type IPParamISP struct {
	Name     string
	Operator IPISPOperators
}

type IPISPOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamPTR struct {
	Name     string
	Operator IPPTROperators
}

type IPPTROperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamASOrg struct {
	Name     string
	Operator IPASOrgOperators
}

type IPASOrgOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamASNum struct {
	Name     string
	Operator IPASNumOperators
}

type IPASNumOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamGeoCity struct {
	Name     string
	Operator IPGeoCityOperators
}

type IPGeoCityOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamGeoCountry struct {
	Name     string
	Operator IPGeoCountryOperators
}

type IPGeoCountryOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamGeoCountryISOCode struct {
	Name     string
	Operator IPGeoCountryISOCodeOperators
}

type IPGeoCountryISOCodeOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamTechnologyCPE struct {
	Name     string
	Operator IPTechnologyCPEOperators
}

type IPTechnologyCPEOperators struct {
	Equal    string
	Contains string
}

type IPParamPortTechnologyName struct {
	Name     string
	Operator IPPortTechnologyNameOperators
}

type IPPortTechnologyNameOperators struct {
	Equal    string
	Contains string
}

type IPParamPortTechnologyVersion struct {
	Name     string
	Operator IPPortTechnologyVersionOperators
}

type IPPortTechnologyVersionOperators struct {
	Equal string
	Gte   string
	Lte   string
}

type IPParamOpenPort struct {
	Name     string
	Operator IPOpenPortOperators
}

type IPOpenPortOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamPortCVEID struct {
	Name     string
	Operator IPPortCVEIDOperators
}

type IPPortCVEIDOperators struct {
	Equal string
}

type IPParamPortBanner struct {
	Name     string
	Operator IPPortBannerOperators
}

type IPPortBannerOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type IPParamPortService struct {
	Name     string
	Operator IPPortServiceOperators
}

type IPPortServiceOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractDescription struct {
	Name     string
	Operator IPHTTPExtractDescriptionOperators
}

type IPHTTPExtractDescriptionOperators struct {
	Equal      string
	StartsWith string
	Exists     string
	NotExists  string
}

type IPParamHTTPExtractTitle struct {
	Name     string
	Operator IPHTTPExtractTitleOperators
}

type IPHTTPExtractTitleOperators struct {
	Equal      string
	StartsWith string
	Exists     string
	NotExists  string
}

type IPParamHTTPExtractEmail struct {
	Name     string
	Operator IPHTTPExtractEmailOperators
}

type IPHTTPExtractEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractRobots struct {
	Name     string
	Operator IPHTTPExtractRobotsOperators
}

type IPHTTPExtractRobotsOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractStyles struct {
	Name     string
	Operator IPHTTPExtractStylesOperators
}

type IPHTTPExtractStylesOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractScripts struct {
	Name     string
	Operator IPHTTPExtractScriptsOperators
}

type IPHTTPExtractScriptsOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractTrackerAdsenseID struct {
	Name     string
	Operator IPHTTPExtractTrackerAdsenseIDOperators
}

type IPHTTPExtractTrackerAdsenseIDOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractMetaName struct {
	Name     string
	Operator IPHTTPExtractMetaNameOperators
}

type IPHTTPExtractMetaNameOperators struct {
	Equal    string
	Contains string
}

type IPParamHTTPExtractMetaValue struct {
	Name     string
	Operator IPHTTPExtractMetaValueOperators
}

type IPHTTPExtractMetaValueOperators struct {
	Equal    string
	Contains string
}

type IPParamHTTPExtractLinkHost struct {
	Name     string
	Operator IPHTTPExtractLinkHostOperators
}

type IPHTTPExtractLinkHostOperators struct {
	Contains   string
	StartsWith string
	EndsWith   string
}

type IPParamHTTPExtractStatusCode struct {
	Name     string
	Operator IPHTTPExtractStatusCodeOperators
}

type IPHTTPExtractStatusCodeOperators struct {
	Equal     string
	Gte       string
	Lte       string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractFaviconSHA256 struct {
	Name     string
	Operator IPHTTPExtractFaviconSHA256Operators
}

type IPHTTPExtractFaviconSHA256Operators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractHeadersName struct {
	Name     string
	Operator IPHTTPExtractHeadersNameOperators
}

type IPHTTPExtractHeadersNameOperators struct {
	Equal    string
	Contains string
}

type IPParamHTTPExtractHeadersValue struct {
	Name     string
	Operator IPHTTPExtractHeadersValueOperators
}

type IPHTTPExtractHeadersValueOperators struct {
	Equal    string
	Contains string
}

type IPParamHTTPExtractLinkURL struct {
	Name     string
	Operator IPHTTPExtractLinkURLOperators
}

type IPHTTPExtractLinkURLOperators struct {
	Contains   string
	StartsWith string
	EndsWith   string
}

type IPParamHTTPExtractFinalRedirectURL struct {
	Name     string
	Operator IPHTTPExtractFinalRedirectURLOperators
}

type IPHTTPExtractFinalRedirectURLOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractTrackerGoogleAnalyticsKey struct {
	Name     string
	Operator IPHTTPExtractTrackerGoogleAnalyticsKeyOperators
}

type IPHTTPExtractTrackerGoogleAnalyticsKeyOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractTrackerGooglePlayApp struct {
	Name     string
	Operator IPHTTPExtractTrackerGooglePlayAppOperators
}

type IPHTTPExtractTrackerGooglePlayAppOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractTrackerAppleItunesApp struct {
	Name     string
	Operator IPHTTPExtractTrackerAppleItunesAppOperators
}

type IPHTTPExtractTrackerAppleItunesAppOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamHTTPExtractTrackerGoogleSiteVerification struct {
	Name     string
	Operator IPHTTPExtractTrackerGoogleSiteVerificationOperators
}

type IPHTTPExtractTrackerGoogleSiteVerificationOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type IPParamAbusesReportsNum struct {
	Name     string
	Operator IPAbusesReportsNumOperators
}

type IPAbusesReportsNumOperators struct {
	Gte string
	Lte string
}

type IPParamAbusesReportsComment struct {
	Name     string
	Operator IPAbusesReportsCommentOperators
}

type IPAbusesReportsCommentOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamAbusesConfidenceScore struct {
	Name     string
	Operator IPAbusesConfidenceScoreOperators
}

type IPAbusesConfidenceScoreOperators struct {
	Gte string
	Lte string
}

type IPParamAbusesCategoryName struct {
	Name     string
	Operator IPAbusesCategoryNameOperators
}

type IPAbusesCategoryNameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type IPParamAbusesIsWhitelistStrong struct {
	Name     string
	Operator IPAbusesIsWhitelistStrongOperators
}

type IPAbusesIsWhitelistStrongOperators struct {
	Equal string
}

type IPParamAbusesIsWhitelistWeak struct {
	Name     string
	Operator IPAbusesIsWhitelistWeakOperators
}

type IPAbusesIsWhitelistWeakOperators struct {
	Equal string
}

type IPParamSecurityScore struct {
	Name     string
	Operator IPSecurityScoreOperators
}

type IPSecurityScoreOperators struct {
	Gte string
	Lte string
}
