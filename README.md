A tool in Go Lang to utilize [Whoisfreaks](https://whoisfreaks.com/)' services like [whois](https://whoisfreaks.com/products/whois-api.html), [DNS](https://whoisfreaks.com/products/dns-checker-api.html), [SSL](https://whoisfreaks.com/products/ssl-certificate-api.html) and [Domain Availability](https://whoisfreaks.com/products/domain-availability-api.html) api.

### Index
*   [Installation](#installation) 
*   [Command line interface](#command-line-interface-cli)
    -   [Authentication](#authentication)  
    -   [Whois lookup](#whois-lookup-using-cli)  
    -   [Domain availability lookup](#domain-availability-lookup-using-cli)  
    -   [DNS lookup](#dns-lookup-using-cli)  
    -   [SSL lookup](#ssl-lookup-using-cli) 
*   [Software development kit](#software-development-kit-sdk)
    -   [Whois lookup](#whois-lookup)  
    -   [Domain availability lookup](#domain-availability-lookup)  
    -   [DNS lookup](#dns-lookup)  
    -   [SSL lookup](#ssl-lookup)  

### INSTALLATION
For Installation go must be installed in your machine;

```
go get -u github.com/WhoisFreaks/whoisfreaks
```
OR
```
go install github.com/WhoisFreaks/whoisfreaks
```
##### using source code
if you want to build on your own you can follow this;
```
git clone https://github.com/WhoisFreaks/whoisfreaks.git
```
go to the cloned directory and build the project using;
```
go build
```
and move the binary to the your machine's go path.
Note: After Installation you can use following guide to perform any sort of lookup.

You can use the services using command line as well as SDK.

## COMMAND LINE INTERFACE (CLI)


### Authentication
For utilizing whoisfreaks' services you must [sign up](https://whoisfreaks.com/signup.html) and get a API_KEY and then add that key to your machine which will be used for authentication.
```
export WHOISFREAKS_API_KEY="..."
```

### Whois Lookup using CLI
With whoisfreaks's Whois tool you can perform;
* [Live whois lookup](https://whoisfreaks.com/tools/whois/lookup) for fetching whois information from their respective servers.
* [Historical whois lookup](https://whoisfreaks.com/tools/whois/history/lookup) for fetching historical data from whoisfreaks' database.
* [Reverse whois lookup](https://whoisfreaks.com/tools/whois/reverse/search) to fetch data from database based on keyword, company name, owner name or email.

##### Live whois Lookup
For live whois lookup you can use;

```
whoisfreaks -whois -live -domain whoisfreaks.com
```

to perform bulk live lookup you can use following code where you will provide multiple domains that are comma separated.

```
whoisfreaks -whois -live -bulk -domains whoisfreaks.com,ipgeolocation.io
```
Note: Maximum 100 domains will be served if you want to perform bulk lookup with more then 100 domains you can visit [BULK WHOIS LOOKUP](https://whoisfreaks.com/products/bulk-whois-lookup.html).

##### Historical whois Lookup
For historical whois lookup you can use;

```
whoisfreaks -whois -historical -domain whoisfreaks.com
```

##### Reverse whois Lookup
For reverse whois lookup based on keyword you can use;

```
whoisfreaks -whois -reverse -keyword whoisfreaks
```

For reverse whois lookup based on owner name you can use;

```
whoisfreaks -whois -reverse -owner ejaz_ahmed
```

For reverse whois lookup based on company name you can use;

```
whoisfreaks -whois -reverse -company jfreak
```

For reverse whois lookup based on email you can use;

```
whoisfreaks -whois -reverse -email ejaz_ahmed@outlook.com
```

with any of the above reverse queries you can use some additional flags like;
* ``-mini`` to get a mini whois response that contains create_date, update_date, expiry_date and few more.
* ``-page n`` to get control over pages if there are more then 100 records as each page contains 100 records. By default you will get first page.


### Domain Availability Lookup using CLI

With whoisfreaks' domain [availability lookup tool](https://tools.whoisfreaks.com/domainavailability) you can check for any domain if it is available for registration or not. You can use following;
```
whoisfreaks -domainAvailability -domain whoisfreaks.com
```  
* use ``-sug`` to get default five suggestions.
* use ``-count`` if you want any number of suggestions but greater then five.

similarly, to perform bulk Domain availability lookup you can use;
```
whoisfreaks -domainAvailability -bulk -domains whoisfreaks.com,ipgeolocation.io
```
Note: suggestions flag is only available with single domain availability lookup.


### DNS Lookup using CLI
With whoisfreaks's DNS tool you can perform;
* [Live DNS lookup](https://tools.whoisfreaks.com/dnslive) for fetching dns information from their respective servers.
* [Historical DNS lookup](https://tools.whoisfreaks.com/dnshistorical) for fetching historical data from whoisfreaks' database.
* [Reverse DNS lookup](https://tools.whoisfreaks.com/dnsreverse) to fetch data from database based on values like `8.8.8.8` for reverse A records and `mx.zoho.com` for reverse mx records.

##### Live DNS Lookup
For live DNS lookup you can use following;

```
whoisfreaks -dns -live -domain whoisfreaks.com -dnsType all
```
* use ``-dnsType a`` for A records, if you want any other record use name of that record. If you want multiple use a comma separated list like `` -dnsType a,mx,ns``

##### Historical DNS Lookup
For historical DNS lookup you can use;

```
whoisfreaks -dns -historical -domain whoisfreaks.com -dnsType all
```
* use ``-dnsType a`` for A records, if you want any other record use name of that record. If you want multiple use a comma separated list like ``-dnsType a,mx,ns``
  
##### Reverse DNS Lookup
For reverse DNS lookup based on specific dns Record you can use;

```
whoisfreaks -dns -reverse -value 8.8.8.8 -dnsType a
```

* ``-page n`` to get control over pages if there are more then 100 records as each page contains 100 records. By default you will get first page.


### SSL Lookup using CLI
SSL cert lookup tool allows you to fetch Secure Sockets Layer (SSL) certificate along with its complete SSL cert chain.

* [Live SSL Lookup](https://tools.whoisfreaks.com/ssllive) for fetching live and up-to-date information.

##### live SSL Lookup
For live SSL lookup you can use following;
```
whoisfreaks -ssl -live -domain whoisfreaks.com
```
some other flags are also provided like;

* ``-chain`` to get chain of all domain ssl certificates sorted from end-user to root.
* ``-raw`` to get the raw openssl response of the domain.

## SOFTWARE DEVELOPMENT KIT (SDK)

### WHOIS Lookup
This package is for performing any type of WHOIS lookup, such as live, historical, or reverse lookups.

#### Index

[WHOIS Authentication Setup](#whois-authentication-setup)  
[WHOIS Bulk Live Lookup](#whois-bulk-live-lookup)  
[WHOIS Historical Lookup](#whois-historical-lookup)  
[WHOIS Live Lookup](#whois-live-lookup)  
[WHOIS Reverse Lookup](#whois-reverse-lookup)  
[WHOIS Mini Reverse Lookup](#whois-mini-reverse-lookup)  

#### Functions

##### [WHOIS Authentication Setup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/whois/live.go#L23)
```
func SetAPIKey(key string)
```

SetAPIKey sets the global API key to the specified value. It allows you to configure the API key used for authentication in your application.

Example usage: 

```
whois.SetAPIKey("your_api_key_here")
```

Parameters:

*   key: A string representing the API key to be set.


##### [WHOIS Bulk Live Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/whois/live.go#L91)
```
func GetBulkLiveResponse(domains []string) (*modal.BulkDomainInfo, *modal.Error)
```

GetBulkLiveResponse retrieves live WHOIS information for multiple domains in bulk using the WhoisFreaks API.

Parameters:

*   domains: A slice of strings containing the domain names for which live WHOIS information is requested.

Returns:

*   *modal.BulkDomainInfo: A pointer to a BulkDomainInfo struct containing live domain information for the bulk request.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example usage: 

```
domains := []string{"example1.com", "example2.com", "example3.com"}
bulkDomainInfo, err := GetBulkLiveResponse(domains)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Bulk Live Domain Info:", bulkDomainInfo)
```

##### [WHOIS Historical Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/whois/historical.go#L28)
```
func GetHistoricalResponse(domain string) (*modal.HistoricalDomainInfo, *modal.Error)
```

GetHistoricalResponse retrieves historical WHOIS information using the WhoisFreaks API.

Parameters:

*   domain: The domain name for which historical WHOIS information is requested (e.g., "example.com").

Returns:

*   *modal.HistoricalDomainInfo: A pointer to a HistoricalDomainInfo struct containing historical domain information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example usage:

```
historicalInfo, err := whois.GetHistoricalResponse("example.com")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Historical Domain Info:", historicalInfo)
```

##### [WHOIS Live Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/whois/live.go#L43)

```
func GetLiveResponse(domain string) (*modal.DomainInfo, *modal.Error)
```

GetLiveResponse retrieves live WHOIS information for a specific domain using the WhoisFreaks API.

Parameters:

*   domain: The domain name for which live WHOIS information is requested (e.g., "example.com").

Returns:

*   *modal.DomainInfo: A pointer to a DomainInfo struct containing live domain information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example usage:
```
domainInfo, err := whois.GetLiveResponse("example.com")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Live Domain Info:", domainInfo)
```

##### [WHOIS Reverse Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/whois/reverse.go#L95)

```
func GetReverseResponse(keyword, email, company, owner, page string) (*modal.ReverseDomainInfo, *modal.Error)
```

GetReverseResponse performs a reverse whois lookup using the WhoisFreaks API. 

Parameters:

*   keyword: The keyword to search for in domain records.
*   email: The email address to search for in domain records.
*   company: The company name to search for in domain records.
*   owner: The owner name to search for in domain records.
*   page: The optional page number for paginated results. Leave empty for the first page.


Returns:

*   *modal.ReverseDomainInfo: A pointer to a ReverseDomainInfo struct containing reverse whois information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.


Example usage:
```
reverseDomainInfo, err := whois.GetReverseResponse("example", "email@example.com", "", "", "1")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Reverse Domain Info:", reverseDomainInfo)
```


##### [WHOIS Mini Reverse Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/whois/reverse.go#L31)

```
func GetReverseMiniResponse(keyword, email, company, owner, page string) (*modal.ReverseMiniDomainInfo, *modal.Error)
```

GetReverseMiniResponse performs a reverse whois lookup in mini mode using the WhoisFreaks API. 

Parameters:

*   keyword: The keyword to search for in domain records.
*   email: The email address to search for in domain records.
*   company: The company name to search for in domain records.
*   owner: The owner name to search for in domain records.
*   page: The optional page number for paginated results. Leave empty for the first page.



Returns:

*   *modal.ReverseMiniDomainInfo: A pointer to a ReverseMiniDomainInfo struct containing reverse whois information in mini mode.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.



Example usage:
```
reverseMiniDomainInfo, err := whois.GetReverseMiniResponse("example", "email@example.com", "", "", "1")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Reverse Mini Domain Info:", reverseMiniDomainInfo)
```


### Domain Availability Lookup
This package is for performing any type of Domain Availability lookup.  

#### Index

[Domain Availability Authentication Setup](#domain-availability-authentication-setup)  
[Bulk Domain Availability Lookup](#bulk-domain-availability-lookup)  
[Domain Availability Check](#domain-availability-check)  
[Domain Availability Check and Suggestion](#domain-availability-check-and-suggestion)  

#### Functions


##### [Domain Availability Authentication Setup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/domainavailability/live.go#L24)

```
func SetAPIKey(key string)
```

SetAPIKey sets the global API key to the specified value. It allows you to configure the API key used for authentication in your application.

Example usage:
```
domainavailability.SetAPIKey("your_api_key_here")
```

Parameters:

*   key: A string representing the API key to be set.


##### [Bulk Domain Availability Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/domainavailability/live.go#L151)

```
func Bulk(domains []string) (*modal.BulkDomainAvailability, *modal.Error)
```

Bulk performs a bulk domain availability check using the WhoisFreaks API. It checks the availability of multiple domain names in bulk.

Parameters:

*   domains: A slice of strings containing the domain names to be checked for availability.

Returns:

*   *modal.BulkDomainAvailability: A pointer to a BulkDomainAvailability struct containing bulk domain availability information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example usage: 

```
domains := []string{"example1.com", "example2.com", "example3.com"}
bulkDomainAvailability, err := domainavailability.Bulk(domains)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Bulk Domain Availability Info:", bulkDomainAvailability)
```

##### [Domain Availability Check](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/domainavailability/live.go#L45)

```
func Check(domain string) (*modal.DomainAvailability, *modal.Error)
```

Check performs a domain availability check using the WhoisFreaks API. It checks whether a specific domain name is available for registration.

Parameters:

*   domain: The domain name to be checked for availability (e.g., "example.com").

Returns:

*   *modal.DomainAvailability: A pointer to a DomainAvailability struct containing domain availability information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example usage:

```
domainAvailability, err := domainavailability.Check("example.com")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Domain Availability Info:", domainAvailability)
```

##### [Domain Availability Check and Suggestion](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/domainavailability/live.go#L95)
```
func CheckAndSuggest(domain string, sug bool, count string) (*modal.BulkDomainAvailability, *modal.Error)
```

CheckAndSuggest performs a domain availability check and suggests similar domain names using the WhoisFreaks API.

Parameters:

*   domain: The domain name to be checked for availability (e.g., "example.com").
*   sug: A boolean flag indicating whether to suggest similar domain names.
*   count: The number of suggested domain names to retrieve (valid only if sug is true).

Returns:

*   *modal.BulkDomainAvailability: A pointer to a BulkDomainAvailability struct containing bulk domain availability information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example usage:

```
bulkDomainAvailability, err := domainavailability.CheckAndSuggest("example.com", true, "5")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Bulk Domain Availability Info:", bulkDomainAvailability)
```

##### Types

This section is empty.


### DNS Lookup
This package is for performing any type of DNS lookup, such as live, historical, or reverse lookups. 

#### Index

[DNS Authentication Setup](#dns-authentication-setup)  
[DNS Historical Lookup](#dns-historical-lookup)  
[DNS Live Lookup](#dns-live-lookup)  
[DNS Reverse Lookup](#dns-reverse-lookup)  

#### Functions

##### [DNS Authentication Setup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/dns/live.go#L22)
```
func SetAPIKey(key string)
```

SetAPIKey sets the global API key to the specified value. It allows you to configure the API key used for authentication in your application. 

Example usage: 
```
dns.SetAPIKey("your_api_key_here")
```

Parameters:
*   key: A string representing the API key to be set.

##### [DNS Historical Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/dns/historical.go#L31)
```
func GetHistoricalResponse(dnsType, domain, page string) (*modal.HistoricalDnsInfo, *modal.Error)
```
``GetHistoricalResponse`` performs a historical DNS lookup using the WhoisFreaks API. It retrieves historical DNS information for a specific domain and DNS type. 

Parameters:
*   dnsType: The type of DNS record to look up (e.g., "A", "MX", "CNAME").
*   domain: The domain name for which historical DNS information is requested.
*   page: The optional page number for paginated results. Leave empty for the first page.

Returns:
*   *modal.HistoricalDnsInfo: A pointer to a HistoricalDnsInfo struct containing historical DNS information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example Usage:

```
historicalDnsInfo, err := dns.GetHistoricalResponse("A", "example.com", "1")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Historical DNS Info:", historicalDnsInfo)
```

##### [DNS Live Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/dns/live.go#L44)
```
func GetLiveResponse(dnsType, domain string) (*modal.DNSInfo, *modal.Error)
```
``GetLiveResponse`` performs a live DNS lookup using the WhoisFreaks API. It retrieves real-time DNS information for a specific domain and DNS type. 

Parameters:
*   dnsType: The type of DNS record to look up (e.g., "A", "MX", "CNAME").
*   domain: The domain name for which live DNS information is requested.

Returns:
*   *modal.DNSInfo: A pointer to a DNSInfo struct containing live DNS information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example Usage:
```
dnsInfo, err := dns.GetLiveResponse("a", "example.com")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Live DNS Info:", dnsInfo)
```

##### [DNS Reverse Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/dns/reverse.go#L37)
```
func GetReverseResponse(dnsType, value, page string) (*modal.ReverseDnsInfo, *modal.Error)
```

```
reverseDnsInfo, err := dns.GetReverseResponse("mx", "mx.zoho.com", "1")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Reverse DNS Info:", reverseDnsInfo)
```


### SSL Lookup
This package is for performing any type of SSL lookup. 

#### Index

[SSL Authentication Setup](#ssl-authentication-setup)  
[SSL Live Lookup](#ssl-live-lookup)  

#### Functions


##### [SSL Authentication Setup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/ssl/live.go#L23)

```
func SetAPIKey(key string)
```

SetAPIKey sets the global API key to the specified value. It allows you to configure the API key used for authentication in your application.

Example usage:
```
ssl.SetAPIKey("your_api_key_here")
```

Parameters:

*   key: A string representing the API key to be set.


##### [SSL Live Lookup](https://github.com/WhoisFreaks/whoisfreaks/blob/v1.0.2/ssl/live.go#L45)
```
func GetLiveResponse(domain string, chain bool, raw bool) (*modal.DomainSslInfo, *modal.Error)
```

GetLiveResponse retrieves live SSL information for a specific domain using the WhoisFreaks API.

Parameters:

*   domain: The domain name for which live SSL information is requested (e.g., "example.com").
*   chain: A boolean flag indicating whether to include SSL certificate chain information.
*   raw: A boolean flag indicating whether to include raw SSL certificate information.

Returns:

*   *modal.DomainSslInfo: A pointer to a DomainSslInfo struct containing live SSL information.
*   *modal.Error: A pointer to an Error struct if there is an API error, or nil if the request is successful.

Example usage:

```
sslInfo, err := ssl.GetLiveResponse("example.com", true, false)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Live SSL Info:", sslInfo)
```
