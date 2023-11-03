package modal

// ReverseMiniDomainInfo represents reverse domain information in a compact format including pagination details.
type ReverseMiniDomainInfo struct {
	// TotalResult represents the total number of results available in the reverse domain query.
	TotalResult int `json:"total_result"`
	// TotalPages represents the total number of pages available in the reverse domain query.
	TotalPages int `json:"total_pages"`
	// CurrentPage represents the current page number of the reverse domain query results.
	CurrentPage int `json:"current_page"`
	// HistoricalData is a slice of MiniDomainInfo representing historical domain information in a minimal format.

	HistoricalData []MiniDomainInfo `json:"whois_domains_historical"`
}
