package modal

type ResellerContact struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	Phone        string `json:"phone,omitempty"`
}
