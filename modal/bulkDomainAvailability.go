package modal

type BulkDomainAvailability struct {
	DomainAvailabilityWithSuggestion []DomainAvailability `json:"domain_available_response,omitempty"`
	BulkDomainAvailabilityList       []DomainAvailability `json:"bulk_domain_availability_response,omitempty"`
}
