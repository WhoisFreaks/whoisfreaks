package modal

// DomainSslInfo represents SSL certificate information for a domain including domain name, query time,
type DomainSslInfo struct {
	// DomainName represents the domain name for which SSL certificate information is retrieved.
	DomainName string `json:"domainName,omitempty"`

	// QueryTime represents the time of the SSL certificate information query.
	QueryTime string `json:"queryTime,omitempty"`

	// SslCertificates is a slice of CertificateInfo representing SSL certificates associated with the domain.
	SslCertificates []CertificateInfo `json:"sslCertificates,omitempty"`

	// SslRaw contains raw SSL information in PEM format.
	SslRaw string `json:"sslRaw,omitempty"`
}

// CertificateInfo represents detailed information about an SSL certificate including chain order, authentication type,
// validity dates, serial number, signature algorithm, subject, issuer, public key, extensions, and raw PEM information.
type CertificateInfo struct {
	// ChainOrder represents the order of the certificate in the certificate chain.
	ChainOrder string `json:"chainOrder,omitempty"`

	// AuthenticationType represents the type of authentication used in the certificate.
	AuthenticationType string `json:"authenticationType,omitempty"`

	// ValidityStartDate represents the start date of the certificate's validity period.
	ValidityStartDate string `json:"validityStartDate,omitempty"`

	// ValidityEndDate represents the end date of the certificate's validity period.
	ValidityEndDate string `json:"validityEndDate,omitempty"`

	// SerialNumber represents the serial number of the certificate.
	SerialNumber string `json:"serialNumber,omitempty"`

	// SignatureAlgorithm represents the algorithm used for the certificate's digital signature.
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty"`

	// Subject represents the subject details of the certificate.
	Subject UnitInfo `json:"subject,omitempty"`

	// Issuer represents the issuer details of the certificate.
	Issuer UnitInfo `json:"issuer,omitempty"`

	// PublicKey represents the public key information of the certificate.
	PublicKey PublicKeyInfo `json:"publicKey,omitempty"`

	// Extensions represents various extensions associated with the certificate.
	Extensions ExtensionsInfo `json:"extensions,omitempty"`

	// PemRaw contains the raw PEM-encoded information of the certificate.
	PemRaw string `json:"pemRaw,omitempty"`
}

// PublicKeyInfo represents the public key information including key size, key algorithm, and raw PEM information.
type PublicKeyInfo struct {
	// KeySize represents the size of the public key in bits.
	KeySize string `json:"keySize,omitempty"`

	// KeyAlgorithm represents the algorithm used for the public key.
	KeyAlgorithm string `json:"keyAlgorithm,omitempty"`

	// PemRaw contains the raw PEM-encoded public key information.
	PemRaw string `json:"pemRaw,omitempty"`
}

// ExtensionsInfo represents various extensions associated with an SSL certificate including authority key identifier,
// subject key identifier, key usages, extended key usages, CRL distribution points, authority info access, subject alternative names,
// and certificate policies.
type ExtensionsInfo struct {
	// AuthorityKeyIdentifier represents the authority key identifier extension value.
	AuthorityKeyIdentifier string `json:"authorityKeyIdentifier,omitempty"`

	// SubjectKeyIdentifier represents the subject key identifier extension value.
	SubjectKeyIdentifier string `json:"subjectKeyIdentifier,omitempty"`

	// KeyUsages represents the key usages extension values.
	KeyUsages []string `json:"keyUsages,omitempty"`

	// ExtendedKeyUsages represents the extended key usages extension values.
	ExtendedKeyUsages []string `json:"extendedKeyUsages,omitempty"`

	// CrlDistributionPoints represents the CRL distribution points extension values.
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty"`

	// AuthorityInfoAccess represents the authority information access extension values.
	AuthorityInfoAccess AuthorityInfo `json:"authorityInfoAccess,omitempty"`

	// SubjectAlternativeNames represents the subject alternative names extension values.
	SubjectAlternativeNames AlternateNamesInfo `json:"subjectAlternativeNames,omitempty"`

	// CertificatePolicies represents the certificate policies extension values.
	CertificatePolicies []CertificatePoliciesInfo `json:"certificatePolicies,omitempty"`
}

