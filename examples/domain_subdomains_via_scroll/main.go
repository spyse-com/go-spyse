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

	svc := spyse.NewDomainService(client)

	//Dot before domain name is important because we search any domains that name ends with ".spyse.com"
	//You can test this example on domains with many subdomains, e.g. "wix.com", but notice that every ScrollSearch calls count as 1 request to API
	var domainName = ".spyse.com"
	var subdomainsSearch spyse.QueryBuilder

	//We add QueryParam to QueryBuilder, notice that all QueryParam will act with operator AND.
	subdomainsSearch.AppendParam(spyse.QueryParam{
		Name:     svc.Params().Name.Name,
		Operator: svc.Params().Name.Operator.EndsWith,
		Value:    domainName,
	})

	//With every ScrollSearch you obtain SearchID and you must use it for next ScrollSearch
	searchResults, err := svc.ScrollSearch(
		context.Background(), subdomainsSearch.Query, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	for len(searchResults.Items) > 0 {
		//Notice that we process results before new ScrollSearch, because with the first call we obtain results too
		for _, result := range searchResults.Items {
			fmt.Println(result.Name)
		}
		searchResults, err = svc.ScrollSearch(
			context.Background(), subdomainsSearch.Query, searchResults.SearchID)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}
