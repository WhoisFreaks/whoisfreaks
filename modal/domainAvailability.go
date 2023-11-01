package modal

type DomainAvailability struct {
	Domain             string `json:"domain,omitempty"`
	DomainAvailability bool   `json:"domainAvailability"`
}
