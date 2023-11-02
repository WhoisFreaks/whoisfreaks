package utility

import (
	"fmt"
	"log"
	"os"
)

func DoParameterValidation(whois, dns, domainavailability, ssl, live, historical, reverse bool) (bool, string) {
	apikey, present := os.LookupEnv("WHOISFREAKS_API_KEY")
	if present != true {
		log.Fatal("Are you sure you've set 'WHOISFREAKS_API_KEY' environmental variable? Because it's current value is: " + apikey)
	}

	if !(whois) && !(dns) && !(domainavailability) && !(ssl) {
		fmt.Println("You must use one of the following as per your need;\n" +
			"    -whois to perform a Whois Lookup.\n    -dns to perform a DNS lookup.\n    -domainavailability to perform a Domain Availability lookup.\n" +
			"    -ssl to perform a SSL Lookup")
		return false, apikey
	}

	if whois && !(live || historical || reverse) {
		fmt.Println("For whois query, specify at least one of the following;\n" +
			"    -live to perform live query. \n    -historical to perform historical query.\n    -reverse to perform a reverse query.")
		return false, apikey
	}

	if dns && !(live || historical || reverse) {
		fmt.Println("For dns query, specify at least one of the following;\n" +
			"    -live to perform live query. \n    -historical to perform historical query.\n    -reverse to perform a reverse query.")
		return false, apikey
	}

	if ssl && !live {
		fmt.Println("For ssl query, specify following:\n" +
			"    -live to perform live query")
		return false, apikey
	}
	return true, apikey
}
