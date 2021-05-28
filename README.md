# go-spyse-v4

[![codecov](https://codecov.io/gh/spyse-com/go-spyse/branch/master/graph/badge.svg)](https://codecov.io/gh/spyse-com/go-spyse)

The official client for Spyse API v4, written in Golang.

Spyse ([https://spyse.com/](https://spyse.com/)) provides a token-based authentication API. Access tokens are available to registered users on their [account page](https://spyse.com/user).

For more information please check the [API Reference](https://spyse-dev.readme.io/reference/quick-start).

## Installation

```bash
go get github.com/spyse-com/go-spyse-v4
```

## Usage

```golang

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
go test ./...
```

Run tests and create code coverage report:
```bash
go test ./... -race -coverprofile=coverage.txt -covermode=atomic
```

## License

Distributed under the MIT License. See [LICENSE](./LICENSE.md) for more information.


## Contact

For any proposals and questions, please write at [contact@spyse.com](contact@spyse.com).