package modal

type HistoricalDomainInfo struct {
	Status                 bool         `json:"status,omitempty"`
	Whois                  string       `json:"whois,omitempty"`
	TotalRecords           string       `json:"total_records,omitempty"`
	WhoisDomainsHistorical []DomainInfo `json:"whois_domains_historical,omitempty"`
}
