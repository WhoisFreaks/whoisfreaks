package modal

// HistoricalDnsInfo represents historical DNS information including pagination details.
type HistoricalDnsInfo struct {
	// TotalPages represents the total number of pages available in the historical DNS records.
	TotalPages int `json:"totalPages,omitempty"`
	// CurrentPage represents the current page number of the historical DNS records.
	CurrentPage int `json:"currentPage,omitempty"`
	// TotalRecords represents the total number of historical DNS records available.
	TotalRecords int `json:"totalRecords,omitempty"`
	// HistoricalDns represents a slice of DNSInfo containing historical DNS records.
	HistoricalDns []DNSInfo `json:"historicalDnsRecords,omitempty"`
}
