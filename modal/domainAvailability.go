package modal

// DomainAvailability represents the availability status of a domain name.
type DomainAvailability struct {
	// Domain is the name of the domain.
	Domain string `json:"domain,omitempty"`
	// DomainAvailability indicates whether the domain is available (true) or not (false).
	DomainAvailability bool `json:"domainAvailability"`
}
