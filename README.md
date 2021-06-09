# Spyse API wrapper for Go

The official wrapper for Spyse API, written in Golang, aimed to help developers build their
integrations with Spyse.

[Spyse](https://spyse.com/) is the most complete Internet assets registry for every cybersecurity
professional. Examples of data Spyse delivers:
* List of 300+ most popular open ports found on 3.5 Billion publicly accessible IPv4 hosts.
* Technologies used on 300+ most popular open ports and IP addresses and domains using a particular technology.
* Security score for each IP host and website, calculated based on the found vulnerabilities.
* List of websites hosted on each IPv4 host.
* DNS and WHOIS records of the domain names.
* SSL certificates provided by the website hosts.
* Structured content of the website homepages.
* Abuse reports associated with IPv4 hosts.
* Organizations and industries associated with the domain names.
* Email addresses found during the Internet scanning, associated with a domain name.

More information about the data Spyse collects is available on the [Our data](https://spyse.com/our-data) page. 

Spyse provides an API accessible via **token-based authentication**. 
Access tokens are **available only for registered users** on their [account page](https://spyse.com/user).

For more information about the API, please check the [API Reference](https://spyse-dev.readme.io/reference/quick-start).

## Installation

```bash
go get github.com/spyse-com/go-spyse
```

## Examples

In this section we provide the output for some [examples](./examples) of the library usage.

Before running the examples, set the access token in your environment. 
Access tokens are available to registered users on their account page:
```bash
export ACCESS_TOKEN=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
```

[Autonomous system example](./examples/as/main.go) output:
```bash
$ go run examples/as/main.go --access_token=$ACCESS_TOKEN

Details about ASN 10000
  Nagasaki Cable Media Inc.
  18 IPv4 prefix(es)
  1 IPv6 prefix(es)

Search result by ASN 21000
  Atos Infogerance SAS
  9 IPv4 prefix(es)
  1 IPv6 prefix(es)
```

[Domains example](./examples/domain/main.go) output:
```bash
$ go run examples/domain/main.go --access_token=$ACCESS_TOKEN

Details about tesla.com
  Updated at: 2021-05-14T17:18:29.070387768Z
  Website title: Electric Cars, Solar & Clean Energy | Tesla
  Alexa rank: 1439
  Certificate fingerprint: 75445f6a9bcf9ec...
  Certificate subject org: Tesla, Inc.
  Certificate issuer org: DigiCert Inc

Subdomains (55) :
  - static-assets.tesla.com
  - itanswers.tesla.com
  - api-toolbox.tesla.com
    --- and 52 more ---

Scroll search for domain names with different suffixes (tesla.com)
  Search ID: 646d69747269692...
  Total items: 268
  Results on the first page (100):
  - tesla.ac
  - tesla.africa
  - tesla.agency
    --- and 97 more ---
  Results on the second page (100):
  - tesla.eu
  - tesla.exchange
  - tesla.expert
    --- and 97 more ---
  Results on the third page (68):
  - tesla.pl
  - tesla.place
  - tesla.plus
    --- and 65 more ---
```

[IPs example](./examples/ip/main.go) output:
```bash
$ go run examples/ip/main.go --access_token=$ACCESS_TOKEN

Details about 91.210.36.26
  Updated at: 2021-05-31T20:14:31.965867941Z
  Country: Ukraine
  Abuse confidence score: 0
  Number of abuses: 0
  CIDR: 91.210.36.0/22
  Open ports (9)
  - 587
  - 993
  - 8083
    --- and 6 more ---
  Affected by 0 CVEs
  PTR record: www.idatacenter.com.ua
  Technologies detected (12)
    - Exim smtpd, version 4.84_2 on port 2525
    - UW Imap, version  on port 993
    - Exim smtpd, version 4.84_2 on port 465
      --- and 9 more ---

IPs with open port 9200 (1791024) :
  - 5x.xxx.xxx.xx5 (United States)  // redacted
  - 4x.xxx.xxx.xx5 (China)          // redacted
  - 4x.xxx.xxx.xx8 (China)          // redacted
    --- and 1791021 more ---

Scroll search for IPs in United States
  Search ID: 646d69747269692...
  Total items: 1564359170
  Results on the first page (100):
  - 2.16.33.76
  - 2.16.40.0
  - 2.16.40.1
    --- and 97 more ---
```

[DNS history example](./examples/history/main.go) output:
```bash
$ go run examples/history/main.go --access_token=$ACCESS_TOKEN

DNS history for domain google.com

  A record         First seen  Last seen
  172.217.26.46    2020-09-18  2020-09-18
  172.217.161.174  2020-09-25  2020-09-25
  172.217.194.100  2020-09-26  2020-09-26
  --- and 8 more ---

  NS record        First seen  Last seen
  ns1.google.com   2020-09-18  2021-05-15
  ns2.google.com   2020-09-18  2021-05-15
  ns3.google.com   2020-09-18  2021-05-15
  --- and 1 more ---

  MX record                First seen  Last seen
  alt1.aspmx.l.google.com  2020-09-18  2021-05-15
  alt2.aspmx.l.google.com  2020-09-18  2021-05-15
  alt3.aspmx.l.google.com  2020-09-18  2021-05-15
  --- and 2 more ---
```

[Emails search example](./examples/email/main.go) output:
```bash
$ go run examples/email/main.go --access_token=$ACCESS_TOKEN

All emails seen with domain tesla.com (32)

  Email                              Last seen                          At location              Location type
  muskateer@tesla.com                2020-10-02T19:01:36.668584112Z     www.outfuel.com          site
  Press@tesla.com                    2020-10-08T03:20:15.930655717Z     pr-journal.biz           site
  ServiceManualFeedback@tesla.com    2020-10-02T14:51:04.72981981Z      services.teslamotors.com  site
  elon@tesla.com                     2020-10-04T18:14:47.519434163Z     developers.madkudu.com   site
  elon@tesla.com                     2021-05-14T15:31:11.654391896Z     35.246.121.122           port
  elon@tesla.com                     2021-05-14T15:28:24.674757356Z     35.241.5.200             port
  elon@tesla.com                     2021-06-01T22:29:02.042521863Z     35.224.127.164           port
  elon@tesla.com                     2021-05-15T01:19:40.450940427Z     95.213.235.253           port
  elon@tesla.com                     2021-06-02T07:04:15.669265915Z     162.0.229.79             port
  elon@tesla.com                     2021-06-01T00:29:05.027316565Z     192.241.194.167          port
  elon@tesla.com                     2021-05-31T19:46:47.608593243Z     45.55.110.167            port
  elon@tesla.com                     2021-05-16T22:48:44.161414374Z     154.16.245.28            port
  elon@tesla.com                     2021-05-31T18:08:14.869611061Z     165.227.63.48            port
  elon@tesla.com                     2021-05-31T21:39:19.937729111Z     165.227.193.116          port
  elon@tesla.com                     2021-06-01T19:39:25.015716367Z     37.139.0.24              port
  elon@tesla.com                     2021-06-02T07:22:56.090613983Z     3.122.118.222            port
  ServicePortalFeedback@tesla.com    2020-10-02T14:51:04.72981981Z      services.teslamotors.com  site
  forums@tesla.com                   2020-10-16T02:42:09.446928549Z     4a.5n0hrvl.cn            site
  elon4president@tesla.com           n/a                                n/a                      n/a
  gfb-infra@tesla.com                2021-06-01T18:24:10.64127371Z      52.29.176.153            port
  gfb-infra@tesla.com                2021-06-01T19:17:03.218754564Z     35.158.26.10             port
  gfb-infra@tesla.com                2021-06-02T08:00:05.986366286Z     18.196.255.226           port
  referralprogram@tesla.com          2021-05-16T05:30:02.596746783Z     www.bytenet.com          site
  referralprogram@tesla.com          2021-06-02T12:24:24.1898908Z       35.209.235.127           port
  ceo@tesla.com                      2020-10-01T11:43:01.819847842Z     www.the-schweiger.com    site
  --- and 22 more ---
```

## Covered API endpoints
 
**Autonomous Systems Methods**

- [x] Details: /as/{asn}
- [x] Search: /as/search
- [x] Scroll search: /as/scroll/search
- [x] Search count: /as/search/count

**SSL/TLS Certificates Methods**

- [x] Details: /certificate/{fingerprint_sha256}
- [x] Search: /certificate/search
- [x] Scroll search: /certificate/scroll/search
- [x] Search count: /certificate/search/count

**CVEs Methods**

- [x] Details: /cve/{cve_id}
- [x] Search: /cve/search
- [x] Scroll search: /cve/scroll/search
- [x] Search count: /cve/search/count

**Domains Methods**

- [x] Details: /domain/{domain_name}
- [x] Search: /domain/search
- [x] Scroll search: /domain/scroll/search
- [x] Search count: /domain/search/count

**Emails Methods**

- [x] Details: /email/{email}
- [x] Search: /email/search
- [x] Scroll search: /email/scroll/search
- [x] Search count: /email/search/count

**IPs Methods**

- [x] Details: /ip/{ip_address}
- [x] Search: /ip/search
- [x] Scroll search: /ip/scroll/search
- [x] Search count: /ip/search/count

**Bulk search Methods**

- [x] Domains lookup: /bulk-search/domain
- [x] IPs lookup: /bulk-search/ip

**History Methods**

- [x] DNS history: /history/dns/{dns_type}/{domain_name}
- [x] WHOIS history: /history/domain-whois/{domain_name}

## Testing

Run tests:
```bash
go test $(go list ./... | grep -v /examples/)
```

Run tests and create code coverage report:
```bash
go test $(go list ./... | grep -v /examples/) -race -coverprofile=coverage.txt -covermode=atomic
```

## License

Distributed under the MIT License. See [LICENSE](./LICENSE.md) for more information.

## Contact

For any proposals and questions, please write at [contact@spyse.com](contact@spyse.com).