package modal

type MiniDomainInfo struct {
	Num         int    `json:"num,omitempty"`
	DomainName  string `json:"domain_name,omitempty"`
	QueryTime   string `json:"query_time,omitempty"`
	CreateDate  string `json:"create_date,omitempty"`
	UpdateDate  string `json:"update_date,omitempty"`
	ExpiryDate  string `json:"expiry_date,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
}
