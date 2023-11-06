package modal

// ReverseDnsInfo represents reverse DNS information including pagination details.
type ReverseDnsInfo struct {
	// TotalPages represents the total number of pages available in the reverse DNS records.
	TotalPages int `json:"totalPages,omitempty"`
	// CurrentPage represents the current page number of the reverse DNS records.
	CurrentPage int `json:"currentPage,omitempty"`
	// TotalRecords represents the total number of reverse DNS records available.
	TotalRecords int `json:"totalRecords,omitempty"`
	// ReverseDns represents a slice of DNSInfo containing reverse DNS records.
	ReverseDns []DNSInfo `json:"reverseDnsRecords,omitempty"`
}
