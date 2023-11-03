package main

import (
	"flag"
	"os"
	"strings"
	"log"
	"fmt"
	"encoding/json"

	"github.com/WhoisFreaks/whoisfreaks/dns"

	"github.com/WhoisFreaks/whoisfreaks/domainavailability"

	"github.com/WhoisFreaks/whoisfreaks/whois"

	"github.com/WhoisFreaks/whoisfreaks/ssl"
)

func main() {
	whoisPtr := flag.Bool("whois", false, "For specifying a whois query. With this flag you must choose a flag from -live -historical -reverse")
	dnsPtr := flag.Bool("dns", false, "For specifying a dns query. With this flag you must choose a flag from -live -historical -reverse")
	domainAvailabilityPtr := flag.Bool("domainAvailability", false, "For specifying a domain Availability query. With this flag you can also choose the bulk query by mentioning -bulk.")
	sslPtr := flag.Bool("ssl", false, "For specifying a ssl certificate extraction query. With this flag you can choose -live for performing a live query.")

	livePtr := flag.Bool("live", false, "For specifying a live query.")
	historicalPtr := flag.Bool("historical", false, "For specifying a historical query.")
	reversePtr := flag.Bool("reverse", false, "For specifying a reverse query.")

	mini := flag.Bool("mini", false, "Used for getting mini whois Responses that contains fewer fields. It is only useful with -whois -reverse.")
	bulk := flag.Bool("bulk", false, "For specifying a bulk query. With this flag you can use -domains and pass multiple domains with comma Separated policy.")
	sug := flag.Bool("sug", false, "For activating the suggestions during performing domain Availability query")
	chain := flag.Bool("chain", false, "For specifying chain certificate extraction. This flag is compatible with -ssl -live.")
	raw := flag.Bool("raw", false, "For specifying raw certificate extraction. This flag is compatible with -ssl -live.")
	count := flag.String("count", "", "for specifying the number of suggestions, useful with historical query when suggestions are activated.")
	domain := flag.String("domain", "", "Domain Name for which you need to get whois Information. Only for live and historical query.")
	domains := flag.String("domains", "", "Domain Names for bulk query. Only for live Whois and Domain Availability query. Must be used with -bulk flag.")
	company_name := flag.String("company", "", "Company Name for which you need to get whois Information. Only with reverse whois query.")
	keyword := flag.String("keyword", "", "Keyword for which you need to get whois Information. Only with reverse whois query.")
	owner_name := flag.String("owner", "", "Owner Name for which you need to get whois Information. Only with reverse whois query.")
	email := flag.String("email", "", "Email for which you need to get whois Information. Only with reverse whois query.")
	dnsType := flag.String("dnsType", "", "Dns Record type for which you need to get DNS Information. use 'all' for retrieving all types of records. Must be used with dns query of any type.")
	value := flag.String("value", "", "Value of record type for which you need to get DNS Information. e.g. 8.8.8.8 for ARecord. Must be used with reverse dns query.")
	page := flag.String("page", "", "Required Page Number. Each page contains 100 records. Useful with reverse and historical queries when there are a lot of records.")

	flag.Parse()

	if flag.NFlag() == 0 {
		printStarter()
	}

	validated, apiKey := doParameterValidation(*whoisPtr, *dnsPtr, *domainAvailabilityPtr, *sslPtr, *livePtr, *historicalPtr, *reversePtr)
	if !validated {
		os.Exit(1)
	}

	// For Sending Query to the requested servers.
	if *whoisPtr {
		if *livePtr {
			if *bulk {
				domainInfo, errorInfo := whois.GetBulkLiveResponse(strings.Split(*domains, ","), apiKey)
				if errorInfo != nil {
					printError(errorInfo)
					return
				}
				printInfo(domainInfo)
			} else {
				domainInfo, errorInfo := whois.GetLiveResponse(*domain, apiKey)
				if errorInfo != nil {
					printError(errorInfo)
					return
				}
				printInfo(domainInfo)
			}
		} else if *historicalPtr {
			domainInfo, errorInfo := whois.GetHistoricalResponse(*domain, apiKey)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(domainInfo)
		} else if *reversePtr {
			if *mini {
				domainInfo, errorInfo := whois.GetReverseMiniResponse(*keyword, *email, *company_name, *owner_name, apiKey, *page)
				if errorInfo != nil {
					printError(errorInfo)
					return
				}
				printInfo(domainInfo)
			} else {
				domainInfo, errorInfo := whois.GetReverseResponse(*keyword, *email, *company_name, *owner_name, apiKey, *page)
				if errorInfo != nil {
					printError(errorInfo)
					return
				}
				printInfo(domainInfo)
			}
		}
	} else if *dnsPtr {
		if *livePtr {
			dnsInfo, errorInfo := dns.GetLiveResponse(*dnsType, *domain, apiKey)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(dnsInfo)
		} else if *historicalPtr {
			historicalDnsInfo, errorInfo := dns.GetHistoricalResponse(*dnsType, *domain, *page, apiKey)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(historicalDnsInfo)
		} else if *reversePtr {
			reverseDnsInfo, errorInfo := dns.GetReverseResponse(*dnsType, *value, *page, apiKey)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(reverseDnsInfo)
		}
	} else if *domainAvailabilityPtr {
		if *bulk {
			bulkDomainavailabilityInfo, errorInfo := domainavailability.Bulk(strings.Split(*domains, ","), apiKey)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(bulkDomainavailabilityInfo)
		} else if *sug {
			domainavailabilityInfo, errorInfo := domainavailability.CheckAndSuggest(*domain, apiKey, *sug, *count)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(domainavailabilityInfo)
		} else {
			bulkDomainavailabilityInfo, errorInfo := domainavailability.Check(*domain, apiKey)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(bulkDomainavailabilityInfo)
		}
	} else if *sslPtr {
		if *livePtr {
			sslInfo, errorInfo := ssl.GetLiveResponse(*domain, apiKey, *chain, *raw)
			if errorInfo != nil {
				printError(errorInfo)
				return
			}
			printInfo(sslInfo)

		}
	}

}

func doParameterValidation(whois, dns, domainavailability, ssl, live, historical, reverse bool) (bool, string) {
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


func printError(errorInfo interface{}) {
	errorJSON, err := json.MarshalIndent(errorInfo, "", "    ")
	if err != nil {
		log.Fatal("Error while Marshaling the Error Object:", err)
		return
	}
	fmt.Println(string(errorJSON))
}

func printInfo(infoObject interface{}) {
	infoJSON, err := json.MarshalIndent(infoObject, "", "    ")
	if err != nil {
		log.Fatal("Error while Marshaling the Object:", err)
		return
	}
	fmt.Println(string(infoJSON))
}

func printStarter() {
	fmt.Println("Use the following flags according to your requirements to fully utilize WhoisFreaks' services.")
	fmt.Println()
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("	-%s    :    %s\n", f.Name, f.Usage)
	})
	fmt.Println()
	fmt.Println("For Detailed Usage you can consult with README.md.")
	os.Exit(0)
}
