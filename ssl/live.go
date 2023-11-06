// This package is for performing any type of SSL lookup.
package ssl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// apiKey is a global variable in ssl module that stores the API key used for authentication.
var apiKey string

// SetAPIKey sets the global API key to the specified value.
// It allows you to configure the API key used for authentication in your application.
//
// Example usage:
//   ssl.SetAPIKey("your_api_key_here")
//
// Parameters:
//   - key: A string representing the API key to be set.
func SetAPIKey(key string) {
	apiKey = key
}

// GetLiveResponse retrieves live SSL information for a specific domain using the WhoisFreaks API.
//
// Parameters:
//   - domain: The domain name for which live SSL information is requested (e.g., "example.com").
//   - chain: A boolean flag indicating whether to include SSL certificate chain information.
//   - raw: A boolean flag indicating whether to include raw SSL certificate information.
//
// Returns:
//   - *modal.DomainSslInfo: A pointer to a DomainSslInfo struct containing live SSL information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   sslInfo, err := ssl.GetLiveResponse("example.com", true, false)
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Live SSL Info:", sslInfo)
func GetLiveResponse(domain string, chain bool, raw bool) (*modal.DomainSslInfo, *modal.Error) {

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
