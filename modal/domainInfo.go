package modal

type DomainInfo struct {
	DomainExt             string               `json:"-"`
	StatusInfo            string               `json:"-"`
	Status                bool                 `json:"status,omitempty"`
	DomainName            string               `json:"domain_name,omitempty"`
	QueryTime             string               `json:"query_time,omitempty"`
	WhoisServer           string               `json:"whois_server,omitempty"`
	DomainRegistered      string               `json:"domain_registered,omitempty"`
	CreateDate            string               `json:"create_date,omitempty"`
	UpdateDate            string               `json:"update_date,omitempty"`
	ExpiryDate            string               `json:"expiry_date,omitempty"`
	DomainRegistrar       RegistrarInformation `json:"domain_registrar,omitempty"`
	ResellerContact       ResellerContact      `json:"reseller_contact,omitempty"`
	RegistrantContact     PersonalInformation  `json:"registrant_contact,omitempty"`
	AdministrativeContact PersonalInformation  `json:"administrative_contact,omitempty"`
	TechnicalContact      PersonalInformation  `json:"technical_contact,omitempty"`
	BillingContact        PersonalInformation  `json:"billing_contact,omitempty"`
	NameServers           []string             `json:"name_servers,omitempty"`
	DomainStatus          []string             `json:"domain_status,omitempty"`
	RegistryData          RegistryData         `json:"registry_data,omitempty"`
}

type PersonalInformation struct {
	Name           string `json:"name,omitempty"`
	Company        string `json:"company,omitempty"`
	Street         string `json:"street,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	CountryName    string `json:"country_name,omitempty"`
	CountryCode    string `json:"country_code,omitempty"`
	EmailAddress   string `json:"email_address,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Fax            string `json:"fax,omitempty"`
	MailingAddress string `json:"mailing_address,omitempty"`
}

type ResellerContact struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	Phone        string `json:"phone,omitempty"`
}

type RegistryData struct {
	DomainName       string                  `json:"domain_name,omitempty"`
	QueryTime        string                  `json:"query_time,omitempty"`
	WhoisServer      string                  `json:"whois_server,omitempty"`
	DomainRegistered string                  `json:"domain_registered,omitempty"`
	CreateDate       string                  `json:"create_date,omitempty"`
	UpdateDate       string                  `json:"update_date,omitempty"`
	ExpiryDate       string                  `json:"expiry_date,omitempty"`
	DomainRegistrar  DomainRegistrarRegistry `json:"domain_registrar,omitempty"`
	NameServers      []string                `json:"name_servers,omitempty"`
	DomainStatus     []string                `json:"domain_status,omitempty"`
}

type DomainRegistrarRegistry struct {
	RegistrarInformation
}

type RegistrarInformation struct {
	IanaID        string `json:"iana_id,omitempty"`
	RegistrarName string `json:"registrar_name,omitempty"`
	WhoisServer   string `json:"whois_server,omitempty"`
	WebsiteURL    string `json:"website_url,omitempty"`
	EmailAddress  string `json:"email_address,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
}
