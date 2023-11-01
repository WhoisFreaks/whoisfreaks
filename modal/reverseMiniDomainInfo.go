package modal

type ReverseMiniDomainInfo struct {
	TotalResult    int              `json:"total_result"`
	TotalPages     int              `json:"total_pages"`
	CurrentPage    int              `json:"current_page"`
	HistoricalData []MiniDomainInfo `json:"whois_domains_historical"`
}
