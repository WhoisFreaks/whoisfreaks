package main

import (
	"flag"
	"os"
	"strings"
	"whoisfreaks/dns"
	domainavailability "whoisfreaks/domainAvailability"
	"whoisfreaks/utility"
	"whoisfreaks/whois"
)

func main() {
	whoisPtr := flag.Bool("whois", false, "for specifying a whois query")
	dnsPtr := flag.Bool("dns", false, "for specifying a dns query")
	domainAvailabilityPtr := flag.Bool("domainAvailability", false, "for specifying a domain Availability query")

	livePtr := flag.Bool("live", false, "for specifying a live query of either whois or dns service.")
	historicalPtr := flag.Bool("historical", false, "for specifying a historical query of either whois or dns service.")
	reversePtr := flag.Bool("reverse", false, "for specifying a reverse query of either whois or dns service.")

	mini := flag.Bool("mini", false, "used for returning mini whois Responses that contains fewer fields. It is only useful with reverse whois query.")
	bulk := flag.Bool("bulk", false, "for specifying a bulk query of whois service.")
	sug := flag.Bool("sug", false, "for activating the suggestions.")
	count := flag.String("count", "", "for specifying the number of suggestions.")

	var domain string
	flag.StringVar(&domain, "domain", "", "Domain Name for which you need to get whois Information. Only for live and historical whois query.")

	var domains string
	flag.StringVar(&domains, "domains", "", "Domain Names for bulk query. Only for live Whois and Domain Availability query.")

	var company_name string
	flag.StringVar(&company_name, "company", "", "Company Name for which you need to get whois Information. only with reverse whois query.")

	var keyword string
	flag.StringVar(&keyword, "keyword", "", "Keyword for which you need to get whois Information. only with reverse whois query.")

	var owner_name string
	flag.StringVar(&owner_name, "owner", "", "Owner Name for which you need to get whois Information. only with reverse whois query.")

	var email string
	flag.StringVar(&email, "email", "", "Email for which you need to get whois Information. only with reverse whois query.")

	var dnsType string
	flag.StringVar(&dnsType, "dnsType", "", "Dns Record type for which you need to get DNS Information. use 'all' for retrieving all types of records.")

	var value string
	flag.StringVar(&value, "value", "", "Value of record type for which you need to get DNS Information. e.g. 8.8.8.8 for ARecord")

	var page string
	flag.StringVar(&page, "page", "", "Required Page Number. Each page contains 100 records.")

	flag.Parse()

	validated, apiKey := utility.DoParameterValidation(*whoisPtr, *dnsPtr, *domainAvailabilityPtr, *livePtr, *historicalPtr, *reversePtr)
	if !validated {
		os.Exit(1)
	}

	// For Sending Query to the requested servers.
	if *whoisPtr {
		if *livePtr {
			if *bulk {
				domainInfo, errorInfo := whois.GetBulkLiveResponse(strings.Split(domains, ","), apiKey)
				if errorInfo != nil {
					utility.PrintError(errorInfo)
					return
				}
				utility.PrintInfo(domainInfo)
			} else {
				domainInfo, errorInfo := whois.GetLiveResponse(domain, apiKey)
				if errorInfo != nil {
					utility.PrintError(errorInfo)
					return
				}
				utility.PrintInfo(domainInfo)
			}
		} else if *historicalPtr {
			domainInfo, errorInfo := whois.GetHistoricalResponse(domain, apiKey)
			if errorInfo != nil {
				utility.PrintError(errorInfo)
				return
			}
			utility.PrintInfo(domainInfo)
		} else if *reversePtr {
			if *mini {
				domainInfo, errorInfo := whois.GetReverseMiniResponse(keyword, email, company_name, owner_name, apiKey, page)
				if errorInfo != nil {
					utility.PrintError(errorInfo)
					return
				}
				utility.PrintInfo(domainInfo)
			} else {
				domainInfo, errorInfo := whois.GetReverseResponse(keyword, email, company_name, owner_name, apiKey, page)
				if errorInfo != nil {
					utility.PrintError(errorInfo)
					return
				}
				utility.PrintInfo(domainInfo)
			}
		} else {
		}
	} else if *dnsPtr {
		if *livePtr {
			dnsInfo, errorInfo := dns.GetLiveResponse(dnsType, domain, apiKey)
			if errorInfo != nil {
				utility.PrintError(errorInfo)
				return
			}
			utility.PrintInfo(dnsInfo)
		} else if *historicalPtr {
			historicalDnsInfo, errorInfo := dns.GetHistoricalResponse(dnsType, domain, page, apiKey)
			if errorInfo != nil {
				utility.PrintError(errorInfo)
				return
			}
			utility.PrintInfo(historicalDnsInfo)
		} else if *reversePtr {
			reverseDnsInfo, errorInfo := dns.GetReverseResponse(dnsType, value, page, apiKey)
			if errorInfo != nil {
				utility.PrintError(errorInfo)
				return
			}
			utility.PrintInfo(reverseDnsInfo)
		} else {

		}
	} else if *domainAvailabilityPtr {
		if *bulk {
			bulkDomainavailabilityInfo, errorInfo := domainavailability.Bulk(strings.Split(domains, ","), apiKey)
			if errorInfo != nil {
				utility.PrintError(errorInfo)
				return
			}
			utility.PrintInfo(bulkDomainavailabilityInfo)
		} else {
			if *sug {
				domainavailabilityInfo, errorInfo := domainavailability.CheckAndSuggest(domain, apiKey, *sug, *count)
				if errorInfo != nil {
					utility.PrintError(errorInfo)
					return
				}
				utility.PrintInfo(domainavailabilityInfo)
			} else {
				bulkDomainavailabilityInfo, errorInfo := domainavailability.Check(domain, apiKey)
				if errorInfo != nil {
					utility.PrintError(errorInfo)
					return
				}
				utility.PrintInfo(bulkDomainavailabilityInfo)
			}
		}
	}

}
