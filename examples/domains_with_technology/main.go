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
	//You can see list of all technologies on spyse.com
	var technologyName = "WordPress"
	var domainsWithTechSearch spyse.QueryBuilder

	//We add QueryParam to QueryBuilder, notice that all QueryParam will act with operator AND.
	//We use AppendGroupParams to search technology name and version in pair.
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

	//And we can exclude subdomains
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

	//With every ScrollSearch you obtain SearchID and you must use it for next ScrollSearch
	searchResults, err := svc.ScrollSearch(
		context.Background(), domainsWithTechSearch.Query, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Domains list:")
	for len(searchResults.Items) > 0 {
		//Notice that we process results before new ScrollSearch, because with the first call we obtain results too
		for _, result := range searchResults.Items {
			fmt.Println(result.Name)
		}
		//Remove this break to gather all domains, BUT Spyse will count every ScrollSearch as 1 API request.
		break
		searchResults, err = svc.ScrollSearch(
			context.Background(), domainsWithTechSearch.Query, searchResults.SearchID)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}
