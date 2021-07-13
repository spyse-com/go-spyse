package spyse

func (s *DomainService) Params() DomainParams {
	return DomainParams{
		Name: DomainParamName{
			Name: "name",
			Operator: DomainNameOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
				EndsWith:   OperatorEndsWith,
				NotEqual:   OperatorNotEqual,
			},
		},
		AdsenseID: DomainParamAdsenseID{
			Name: "http_extract_tracker_adsense_id",
			Operator: DomainAdsenseIDOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		AlexaRank: DomainParamAlexaRank{
			Name: "alexa_rank",
			Operator: DomainAlexaRankOperators{
				Equal:     OperatorEqual,
				Gte:       OperatorGreaterThanOrEqual,
				Lte:       OperatorLessThanOrEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSA: DomainParamDNSA{
			Name: "dns_a",
			Operator: DomainDNSAOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSAAAA: DomainParamDNSAAAA{
			Name: "dns_aaaa",
			Operator: DomainDNSAAAAOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSNS: DomainParamDNSNS{
			Name: "dns_ns",
			Operator: DomainDNSNSOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
				Contains:   OperatorContains,
				Exists:     OperatorExists,
				NotExists:  OperatorNotExists,
			},
		},
		DNSMX: DomainParamDNSMX{
			Name: "dns_mx",
			Operator: DomainDNSMXOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
				Contains:   OperatorContains,
				Exists:     OperatorExists,
				NotExists:  OperatorNotExists,
			},
		},
		DNSTXT: DomainParamDNSTXT{
			Name: "dns_txt",
			Operator: DomainDNSTXTOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSCAA: DomainParamDNSCAA{
			Name: "dns_caa",
			Operator: DomainDNSCAAOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSCNAME: DomainParamDNSCNAME{
			Name: "dns_cname",
			Operator: DomainDNSCNAMEOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFRAW: DomainParamDNSSPFRAW{
			Name: "dns_spf_raw",
			Operator: DomainDNSSPFRAWOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFVersion: DomainParamDNSSPFVersion{
			Name: "dns_spf_version",
			Operator: DomainDNSSPFVersionOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFErrorsTarget: DomainParamDNSSPFErrorsTarget{
			Name: "dns_spf_errors_target",
			Operator: DomainDNSSPFErrorsTargetOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFModifiersName: DomainParamDNSSPFModifiersName{
			Name: "dns_spf_modifiers_name",
			Operator: DomainDNSSPFModifiersNameOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFMechanismsName: DomainParamDNSSPFMechanismsName{
			Name: "dns_spf_mechanisms_name",
			Operator: DomainDNSSPFMechanismsNameOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFModifiersValue: DomainParamDNSSPFModifiersValue{
			Name: "dns_spf_modifiers_value",
			Operator: DomainDNSSPFModifiersValueOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFMechanismsValue: DomainParamDNSSPFMechanismsValue{
			Name: "dns_spf_mechanisms_value",
			Operator: DomainDNSSPFMechanismsValueOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFErrorsDescription: DomainParamDNSSPFErrorsDescription{
			Name: "dns_spf_errors_description",
			Operator: DomainDNSSPFErrorsDescriptionOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		DNSSPFMechanismsQualifier: DomainParamDNSSPFMechanismsQualifier{
			Name: "dns_spf_mechanisms_qualifier",
			Operator: DomainDNSSPFMechanismsQualifierOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTitle: DomainParamHTTPExtractTitle{
			Name: "http_extract_title",
			Operator: DomainHTTPExtractTitleOperators{
				Equal:       OperatorEqual,
				StartsWith:  OperatorStartsWith,
				Contains:    OperatorContains,
				NotContains: OperatorNotContains,
				Exists:      OperatorExists,
				NotExists:   OperatorNotExists,
			},
		},
		HTTPExtractEmail: DomainParamHTTPExtractEmail{
			Name: "http_extract_email",
			Operator: DomainHTTPExtractEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractRobots: DomainParamHTTPExtractRobots{
			Name: "http_extract_robots",
			Operator: DomainHTTPExtractRobotsOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractStyles: DomainParamHTTPExtractStyles{
			Name: "http_extract_styles",
			Operator: DomainHTTPExtractStylesOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractScripts: DomainParamHTTPExtractScripts{
			Name: "http_extract_scripts",
			Operator: DomainHTTPExtractScriptsOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractMetaName: DomainParamHTTPExtractMetaName{
			Name: "http_extract_meta_name",
			Operator: DomainHTTPExtractMetaNameOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractLinkHost: DomainParamHTTPExtractLinkHost{
			Name: "http_extract_link_host",
			Operator: DomainHTTPExtractLinkHostOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
				EndsWith:   OperatorEndsWith,
			},
		},
		HTTPExtractMetaValue: DomainParamHTTPExtractMetaValue{
			Name: "http_extract_meta_value",
			Operator: DomainHTTPExtractMetaValueOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractFaviconURI: DomainParamHTTPExtractFaviconURI{
			Name: "http_extract_favicon_uri",
			Operator: DomainHTTPExtractFaviconURIOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractStatusCode: DomainParamHTTPExtractStatusCode{
			Name: "http_extract_status_code",
			Operator: DomainHTTPExtractStatusCodeOperators{
				Equal:     OperatorEqual,
				Gte:       OperatorGreaterThanOrEqual,
				Lte:       OperatorLessThanOrEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractFaviconSHA256: DomainParamHTTPExtractFaviconSHA256{
			Name: "http_extract_favicon_sha256",
			Operator: DomainHTTPExtractFaviconSHA256Operators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractHeadersName: DomainParamHTTPExtractHeadersName{
			Name: "http_extract_headers_name",
			Operator: DomainHTTPExtractHeadersNameOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractDescription: DomainParamHTTPExtractDescription{
			Name: "http_extract_description",
			Operator: DomainHTTPExtractDescriptionOperators{
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractHeadersValue: DomainParamHTTPExtractHeadersValue{
			Name: "http_extract_headers_value",
			Operator: DomainHTTPExtractHeadersValueOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		HTTPExtractLinkURL: DomainParamHTTPExtractLinkURL{
			Name: "http_extract_link_url",
			Operator: DomainHTTPExtractLinkURLOperators{
				Equal:      OperatorEqual,
				StartsWith: OperatorStartsWith,
				EndsWith:   OperatorEndsWith,
			},
		},
		HTTPExtractFinalRedirectURL: DomainParamHTTPExtractFinalRedirectURL{
			Name: "http_extract_final_redirect_url",
			Operator: DomainHTTPExtractFinalRedirectURLOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CVEID: DomainParamCVEID{
			Name: "cve_id",
			Operator: DomainCVEIDOperators{
				Equal: OperatorEqual,
			},
		},
		CVESeverity: DomainParamCVESeverity{
			Name: "cve_severity",
			Operator: DomainCVESeverityOperators{
				Equal: OperatorEqual,
			},
		},
		TechnologyName: DomainParamTechnologyName{
			Name: "technology_name",
			Operator: DomainTechnologyNameOperators{
				Equal:    OperatorEqual,
				Contains: OperatorContains,
			},
		},
		TechnologyVersion: DomainParamTechnologyVersion{
			Name: "technology_version",
			Operator: DomainTechnologyVersionOperators{
				Equal: OperatorEqual,
				Gte:   OperatorGreaterThanOrEqual,
				Lte:   OperatorLessThanOrEqual,
			},
		},
		CertificateSHA256: DomainParamCertificateSHA256{
			Name: "certificate_sha256",
			Operator: DomainCertificateSHA256Operators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateVersion: DomainParamCertificateVersion{
			Name: "certificate_version",
			Operator: DomainCertificateVersionOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WhoisRegistrarWhoisServer: DomainParamWhoisRegistrarWhoisServer{
			Name: "whois_registrar_whois_server",
			Operator: DomainWhoisRegistrarWhoisServerOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WhoisRegistrantOrg: DomainParamWhoisRegistrantOrg{
			Name: "whois_registrant_org",
			Operator: DomainWhoisRegistrantOrgOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WhoisRegistrarName: DomainParamWhoisRegistrarName{
			Name: "whois_registrar_name",
			Operator: DomainWhoisRegistrarNameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WhoisRegistrantName: DomainParamWhoisRegistrantName{
			Name: "whois_registrant_name",
			Operator: DomainWhoisRegistrantNameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WhoisRegistrarEmail: DomainParamWhoisRegistrarEmail{
			Name: "whois_registrar_email",
			Operator: DomainWhoisRegistrarEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WhoisRegistrantPhone: DomainParamWhoisRegistrantPhone{
			Name: "whois_registrant_phone",
			Operator: DomainWhoisRegistrantPhoneOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WhoisRegistrantEmail: DomainParamWhoisRegistrantEmail{
			Name: "whois_registrant_email",
			Operator: DomainWhoisRegistrantEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		WithoutSuffix: DomainParamWithoutSuffix{
			Name: "without_suffix",
			Operator: DomainWithoutSuffixOperators{
				Equal: OperatorEqual,
			},
		},
		HTTPExtractTrackerGoogleAnalyticsKey: DomainParamHTTPExtractTrackerGoogleAnalyticsKey{
			Name: "http_extract_tracker_google_analytics_key",
			Operator: DomainHTTPExtractTrackerGoogleAnalyticsKeyOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerGooglePlayApp: DomainParamHTTPExtractTrackerGooglePlayApp{
			Name: "http_extract_tracker_google_play_app",
			Operator: DomainHTTPExtractTrackerGooglePlayAppOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerAppleItunesApp: DomainParamHTTPExtractTrackerAppleItunesApp{
			Name: "http_extract_tracker_apple_itunes_app",
			Operator: DomainHTTPExtractTrackerAppleItunesAppOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		HTTPExtractTrackerGoogleSiteVerification: DomainParamHTTPExtractTrackerGoogleSiteVerification{
			Name: "http_extract_tracker_google_site_verification",
			Operator: DomainHTTPExtractTrackerGoogleSiteVerificationOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateIssuerOrg: DomainParamCertificateIssuerOrg{
			Name: "certificate_issuer_org",
			Operator: DomainCertificateIssuerOrgOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateIssuerCname: DomainParamCertificateIssuerCname{
			Name: "certificate_issuer_cname",
			Operator: DomainCertificateIssuerCnameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateIssuerEmail: DomainParamCertificateIssuerEmail{
			Name: "certificate_issuer_email",
			Operator: DomainCertificateIssuerEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateIssuerOrganizationalUnit: DomainParamCertificateIssuerOrganizationalUnit{
			Name: "certificate_issuer_organizational_unit",
			Operator: DomainCertificateIssuerOrganizationalUnitOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateIssuerCountry: DomainParamCertificateIssuerCountry{
			Name: "certificate_issuer_country",
			Operator: DomainCertificateIssuerCountryOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateIssuerState: DomainParamCertificateIssuerState{
			Name: "certificate_issuer_state",
			Operator: DomainCertificateIssuerStateOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateIssuerLocality: DomainParamCertificateIssuerLocality{
			Name: "certificate_issuer_locality",
			Operator: DomainCertificateIssuerLocalityOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectOrg: DomainParamCertificateSubjectOrg{
			Name: "certificate_subject_org",
			Operator: DomainCertificateSubjectOrgOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectCNAME: DomainParamCertificateSubjectCNAME{
			Name: "certificate_subject_cname",
			Operator: DomainCertificateSubjectCNAMEOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectEmail: DomainParamCertificateSubjectEmail{
			Name: "certificate_subject_email",
			Operator: DomainCertificateSubjectEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectOrganizationalUnit: DomainParamCertificateSubjectOrganizationalUnit{
			Name: "certificate_subject_organizational_unit",
			Operator: DomainCertificateSubjectOrganizationalUnitOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectCountry: DomainParamCertificateSubjectCountry{
			Name: "certificate_subject_country",
			Operator: DomainCertificateSubjectCountryOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectState: DomainParamCertificateSubjectState{
			Name: "certificate_subject_state",
			Operator: DomainCertificateSubjectStateOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectLocality: DomainParamCertificateSubjectLocality{
			Name: "certificate_subject_locality",
			Operator: DomainCertificateSubjectLocalityOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateSubjectSerialNumber: DomainParamCertificateSubjectSerialNumber{
			Name: "certificate_subject_serial_number",
			Operator: DomainCertificateSubjectSerialNumberOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		CertificateValidityEnd: DomainParamCertificateValidityEnd{
			Name: "certificate_validity_end",
			Operator: DomainCertificateValidityEndOperators{
				Equal: OperatorEqual,
				Gte:   OperatorGreaterThanOrEqual,
				Lte:   OperatorLessThanOrEqual,
			},
		},
		GeoCountryISOCode: DomainParamGeoCountryISOCode{
			Name: "geo_country_iso_code",
			Operator: DomainGeoCountryISOCodeOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		GeoCountry: DomainParamGeoCountry{
			Name: "geo_country",
			Operator: DomainGeoCountryOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		GeoCity: DomainParamGeoCity{
			Name: "geo_city",
			Operator: DomainGeoCityOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		ASNum: DomainParamASNum{
			Name: "as_num",
			Operator: DomainASNumOperators{
				Equal:     OperatorEqual,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		OrganizationIndustry: DomainParamOrganizationIndustry{
			Name: "organization_industry",
			Operator: DomainOrganizationIndustryOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		OrganizationEmail: DomainParamOrganizationEmail{
			Name: "organization_email",
			Operator: DomainOrganizationEmailOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		OrganizationName: DomainParamOrganizationName{
			Name: "organization_name",
			Operator: DomainOrganizationNameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		OrganizationLegalName: DomainParamOrganizationLegalName{
			Name: "organization_legal_name",
			Operator: DomainOrganizationLegalNameOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		ISP: DomainParamISP{
			Name: "isp",
			Operator: DomainISPOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		ASOrg: DomainParamASOrg{
			Name: "as_org",
			Operator: DomainASOrgOperators{
				Equal:     OperatorEqual,
				Contains:  OperatorContains,
				Exists:    OperatorExists,
				NotExists: OperatorNotExists,
			},
		},
		IsPTR: DomainParamIsPTR{
			Name: "is_ptr",
			Operator: DomainIsPTROperators{
				Equal: OperatorEqual,
			},
		},
		IsMX: DomainParamIsMX{
			Name: "is_mx",
			Operator: DomainIsMXOperators{
				Equal: OperatorEqual,
			},
		},
		IsNS: DomainParamIsNS{
			Name: "is_ns",
			Operator: DomainIsNSOperators{
				Equal: OperatorEqual,
			},
		},
		IsSubdomain: DomainParamIsSubdomain{
			Name: "is_subdomain",
			Operator: DomainIsSubdomainOperators{
				Equal: OperatorEqual,
			},
		},
		IsCNAME: DomainParamIsCNAME{
			Name: "is_cname",
			Operator: DomainIsCNAMEOperators{
				Equal: OperatorEqual,
			},
		},
	}
}

// DomainParams for Domain search:
//
// All search parameters see at https://spyse-dev.readme.io/reference/domains#domain_search
type DomainParams struct {
	// Name search domains by the domain name. Operators: eq, starts, ends, not_eq.
	Name DomainParamName
	// AdsenseID search domains with the homepage containing the specified Google AdSense identifier.
	AdsenseID DomainParamAdsenseID
	// AlexaRank search by Alexa Rank value.
	AlexaRank DomainParamAlexaRank
	// DNSA search by a specific IPv4 address, stored in DNS A record, or CIDR.
	DNSA DomainParamDNSA
	// DNSAAAA search by a specific IPv6 address stored in DNS AAAA record.
	DNSAAAA DomainParamDNSAAAA
	// DNSNS search by name server (NS) address stored in DNS NS record. Example: ns1.google.com.
	DNSNS DomainParamDNSNS
	// DNSMX search by name server (MX) address stored in DNS MX record. Example: ns1.google.com.
	DNSMX DomainParamDNSMX
	// DNSTXT search by content of DNS TXT record.
	DNSTXT DomainParamDNSTXT
	// DNSCAA search by content of DNS CAA record.
	DNSCAA DomainParamDNSCAA
	// DNSCNAME search by content of DNS CNAME record.
	DNSCNAME DomainParamDNSCNAME
	// DNSSPFRAW search by not parsed SPF record.
	DNSSPFRAW DomainParamDNSSPFRAW
	// DNSSPFVersion search by an SPF record version. F.e., 'spf1'.
	DNSSPFVersion DomainParamDNSSPFVersion
	// DNSSPFErrorsTarget search for mechanism or modifier, which includes an error.
	DNSSPFErrorsTarget DomainParamDNSSPFErrorsTarget
	// DNSSPFModifiersName search by a modifier name. F.e., 'ip4', 'redirect', 'include'.
	DNSSPFModifiersName DomainParamDNSSPFModifiersName
	// DNSSPFMechanismsName search by the name of the mechanism e.g. 'ip4', 'ip6', 'all'.
	DNSSPFMechanismsName DomainParamDNSSPFMechanismsName
	// DNSSPFModifiersValue search by a modifier value, e.g. '_spf.mailhostbox.com', '_spf.yandex.net.
	DNSSPFModifiersValue DomainParamDNSSPFModifiersValue
	// DNSSPFMechanismsValue search by a mechanism value, e.g. '195.201.206.85', '_spf.google.com'.
	DNSSPFMechanismsValue DomainParamDNSSPFMechanismsValue
	// DNSSPFErrorsDescription search by the description of a validation error.
	DNSSPFErrorsDescription DomainParamDNSSPFErrorsDescription
	// DNSSPFMechanismsQualifier search by a qualifier of the mechanism.
	DNSSPFMechanismsQualifier DomainParamDNSSPFMechanismsQualifier
	// HTTPExtractTitle search by content of the HTML title tag.
	HTTPExtractTitle DomainParamHTTPExtractTitle
	// HTTPExtractEmail search by emails found on the web-page.
	HTTPExtractEmail DomainParamHTTPExtractEmail
	// HTTPExtractRobots search by content of the robots.txt file.
	HTTPExtractRobots DomainParamHTTPExtractRobots
	// HTTPExtractStyles search by links found in the href attribute of the link HTML tag.
	HTTPExtractStyles DomainParamHTTPExtractStyles
	// HTTPExtractScripts search by links found in the src attribute of the script HTML tag.
	HTTPExtractScripts DomainParamHTTPExtractScripts
	// HTTPExtractMetaName search by HTML meta tag name.
	HTTPExtractMetaName DomainParamHTTPExtractMetaName
	// HTTPExtractLinkHost search by a set of links found on sites.
	HTTPExtractLinkHost DomainParamHTTPExtractLinkHost
	// HTTPExtractMetaValue search by HTML meta tag value.
	HTTPExtractMetaValue DomainParamHTTPExtractMetaValue
	// HTTPExtractFaviconURI search by website favicon URI.
	HTTPExtractFaviconURI DomainParamHTTPExtractFaviconURI
	// HTTPExtractStatusCode search by HTTP response status code.
	HTTPExtractStatusCode DomainParamHTTPExtractStatusCode
	// HTTPExtractFaviconSHA256 search by SHA256 hash of the site favicon.
	HTTPExtractFaviconSHA256 DomainParamHTTPExtractFaviconSHA256
	// HTTPExtractHeadersName search by HTTP response header name.
	HTTPExtractHeadersName DomainParamHTTPExtractHeadersName
	// HTTPExtractDescription search by content of the HTML description meta tag.
	HTTPExtractDescription DomainParamHTTPExtractDescription
	// HTTPExtractHeadersValue search by HTTP response header value.
	HTTPExtractHeadersValue DomainParamHTTPExtractHeadersValue
	// HTTPExtractLinkURL search by a full link found on sites.
	HTTPExtractLinkURL DomainParamHTTPExtractLinkURL
	// HTTPExtractFinalRedirectURL find hosts that redirect to a specific domain.
	HTTPExtractFinalRedirectURL DomainParamHTTPExtractFinalRedirectURL
	// CVEID search by Common Vulnerability and Exposures (CVE) ID. Example: CVE-2019-16905.
	CVEID DomainParamCVEID
	// CVESeverity search by CVE severity. Options: "high", "medium", "low".
	CVESeverity DomainParamCVESeverity
	// TechnologyName search by technology name. Example: PHP.
	TechnologyName DomainParamTechnologyName
	// TechnologyVersion search by technology version. Example: 5.6.40.
	TechnologyVersion DomainParamTechnologyVersion
	// CertificateSHA256 search by certificate SHA256 fingerprint.
	CertificateSHA256 DomainParamCertificateSHA256
	// CertificateVersion search by a specific SSL/TLS version. Example: SSLv3.
	CertificateVersion DomainParamCertificateVersion
	// WhoisRegistrarWhoisServer search by WHOIS server domain name.
	WhoisRegistrarWhoisServer DomainParamWhoisRegistrarWhoisServer
	// WhoisRegistrantOrg search by the name of organization who has registered the domain name.
	WhoisRegistrantOrg DomainParamWhoisRegistrantOrg
	// WhoisRegistrarName search by registrar name.
	WhoisRegistrarName DomainParamWhoisRegistrarName
	// WhoisRegistrantName search by registrant name.
	WhoisRegistrantName DomainParamWhoisRegistrantName
	// WhoisRegistrarEmail search by registrar email.
	WhoisRegistrarEmail DomainParamWhoisRegistrarEmail
	// WhoisRegistrantPhone search by registrant phone.
	WhoisRegistrantPhone DomainParamWhoisRegistrantPhone
	// WhoisRegistrantEmail search by registrant email.
	WhoisRegistrantEmail DomainParamWhoisRegistrantEmail
	// WithoutSuffix Search by domain name without public suffix.
	WithoutSuffix DomainParamWithoutSuffix
	// HTTPExtractTrackerGoogleAnalyticsKey search domains with the homepage containing the specified
	// Google Analytics key.
	HTTPExtractTrackerGoogleAnalyticsKey DomainParamHTTPExtractTrackerGoogleAnalyticsKey
	// HTTPExtractTrackerGooglePlayApp search domains with the homepage containing the specified Google
	// Play app identifier.
	HTTPExtractTrackerGooglePlayApp DomainParamHTTPExtractTrackerGooglePlayApp
	// HTTPExtractTrackerAppleItunesApp search domains with the homepage containing the specified Apple
	// iTunes app identifier.
	HTTPExtractTrackerAppleItunesApp DomainParamHTTPExtractTrackerAppleItunesApp
	// HTTPExtractTrackerGoogleSiteVerification search domains with the homepage containing the specified
	// Google Site verification identifier.
	HTTPExtractTrackerGoogleSiteVerification DomainParamHTTPExtractTrackerGoogleSiteVerification
	// CertificateIssuerOrg search domains by certificate issuer organization.
	CertificateIssuerOrg DomainParamCertificateIssuerOrg
	// CertificateIssuerCname search domains by certificate issuer CNAME.
	CertificateIssuerCname DomainParamCertificateIssuerCname
	// CertificateIssuerEmail search domains by certificate issuer email.
	CertificateIssuerEmail DomainParamCertificateIssuerEmail
	// CertificateIssuerOrganizationalUnit search domains by certificate issuer organizational unit.
	CertificateIssuerOrganizationalUnit DomainParamCertificateIssuerOrganizationalUnit
	// CertificateIssuerCountry search domains by certificate issuer country.
	CertificateIssuerCountry DomainParamCertificateIssuerCountry
	// CertificateIssuerState search domains by certificate issuer state.
	CertificateIssuerState DomainParamCertificateIssuerState
	// CertificateIssuerLocality search domains by certificate issuer locality.
	CertificateIssuerLocality DomainParamCertificateIssuerLocality
	// CertificateSubjectOrg search domains by certificate subject organization.
	CertificateSubjectOrg DomainParamCertificateSubjectOrg
	// CertificateSubjectCNAME search domains by certificate subject CNAME.
	CertificateSubjectCNAME DomainParamCertificateSubjectCNAME
	// CertificateSubjectEmail search domains by certificate subject email.
	CertificateSubjectEmail DomainParamCertificateSubjectEmail
	// CertificateSubjectOrganizationalUnit search domains by certificate subject organizational unit.
	CertificateSubjectOrganizationalUnit DomainParamCertificateSubjectOrganizationalUnit
	// CertificateSubjectCountry search domains by certificate subject country.
	CertificateSubjectCountry DomainParamCertificateSubjectCountry
	// CertificateSubjectState search domains by certificate subject state.
	CertificateSubjectState DomainParamCertificateSubjectState
	// CertificateSubjectLocality search domains by certificate subject locality.
	CertificateSubjectLocality DomainParamCertificateSubjectLocality
	// CertificateSubjectSerialNumber search domains by certificate subject serial number.
	CertificateSubjectSerialNumber DomainParamCertificateSubjectSerialNumber
	// CertificateValidityEnd search domains by certificate validity end date.
	CertificateValidityEnd DomainParamCertificateValidityEnd
	// GeoCountryISOCode find domains by the ISO code of the country they are located in.
	GeoCountryISOCode DomainParamGeoCountryISOCode
	// GeoCountry find domains by the country they are located in.
	GeoCountry DomainParamGeoCountry
	// GeoCity find domains by the city they are located in.
	GeoCity DomainParamGeoCity
	// ASNum find domains by the autonomous system number.
	ASNum DomainParamASNum
	// OrganizationIndustry find domains by the organization industry.
	OrganizationIndustry DomainParamOrganizationIndustry
	// OrganizationEmail find domains by the organization email.
	OrganizationEmail DomainParamOrganizationEmail
	// OrganizationName find domains by the organization name.
	OrganizationName DomainParamOrganizationName
	// OrganizationLegalName find domains by the organization legal name.
	OrganizationLegalName DomainParamOrganizationLegalName
	// ISP find domains by the autonomous system ISP.
	ISP DomainParamISP
	// ASOrg find domains by the autonomous system organization.
	ASOrg DomainParamASOrg
	// IsPTR find domain names that have been found in any DNS PTR record.
	IsPTR DomainParamIsPTR
	// IsMX find domains that are fully qualified domain names of a mail server.
	IsMX DomainParamIsMX
	// IsNS find domains that are fully qualified domain names of a nameserver.
	IsNS DomainParamIsNS
	// IsSubdomain find domains that are subdomains.
	IsSubdomain DomainParamIsSubdomain
	// IsCNAME find domain names that have been found in any CNAME record.
	IsCNAME DomainParamIsCNAME
}

type DomainParamName struct {
	Name     string
	Operator DomainNameOperators
}

type DomainNameOperators struct {
	Equal      string
	StartsWith string
	EndsWith   string
	NotEqual   string
}

type DomainParamAdsenseID struct {
	Name     string
	Operator DomainAdsenseIDOperators
}

type DomainAdsenseIDOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamAlexaRank struct {
	Name     string
	Operator DomainAlexaRankOperators
}

type DomainAlexaRankOperators struct {
	Equal     string
	Gte       string
	Lte       string
	Exists    string
	NotExists string
}

type DomainParamDNSA struct {
	Name     string
	Operator DomainDNSAOperators
}

type DomainDNSAOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamDNSAAAA struct {
	Name     string
	Operator DomainDNSAAAAOperators
}

type DomainDNSAAAAOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamDNSNS struct {
	Name     string
	Operator DomainDNSNSOperators
}

type DomainDNSNSOperators struct {
	Equal      string
	StartsWith string
	Contains   string
	Exists     string
	NotExists  string
}

type DomainParamDNSMX struct {
	Name     string
	Operator DomainDNSMXOperators
}

type DomainDNSMXOperators struct {
	Equal      string
	StartsWith string
	Contains   string
	Exists     string
	NotExists  string
}

type DomainParamDNSTXT struct {
	Name     string
	Operator DomainDNSTXTOperators
}

type DomainDNSTXTOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSCAA struct {
	Name     string
	Operator DomainDNSCAAOperators
}

type DomainDNSCAAOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSCNAME struct {
	Name     string
	Operator DomainDNSCNAMEOperators
}

type DomainDNSCNAMEOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFRAW struct {
	Name     string
	Operator DomainDNSSPFRAWOperators
}

type DomainDNSSPFRAWOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFVersion struct {
	Name     string
	Operator DomainDNSSPFVersionOperators
}

type DomainDNSSPFVersionOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFErrorsTarget struct {
	Name     string
	Operator DomainDNSSPFErrorsTargetOperators
}

type DomainDNSSPFErrorsTargetOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFModifiersName struct {
	Name     string
	Operator DomainDNSSPFModifiersNameOperators
}

type DomainDNSSPFModifiersNameOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFMechanismsName struct {
	Name     string
	Operator DomainDNSSPFMechanismsNameOperators
}

type DomainDNSSPFMechanismsNameOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFModifiersValue struct {
	Name     string
	Operator DomainDNSSPFModifiersValueOperators
}

type DomainDNSSPFModifiersValueOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFMechanismsValue struct {
	Name     string
	Operator DomainDNSSPFMechanismsValueOperators
}

type DomainDNSSPFMechanismsValueOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFErrorsDescription struct {
	Name     string
	Operator DomainDNSSPFErrorsDescriptionOperators
}

type DomainDNSSPFErrorsDescriptionOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamDNSSPFMechanismsQualifier struct {
	Name     string
	Operator DomainDNSSPFMechanismsQualifierOperators
}

type DomainDNSSPFMechanismsQualifierOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractTitle struct {
	Name     string
	Operator DomainHTTPExtractTitleOperators
}

type DomainHTTPExtractTitleOperators struct {
	Equal       string
	StartsWith  string
	Contains    string
	NotContains string
	Exists      string
	NotExists   string
}

type DomainParamHTTPExtractEmail struct {
	Name     string
	Operator DomainHTTPExtractEmailOperators
}

type DomainHTTPExtractEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractRobots struct {
	Name     string
	Operator DomainHTTPExtractRobotsOperators
}

type DomainHTTPExtractRobotsOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractStyles struct {
	Name     string
	Operator DomainHTTPExtractStylesOperators
}

type DomainHTTPExtractStylesOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractScripts struct {
	Name     string
	Operator DomainHTTPExtractScriptsOperators
}

type DomainHTTPExtractScriptsOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractMetaName struct {
	Name     string
	Operator DomainHTTPExtractMetaNameOperators
}

type DomainHTTPExtractMetaNameOperators struct {
	Equal    string
	Contains string
}

type DomainParamHTTPExtractLinkHost struct {
	Name     string
	Operator DomainHTTPExtractLinkHostOperators
}

type DomainHTTPExtractLinkHostOperators struct {
	Equal      string
	StartsWith string
	EndsWith   string
}

type DomainParamHTTPExtractMetaValue struct {
	Name     string
	Operator DomainHTTPExtractMetaValueOperators
}

type DomainHTTPExtractMetaValueOperators struct {
	Equal    string
	Contains string
}

type DomainParamHTTPExtractFaviconURI struct {
	Name     string
	Operator DomainHTTPExtractFaviconURIOperators
}

type DomainHTTPExtractFaviconURIOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractStatusCode struct {
	Name     string
	Operator DomainHTTPExtractStatusCodeOperators
}

type DomainHTTPExtractStatusCodeOperators struct {
	Equal     string
	Gte       string
	Lte       string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractFaviconSHA256 struct {
	Name     string
	Operator DomainHTTPExtractFaviconSHA256Operators
}

type DomainHTTPExtractFaviconSHA256Operators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractHeadersName struct {
	Name     string
	Operator DomainHTTPExtractHeadersNameOperators
}

type DomainHTTPExtractHeadersNameOperators struct {
	Equal    string
	Contains string
}

type DomainParamHTTPExtractDescription struct {
	Name     string
	Operator DomainHTTPExtractDescriptionOperators
}

type DomainHTTPExtractDescriptionOperators struct {
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractHeadersValue struct {
	Name     string
	Operator DomainHTTPExtractHeadersValueOperators
}

type DomainHTTPExtractHeadersValueOperators struct {
	Equal    string
	Contains string
}

type DomainParamHTTPExtractLinkURL struct {
	Name     string
	Operator DomainHTTPExtractLinkURLOperators
}

type DomainHTTPExtractLinkURLOperators struct {
	Equal      string
	StartsWith string
	EndsWith   string
}

type DomainParamHTTPExtractFinalRedirectURL struct {
	Name     string
	Operator DomainHTTPExtractFinalRedirectURLOperators
}

type DomainHTTPExtractFinalRedirectURLOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCVEID struct {
	Name     string
	Operator DomainCVEIDOperators
}

type DomainCVEIDOperators struct {
	Equal string
}

type DomainParamCVESeverity struct {
	Name     string
	Operator DomainCVESeverityOperators
}

type DomainCVESeverityOperators struct {
	Equal string
}

type DomainParamTechnologyName struct {
	Name     string
	Operator DomainTechnologyNameOperators
}

type DomainTechnologyNameOperators struct {
	Equal    string
	Contains string
}

type DomainParamTechnologyVersion struct {
	Name     string
	Operator DomainTechnologyVersionOperators
}

type DomainTechnologyVersionOperators struct {
	Equal string
	Gte   string
	Lte   string
}

type DomainParamCertificateSHA256 struct {
	Name     string
	Operator DomainCertificateSHA256Operators
}

type DomainCertificateSHA256Operators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamCertificateVersion struct {
	Name     string
	Operator DomainCertificateVersionOperators
}

type DomainCertificateVersionOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamWhoisRegistrarWhoisServer struct {
	Name     string
	Operator DomainWhoisRegistrarWhoisServerOperators
}

type DomainWhoisRegistrarWhoisServerOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamWhoisRegistrantOrg struct {
	Name     string
	Operator DomainWhoisRegistrantOrgOperators
}

type DomainWhoisRegistrantOrgOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamWhoisRegistrarName struct {
	Name     string
	Operator DomainWhoisRegistrarNameOperators
}

type DomainWhoisRegistrarNameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamWhoisRegistrantName struct {
	Name     string
	Operator DomainWhoisRegistrantNameOperators
}

type DomainWhoisRegistrantNameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamWhoisRegistrarEmail struct {
	Name     string
	Operator DomainWhoisRegistrarEmailOperators
}

type DomainWhoisRegistrarEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamWhoisRegistrantPhone struct {
	Name     string
	Operator DomainWhoisRegistrantPhoneOperators
}

type DomainWhoisRegistrantPhoneOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamWhoisRegistrantEmail struct {
	Name     string
	Operator DomainWhoisRegistrantEmailOperators
}

type DomainWhoisRegistrantEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamWithoutSuffix struct {
	Name     string
	Operator DomainWithoutSuffixOperators
}

type DomainWithoutSuffixOperators struct {
	Equal string
}

type DomainParamHTTPExtractTrackerGoogleAnalyticsKey struct {
	Name     string
	Operator DomainHTTPExtractTrackerGoogleAnalyticsKeyOperators
}

type DomainHTTPExtractTrackerGoogleAnalyticsKeyOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractTrackerGooglePlayApp struct {
	Name     string
	Operator DomainHTTPExtractTrackerGooglePlayAppOperators
}

type DomainHTTPExtractTrackerGooglePlayAppOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractTrackerAppleItunesApp struct {
	Name     string
	Operator DomainHTTPExtractTrackerAppleItunesAppOperators
}

type DomainHTTPExtractTrackerAppleItunesAppOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamHTTPExtractTrackerGoogleSiteVerification struct {
	Name     string
	Operator DomainHTTPExtractTrackerGoogleSiteVerificationOperators
}

type DomainHTTPExtractTrackerGoogleSiteVerificationOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamCertificateIssuerOrg struct {
	Name     string
	Operator DomainCertificateIssuerOrgOperators
}

type DomainCertificateIssuerOrgOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateIssuerCname struct {
	Name     string
	Operator DomainCertificateIssuerCnameOperators
}

type DomainCertificateIssuerCnameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateIssuerEmail struct {
	Name     string
	Operator DomainCertificateIssuerEmailOperators
}

type DomainCertificateIssuerEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateIssuerOrganizationalUnit struct {
	Name     string
	Operator DomainCertificateIssuerOrganizationalUnitOperators
}

type DomainCertificateIssuerOrganizationalUnitOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateIssuerCountry struct {
	Name     string
	Operator DomainCertificateIssuerCountryOperators
}

type DomainCertificateIssuerCountryOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamCertificateIssuerState struct {
	Name     string
	Operator DomainCertificateIssuerStateOperators
}

type DomainCertificateIssuerStateOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateIssuerLocality struct {
	Name     string
	Operator DomainCertificateIssuerLocalityOperators
}

type DomainCertificateIssuerLocalityOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectOrg struct {
	Name     string
	Operator DomainCertificateSubjectOrgOperators
}

type DomainCertificateSubjectOrgOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectCNAME struct {
	Name     string
	Operator DomainCertificateSubjectCNAMEOperators
}

type DomainCertificateSubjectCNAMEOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectEmail struct {
	Name     string
	Operator DomainCertificateSubjectEmailOperators
}

type DomainCertificateSubjectEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectOrganizationalUnit struct {
	Name     string
	Operator DomainCertificateSubjectOrganizationalUnitOperators
}

type DomainCertificateSubjectOrganizationalUnitOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectCountry struct {
	Name     string
	Operator DomainCertificateSubjectCountryOperators
}

type DomainCertificateSubjectCountryOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectState struct {
	Name     string
	Operator DomainCertificateSubjectStateOperators
}

type DomainCertificateSubjectStateOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectLocality struct {
	Name     string
	Operator DomainCertificateSubjectLocalityOperators
}

type DomainCertificateSubjectLocalityOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamCertificateSubjectSerialNumber struct {
	Name     string
	Operator DomainCertificateSubjectSerialNumberOperators
}

type DomainCertificateSubjectSerialNumberOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamCertificateValidityEnd struct {
	Name     string
	Operator DomainCertificateValidityEndOperators
}

type DomainCertificateValidityEndOperators struct {
	Equal string
	Gte   string
	Lte   string
}

type DomainParamGeoCountryISOCode struct {
	Name     string
	Operator DomainGeoCountryISOCodeOperators
}

type DomainGeoCountryISOCodeOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamGeoCountry struct {
	Name     string
	Operator DomainGeoCountryOperators
}

type DomainGeoCountryOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamGeoCity struct {
	Name     string
	Operator DomainGeoCityOperators
}

type DomainGeoCityOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamASNum struct {
	Name     string
	Operator DomainASNumOperators
}

type DomainASNumOperators struct {
	Equal     string
	Exists    string
	NotExists string
}

type DomainParamOrganizationIndustry struct {
	Name     string
	Operator DomainOrganizationIndustryOperators
}

type DomainOrganizationIndustryOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamOrganizationEmail struct {
	Name     string
	Operator DomainOrganizationEmailOperators
}

type DomainOrganizationEmailOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamOrganizationName struct {
	Name     string
	Operator DomainOrganizationNameOperators
}

type DomainOrganizationNameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamOrganizationLegalName struct {
	Name     string
	Operator DomainOrganizationLegalNameOperators
}

type DomainOrganizationLegalNameOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamISP struct {
	Name     string
	Operator DomainISPOperators
}

type DomainISPOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamASOrg struct {
	Name     string
	Operator DomainASOrgOperators
}

type DomainASOrgOperators struct {
	Equal     string
	Contains  string
	Exists    string
	NotExists string
}

type DomainParamIsPTR struct {
	Name     string
	Operator DomainIsPTROperators
}

type DomainIsPTROperators struct {
	Equal string
}

type DomainParamIsMX struct {
	Name     string
	Operator DomainIsMXOperators
}

type DomainIsMXOperators struct {
	Equal string
}

type DomainParamIsNS struct {
	Name     string
	Operator DomainIsNSOperators
}

type DomainIsNSOperators struct {
	Equal string
}

type DomainParamIsSubdomain struct {
	Name     string
	Operator DomainIsSubdomainOperators
}

type DomainIsSubdomainOperators struct {
	Equal string
}

type DomainParamIsCNAME struct {
	Name     string
	Operator DomainIsCNAMEOperators
}

type DomainIsCNAMEOperators struct {
	Equal string
}
