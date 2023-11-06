// This package is for performing any type of Domain Availability lookup.
package domainavailability

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// apiKey is a global variable in domainavailability module that stores the API key used for authentication.
var apiKey string

// SetAPIKey sets the global API key to the specified value.
// It allows you to configure the API key used for authentication in your application.
//
// Example usage:
//   domainavailability.SetAPIKey("your_api_key_here")
//
// Parameters:
//   - key: A string representing the API key to be set.
func SetAPIKey(key string) {
	apiKey = key
}

// Check performs a domain availability check using the WhoisFreaks API.
// It checks whether a specific domain name is available for registration.
//
// Parameters:
//   - domain: The domain name to be checked for availability (e.g., "example.com").
//
// Returns:
//   - *modal.DomainAvailability: A pointer to a DomainAvailability struct containing domain availability information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   domainAvailability, err := domainavailability.Check("example.com")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Domain Availability Info:", domainAvailability)
func Check(domain string) (*modal.DomainAvailability, *modal.Error) {

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

// CheckAndSuggest performs a domain availability check and suggests similar domain names
// using the WhoisFreaks API.
//
// Parameters:
//   - domain: The domain name to be checked for availability (e.g., "example.com").
//   - sug: A boolean flag indicating whether to suggest similar domain names.
//   - count: The number of suggested domain names to retrieve (valid only if sug is true).
//
// Returns:
//   - *modal.BulkDomainAvailability: A pointer to a BulkDomainAvailability struct containing bulk domain availability information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   bulkDomainAvailability, err := domainavailability.CheckAndSuggest("example.com", true, "5")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Bulk Domain Availability Info:", bulkDomainAvailability)
func CheckAndSuggest(domain string, sug bool, count string) (*modal.BulkDomainAvailability, *modal.Error) {

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

// Bulk performs a bulk domain availability check using the WhoisFreaks API.
// It checks the availability of multiple domain names in bulk.
//
// Parameters:
//   - domains: A slice of strings containing the domain names to be checked for availability.
//
// Returns:
//   - *modal.BulkDomainAvailability: A pointer to a BulkDomainAvailability struct containing bulk domain availability information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   domains := []string{"example1.com", "example2.com", "example3.com"}
//   bulkDomainAvailability, err := domainavailability.Bulk(domains)
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Bulk Domain Availability Info:", bulkDomainAvailability)
func Bulk(domains []string) (*modal.BulkDomainAvailability, *modal.Error) {

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
