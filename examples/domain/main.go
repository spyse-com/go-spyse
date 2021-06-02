package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	spyse "github.com/spyse-com/go-spyse/pkg"
)

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	var apiBaseUrl = "https://api.spyse.com/v4/data/"

	client, _ := spyse.NewClient(apiBaseUrl, *accessToken, nil)

	var detailsDomain = "tesla.com"
	details, err := client.Domain.Details(context.Background(), detailsDomain)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Details result")
	println("Domain:", details.Name)
	println("Updated at:", details.UpdatedAt)
	println("Website title:", details.Extract.Title)
	println("Alexa rank:", *details.AlexaInfo.Rank)
	println("Certificate fingerprint:", details.CertSummary.FingerprintSHA256)
	println("Certificate subject org:", details.CertSummary.SubjectDN.O)
	println("Certificate issuer org:", details.CertSummary.IssuerDN.O)
	println()

	var searchDomain = ".tesla.com"
	var params = []map[string]spyse.SearchParameter{
		{
			"name": spyse.SearchParameter{ // More search parameters see at https://spyse-dev.readme.io/reference/domains#domain_search
				Operator: spyse.SEARCH_OPERATOR_ENDS_WITH,
				Value:    searchDomain,
			},
		},
	}
	countResults, err := client.Domain.SearchCount(context.Background(), params)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	var limit = 100
	var offset = 0
	searchResults, err := client.Domain.Search(context.Background(), params, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(fmt.Sprintf("Subdomains (%d) :", countResults))
	println("-", searchResults[0].Name)
	println("-", searchResults[1].Name)
	println("-", searchResults[2].Name)
	println(fmt.Sprintf("- and %d more", countResults-3))
}
