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
API tokens are **available only for registered users** on their [account page](https://spyse.com/user).

For more information about the API, please check the [API Reference](https://spyse-dev.readme.io/reference/quick-start).

## Installation

```bash
go get github.com/spyse-com/go-spyse
```

## Quick start

Add import to your project
```go
import spyse "github.com/spyse-com/go-spyse/pkg"
```

Create client using Spyse API base URL `https://api.spyse.com/v4/data/` and your API token (See on [account page](https://spyse.com/user)). 

You can specify HTTP client or pass `nil` for default net.http client.
```go
client, err := spyse.NewClient("https://api.spyse.com/v4/data/", "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX", nil)
```
<sup>It's better to use a secure way to pass your API token to code. Please, do NOT push your token to public repos ;)

Create new service for your case.  For example, domains (more at [examples](./examples/examples.md)):
```go
	svc := spyse.NewDomainService(client)
```

Obtain all data about the domain:
```go
	details, err := svc.Details(context.Background(), "spyse.com")
```

For more see [examples](./examples/examples.md).

## Troubleshooting

If you need help with installation and/or usage, please join our [Discord](https://discord.gg/XqaUP8c) server.

## Errors handling

You can assert default golang error to `spyse.ErrResponse` for more details:
```go
        spyseError := err.(*spyse.ErrResponse)
        // Error message, e.g. "wrong access token provided"
        spyseError.Err.Message
        // Status code for request to API, e.g. 401
        spyseError.Err.Status
        // Error text code, e.g. "unauthorized"
        spyseError.Err.Code
```

You can see all error codes in [pkg/error.go](pkg/error.go)

## Covered API endpoints

**Account**

- [x] Quota: /account/quota
 
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