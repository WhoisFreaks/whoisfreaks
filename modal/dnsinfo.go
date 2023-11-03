package modal

// DNSInfo represents the DNS information for a specific domain.
type DNSInfo struct {
	// DomainName is the domain name for which DNS information is retrieved.
	DomainName string `json:"domainName,omitempty"`
	// QueryTime represents the time of the DNS query.
	QueryTime string `json:"queryTime,omitempty"`
	// DnsTypes is a map containing DNS record types (e.g., "A", "MX") as keys and their counts as values.
	DnsTypes map[string]int `json:"dnsTypes,omitempty"`
	// DnsRecords is a slice containing individual DNSRecord structs representing DNS records associated with the domain.
	DnsRecords []DNSRecord `json:"dnsRecords,omitempty"`
}

type DNSRecord struct {
	// Name is the name of the domain.
	Name string `json:"name,omitempty"`
	// Type is the numeric representation of the DNS record type.
	Type int `json:"type,omitempty"`
	// DNSType is the string representation of the DNS record type (e.g., "A", "MX").
	DNSType string `json:"dnsType,omitempty"`
	// TTL (Time To Live) is the duration that the DNS record is cached.
	TTL int `json:"ttl,omitempty"`
	// RawText is the raw text representation of the DNS record.
	RawText string `json:"rawText,omitempty"`
	// RRSetType is the numeric representation of the RRSet type.
	RRSetType int `json:"rRsetType,omitempty"`
	// Target is the target address associated with the DNS record (e.g., IP address, mail server address)
	Target string `json:"target,omitempty"`
	// Priority is the priority of the DNS record (used in MX records).
	Priority int `json:"priority,omitempty"`
	// Address represents an IP address associated with the DNS record.
	Address string `json:"address,omitempty"`
	// SingleName is a single name associated with the DNS record.
	SingleName string `json:"singleName,omitempty"`
	// Admin is the administrative contact associated with the DNS record.
	Admin string `json:"admin,omitempty"`
	// Host is the host information associated with the DNS record.
	Host string `json:"host,omitempty"`
	// Expire represents the time when the DNS record expires.
	Expire int `json:"expire,omitempty"`

	Minimum int `json:"minimum,omitempty"`
	Refresh int `json:"refresh,omitempty"`
	Retry   int `json:"retry,omitempty"`
	Serial  int `json:"serial,omitempty"`
	// Strings is a slice containing additional strings associated with the DNS record.
	Strings []string `json:"strings,omitempty"`
}
