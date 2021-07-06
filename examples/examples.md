# Examples

Here you can see a few examples of wrapper usage.

If you still have questions, please join our [Discord](https://discord.gg/XqaUP8c) server and feel free to ask.


Before running the examples, set the access token in your environment.
Access tokens are available to registered users on their account page:
```bash
export ACCESS_TOKEN=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
```

Then execute example
```bash
go run main.go --access_token=$ACCESS_TOKEN
```

### Get requests limit and current requests amount


### All data about domain

In [this](./domain_details/main.go) example you can collect data about domain.

### Subdomains for domain

In [this](./domain_subdomains/main.go) example you can gather subdomains for domain.

This example uses Search method that provide only first 10000 (can depend on your subscription plan) results.

### Subdomains for domain using scroll functionality

In [this](./domain_subdomains_via_scroll/main.go) example you can gather subdomains for domain.

This example uses ScrollSearch method that can provide all results unlike Search method.


### Search for domains with different suffixes

In [this](./domains_with_different_suffix/main.go) example you can gather domains with different suffixes.


### Search for domains with technology

In [this](./domains_with_technology/main.go) example you can gather domains with technology and technology version by using scroll functionality.


### Bulk search

In [this](./bulk_search/main.go) example you can gather information about a bunch of domains and ips.

### All data about IP

In [this](./ip_details/main.go) example you can collect data about IP.

### IPs with open port

In [this](./ips_with_open_port/main.go) example you can obtain all IPs with port.

### IPs by geolocation

In [this](./ips_by_country/main.go) example you can obtain all IPs that located in country.

### All data about AS

In [this](./as_details/main.go) example you can obtain data about AS by ASN.

### DNS history for domain

In [this](./domain_history_dns/main.go) example you can obtain historical DNS records for domain.


### Search for emails for domain

In [this](./emails_search/main.go) example you can obtain all emails by domain.


