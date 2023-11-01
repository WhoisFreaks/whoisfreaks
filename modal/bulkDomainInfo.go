package modal

type BulkDomainInfo struct {
	BulkWhoisResponse []DomainInfo `json:"bulk_whois_response"`
}
