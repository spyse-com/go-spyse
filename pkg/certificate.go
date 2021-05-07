package spyse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	CertificateDetailsEndpoint     = "certificate/"
	CertificateSearchEndpoint      = "certificate/search"
	CertificateSearchCountEndpoint = "certificate/search/count"
)

// CertificateService handles SSL/TLS Certificates for the Spyse API.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ssltls-certificates
type CertificateService struct {
	client *Client
}

// Certificate represents certificate information
type Certificate struct {
	Parsed            CertParsed  `json:"parsed,omitempty"`
	Raw               string      `json:"raw,omitempty"`
	FingerprintSHA256 string      `json:"fingerprint_sha256,omitempty"`
	Validation        *Validation `json:"validation,omitempty"`
	UpdatedAt         string      `json:"updated_at,omitempty"`
}

type CertParsed struct {
	Extensions             CertParsedExtensions         `json:"extensions,omitempty"`
	FingerprintMd5         string                       `json:"fingerprint_md5,omitempty"`
	FingerprintSha1        string                       `json:"fingerprint_sha1,omitempty"`
	FingerprintSha256      string                       `json:"fingerprint_sha256,omitempty"`
	Issuer                 CertParsedIssuer             `json:"issuer,omitempty"`
	IssuerDN               string                       `json:"issuer_dn,omitempty"`
	Names                  []string                     `json:"names,omitempty"`
	Redacted               *bool                        `json:"redacted,omitempty"`
	SerialNumber           string                       `json:"serial_number,omitempty"`
	Signature              CertParsedSignature          `json:"signature,omitempty"`
	SignatureAlgorithm     CertParsedSignatureAlgorithm `json:"signature_algorithm,omitempty"`
	SpkiSubjectFingerprint string                       `json:"spki_subject_fingerprint,omitempty"`
	Subject                CertParsedSubject            `json:"subject,omitempty"`
	SubjectDN              string                       `json:"subject_dn,omitempty"`
	SubjectKeyInfo         CertParsedSubjectKeyInfo     `json:"subject_key_info,omitempty"`
	TbsFingerprint         string                       `json:"tbs_fingerprint,omitempty"`
	TbsNoctFingerprint     string                       `json:"tbs_noct_fingerprint,omitempty"`
	ValidationLevel        string                       `json:"validation_level,omitempty"`
	Validity               CertParsedValidity           `json:"validity,omitempty"`
	Version                int64                        `json:"version,omitempty"`
}

type CertParsedExtensions struct {
	AuthorityInfoAccess         CertParsedExtensionsAuthorityInfoAccess `json:"authority_info_access,omitempty"`
	AuthorityKeyID              string                                  `json:"authority_key_id,omitempty"`
	BasicConstraints            CertParsedExtensionsBasicConstraints    `json:"basic_constraints,omitempty"`
	CertificatePolicies         []CertParsedExtensionsCertPolicies      `json:"certificate_policies,omitempty"`
	CrlDistributionPoints       []string                                `json:"crl_distribution_points,omitempty"`
	ExtendedKeyUsage            CertParsedExtensionsExtendedKeyUsage    `json:"extended_key_usage,omitempty"`
	IssuerAltName               CertIssuerAltName                       `json:"issuer_alt_name,omitempty"`
	KeyUsage                    CertParsedExtensionsKeyUsage            `json:"key_usage,omitempty"`
	NameConstraints             []CertParsedExtensionsNameConstraints   `json:"name_constraints,omitempty"`
	SignedCertificateTimestamps []SignedCertificateTimestamps           `json:"signed_certificate_timestamps,omitempty"`
	SubjectAltName              CertParsedExtensionsSubjectAltName      `json:"subject_alt_name,omitempty"`
	SubjectKeyID                string                                  `json:"subject_key_id,omitempty"`
}

type CertParsedExtensionsAuthorityInfoAccess struct {
	IssuerUrls []string `json:"issuer_urls,omitempty"`
	OCSPUrls   []string `json:"ocsp_urls,omitempty"`
}

type CertParsedExtensionsBasicConstraints struct {
	IsCA       *bool `json:"is_ca,omitempty"`
	MaxPathLen *bool `json:"max_path_len,omitempty"`
}

type CertParsedExtensionsCertPolicies struct {
	CPS        []string                           `json:"cps,omitempty"`
	ID         string                             `json:"id,omitempty"`
	UserNotice []ExtensionsCertPoliciesUserNotice `json:"user_notice,omitempty"`
}

type ExtensionsCertPoliciesUserNotice struct {
	ExplicitText    string          `json:"explicit_text,omitempty"`
	NoticeReference NoticeReference `json:"notice_reference,omitempty"`
}

type NoticeReference struct {
	NoticeNumbers int64  `json:"notice_numbers,omitempty"`
	Organization  string `json:"organization,omitempty"`
}

