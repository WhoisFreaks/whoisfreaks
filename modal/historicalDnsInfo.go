package modal

type HistoricalDnsInfo struct {
	TotalPages    int       `json:"totalPages,omitempty"`
	CurrentPage   int       `json:"currentPage,omitempty"`
	TotalRecords  int       `json:"totalRecords,omitempty"`
	HistoricalDns []DNSInfo `json:"historicalDnsRecords,omitempty"`
}
