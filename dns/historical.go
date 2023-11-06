// This package is for performing any type of DNS lookup, such as live, historical, or reverse lookups.
package dns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// GetHistoricalResponse performs a historical DNS lookup using the WhoisFreaks API.
// It retrieves historical DNS information for a specific domain and DNS type.
//
// Parameters:
//   - dnsType: The type of DNS record to look up (e.g., "A", "MX", "CNAME").
//   - domain: The domain name for which historical DNS information is requested.
//   - page: The optional page number for paginated results. Leave empty for the first page.
//
// Returns:
//   - *modal.HistoricalDnsInfo: A pointer to a HistoricalDnsInfo struct containing historical DNS information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//   historicalDnsInfo, err := dns.GetHistoricalResponse("A", "example.com", "1")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Historical DNS Info:", historicalDnsInfo)
func GetHistoricalResponse(dnsType, domain, page string) (*modal.HistoricalDnsInfo, *modal.Error) {

	var historicalDNSURL = "https://api.whoisfreaks.com/v1.0/dns/historical?apiKey=" + apiKey + "&type=" + dnsType + "&domainName=" + domain
	var historicalDnsInfo modal.HistoricalDnsInfo
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

	err = json.NewDecoder(response.Body).Decode(&historicalDnsInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live DNS Information:", err)
		return nil, nil
	}

	return &historicalDnsInfo, nil
}
