// This package is for performing any type of WHOIS lookup, such as live, historical, or reverse lookups.
package whois

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// GetHistoricalResponse retrieves historical WHOIS information using the WhoisFreaks API.
//
// Parameters:
//   - domain: The domain name for which historical WHOIS information is requested (e.g., "example.com").
//   - apiKey: The API key for authenticating the request with the WhoisFreaks API.
//
// Returns:
//   - *modal.HistoricalDomainInfo: A pointer to a HistoricalDomainInfo struct containing historical domain information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   historicalInfo, err := whois.GetHistoricalResponse("example.com", "your_api_key")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Historical Domain Info:", historicalInfo)
func GetHistoricalResponse(domain, apiKey string) (*modal.HistoricalDomainInfo, *modal.Error) {

	var liveWhoisURL = "https://api.whoisfreaks.com/v1.0/whois?apiKey=" + apiKey + "&whois=historical&domainName=" + domain
	var historcalDomainInfo modal.HistoricalDomainInfo
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
			fmt.Println("Error decoding JSON response of Error:", err)
			return nil, nil
		}
		return nil, &errorInfo
	}

	err = json.NewDecoder(response.Body).Decode(&historcalDomainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Historical Domains Information:", err)
		return nil, nil
	}

	return &historcalDomainInfo, nil
}
