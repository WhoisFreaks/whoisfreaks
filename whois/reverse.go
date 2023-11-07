package whois

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WhoisFreaks/whoisfreaks/modal"
)

// GetReverseMiniResponse performs a reverse whois lookup in mini mode using the WhoisFreaks API.
//
// Parameters:
//   - keyword: The keyword to search for in domain records.
//   - email: The email address to search for in domain records.
//   - company: The company name to search for in domain records.
//   - owner: The owner name to search for in domain records.
//   - page: The optional page number for paginated results. Leave empty for the first page.
//
// Returns:
//   - *modal.ReverseMiniDomainInfo: A pointer to a ReverseMiniDomainInfo struct containing reverse whois information in mini mode.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//	 whois.SetAPIKey("your_api_key_here")
//   reverseMiniDomainInfo, err := whois.GetReverseMiniResponse("example", "email@example.com", "", "", "1")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Reverse Mini Domain Info:", reverseMiniDomainInfo)
func GetReverseMiniResponse(keyword, email, company, owner, page string) (*modal.ReverseMiniDomainInfo, *modal.Error) {

	var reverseWhoisURL = "https://api.whoisfreaks.com/v1.0/whois?apiKey=" + apiKey + "&whois=reverse&mode=mini"
	var reverseMiniDomainInfo modal.ReverseMiniDomainInfo
	var errorInfo modal.Error

	if keyword != "" {
		reverseWhoisURL = reverseWhoisURL + "&keyword=" + keyword
	} else if email != "" {
		reverseWhoisURL = reverseWhoisURL + "&email=" + email
	} else if company != "" {
		reverseWhoisURL = reverseWhoisURL + "&company=" + company
	} else if owner != "" {
		reverseWhoisURL = reverseWhoisURL + "&owner=" + owner
	}
	if page != "" {
		reverseWhoisURL = reverseWhoisURL + "&page=" + page
	}

	response, err := http.Get(reverseWhoisURL)
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

	err = json.NewDecoder(response.Body).Decode(&reverseMiniDomainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live Domain Information:", err)
		return nil, nil
	}

	return &reverseMiniDomainInfo, nil
}

// GetReverseResponse performs a reverse whois lookup using the WhoisFreaks API.
//
// Parameters:
//   - keyword: The keyword to search for in domain records.
//   - email: The email address to search for in domain records.
//   - company: The company name to search for in domain records.
//   - owner: The owner name to search for in domain records.
//   - page: The optional page number for paginated results. Leave empty for the first page.
//
// Returns:
//   - *modal.ReverseDomainInfo: A pointer to a ReverseDomainInfo struct containing reverse whois information.
//   - *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.
//
// Example usage:
//	 whois.SetAPIKey("your_api_key_here")
//   reverseDomainInfo, err := whois.GetReverseResponse("example", "email@example.com", "", "", "1")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Println("Reverse Domain Info:", reverseDomainInfo)
func GetReverseResponse(keyword, email, company, owner, page string) (*modal.ReverseDomainInfo, *modal.Error) {

	var reverseWhoisURL = "https://api.whoisfreaks.com/v1.0/whois?apiKey=" + apiKey + "&whois=reverse"
	var reverseDomainInfo modal.ReverseDomainInfo
	var errorInfo modal.Error

	if keyword != "" {
		reverseWhoisURL = reverseWhoisURL + "&keyword=" + keyword
	} else if email != "" {
		reverseWhoisURL = reverseWhoisURL + "&email=" + email
	} else if company != "" {
		reverseWhoisURL = reverseWhoisURL + "&company=" + company
	} else if owner != "" {
		reverseWhoisURL = reverseWhoisURL + "&owner=" + owner
	}
	if page != "" {
		reverseWhoisURL = reverseWhoisURL + "&page=" + page
	}
	response, err := http.Get(reverseWhoisURL)
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

	err = json.NewDecoder(response.Body).Decode(&reverseDomainInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response of Live Domain Information:", err)
		return nil, nil
	}

	return &reverseDomainInfo, nil
}
