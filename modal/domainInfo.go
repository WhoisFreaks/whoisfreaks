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
