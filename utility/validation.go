package utility

import (
	"fmt"
	"log"
	"os"
)

func DoParameterValidation(whois, dns, domainavailability, live, historical, reverse bool) (bool, string) {
	apikey, present := os.LookupEnv("WHOISFREAKS_API_KEY")
	if present != true {
		log.Fatal("Are you sure you've set 'WHOISFREAKS_API_KEY' environmental variable? Because it's current value is: " + apikey)
	}

	if !(whois) && !(dns) && !(domainavailability) {
		fmt.Println("Error: Either -whois or -dns flag must be specified and set to true.")
		return false, apikey
	}

	if whois && !(live || historical || reverse) {
		fmt.Println("Error: For whois query, specify at least one of -live, -historical, or -reverse and set it to true.")
		return false, apikey
	}

	if dns && !(live || historical || reverse) {
		fmt.Println("Error: For dns query, specify at least one of -live, -historical, or -reverse and set it to true.")
		return false, apikey
	}
	return true, apikey
}
