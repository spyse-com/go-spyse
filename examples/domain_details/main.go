package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	spyse "github.com/spyse-com/go-spyse/pkg"
)

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	client, err := spyse.NewClient(*accessToken, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	svc := spyse.NewDomainService(client)

	domainName := "spyse.com"

	details, err := svc.Details(context.Background(), domainName)
	if err != nil {
		log.Fatal(err.Error())
	}
	if details != nil {
		printDomainInfo(details)
	} else {
		log.Println(fmt.Sprintf("Domain %s not found", domainName))
	}

	//	Case with domain that not exist
	domainName = "fakedomainplsdontcreateit.com"
	details, err = svc.Details(context.Background(), domainName)
	if err != nil {
		log.Fatal(err.Error())
	}
	if details != nil {
		printDomainInfo(details)
	} else {
		log.Println(fmt.Sprintf("Domain %s not found", domainName))
	}
}

func printDomainInfo(details *spyse.Domain) {
	fmt.Println("Details about " + details.Name)
	for _, dnsA := range details.DNSRecords.A {
		fmt.Println("DNS A record:", dnsA)
	}
	fmt.Println("Website title:", details.Extract.Title)
	fmt.Println("Alexa rank:", details.AlexaInfo.Rank)
	fmt.Println("Certificate subject org:", details.CertSummary.Subject.O)
	fmt.Println("Certificate issuer org:", details.CertSummary.Issuer.O)
	fmt.Println("Updated at:", details.UpdatedAt)
}
