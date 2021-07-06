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

	//Dot before domain name is important because we search any domains that name ends with ".tesla.com"
	var searchDomain = ".tesla.com"
	var subdomainsSearchParams spyse.QueryBuilder

	//We add QueryParam to QueryBuilder, notice that if you add more QueryParam they will act with operator AND.
	subdomainsSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().Name.Name,
		Operator: svc.Params().Name.Operator.EndsWith,
		Value:    searchDomain,
	})

	//You can count result of search before search itself.
	//It's helpful to iterate over results, or you can skip SearchCount and call Search.
	countResults, err := svc.SearchCount(context.Background(), subdomainsSearchParams.Query)
	if err != nil {
		log.Fatal(err.Error())
	}

	var limit = 100
	var offset = 0
	var searchResults []spyse.Domain
	var domain spyse.Domain
	for ; int64(offset) < countResults; offset += limit {
		// !Notice that you can obtain only first 10000 (can depend on your subscription plan) results using Search method
		searchResults, err = svc.Search(context.Background(), subdomainsSearchParams.Query, limit, offset)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, domain = range searchResults {
			fmt.Println(domain.Name)
		}
	}

}
