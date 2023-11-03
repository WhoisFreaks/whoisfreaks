package modal

// DomainInfo represents detailed information about a domain name, including registration details and contacts.
type DomainInfo struct {
	// DomainExt represents the domain extension (e.g., ".com", ".org"). This field is not included in JSON output.
	DomainExt  string `json:"-"`
	StatusInfo string `json:"-"`
	// Status indicates whether the domain information retrieval was successful (true) or not (false).
	Status bool `json:"status,omitempty"`
	// DomainName is the name of the domain.
	DomainName string `json:"domain_name,omitempty"`
	// QueryTime represents the time of the domain information query.
	QueryTime string `json:"query_time,omitempty"`
	// WhoisServer is the WHOIS server used for domain information retrieval.
	WhoisServer string `json:"whois_server,omitempty"`
	// DomainRegistered indicates wether the domain is registered.
	DomainRegistered string `json:"domain_registered,omitempty"`
	// CreateDate represents the creation date of the domain registration.
	CreateDate string `json:"create_date,omitempty"`
	// UpdateDate represents the last update date of the domain registration details.
	UpdateDate string `json:"update_date,omitempty"`
	// ExpiryDate represents the expiration date of the domain registration.
	ExpiryDate string `json:"expiry_date,omitempty"`
	// DomainRegistrar contains information about the domain registrar.
	DomainRegistrar RegistrarInformation `json:"domain_registrar,omitempty"`
	// ResellerContact contains information about the domain reseller contact.
	ResellerContact ResellerContact `json:"reseller_contact,omitempty"`
	// RegistrantContact contains information about the domain registrant contact.
	RegistrantContact PersonalInformation `json:"registrant_contact,omitempty"`
	// AdministrativeContact contains information about the domain administrative contact.
	AdministrativeContact PersonalInformation `json:"administrative_contact,omitempty"`
	// TechnicalContact contains information about the domain technical contact.
	TechnicalContact PersonalInformation `json:"technical_contact,omitempty"`
	// BillingContact contains information about the domain billing contact.
	BillingContact PersonalInformation `json:"billing_contact,omitempty"`
	// NameServers is a list of domain name servers associated with the domain.
	NameServers []string `json:"name_servers,omitempty"`
	// DomainStatus is a list of status values associated with the domain.
	DomainStatus []string `json:"domain_status,omitempty"`
	// RegistryData contains additional registry-specific information about the domain.
	RegistryData RegistryData `json:"registry_data,omitempty"`
}

// PersonalInformation represents personal details associated with a domain contact.
type PersonalInformation struct {
	// Name represents the name of the domain contact.
	Name string `json:"name,omitempty"`
	// Company represents the company name of the domain contact.
	Company string `json:"company,omitempty"`
	// Street represents the street address of the domain contact.
	Street string `json:"street,omitempty"`
	// City represents the city of the domain contact.
	City string `json:"city,omitempty"`
	// State represents the state of the domain contact.
	State string `json:"state,omitempty"`
	// ZipCode represents the ZIP code of the domain contact.
	ZipCode string `json:"zip_code,omitempty"`
	// CountryName represents the country name of the domain contact.
	CountryName string `json:"country_name,omitempty"`
	// CountryCode represents the country code of the domain contact.
	CountryCode string `json:"country_code,omitempty"`
	// EmailAddress represents the email address of the domain contact.
	EmailAddress string `json:"email_address,omitempty"`
	// Phone represents the phone number of the domain contact.
	Phone string `json:"phone,omitempty"`
	// Fax represents the fax number of the domain contact.
	Fax string `json:"fax,omitempty"`
	// MailingAddress represents the mailing address of the domain contact.
	MailingAddress string `json:"mailing_address,omitempty"`
}

// ResellerContact represents contact information for the domain reseller.
type ResellerContact struct {
	// Name represents the name of the domain reseller.
	Name string `json:"name,omitempty"`
	// EmailAddress represents the email address of the domain reseller.
	EmailAddress string `json:"email_address,omitempty"`
	// Phone represents the phone number of the domain reseller.
	Phone string `json:"phone,omitempty"`
}

// RegistryData represents registry-specific information about a domain.
type RegistryData struct {
	// DomainName is the name of the domain.
	DomainName string `json:"domain_name,omitempty"`
	// QueryTime represents the time of the registry data query.
	QueryTime string `json:"query_time,omitempty"`
	// WhoisServer represents the WHOIS server used for the registry data query.
	WhoisServer string `json:"whois_server,omitempty"`
	// DomainRegistered indicates wether the domain was registered.
	DomainRegistered string `json:"domain_registered,omitempty"`
	// CreateDate represents the creation date of the domain registration.
	CreateDate string `json:"create_date,omitempty"`
	// UpdateDate represents the last update date of the domain registration details.
	UpdateDate string `json:"update_date,omitempty"`
	// ExpiryDate represents the expiration date of the domain registration.
	ExpiryDate string `json:"expiry_date,omitempty"`
	// DomainRegistrar contains information about the domain registrar.
	DomainRegistrar DomainRegistrarRegistry `json:"domain_registrar,omitempty"`
	// NameServers is a list of domain name servers associated with the domain.
	NameServers []string `json:"name_servers,omitempty"`
	// DomainStatus is a list of status values associated with the domain.
	DomainStatus []string `json:"domain_status,omitempty"`
}

// DomainRegistrarRegistry represents registry-specific information about the domain registrar.
type DomainRegistrarRegistry struct {
	RegistrarInformation
}

// RegistrarInformation represents general information about the domain registrar.
type RegistrarInformation struct {
	// IanaID represents the IANA ID of the domain registrar.
	IanaID string `json:"iana_id,omitempty"`
	// RegistrarName represents the name of the domain registrar.
	RegistrarName string `json:"registrar_name,omitempty"`
	// WhoisServer represents the WHOIS server of the domain registrar.
	WhoisServer string `json:"whois_server,omitempty"`
	// WebsiteURL represents the website URL of the domain registrar.
	WebsiteURL string `json:"website_url,omitempty"`
	// EmailAddress represents the email address of the domain registrar.
	EmailAddress string `json:"email_address,omitempty"`
	// PhoneNumber represents the phone number of the domain registrar.
	PhoneNumber string `json:"phone_number,omitempty"`
}
