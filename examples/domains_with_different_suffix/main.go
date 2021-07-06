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

	var domainNameWithoutSuffix = "spyse"
	var domainName = "spyse.com"
	var diffSuffixesSearchParams spyse.QueryBuilder

	//We add QueryParam to QueryBuilder, notice that all QueryParam will act with operator AND.
	diffSuffixesSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().WithoutSuffix.Name,
		Operator: svc.Params().WithoutSuffix.Operator.Equal,
		Value:    domainNameWithoutSuffix,
	})
	//We add NotEqual operator for domain name to exclude tesla.com itself from results, but it's optional
	diffSuffixesSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().Name.Name,
		Operator: svc.Params().Name.Operator.NotEqual,
		Value:    domainName,
	})

	//With every ScrollSearch you obtain SearchID and you must use it for next ScrollSearch
	searchResults, err := svc.ScrollSearch(
		context.Background(), diffSuffixesSearchParams.Query, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	for len(searchResults.Items) > 0 {
		//Notice that we process results before new ScrollSearch, because with the first call we obtain results too
		for _, result := range searchResults.Items {
			fmt.Println(result.Name)
		}
		searchResults, err = svc.ScrollSearch(
			context.Background(), diffSuffixesSearchParams.Query, searchResults.SearchID)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}
