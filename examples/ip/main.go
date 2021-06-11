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
	var examplesToPrint = 3

	client, _ := spyse.NewClient(apiBaseUrl, *accessToken, nil)

	svc := spyse.NewIPService(client)

	var detailsIP = "91.210.36.26"
	details, err := svc.Details(context.Background(), detailsIP)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println()
	println("Details about " + details.IPAddress)
	println(fmt.Sprintf("  Updated at: %s", details.UpdatedAt))
	println(fmt.Sprintf("  Country: %s", details.GEOInfo.Country))
	println(fmt.Sprintf("  Abuse confidence score: %d", details.Abuses.Score))
	println(fmt.Sprintf("  Number of abuses: %d", details.Abuses.ReportsNum))
	println(fmt.Sprintf("  CIDR: %s", details.CIDR))
	println(fmt.Sprintf("  Open ports (%d)", len(details.Ports)))
	if len(details.Ports) >= examplesToPrint {
		for i := 0; i < examplesToPrint; i++ {
			println(fmt.Sprintf("  - %d", details.Ports[i].Port))
		}
		println(fmt.Sprintf("    --- and %d more ---", len(details.Ports)-examplesToPrint))
	}
	println(fmt.Sprintf("  Affected by %d CVEs", len(details.CVEList)))
	println(fmt.Sprintf("  PTR record: %s", details.PtrRecord.Value))
	println(fmt.Sprintf("  Technologies detected (%d)", len(details.Technologies)))
	if len(details.Technologies) >= examplesToPrint {
		for i := 0; i < 12; i++ {
			println(fmt.Sprintf("  - %s, version %s on port %d",
				details.Technologies[i].Name,
				details.Technologies[i].Version,
				details.Technologies[i].Port,
			))
		}
		println(fmt.Sprintf("      --- and %d more ---", len(details.Technologies)-examplesToPrint))
	}
	println()

	var searchPort = "9200" // Elasticsearch port
	var ipSearchParams spyse.QueryBuilder

	ipSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().OpenPort.Name,
		Operator: svc.Params().OpenPort.Operator.Equal,
		Value:    searchPort,
	})

	countResults, err := svc.SearchCount(context.Background(), ipSearchParams.Query)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	var limit = 100
	var offset = 0
	searchResults, err := svc.Search(context.Background(), ipSearchParams.Query, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(fmt.Sprintf("IPs with open port %s (%d) :", searchPort, countResults))
	if len(searchResults) >= examplesToPrint {
		for i := 0; i < examplesToPrint; i++ {
			println(fmt.Sprintf("  - %s (%s)", searchResults[i].IPAddress, searchResults[i].GEOInfo.Country))
		}
		println(fmt.Sprintf("    and %d more", int(countResults)-examplesToPrint))
	}
	println()

	var usCountryName = "United States"
	var usHostsSearchParams spyse.QueryBuilder

	usHostsSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().GeoCountry.Name,
		Operator: svc.Params().GeoCountry.Operator.Equal,
		Value:    usCountryName,
	})

	scrollSearchResults, err := svc.ScrollSearch(
		context.Background(), usHostsSearchParams.Query, "")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(fmt.Sprintf("Scroll search for IPs in %s", usCountryName))
	println("  Search ID:", scrollSearchResults.SearchID[0:15]+"...")
	println("  Total items:", scrollSearchResults.TotalItems)

	if len(scrollSearchResults.Items) >= examplesToPrint {
		println(fmt.Sprintf("  Results on the first page (%d):", len(scrollSearchResults.Items)))
		for i := 0; i < examplesToPrint; i++ {
			println("  - " + scrollSearchResults.Items[i].IPAddress)
		}
		println(fmt.Sprintf("    and %d more", len(scrollSearchResults.Items)-examplesToPrint))
	}
}
