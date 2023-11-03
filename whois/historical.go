package whois

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

func GetHistoricalResponse(domain, apiKey string) (*modal.HistoricalDomainInfo, *modal.Error) {

	var liveWhoisURL = "https://api.whoisfreaks.com/v1.0/whois?apiKey=" + apiKey + "&whois=historical&domainName=" + domain
	var historcalDomainInfo modal.HistoricalDomainInfo
	var errorInfo modal.Error

	response, err := http.Get(liveWhoisURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, nil
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&errorInfo)
		if err != nil {
			fmt.Println("Error decoding JSON response of Error:", err)
			return nil, nil
		}
		return nil, &errorInfo
	}

	err = json.NewDecoder(response.Body).Decode(&historcalDomainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Historical Domains Information:", err)
		return nil, nil
	}

	return &historcalDomainInfo, nil
}
