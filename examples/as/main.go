package main

import (
	"context"
	"flag"
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
	println("Details result")
	println("ASN", details.Number,
		"-", details.Organization,
		"-", len(details.IPv4Prefixes), "IPv4 prefixes",
		"-", len(details.IPv6Prefixes), "IPv6 prefixes",
	)
	println()

	var searchAsn = "21000"
	var limit = 1
	var offset = 0
	var params = []map[string]spyse.SearchParameter{
		{
			"asn": spyse.SearchParameter{ // More search parameters see at https://spyse-dev.readme.io/reference/autonomous-systems#as_search
				Operator: spyse.SEARCH_OPERATOR_EQUAL,
				Value:    searchAsn,
			},
		},
	}
	searchResults, err := client.AS.Search(context.Background(), params, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Search result")
	if len(searchResults) > 0 {
		println("ASN", searchResults[0].Number,
			"-", searchResults[0].Organization,
			"-", len(searchResults[0].IPv4Prefixes), "IPv4 prefixes",
			"-", len(searchResults[0].IPv6Prefixes), "IPv6 prefixes",
		)
	} else {
		println("No data")
	}
}
