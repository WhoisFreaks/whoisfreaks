package modal

// MiniDomainInfo represents a minimal set of domain information including numeric identifier (Num),
type MiniDomainInfo struct {
	// Num is the numeric identifier for the domain.
	Num int `json:"num,omitempty"`
	// DomainName is the name of the domain.
	DomainName string `json:"domain_name,omitempty"`
	// QueryTime represents the time of the domain information query.
	QueryTime string `json:"query_time,omitempty"`
	// CreateDate represents the creation date of the domain registration.
	CreateDate string `json:"create_date,omitempty"`
	// UpdateDate represents the last update date of the domain registration details.
	UpdateDate string `json:"update_date,omitempty"`
	// ExpiryDate represents the expiration date of the domain registration.
	ExpiryDate string `json:"expiry_date,omitempty"`
	// Name represents the name of the registrant associated with the domain.
	Name string `json:"name,omitempty"`
	// Email represents the email address of the registrant associated with the domain.
	Email string `json:"email,omitempty"`
	// CompanyName represents the company name of the registrant associated with the domain.
	CompanyName string `json:"company_name,omitempty"`
}
