package domainavailability

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

func Check(domain, apiKey string) (*modal.DomainAvailability, *modal.Error) {

	var domainAvailabilityURL = "https://api.whoisfreaks.com/v1.0/domain/availability?apiKey=" + apiKey + "&domain=" + domain
	var domainAvailability modal.DomainAvailability
	var errorInfo modal.Error

	response, err := http.Get(domainAvailabilityURL)
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

	err = json.NewDecoder(response.Body).Decode(&domainAvailability)
	if err != nil {
		fmt.Println("Error decoding JSON response of Domain Availability Information:", err)
		return nil, nil
	}

	return &domainAvailability, nil
}

func CheckAndSuggest(domain string, apiKey string, sug bool, count string) (*modal.BulkDomainAvailability, *modal.Error) {

	var domainAvailabilityURL = "https://api.whoisfreaks.com/v1.0/domain/availability?apiKey=" + apiKey + "&domain=" + domain
	var bulkDomainAvailability modal.BulkDomainAvailability
	var errorInfo modal.Error

	if sug {
		domainAvailabilityURL = domainAvailabilityURL + "&sug=true"
	}
	if sug && count != "" {
		domainAvailabilityURL = domainAvailabilityURL + "&count=" + count
	}

	response, err := http.Get(domainAvailabilityURL)
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

	err = json.NewDecoder(response.Body).Decode(&bulkDomainAvailability)
	if err != nil {
		fmt.Println("Error decoding JSON response of Domain Availability Information:", err)
		return nil, nil
	}

	return &bulkDomainAvailability, nil
}

func Bulk(domains []string, apiKey string) (*modal.BulkDomainAvailability, *modal.Error) {

	var domainAvailabilityURL = "https://api.whoisfreaks.com/v1.0/domain/availability?apiKey=" + apiKey
	var bulkDomainAvailability modal.BulkDomainAvailability
	var errorInfo modal.Error

	response, err := http.Post(domainAvailabilityURL, "application/json", bytes.NewBuffer(modal.GetRequestBodyForBulkAvailability(domains)))
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

	err = json.NewDecoder(response.Body).Decode(&bulkDomainAvailability)
	if err != nil {
		fmt.Println("Error decoding JSON response of Domain Availability Information:", err)
		return nil, nil
	}

	return &bulkDomainAvailability, nil
}
