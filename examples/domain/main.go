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
	var subdomainsSearchParams spyse.QueryBuilder

	subdomainsSearchParams.AppendParam(spyse.QueryParam{
		Name:     spyse.DomainParamName,
		Operator: spyse.OperatorEndsWith,
		Value:    searchDomain,
	})

	countResults, err := client.Domain.SearchCount(context.Background(), subdomainsSearchParams.Query)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	var limit = 100
	var offset = 0
	var examplesToPrint = 3
	searchResults, err := client.Domain.Search(context.Background(), subdomainsSearchParams.Query, limit, offset)
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
	var diffSuffixesSearchParams spyse.QueryBuilder

	diffSuffixesSearchParams.AppendParam(spyse.QueryParam{
		Name:     spyse.DomainParamWithoutSuffix,
		Operator: spyse.OperatorEqual,
		Value:    teslaRootDomainNameWithoutTld,
	})
	diffSuffixesSearchParams.AppendParam(spyse.QueryParam{
		Name:     spyse.DomainParamName,
		Operator: spyse.OperatorNotEqual,
		Value:    teslaRootDomainName,
	})
	scrollSearchResultsPageOne, err := client.Domain.ScrollSearch(
		context.Background(), diffSuffixesSearchParams.Query, "")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	scrollSearchResultsPageTwo, err := client.Domain.ScrollSearch(
		context.Background(), diffSuffixesSearchParams.Query, scrollSearchResultsPageOne.SearchID)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	scrollSearchResultsPageThree, err := client.Domain.ScrollSearch(
		context.Background(), diffSuffixesSearchParams.Query, scrollSearchResultsPageOne.SearchID)
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
