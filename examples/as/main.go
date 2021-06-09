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

	var detailsAsn = "10000"
	details, err := client.AS.Details(context.Background(), detailsAsn)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(fmt.Sprintf("Details about ASN %d", details.Number))
	println(fmt.Sprintf("  %s", details.Organization))
	println(fmt.Sprintf("  %d IPv4 prefix(es)", len(details.IPv4Prefixes)))
	println(fmt.Sprintf("  %d IPv6 prefix(es)", len(details.IPv6Prefixes)))
	println()

	var searchAsn = "21000"
	var limit = 1
	var offset = 0
	var params = []map[string]spyse.SearchParameter{
		{
			"asn": spyse.SearchParameter{ // More search parameters see at https://spyse-dev.readme.io/reference/autonomous-systems#as_search
				Operator: spyse.OperatorEqual,
				Value:    searchAsn,
			},
		},
	}
	searchResults, err := client.AS.Search(context.Background(), params, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(fmt.Sprintf("Search result by ASN %d", searchResults[0].Number))
	if len(searchResults) > 0 {
		println(fmt.Sprintf("  %s", searchResults[0].Organization))
		println(fmt.Sprintf("  %d IPv4 prefix(es)", len(searchResults[0].IPv4Prefixes)))
		println(fmt.Sprintf("  %d IPv6 prefix(es)", len(searchResults[0].IPv6Prefixes)))
	} else {
		println("No data")
	}
}
