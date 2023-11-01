package modal

type DNSInfo struct {
	DomainName string         `json:"domainName,omitempty"`
	QueryTime  string         `json:"queryTime,omitempty"`
	DnsTypes   map[string]int `json:"dnsTypes,omitempty"`
	DnsRecords []DNSRecord    `json:"dnsRecords,omitempty"`
}
