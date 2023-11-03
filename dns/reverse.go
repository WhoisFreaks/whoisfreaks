package dns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// GetReverseResponse performs a reverse DNS lookup using the WhoisFreaks API.
// It retrieves DNS information for a specific IP address or value.
//
// Parameters:
//   - dnsType: The type of DNS record to look up (e.g., "a").
//   - value: The IP address or value for which reverse DNS information is requested.
//   - page: The optional page number for paginated results. Leave empty for the first page.
//   - apiKey: The API key for authenticating the request with the WhoisFreaks API.
//
// Returns:
//   - *modal.ReverseDnsInfo: A pointer to a ReverseDnsInfo struct containing reverse DNS information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   reverseDnsInfo, err := dns.GetReverseResponse("a", "8.8.8.8", "1", "your_api_key")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Reverse DNS Info:", reverseDnsInfo)

//   reverseDnsInfo, err := dns.GetReverseResponse("mx", "mx.zoho.com", "1", "your_api_key")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Reverse DNS Info:", reverseDnsInfo)
func GetReverseResponse(dnsType, value, page, apiKey string) (*modal.ReverseDnsInfo, *modal.Error) {

	var historicalDNSURL = "https://api.whoisfreaks.com/v1.0/dns/reverse?apiKey=" + apiKey + "&type=" + dnsType + "&value=" + value
	var reverseDnsInfo modal.ReverseDnsInfo
	var errorInfo modal.Error

	if page != "" {
		historicalDNSURL = historicalDNSURL + "&page=" + page
	}

	response, err := http.Get(historicalDNSURL)
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

	err = json.NewDecoder(response.Body).Decode(&reverseDnsInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live DNS Information:", err)
		return nil, nil
	}

	return &reverseDnsInfo, nil
}