type CertParsedExtensionsExtendedKeyUsage struct {
	Any                            *bool `json:"any,omitempty"`
	AppleIChatEncryption           *bool `json:"apple_ichat_encryption,omitempty"`
	AppleIChatSigning              *bool `json:"apple_ichat_signing,omitempty"`
	AppleSystemIdentity            *bool `json:"apple_system_identity,omitempty"`
	ClientAuth                     *bool `json:"client_auth,omitempty"`
	CodeSigning                    *bool `json:"code_signing,omitempty"`
	DVCS                           *bool `json:"dvcs,omitempty"`
	EapOverLan                     *bool `json:"eap_over_lan,omitempty"`
	EapOverPPP                     *bool `json:"eap_over_ppp,omitempty"`
	EmailProtection                *bool `json:"email_protection,omitempty"`
	IpsecEndSystem                 *bool `json:"ipsec_end_system,omitempty"`
	IpsecIntermediateSystemUsage   *bool `json:"ipsec_intermediate_system_usage,omitempty"`
	IpsecTunnel                    *bool `json:"ipsec_tunnel,omitempty"`
	IpsecUser                      *bool `json:"ipsec_user,omitempty"`
	MicrosoftCAExchange            *bool `json:"microsoft_ca_exchange,omitempty"`
	MicrosoftCertTrustListSigning  *bool `json:"microsoft_cert_trust_list_signing,omitempty"`
	MicrosoftDocumentSigning       *bool `json:"microsoft_document_signing,omitempty"`
	MicrosoftDRM                   *bool `json:"microsoft_drm,omitempty"`
	MicrosoftEFSRecovery           *bool `json:"microsoft_efs_recovery,omitempty"`
	MicrosoftEmbeddedNTCrypto      *bool `json:"microsoft_embedded_nt_crypto,omitempty"`
	MicrosoftEncryptedFileSystem   *bool `json:"microsoft_encrypted_file_system,omitempty"`
	MicrosoftEnrollmentAgent       *bool `json:"microsoft_enrollment_agent,omitempty"`
	MicrosoftKernelModeCodeSigning *bool `json:"microsoft_kernel_mode_code_signing,omitempty"`
	MicrosoftKeyRecovery21         *bool `json:"microsoft_key_recovery_21,omitempty"`
	MicrosoftKeyRecovery3          *bool `json:"microsoft_key_recovery_3,omitempty"`
	MicrosoftLifetimeSigning       *bool `json:"microsoft_lifetime_signing,omitempty"`
	MicrosoftNT5Crypto             *bool `json:"microsoft_nt5_crypto,omitempty"`
	MicrosoftOEMWHQLCrypto         *bool `json:"microsoft_oem_whql_crypto,omitempty"`
	MicrosoftQualifiedSubordinate  *bool `json:"microsoft_qualified_subordinate,omitempty"`
	MicrosoftRootListSigner        *bool `json:"microsoft_root_list_signer,omitempty"`
	MicrosoftServerGatedCrypto     *bool `json:"microsoft_server_gated_crypto,omitempty"`
	MicrosoftSmartcardLogon        *bool `json:"microsoft_smartcard_logon,omitempty"`
	MicrosoftSystemHealth          *bool `json:"microsoft_system_health,omitempty"`
	MicrosoftTimestampSigning      *bool `json:"microsoft_timestamp_signing,omitempty"`
	MicrosoftWHQLCrypto            *bool `json:"microsoft_whql_crypto,omitempty"`
	SbgpCertAAServiceAuth          *bool `json:"sbgp_cert_aa_service_auth,omitempty"`
	ServerAuth                     *bool `json:"server_auth,omitempty"`
	TimeStamping                   *bool `json:"time_stamping,omitempty"`
}

type CertIssuerAltName struct {
	DirectoryNames             []CertIssuerAltNameDirectoryNames `json:"directory_names,omitempty"`
	DNSNames                   []string                          `json:"dns_names,omitempty"`
	EmailAddresses             []string                          `json:"email_addresses,omitempty"`
	IPAddresses                []string                          `json:"ip_addresses,omitempty"`
	OtherNames                 []CertAltNameOtherNames           `json:"other_names,omitempty"`
	RegisteredIDs              []string                          `json:"registered_ids,omitempty"`
	UniformResourceIdentifiers []string                          `json:"uniform_resource_identifiers,omitempty"`
}

type CertIssuerAltNameDirectoryNames struct {
	CommonName         string `json:"common_name,omitempty"`
	Country            string `json:"country,omitempty"`
	DomainComponent    string `json:"domain_component,omitempty"`
	Locality           string `json:"locality,omitempty"`
	Organization       string `json:"organization,omitempty"`
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	SerialNumber       string `json:"serial_number,omitempty"`
	StreetAddress      string `json:"street_address,omitempty"`
}

