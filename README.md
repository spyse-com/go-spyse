# go-spyse

The Spyse API uses token-based authentication, which means that developers must pass their API token parameter with every request.

Registered users can find the API token on the [account page](https://spyse.com/user).

### Usage

### Implemented REST API

#### Autonomous Systems Methods
- [x] Details: /as/{asn}
- [x] Search: /as/search
- [x] Search count: /as/search/count

#### SSL/TLS Certificates Methods
- [x] Details: /certificate/{fingerprint_sha256}
- [x] Search: /certificate/search
- [x] Search count: /certificate/search/count

#### CVEs Methods
- [x] Details: /cve/{cve_id}
- [x] Search: /cve/search
- [x] Search count: /cve/search/count

#### Domains Methods
- [x] Details: /domain/{domain_name}
- [x] Search: /domain/search
- [x] Search count: /domain/search/count

#### Emails Methods
- [x] Details: /email/{email}
- [x] Search: /email/search
- [x] Search count: /email/search/count

#### IPs Methods
- [x] Details: /ip/{ip_address}
- [x] Search: /ip/search
- [x] Search count: /ip/search/count

### Links
* [spyse.com](https://spyse.com)
* [API Documentation](https://spyse-dev.readme.io/reference/quick-start)