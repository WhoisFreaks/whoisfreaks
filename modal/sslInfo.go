package modal

type DomainSslInfo struct {
	DomainName      string            `json:"domainName,omitempty"`
	QueryTime       string            `json:"queryTime,omitempty"`
	SslCertificates []CertificateInfo `json:"sslCertificates,omitempty"`
	SslRaw          string            `json:"sslRaw,omitempty"`
}

type CertificateInfo struct {
	ChainOrder         string         `json:"chainOrder,omitempty"`
	AuthenticationType string         `json:"authenticationType,omitempty"`
	ValidityStartDate  string         `json:"validityStartDate,omitempty"`
	ValidityEndDate    string         `json:"validityEndDate,omitempty"`
	SerialNumber       string         `json:"serialNumber,omitempty"`
	SignatureAlgorithm string         `json:"signatureAlgorithm,omitempty"`
	Subject            UnitInfo       `json:"subject,omitempty"`
	Issuer             UnitInfo       `json:"issuer,omitempty"`
	PublicKey          PublicKeyInfo  `json:"publicKey,omitempty"`
	Extensions         ExtensionsInfo `json:"extensions,omitempty"`
	PemRaw             string         `json:"pemRaw,omitempty"`
}

type PublicKeyInfo struct {
	KeySize      string `json:"keySize,omitempty"`
	KeyAlgorithm string `json:"keyAlgorithm,omitempty"`
	PemRaw       string `json:"pemRaw,omitempty"`
}

type ExtensionsInfo struct {
	AuthorityKeyIdentifier  string                    `json:"authorityKeyIdentifier,omitempty"`
	SubjectKeyIdentifier    string                    `json:"subjectKeyIdentifier,omitempty"`
	KeyUsages               []string                  `json:"keyUsages,omitempty"`
	ExtendedKeyUsages       []string                  `json:"extendedKeyUsages,omitempty"`
	CrlDistributionPoints   []string                  `json:"crlDistributionPoints,omitempty"`
	AuthorityInfoAccess     AuthorityInfo             `json:"authorityInfoAccess,omitempty"`
	SubjectAlternativeNames AlternateNamesInfo        `json:"subjectAlternativeNames,omitempty"`
	CertificatePolicies     []CertificatePoliciesInfo `json:"certificatePolicies,omitempty"`
}

type CertificatePoliciesInfo struct {
	PolicyID        string              `json:"policyId,omitempty"`
	PolicyQualifier PolicyQualifierInfo `json:"policyQualifier,omitempty"`
}

type PolicyQualifierInfo struct {
	Oid        string         `json:"oid,omitempty"`
	CpsUri     string         `json:"cpsUri,omitempty"`
	UserNotice UserNoticeInfo `json:"userNotice,omitempty"`
}

type UserNoticeInfo struct {
	ExplicitText string        `json:"explicitText,omitempty"`
	NoticeRef    NoticeRefInfo `json:"noticeRef,omitempty"`
}

type NoticeRefInfo struct {
	Organization  string `json:"organization,omitempty"`
	NoticeNumbers string `json:"noticeNumbers,omitempty"`
}

type UnitInfo struct {
	CommonName         string `json:"commonName,omitempty"`
	Organization       string `json:"organization,omitempty"`
	OrganizationalUnit string `json:"organizationalUnit,omitempty"`
	Locality           string `json:"locality,omitempty"`
	State              string `json:"state,omitempty"`
	Country            string `json:"country,omitempty"`
	IncCountry         string `json:"incCountry,omitempty"`
	IncState           string `json:"incState,omitempty"`
	BusinessCategory   string `json:"businessCategory,omitempty"`
	Street             string `json:"street,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	SerialNumber       string `json:"serialNumber,omitempty"`
}

type AlternateNamesInfo struct {
	DnsNames       []string `json:"dnsNames,omitempty"`
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	IpAddresses    []string `json:"ipAddresses,omitempty"`
	Uris           []string `json:"uris,omitempty"`
}

type AuthorityInfo struct {
	Issuers []string `json:"issuers,omitempty"`
	Ocsp    []string `json:"ocsp,omitempty"`
}
