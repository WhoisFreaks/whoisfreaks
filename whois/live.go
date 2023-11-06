package whois

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// apiKey is a global variable in whois module that stores the API key used for authentication.
var apiKey string

// SetAPIKey sets the global API key to the specified value.
// It allows you to configure the API key used for authentication in your application.
//
// Example usage:
//   whois.SetAPIKey("your_api_key_here")
//
// Parameters:
//   - key: A string representing the API key to be set.
func SetAPIKey(key string) {
	apiKey = key
}

// GetLiveResponse retrieves live WHOIS information for a specific domain using the WhoisFreaks API.
//
// Parameters:
//   - domain: The domain name for which live WHOIS information is requested (e.g., "example.com").
//
// Returns:
//   - *modal.DomainInfo: A pointer to a DomainInfo struct containing live domain information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   domainInfo, err := whois.GetLiveResponse("example.com")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Live Domain Info:", domainInfo)
func GetLiveResponse(domain string) (*modal.DomainInfo, *modal.Error) {

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

// GetBulkLiveResponse retrieves live WHOIS information for multiple domains in bulk using the WhoisFreaks API.
//
// Parameters:
//   - domains: A slice of strings containing the domain names for which live WHOIS information is requested.
//
// Returns:
//   - *modal.BulkDomainInfo: A pointer to a BulkDomainInfo struct containing live domain information for the bulk request.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   domains := []string{"example1.com", "example2.com", "example3.com"}
//   bulkDomainInfo, err := GetBulkLiveResponse(domains)
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Bulk Live Domain Info:", bulkDomainInfo)
func GetBulkLiveResponse(domains []string) (*modal.BulkDomainInfo, *modal.Error) {
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
