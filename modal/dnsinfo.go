package modal

type DNSInfo struct {
	DomainName string         `json:"domainName,omitempty"`
	QueryTime  string         `json:"queryTime,omitempty"`
	DnsTypes   map[string]int `json:"dnsTypes,omitempty"`
	DnsRecords []DNSRecord    `json:"dnsRecords,omitempty"`
}

type DNSRecord struct {
	Name       string   `json:"name,omitempty"`
	Type       int      `json:"type,omitempty"`
	DNSType    string   `json:"dnsType,omitempty"`
	TTL        int      `json:"ttl,omitempty"`
	RawText    string   `json:"rawText,omitempty"`
	RRSetType  int      `json:"rRsetType,omitempty"`
	Target     string   `json:"target,omitempty"`
	Priority   int      `json:"priority,omitempty"`
	Address    string   `json:"address,omitempty"`
	SingleName string   `json:"singleName,omitempty"`
	Admin      string   `json:"admin,omitempty"`
	Host       string   `json:"host,omitempty"`
	Expire     int      `json:"expire,omitempty"`
	Minimum    int      `json:"minimum,omitempty"`
	Refresh    int      `json:"refresh,omitempty"`
	Retry      int      `json:"retry,omitempty"`
	Serial     int      `json:"serial,omitempty"`
	Strings    []string `json:"strings,omitempty"`
}
