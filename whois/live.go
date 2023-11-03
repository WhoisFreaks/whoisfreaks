package whois

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// GetLiveResponse retrieves live WHOIS information for a specific domain using the WhoisFreaks API.
//
// Parameters:
//   - domain: The domain name for which live WHOIS information is requested (e.g., "example.com").
//   - apiKey: The API key for authenticating the request with the WhoisFreaks API.
//
// Returns:
//   - *modal.DomainInfo: A pointer to a DomainInfo struct containing live domain information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   domainInfo, err := whois.GetLiveResponse("example.com", "your_api_key")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Live Domain Info:", domainInfo)
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

// GetBulkLiveResponse retrieves live WHOIS information for multiple domains in bulk using the WhoisFreaks API.
//
// Parameters:
//   - domains: A slice of strings containing the domain names for which live WHOIS information is requested.
//   - apiKey: The API key for authenticating the request with the WhoisFreaks API.
//
// Returns:
//   - *modal.BulkDomainInfo: A pointer to a BulkDomainInfo struct containing live domain information for the bulk request.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   domains := []string{"example1.com", "example2.com", "example3.com"}
//   bulkDomainInfo, err := GetBulkLiveResponse(domains, "your_api_key")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Bulk Live Domain Info:", bulkDomainInfo)
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
