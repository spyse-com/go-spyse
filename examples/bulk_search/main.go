package main

import (
	"context"
	"flag"
	"fmt"
	spyse "github.com/spyse-com/go-spyse/pkg"
	"log"
)

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	var apiBaseUrl = "https://api.spyse.com/v4/data/"
	client, err := spyse.NewClient(apiBaseUrl, *accessToken, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	svc := spyse.NewBulkService(client)

	var domainNames = []string{"example.com", "tesla.com", "google.com", "some-nonexistent-domain.io"}
	domainSearchResults, err := svc.Domain(context.Background(), domainNames)
	if err != nil {
		log.Fatal(err.Error())

	}
	for _, result := range domainSearchResults {
		fmt.Println(fmt.Sprintf("Domain %s, DNS A recod: %s, HTTP status code: %d, Title: %s", result.Name, result.DNSRecords.A[0], *result.Extract.HTTPStatusCode, result.Extract.Title))
	}
	var ipAddresses = []string{"8.8.8.8", "104.22.58.132", "93.184.216.34", "199.66.11.62", "172.217.4.78"}
	ipSearchResults, err := svc.IP(context.Background(), ipAddresses)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, result := range ipSearchResults {
		fmt.Println(fmt.Sprintf("IP %s, open ports: %d, Country: %s, AS number: %d", result.IPAddress, len(result.Ports), result.GEOInfo.Country, *result.ISPInfo.ASNum))
	}
}
