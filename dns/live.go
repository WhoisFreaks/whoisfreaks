package dns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// apiKey is a global variable in dns module that stores the API key used for authentication.
var apiKey string

// SetAPIKey sets the global API key to the specified value.
// It allows you to configure the API key used for authentication in your application.
//
// Parameters:
//   - key: A string representing the API key to be set.
//
// Example usage:
//   dns.SetAPIKey("your_api_key_here")
//
func SetAPIKey(key string) {
	apiKey = key
}

// GetLiveResponse performs a live DNS lookup using the WhoisFreaks API.
// It retrieves real-time DNS information for a specific domain and DNS type.
//
// Parameters:
//   - dnsType: The type of DNS record to look up (e.g., "A", "MX", "CNAME"). If you want all records use "all" and if you want multiple records use this "a,aaaa".
//   - domain: The domain name for which live DNS information is requested.
//
// Returns:
//   - *modal.DNSInfo: A pointer to a DNSInfo struct containing live DNS information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//	 dns.SetAPIKey("your_api_key_here")
//   dnsInfo, err := dns.GetLiveResponse("a", "example.com")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Live DNS Info:", dnsInfo)
func GetLiveResponse(dnsType, domain string) (*modal.DNSInfo, *modal.Error) {

	var liveDNSURL = "https://api.whoisfreaks.com/v1.0/dns/live?apiKey=" + apiKey + "&type=" + dnsType + "&domainName=" + domain
	var dnsInfo modal.DNSInfo
	var errorInfo modal.Error

	response, err := http.Get(liveDNSURL)
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

	err = json.NewDecoder(response.Body).Decode(&dnsInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live DNS Information:", err)
		return nil, nil
	}

	return &dnsInfo, nil
}
