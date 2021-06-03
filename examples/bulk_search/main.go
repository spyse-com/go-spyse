package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	spyse "github.com/spyse-com/go-spyse/pkg"
)

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	var apiBaseUrl = "https://api.spyse.com/v4/data/"

	client, _ := spyse.NewClient(apiBaseUrl, *accessToken, nil)

	var domainNames = []string{"example.com", "tesla.com", "google.com", "some-nonexistent-domain.io"}
	domainSearchResults, err := client.BulkSearch.Domain(context.Background(), domainNames)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println()
	println(fmt.Sprintf("Domain bulk search results"))
	println(fmt.Sprintf("  Sent domains (%d):", len(domainNames)))
	for i := 0; i < len(domainNames); i++ {
		println(fmt.Sprintf("  - %s", domainNames[i]))
	}
	println()
	println(fmt.Sprintf("  Received datasets (%d):", len(domainSearchResults)))
	for i := 0; i < len(domainSearchResults); i++ {
		println(fmt.Sprintf("  - %s", domainSearchResults[i].Name))
		if len(domainSearchResults[i].DNSRecords.A) > 0 {
			println(fmt.Sprintf("    A records: %s", strings.Join(domainSearchResults[i].DNSRecords.A, ", ")))
		}
		if len(domainSearchResults[i].DNSRecords.NS) > 0 {
			println(fmt.Sprintf("    NS records: %s", strings.Join(domainSearchResults[i].DNSRecords.NS, ", ")))
		}
		if len(domainSearchResults[i].DNSRecords.MX) > 0 {
			println(fmt.Sprintf("    MX records: %s", strings.Join(domainSearchResults[i].DNSRecords.MX, ", ")))
		}
		println()
	}

	var ipAddresses = []string{"93.184.216.34", "199.66.11.62", "172.217.4.78"}
	ipSearchResults, err := client.BulkSearch.IP(context.Background(), ipAddresses)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(fmt.Sprintf("IP bulk search results"))
	println(fmt.Sprintf("  Sent IPs (%d):", len(ipAddresses)))
	for i := 0; i < len(ipAddresses); i++ {
		println(fmt.Sprintf("  - %s", ipAddresses[i]))
	}
	println()
	println(fmt.Sprintf("  Received datasets (%d):", len(ipSearchResults)))
	for i := 0; i < len(ipSearchResults); i++ {
		println(fmt.Sprintf("  - %s", ipSearchResults[i].IPAddress))
		if len(ipSearchResults[i].GEOInfo.Country) > 0 {
			println(fmt.Sprintf("    Country: %s", ipSearchResults[i].GEOInfo.Country))
		}
		if len(ipSearchResults[i].PtrRecord.Value) > 0 {
			println(fmt.Sprintf("    PTR record: %s", ipSearchResults[i].PtrRecord.Value))
		}
		println()
	}
}
