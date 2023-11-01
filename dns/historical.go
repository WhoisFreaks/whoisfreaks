package dns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"whoisfreaks/modal"
)

func GetHistoricalResponse(dnsType, domain, page, apiKey string) (*modal.HistoricalDnsInfo, *modal.Error) {

	var historicalDNSURL = "https://api.whoisfreaks.com/v1.0/dns/historical?apiKey=" + apiKey + "&type=" + dnsType + "&domainName=" + domain
	var historicalDnsInfo modal.HistoricalDnsInfo
	var errorInfo modal.Error

	if page != "" {
		historicalDNSURL = historicalDNSURL + "&page=" + page
	}

	response, err := http.Get(historicalDNSURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, nil
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&errorInfo)
		if err != nil {
			fmt.Println("Error decoding JSON response of Error Information:", err)
			return nil, nil
		}
		return nil, &errorInfo
	}

	err = json.NewDecoder(response.Body).Decode(&historicalDnsInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live DNS Information:", err)
		return nil, nil
	}

	return &historicalDnsInfo, nil
}
