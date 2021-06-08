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

	client, _ := spyse.NewClient(*accessToken, nil)

	var teslaRootDomainName = "tesla.com"
	details, err := client.Domain.Details(context.Background(), teslaRootDomainName)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println()
	println("Details about " + details.Name)
	println("  Website title:", details.Extract.Title)
	println("  Alexa rank:", *details.AlexaInfo.Rank)
	println("  Certificate fingerprint:", details.CertSummary.FingerprintSHA256[:15]+"...")
	println("  Certificate subject org:", details.CertSummary.SubjectDN.O)
	println("  Certificate issuer org:", details.CertSummary.IssuerDN.O)
	println("  Updated at:", details.UpdatedAt)
	println()

	var searchDomain = ".tesla.com"
	var subdomainsSearchParams = []map[string]spyse.SearchParameter{
		{
			// More search parameters see at https://spyse-dev.readme.io/reference/domains#domain_search
			"name": spyse.SearchParameter{
				Operator: spyse.SearchOperatorEndsWith,
				Value:    searchDomain,
			},
		},
	}

	countResults, err := client.Domain.SearchCount(context.Background(), subdomainsSearchParams)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	var limit = 100
	var offset = 0
	var examplesToPrint = 3
	searchResults, err := client.Domain.Search(context.Background(), subdomainsSearchParams, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(fmt.Sprintf("Subdomains (%d) :", countResults))
	for i := 0; i < examplesToPrint; i++ {
		println("  - " + searchResults[i].Name)
	}
	println(fmt.Sprintf("    --- and %d more ---", countResults-int64(examplesToPrint)))
	println()

	var teslaRootDomainNameWithoutTld = "tesla"
	var diffSuffixesSearchParams = []map[string]spyse.SearchParameter{
		{
			// More search parameters see at https://spyse-dev.readme.io/reference/domains#domain_search
			"without_suffix": spyse.SearchParameter{
				Operator: spyse.SearchOperatorEqual,
				Value:    teslaRootDomainNameWithoutTld,
			},
			"name": spyse.SearchParameter{
				Operator: spyse.SearchOperatorNotEqual,
				Value:    teslaRootDomainName,
			},
		},
	}
	scrollSearchResultsPageOne, err := client.Domain.ScrollSearch(
		context.Background(), diffSuffixesSearchParams, "")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	scrollSearchResultsPageTwo, err := client.Domain.ScrollSearch(
		context.Background(), diffSuffixesSearchParams, scrollSearchResultsPageOne.SearchID)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	scrollSearchResultsPageThree, err := client.Domain.ScrollSearch(
		context.Background(), diffSuffixesSearchParams, scrollSearchResultsPageOne.SearchID)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(fmt.Sprintf("Scroll search for domain names with different suffixes (%s)", teslaRootDomainName))
	println("  Search ID:", scrollSearchResultsPageOne.SearchID[0:15]+"...")
	println("  Total items:", scrollSearchResultsPageOne.TotalItems)
	println(fmt.Sprintf("  Results on the first page (%d):", len(scrollSearchResultsPageOne.Items)))
	for i := 0; i < examplesToPrint; i++ {
		println("  - " + scrollSearchResultsPageOne.Items[i].Name)
	}
	println(fmt.Sprintf("    --- and %d more ---", len(scrollSearchResultsPageOne.Items)-examplesToPrint))
	println(fmt.Sprintf("  Results on the second page (%d):", len(scrollSearchResultsPageTwo.Items)))
	for i := 0; i < examplesToPrint; i++ {
		println("  - " + scrollSearchResultsPageTwo.Items[i].Name)
	}
	println(fmt.Sprintf("    --- and %d more ---", len(scrollSearchResultsPageTwo.Items)-examplesToPrint))
	println(fmt.Sprintf("  Results on the third page (%d):", len(scrollSearchResultsPageThree.Items)))
	for i := 0; i < examplesToPrint; i++ {
		println("  - " + scrollSearchResultsPageThree.Items[i].Name)
	}
	println(fmt.Sprintf("    --- and %d more ---", len(scrollSearchResultsPageThree.Items)-examplesToPrint))
}
