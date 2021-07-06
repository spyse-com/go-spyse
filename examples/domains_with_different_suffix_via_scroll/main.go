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

	var domainNameWithoutSuffix = "spyse"
	var domainName = "spyse.com"
	var diffSuffixesSearchParams spyse.QueryBuilder

	diffSuffixesSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().WithoutSuffix.Name,
		Operator: svc.Params().WithoutSuffix.Operator.Equal,
		Value:    domainNameWithoutSuffix,
	})
	diffSuffixesSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().Name.Name,
		Operator: svc.Params().Name.Operator.NotEqual,
		Value:    domainName,
	})

	searchResults, err := svc.ScrollSearch(
		context.Background(), diffSuffixesSearchParams.Query, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	for len(searchResults.Items) > 0 {
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
