package dns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

func GetReverseResponse(dnsType, value, page, apiKey string) (*modal.ReverseDnsInfo, *modal.Error) {

	var historicalDNSURL = "https://api.whoisfreaks.com/v1.0/dns/reverse?apiKey=" + apiKey + "&type=" + dnsType + "&value=" + value
	var reverseDnsInfo modal.ReverseDnsInfo
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

	err = json.NewDecoder(response.Body).Decode(&reverseDnsInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live DNS Information:", err)
		return nil, nil
	}

	return &reverseDnsInfo, nil
}
