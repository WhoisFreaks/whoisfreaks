package modal

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
