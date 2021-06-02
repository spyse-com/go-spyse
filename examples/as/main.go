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
	asResponse, err := client.AS.Details(context.Background(), detailsAsn)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Details result")
	println("ASN", asResponse.Number,
		"-", asResponse.Organization,
		"-", len(asResponse.IPv4Prefixes), "IPv4 prefixes",
		"-", len(asResponse.IPv6Prefixes), "IPv6 prefixes",
	)
	println()

	var searchAsn = "21000"
	var limit = 1
	var offset = 0
	var params = []map[string]spyse.SearchParameter{
		{
			"asn": spyse.SearchParameter{
				Operator: "eq",
				Value:    searchAsn,
			},
		},
	}
	asSearchResponse, err := client.AS.Search(context.Background(), params, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Search result")
	if len(asSearchResponse) > 0 {
		println("ASN", asSearchResponse[0].Number,
			"-", asSearchResponse[0].Organization,
			"-", len(asSearchResponse[0].IPv4Prefixes), "IPv4 prefixes",
			"-", len(asSearchResponse[0].IPv6Prefixes), "IPv6 prefixes",
		)
	} else {
		println("No data")
	}
}
