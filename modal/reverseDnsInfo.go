package modal

type ReverseDnsInfo struct {
	TotalPages   int       `json:"totalPages,omitempty"`
	CurrentPage  int       `json:"currentPage,omitempty"`
	TotalRecords int       `json:"totalRecords,omitempty"`
	ReverseDns   []DNSInfo `json:"reverseDnsRecords,omitempty"`
}
