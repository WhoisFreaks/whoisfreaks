// This package is for performing any type of SSL lookup.
package ssl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// GetLiveResponse retrieves live SSL information for a specific domain using the WhoisFreaks API.
//
// Parameters:
//   - domain: The domain name for which live SSL information is requested (e.g., "example.com").
//   - apiKey: The API key for authenticating the request with the WhoisFreaks API.
//   - chain: A boolean flag indicating whether to include SSL certificate chain information.
//   - raw: A boolean flag indicating whether to include raw SSL certificate information.
//
// Returns:
//   - *modal.DomainSslInfo: A pointer to a DomainSslInfo struct containing live SSL information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   sslInfo, err := ssl.GetLiveResponse("example.com", "your_api_key", true, false)
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Live SSL Info:", sslInfo)
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
