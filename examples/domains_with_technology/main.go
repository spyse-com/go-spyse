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

	var technologyName = "WordPress"
	var domainsWithTechSearch spyse.QueryBuilder

	domainsWithTechSearch.AppendParam(
		spyse.QueryParam{
			Name:     svc.Params().TechnologyName.Name,
			Operator: svc.Params().TechnologyName.Operator.Equal,
			Value:    technologyName,
		},
		spyse.QueryParam{
			Name:     svc.Params().TechnologyVersion.Name,
			Operator: svc.Params().TechnologyVersion.Operator.Equal,
			Value:    "4.6.1",
		},
	)

	domainsWithTechSearch.AppendParam(spyse.QueryParam{
		Name:     svc.Params().IsSubdomain.Name,
		Operator: svc.Params().IsSubdomain.Operator.Equal,
		Value:    "false",
	})

	countResults, err := svc.SearchCount(context.Background(), domainsWithTechSearch.Query)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Domains with technology:", countResults)

	searchResults, err := svc.ScrollSearch(
		context.Background(), domainsWithTechSearch.Query, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Domains list:")
	for len(searchResults.Items) > 0 {
		for _, result := range searchResults.Items {
			fmt.Println(result.Name)
		}
		//Remove this break to fetch all domains, but Spyse will count every ScrollSearch as 1 API request.
		break
		searchResults, err = svc.ScrollSearch(
			context.Background(), domainsWithTechSearch.Query, searchResults.SearchID)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}
