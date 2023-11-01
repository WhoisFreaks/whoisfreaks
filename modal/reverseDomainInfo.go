package modal

type ReverseDomainInfo struct {
	TotalResult      int             `json:"Total_Result"`
	TotalPages       int             `json:"Total_Pages"`
	CurrentPage      int             `json:"Current_Page"`
	WhoisDomainsHist []DomainInfoApi `json:"whois_domains_historical"`
}

type DomainInfoApi struct {
	Num    int  `json:"num,omitempty"`    // Use pointers to allow nil values
	Status bool `json:"status,omitempty"` // Marked with omitempty to exclude if false
	DomainInfo
}
