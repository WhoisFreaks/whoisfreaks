package modal

// BulkDomainInfo represents the response structure for bulk WHOIS information retrieval using WhoisFreaks API.
// It contains a list of DomainInfo structs representing individual domain information for bulk domain queries.
type BulkDomainInfo struct {

	// BulkWhoisResponse contains a list of DomainInfo structs representing individual domain information
	BulkWhoisResponse []DomainInfo `json:"bulk_whois_response"`
}
