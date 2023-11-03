package dns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

func GetLiveResponse(dnsType, domain, apiKey string) (*modal.DNSInfo, *modal.Error) {

	var liveDNSURL = "https://api.whoisfreaks.com/v1.0/dns/live?apiKey=" + apiKey + "&type=" + dnsType + "&domainName=" + domain
	var dnsInfo modal.DNSInfo
	var errorInfo modal.Error

	response, err := http.Get(liveDNSURL)
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

	err = json.NewDecoder(response.Body).Decode(&dnsInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live DNS Information:", err)
		return nil, nil
	}

	return &dnsInfo, nil
}