// CertificatePoliciesInfo represents the certificate policies extension values including policy ID and policy qualifier details.
type CertificatePoliciesInfo struct {
	// PolicyID represents the policy identifier for the certificate policy.
	PolicyID string `json:"policyId,omitempty"`

	// PolicyQualifier represents the policy qualifier information.
	PolicyQualifier PolicyQualifierInfo `json:"policyQualifier,omitempty"`
}

// PolicyQualifierInfo represents the policy qualifier information including OID, CPS URI, and user notice details.
type PolicyQualifierInfo struct {
	// Oid represents the object identifier for the policy qualifier.
	Oid string `json:"oid,omitempty"`

	// CpsUri represents the certification practice statement URI for the policy qualifier.
	CpsUri string `json:"cpsUri,omitempty"`

	// UserNotice represents the user notice information for the policy qualifier.
	UserNotice UserNoticeInfo `json:"userNotice,omitempty"`
}

// UserNoticeInfo represents the user notice information including explicit text and notice reference details.
type UserNoticeInfo struct {
	// ExplicitText represents the explicit text associated with the user notice.
	ExplicitText string `json:"explicitText,omitempty"`

	// NoticeRef represents the notice reference information.
	NoticeRef NoticeRefInfo `json:"noticeRef,omitempty"`
}

// NoticeRefInfo represents the notice reference information including organization and notice numbers.
type NoticeRefInfo struct {
	// Organization represents the organization associated with the notice reference.
	Organization string `json:"organization,omitempty"`

	// NoticeNumbers represents the notice numbers associated with the notice reference.
	NoticeNumbers string `json:"noticeNumbers,omitempty"`
}

// UnitInfo represents the unit information including common name, organization, organizational unit, locality,
// state, country, inclusive country, inclusive state, business category, street, postal code, and serial number.
type UnitInfo struct {
	// CommonName represents the common name associated with the unit.
	CommonName string `json:"commonName,omitempty"`

	// Organization represents the organization name associated with the unit.
	Organization string `json:"organization,omitempty"`

	// OrganizationalUnit represents the organizational unit name associated with the unit.
	OrganizationalUnit string `json:"organizationalUnit,omitempty"`

	// Locality represents the locality associated with the unit.
	Locality string `json:"locality,omitempty"`

	// State represents the state associated with the unit.
	State string `json:"state,omitempty"`

	// Country represents the country associated with the unit.
	Country string `json:"country,omitempty"`

	// IncCountry represents the inclusive country associated with the unit.
	IncCountry string `json:"incCountry,omitempty"`

	// IncState represents the inclusive state associated with the unit.
	IncState string `json:"incState,omitempty"`

	// BusinessCategory represents the business category associated with the unit.
	BusinessCategory string `json:"businessCategory,omitempty"`

	// Street represents the street address associated with the unit.
	Street string `json:"street,omitempty"`

	// PostalCode represents the postal code associated with the unit.
	PostalCode string `json:"postalCode,omitempty"`

	// SerialNumber represents the serial number associated with the unit.
	SerialNumber string `json:"serialNumber,omitempty"`
}

// AlternateNamesInfo represents the alternate names associated with an SSL certificate including DNS names, email addresses,
// IP addresses, and URIs.
type AlternateNamesInfo struct {
	// DnsNames represents the DNS names associated with the alternate names.
	DnsNames []string `json:"dnsNames,omitempty"`

	// EmailAddresses represents the email addresses associated with the alternate names.
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// IpAddresses represents the IP addresses associated with the alternate names.
	IpAddresses []string `json:"ipAddresses,omitempty"`

	// Uris represents the URIs associated with the alternate names.
	Uris []string `json:"uris,omitempty"`
}

// AuthorityInfo represents the authority information associated with an SSL certificate including issuers and OCSP URIs.
type AuthorityInfo struct {
	// Issuers represents the issuers associated with the authority information.
	Issuers []string `json:"issuers,omitempty"`

	// Ocsp represents the OCSP URIs associated with the authority information.
	Ocsp []string `json:"ocsp,omitempty"`
}
