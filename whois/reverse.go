package whois

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Usama015/whoisfreaks/modal"
)

func GetReverseMiniResponse(keyword, email, company, owner, apiKey, page string) (*modal.ReverseMiniDomainInfo, *modal.Error) {

	var reverseWhoisURL = "https://api.whoisfreaks.com/v1.0/whois?apiKey=" + apiKey + "&whois=reverse&mode=mini"
	var reverseMiniDomainInfo modal.ReverseMiniDomainInfo
	var errorInfo modal.Error

	if keyword != "" {
		reverseWhoisURL = reverseWhoisURL + "&keyword=" + keyword
	} else if email != "" {
		reverseWhoisURL = reverseWhoisURL + "&email=" + email
	} else if company != "" {
		reverseWhoisURL = reverseWhoisURL + "&company=" + company
	} else if owner != "" {
		reverseWhoisURL = reverseWhoisURL + "&owner=" + owner
	}
	if page != "" {
		reverseWhoisURL = reverseWhoisURL + "&page=" + page
	}

	response, err := http.Get(reverseWhoisURL)
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

	err = json.NewDecoder(response.Body).Decode(&reverseMiniDomainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live Domain Information:", err)
		return nil, nil
	}

	return &reverseMiniDomainInfo, nil
}

func GetReverseResponse(keyword, email, company, owner, apiKey, page string) (*modal.ReverseDomainInfo, *modal.Error) {

	var reverseWhoisURL = "https://api.whoisfreaks.com/v1.0/whois?apiKey=" + apiKey + "&whois=reverse"
	var reverseDomainInfo modal.ReverseDomainInfo
	var errorInfo modal.Error

	if keyword != "" {
		reverseWhoisURL = reverseWhoisURL + "&keyword=" + keyword
	} else if email != "" {
		reverseWhoisURL = reverseWhoisURL + "&email=" + email
	} else if company != "" {
		reverseWhoisURL = reverseWhoisURL + "&company=" + company
	} else if owner != "" {
		reverseWhoisURL = reverseWhoisURL + "&owner=" + owner
	}
	if page != "" {
		reverseWhoisURL = reverseWhoisURL + "&page=" + page
	}
	response, err := http.Get(reverseWhoisURL)
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

	err = json.NewDecoder(response.Body).Decode(&reverseDomainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live Domain Information:", err)
		return nil, nil
	}

	return &reverseDomainInfo, nil
}
