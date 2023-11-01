package modal

import (
	"encoding/json"
	"fmt"
)

type RequestBodyForBulkAvailability struct {
	DomainNames []string `json:"domainNames,omitempty"`
}

func GetRequestBodyForBulkAvailability(domains []string) []byte {
	requestBody := RequestBodyForBulkAvailability{
		DomainNames: domains,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil
	}

	return jsonBody
}
