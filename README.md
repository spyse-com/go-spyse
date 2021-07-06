# Spyse API wrapper for Go

The official wrapper for [spyse.com](https://spyse.com/) API, written in Golang, aimed to help developers build their
integrations with Spyse.

[Spyse](https://spyse.com/) is the most complete Internet assets search engine for every cybersecurity
professional.

Examples of data Spyse delivers:

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
```go
// Add import
import "github.com/spyse-com/go-spyse/pkg"
// ...

// Use your API key to init the client
client, err := spyse.NewClient("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX", nil)

// Create a new service suitable for your case
svc := spyse.NewDomainService(client)

// Fetch all information about the domain
details, err := svc.Details(context.Background(), "tesla.com")
// ...
```

## Examples

Account:
- [Check your API quotas](./examples/account/main.go)

Target info:
- [IP details](./examples/ip_details/main.go)
- [Domain details](./examples/domain_details/main.go)
- [Autonomous System details](./examples/as_details/main.go)

Search with params (up to 10 000 results):
- [Subdomains lookup](./examples/domain_subdomains/main.go)
- [Search domains by technology](./examples/domains_with_technology/main.go)
- [Search emails by domain name](./examples/emails_search/main.go)
- [Search IPv4 hosts with specific open port](./examples/ips_with_open_port/main.go)
- [Search IPv4 hosts by geolocation](./examples/ips_by_country/main.go)

Scroll search (unlimited results):
- [Subdomains lookup using scroll API](./examples/domain_subdomains_via_scroll/main.go)
- [Search domains by with different suffixes](examples/domains_with_different_suffix_via_scroll/main.go)


Historical records:
- [Domain DNS history](./examples/domain_history_dns/main.go)

Bulk Search:
- [Bulk search domains / IPv4 hosts](./examples/bulk_search/main.go)

Note: You need to pass access_token as an argument to run any example:
```bash
go run ./examples/domain_details/main.go --access_token=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
```


## Errors handling
To properly handle Spyse errors, assert them to `spyse.ErrResponse` and then check the "Code" field.
```go
spyseError := err.(*spyse.ErrResponse)
// Error message, e.g. "wrong access token provided"
println(spyseError.Err.Message)
// Status code for request to API, e.g. 401
println(spyseError.Err.Status)
// Error text code, e.g. "unauthorized"
println(spyseError.Err.Code)

// Check for "limit reached" error
if spyseError.Err.Code == spyse.CodeRequestsLimitReached {
// ...
}
```

A list of error codes can be found in [pkg/error.go](pkg/error.go)

## Testing

Run tests:
```bash
go test $(go list ./... | grep -v /examples/)
```

Run tests and create code coverage report:
```bash
go test $(go list ./... | grep -v /examples/) -race -coverprofile=coverage.txt -covermode=atomic
```

## Covered API endpoints

All the available API methods are fully supported.


## License

Distributed under the MIT License. See [LICENSE](./LICENSE.md) for more information.

## Troubleshooting and contacts

For any proposals and questions, please write at:

- Email: [contact@spyse.com](contact@spyse.com)
- Discord: [channel](https://discord.gg/XqaUP8c)
- Twitter: [@scanpatch](https://twitter.com/scanpatch), [@MrMristov](https://twitter.com/MrMristov)
