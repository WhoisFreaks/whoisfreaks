package ssl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Usama015/whoisfreaks/modal"
)

func GetLiveResponse(domain string, apiKey string, chain bool, raw bool) (*modal.DomainSslInfo, *modal.Error) {

	var liveSslURL = "https://api.whoisfreaks.com/v1.0/ssl/live?apiKey=" + apiKey + "&domainName=" + domain
	var sslInfo modal.DomainSslInfo
	var errorInfo modal.Error

	if chain {
		liveSslURL = liveSslURL + "&chain=true"
	} else if raw {
		liveSslURL = liveSslURL + "&sslRaw=true"
	}

	response, err := http.Get(liveSslURL)
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

	err = json.NewDecoder(response.Body).Decode(&sslInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live Domain Information:", err)
		return nil, nil
	}

	return &sslInfo, nil
}
