package modal

// HistoricalDomainInfo represents historical WHOIS information including status, WHOIS data,
type HistoricalDomainInfo struct {
	// Status indicates whether the historical domain information retrieval was successful (true) or not (false).
	Status bool `json:"status,omitempty"`
	// Whois represents the type of WHOIS query.
	Whois string `json:"whois,omitempty"`
	// TotalRecords represents the total number of historical domain records available.
	TotalRecords string `json:"total_records,omitempty"`
	// WhoisDomainsHistorical is a slice of DomainInfo containing historical domain details.
	WhoisDomainsHistorical []DomainInfo `json:"whois_domains_historical,omitempty"`
}
