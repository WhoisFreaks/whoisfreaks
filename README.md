A tool in Go Lang to utilize [Whoisfreaks](https://whoisfreaks.com/)' services like [whois](https://whoisfreaks.com/products/whois-api.html), [DNS](https://whoisfreaks.com/products/dns-checker-api.html), [SSL](https://whoisfreaks.com/products/ssl-certificate-api.html) and [Domain Availability](https://whoisfreaks.com/products/domain-availability-api.html) api.

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

### Authentication
For utilizing whoisfreaks' services you must [sign up](https://whoisfreaks.com/signup.html) and get a API_KEY and then add that key to your machine which will be used for authentication.
```
export WHOISFREAKS_API_KEY="..."
```

### Whois Lookup
With whoisfreaks's Whois tool you can perform;
* [Live whois lookup](https://whoisfreaks.com/tools/whois/lookup) for fetching whois information from their respective servers.
* [Historical whois lookup](https://whoisfreaks.com/tools/whois/history/lookup) for fetching historical data from whoisfreaks' database.
* [Reverse whois lookup](https://whoisfreaks.com/tools/whois/reverse/search) to fetch data from database based on keyword, company name, owner name or email.

##### Live whois Lookup
For live whois lookup you can use;

```
$ whoisfreaks -whois -live -domain whoisfreaks.com
```

to perform bulk live lookup you can use following code where you will provide multiple domains that are comma separated.

```
whoisfreaks -whois -live -bulk -domains whoisfreaks.com,ipgeolocation.io
```
Note: Maximum 100 domains will be served if you want to perform bulk lookup with more then 100 domains you can visit [BULK WHOIS LOOKUP](https://whoisfreaks.com/products/bulk-whois-lookup.html).

##### Historical whois Lookup
For historical whois lookup you can use;

```
$ whoisfreaks -whois -historical -domain whoisfreaks.com
```

##### Reverse whois Lookup
For reverse whois lookup based on keyword you can use;

```
$ whoisfreaks -whois -reverse -keyword whoisfreaks
```

For reverse whois lookup based on owner name you can use;

```
$ whoisfreaks -whois -reverse -owner ejaz_ahmed
```

For reverse whois lookup based on company name you can use;

```
$ whoisfreaks -whois -reverse -company jfreak
```

For reverse whois lookup based on email you can use;

```
$ whoisfreaks -whois -reverse -email ejaz_ahmed@outlook.com
```

with any of the above reverse queries you can use some additional flags like;
* ``-mini`` to get a mini whois response that contains create_date, update_date, expiry_date and few more.
* ``-page n`` to get control over pages if there are more then 100 records as each page contains 100 records. By default you will get first page.


### Domain Availability Lookup

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


### DNS Lookup
With whoisfreaks's DNS tool you can perform;
* [Live DNS lookup](https://tools.whoisfreaks.com/dnslive) for fetching dns information from their respective servers.
* [Historical DNS lookup](https://tools.whoisfreaks.com/dnshistorical) for fetching historical data from whoisfreaks' database.
* [Reverse DNS lookup](https://tools.whoisfreaks.com/dnsreverse) to fetch data from database based on values like `8.8.8.8` for reverse A records and `mx.zoho.com` for reverse mx records.

##### Live DNS Lookup
For live DNS lookup you can use following;

```
$ whoisfreaks -dns -live -domain whoisfreaks.com -dnsType all
```
* use ``-dnsType a`` for A records, if you want any other record use name of that record. If you want multiple use a comma separated list like `` -dnsType a,mx,ns``

##### Historical DNS Lookup
For historical DNS lookup you can use;

```
$ whoisfreaks -dns -historical -domain whoisfreaks.com -dnsType all
```
* use ``-dnsType a`` for A records, if you want any other record use name of that record. If you want multiple use a comma separated list like ``-dnsType a,mx,ns``
  
##### Reverse DNS Lookup
For reverse DNS lookup based on specific dns Record you can use;

```
$ whoisfreaks -dns -reverse -value 8.8.8.8 -dnsType a
```

* ``-page n`` to get control over pages if there are more then 100 records as each page contains 100 records. By default you will get first page.


### SSL Lookup
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
