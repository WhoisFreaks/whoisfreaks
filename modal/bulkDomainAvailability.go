// This package contains Modals or types that will be useful for analyzing response from different available services.

package modal

// BulkDomainAvailability represents the response structure for bulk domain availability checks.
// It contains information about domain availability with suggestions and bulk domain availability responses.
type BulkDomainAvailability struct {
	// DomainAvailabilityWithSuggestion contains a list of DomainAvailability structs representing individual domain availability.
	DomainAvailabilityWithSuggestion []DomainAvailability `json:"domain_available_response,omitempty"`

	// BulkDomainAvailabilityList contains a list of DomainAvailability structs representing bulk domain availability responses.
	BulkDomainAvailabilityList []DomainAvailability `json:"bulk_domain_availability_response,omitempty"`
}
