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

	svc := spyse.NewDomainService(client)

	//Dot before the domain name is important because search fetch any domains that end with ".tesla.com"
	var searchDomain = ".tesla.com"
	var subdomainsSearchParams spyse.QueryBuilder

	subdomainsSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().Name.Name,
		Operator: svc.Params().Name.Operator.EndsWith,
		Value:    searchDomain,
	})

	countResults, err := svc.SearchCount(context.Background(), subdomainsSearchParams.Query)
	if err != nil {
		log.Fatal(err.Error())
	}

	var limit = 100
	var offset = 0
	var searchResults []spyse.Domain
	var domain spyse.Domain
	for ; int64(offset) < countResults; offset += limit {
		//Notice that you can fetch only the first 10000 (can depend on your subscription plan) results using the Search method
		searchResults, err = svc.Search(context.Background(), subdomainsSearchParams.Query, limit, offset)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, domain = range searchResults {
			fmt.Println(domain.Name)
		}
	}

}