type CertAltNameOtherNames struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type CertParsedExtensionsKeyUsage struct {
	CertificateSign   *bool `json:"certificate_sign,omitempty"`
	ContentCommitment *bool `json:"content_commitment,omitempty"`
	CRLSign           *bool `json:"crl_sign,omitempty"`
	DataEncipherment  *bool `json:"data_encipherment,omitempty"`
	DecipherOnly      *bool `json:"decipher_only,omitempty"`
	DigitalSignature  *bool `json:"digital_signature,omitempty"`
	EncipherOnly      *bool `json:"encipher_only,omitempty"`
	KeyAgreement      *bool `json:"key_agreement,omitempty"`
	KeyEncipherment   *bool `json:"key_encipherment,omitempty"`
	Value             int64 `json:"value,omitempty"`
}

type CertParsedExtensionsNameConstraints struct {
	Critical                *bool    `json:"critical,omitempty"`
	PermittedEmailAddresses []string `json:"permitted_email_addresses,omitempty"`
	PermittedNames          []string `json:"permitted_names,omitempty"`
}

type SignedCertificateTimestamps struct {
	LogID     string `json:"log_id,omitempty"`
	Signature string `json:"signature,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Version   int64  `json:"version,omitempty"`
}

type CertParsedExtensionsSubjectAltName struct {
	DirectoryNames             []CertSubjectAltNameDirectoryNames `json:"directory_names,omitempty"`
	DNSNames                   []string                           `json:"dns_names,omitempty"`
	EmailAddresses             []string                           `json:"email_addresses,omitempty"`
	IPAddresses                []string                           `json:"ip_addresses,omitempty"`
	OtherNames                 []CertAltNameOtherNames            `json:"other_names,omitempty"`
	RegisteredIDs              []string                           `json:"registered_ids,omitempty"`
	UniformResourceIdentifiers []string                           `json:"uniform_resource_identifiers,omitempty"`
}

type CertSubjectAltNameDirectoryNames struct {
	CommonName         string `json:"common_name,omitempty"`
	Country            string `json:"country,omitempty"`
	DomainComponent    string `json:"domain_component,omitempty"`
	EmailAddress       string `json:"email_address,omitempty"`
	Locality           string `json:"locality,omitempty"`
	Organization       string `json:"organization,omitempty"`
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	Province           string `json:"province,omitempty"`
	SerialNumber       string `json:"serial_number,omitempty"`
	Surname            string `json:"surname,omitempty"`
}

type CertParsedIssuer struct {
	CommonName           []string `json:"common_name,omitempty"`
	Country              []string `json:"country,omitempty"`
	DomainComponent      []string `json:"domain_component,omitempty"`
	EmailAddress         []string `json:"email_address,omitempty"`
	GivenName            []string `json:"given_name,omitempty"`
	JurisdictionCountry  []string `json:"jurisdiction_country,omitempty"`
	JurisdictionLocality []string `json:"jurisdiction_locality,omitempty"`
	JurisdictionProvince []string `json:"jurisdiction_province,omitempty"`
	Locality             []string `json:"locality,omitempty"`
	Organization         []string `json:"organization,omitempty"`
	OrganizationalUnit   []string `json:"organizational_unit,omitempty"`
	PostalCode           []string `json:"postal_code,omitempty"`
	Province             []string `json:"province,omitempty"`
	SerialNumber         []string `json:"serial_number,omitempty"`
	StreetAddress        []string `json:"street_address,omitempty"`
	Surname              []string `json:"surname,omitempty"`
}

type CertParsedSignatureAlgorithm struct {
	Name string `json:"name,omitempty"`
	OID  string `json:"oid,omitempty"`
}

type CertParsedSignature struct {
	SelfSigned         *bool                        `json:"self_signed,omitempty"`
	SignatureAlgorithm CertParsedSignatureAlgorithm `json:"signature_algorithm,omitempty"`
	Valid              *bool                        `json:"valid,omitempty"`
	Value              string                       `json:"value,omitempty"`
}

type CertParsedSubject struct {
	CommonName           []string `json:"common_name,omitempty"`
	Country              []string `json:"country,omitempty"`
	DomainComponent      []string `json:"domain_component,omitempty"`
	EmailAddress         []string `json:"email_address,omitempty"`
	GivenName            []string `json:"given_name,omitempty"`
	JurisdictionCountry  []string `json:"jurisdiction_country,omitempty"`
	JurisdictionLocality []string `json:"jurisdiction_locality,omitempty"`
	JurisdictionProvince []string `json:"jurisdiction_province,omitempty"`
	Locality             []string `json:"locality,omitempty"`
	Organization         []string `json:"organization,omitempty"`
	OrganizationalUnit   []string `json:"organizational_unit,omitempty"`
	PostalCode           []string `json:"postal_code,omitempty"`
	Province             []string `json:"province,omitempty"`
	SerialNumber         []string `json:"serial_number,omitempty"`
	StreetAddress        []string `json:"street_address,omitempty"`
	Surname              []string `json:"surname,omitempty"`
}

type EcdsaPublicKey struct {
	B      string `json:"b,omitempty"`
	Curve  string `json:"curve,omitempty"`
	Gx     string `json:"gx,omitempty"`
	Gy     string `json:"gy,omitempty"`
	Length int64  `json:"length,omitempty"`
	N      string `json:"n,omitempty"`
	P      string `json:"p,omitempty"`
	Pub    string `json:"pub,omitempty"`
	X      string `json:"x,omitempty"`
	Y      string `json:"y,omitempty"`
}

type KeyAlgorithm struct {
	Name string `json:"name,omitempty"`
}

type RSAPublicKey struct {
	Exponent int64  `json:"exponent,omitempty"`
	Length   int64  `json:"length,omitempty"`
	Modulus  string `json:"modulus,omitempty"`
}

type CertParsedSubjectKeyInfo struct {
	EcdsaPublicKey    EcdsaPublicKey `json:"ecdsa_public_key,omitempty"`
	FingerprintSha256 string         `json:"fingerprint_sha256,omitempty"`
	KeyAlgorithm      KeyAlgorithm   `json:"key_algorithm,omitempty"`
	RSAPublicKey      RSAPublicKey   `json:"rsa_public_key,omitempty"`
}

type CertParsedValidity struct {
	Status string `json:"status,omitempty"`
	End    string `json:"end,omitempty"`
	Length int64  `json:"length,omitempty"`
	Start  string `json:"start,omitempty"`
}

type Validation struct {
	Reason string `json:"reason,omitempty"`
	Valid  *bool  `json:"is_valid,omitempty"`
}

// Details returns a full representation of the Certificate for the given fingerprintSHA256.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ssltls-certificates#certificate
func (s *CertificateService) Details(ctx context.Context, fingerprintSHA256 string) (*Certificate, error) {
	refURI := fmt.Sprintf(CertificateDetailsEndpoint+"%s", fingerprintSHA256)
	req, err := s.client.NewRequest(ctx, http.MethodGet, refURI, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &Certificate{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	if len(resp.Data.Items) > 0 {
		return resp.Data.Items[0].(*Certificate), nil
	}

	return nil, nil
}

// Search returns a list of Certificates that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ssltls-certificates#certificate_search
func (s *CertificateService) Search(
	ctx context.Context,
	filters []map[string]Filter,
	limit, offset int,
) ([]*Certificate, error) {
	body, err := json.Marshal(
		SearchRequest{
			SearchParams: filters,
			PaginatedRequest: PaginatedRequest{
				Size: limit,
				From: offset,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, CertificateSearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &Certificate{})
	if err != nil {
		return nil, NewSpyseError(err)
	}

	var items []*Certificate

	if len(resp.Data.Items) > 0 {
		for _, i := range resp.Data.Items {
			items = append(items, i.(*Certificate))
		}

		return items, nil
	}

	return nil, nil
}

// SearchCount returns a count of Certificates that match the specified filters.
//
// Spyse API docs: https://spyse-dev.readme.io/reference/ssltls-certificates#certificate_search_count
func (s *CertificateService) SearchCount(ctx context.Context, filters []map[string]Filter) (int64, error) {
	body, err := json.Marshal(SearchRequest{SearchParams: filters})
	if err != nil {
		return 0, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, CertificateSearchCountEndpoint, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req, &TotalCountResponseData{})
	if err != nil {
		return 0, NewSpyseError(err)
	}

	return *resp.Data.TotalCount, nil
}

// SearchAll returns a list of Certificates that match the specified filters.
func (s *CertificateService) SearchAll(
	ctx context.Context,
	filters []map[string]Filter,
) (items []*Certificate, err error) {
	var from int

	for {
		body, err := json.Marshal(
			SearchRequest{
				SearchParams: filters,
				PaginatedRequest: PaginatedRequest{
					Size: MaxSearchSize,
					From: from,
				},
			},
		)
		if err != nil {
			return nil, err
		}

		req, err := s.client.NewRequest(ctx, http.MethodPost, CertificateSearchEndpoint, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}

		resp, err := s.client.Do(req, &Certificate{})
		if err != nil {
			return nil, NewSpyseError(err)
		}

		if len(resp.Data.Items) > 0 {
			for _, i := range resp.Data.Items {
				items = append(items, i.(*Certificate))
			}
			from += MaxSearchSize
			if from >= MaxTotalItems || len(resp.Data.Items) < MaxSearchSize {
				break
			}
			continue
		}
		break
	}
	return
}
