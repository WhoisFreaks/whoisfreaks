package modal

type RegistrarInformation struct {
	IanaID        string `json:"iana_id,omitempty"`
	RegistrarName string `json:"registrar_name,omitempty"`
	WhoisServer   string `json:"whois_server,omitempty"`
	WebsiteURL    string `json:"website_url,omitempty"`
	EmailAddress  string `json:"email_address,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
}
