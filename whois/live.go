package whois

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"whoisfreaks/modal"
)

func GetLiveResponse(domain, apiKey string) (*modal.DomainInfo, *modal.Error) {

	var liveWhoisURL = "https://api.whoisfreaks.com/v1.0/whois?apiKey=" + apiKey + "&whois=live&domainName=" + domain
	var domainInfo modal.DomainInfo
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
			fmt.Println("Error decoding JSON response of Error Information:", err)
			return nil, nil
		}
		return nil, &errorInfo
	}

	err = json.NewDecoder(response.Body).Decode(&domainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live Domain Information:", err)
		return nil, nil
	}

	return &domainInfo, nil
}

func GetBulkLiveResponse(domains []string, apiKey string) (*modal.BulkDomainInfo, *modal.Error) {
	var bulkWhoisURL = "https://api.whoisfreaks.com/v1.0/bulkwhois?apiKey=" + apiKey
	var bulkDomainInfo modal.BulkDomainInfo
	var errorInfo modal.Error

	response, err := http.Post(bulkWhoisURL, "application/json", bytes.NewBuffer(modal.GetRequestBodyForBulkAvailability(domains)))
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

	err = json.NewDecoder(response.Body).Decode(&bulkDomainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live Domain Information:", err)
		return nil, nil
	}

	return &bulkDomainInfo, nil
}
