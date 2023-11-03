package modal

// ReverseDomainInfo represents reverse domain information including pagination details.
type ReverseDomainInfo struct {
	// TotalResult represents the total number of results available in the reverse domain query.
	TotalResult int `json:"Total_Result"`
	// TotalPages represents the total number of pages available in the reverse domain query.
	TotalPages int `json:"Total_Pages"`
	// CurrentPage represents the current page number of the reverse domain query results.
	CurrentPage int `json:"Current_Page"`
	// WhoisDomainsHist is a slice of DomainInfoApi representing historical domain information.
	WhoisDomainsHist []DomainInfoApi `json:"whois_domains_historical"`
}

// DomainInfoApi represents domain information for the API response including a numeric identifier (Num),
type DomainInfoApi struct {
	// Num represents the numeric identifier for the domain.
	Num int `json:"num,omitempty"`
	// Status indicates the status of the domain. If false, the domain is marked as inactive.
	Status bool `json:"status,omitempty"`
	// DomainInfo represents detailed domain information including WHOIS data, registrar details, and more.
	DomainInfo
}
