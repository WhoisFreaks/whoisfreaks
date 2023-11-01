package modal

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
