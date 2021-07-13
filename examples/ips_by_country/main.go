package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spyse-com/go-spyse/pkg"
	"log"
)

func main() {

	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	client, err := spyse.NewClient(*accessToken, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	svc := spyse.NewIPService(client)

	var usCountryName = "United States"
	var usHostsSearchParams spyse.QueryBuilder

	usHostsSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().GeoCountry.Name,
		Operator: svc.Params().GeoCountry.Operator.Equal,
		Value:    usCountryName,
	})

	countResults, err := svc.SearchCount(context.Background(), usHostsSearchParams.Query)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(fmt.Sprintf("IPs with %s port: %d", usCountryName, countResults))

	var limit = 100
	var offset = 0
	var searchResults []spyse.IP
	var ip spyse.IP
	for ; int64(offset) < countResults; offset += limit {
		//Notice that you can fetch only first 10000 (can depend on your subscription plan) results using Search method
		searchResults, err = svc.Search(context.Background(), usHostsSearchParams.Query, limit, offset)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, ip = range searchResults {
			fmt.Println(ip.IPAddress)
		}
		//Remove break if you want to see all IPs in this example, BUT notice that every Search call count as 1 request to API
		break
	}
}
