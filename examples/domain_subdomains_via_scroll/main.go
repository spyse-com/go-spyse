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
	//You can test this example on domains with many subdomains, e.g. "wix.com" but notice that every ScrollSearch call count as 1 request to API
	var domainName = ".spyse.com"
	var subdomainsSearch spyse.QueryBuilder

	subdomainsSearch.AppendParam(spyse.QueryParam{
		Name:     svc.Params().Name.Name,
		Operator: svc.Params().Name.Operator.EndsWith,
		Value:    domainName,
	})

	searchResults, err := svc.ScrollSearch(
		context.Background(), subdomainsSearch.Query, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	for len(searchResults.Items) > 0 {
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
