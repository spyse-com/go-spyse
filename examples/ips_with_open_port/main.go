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
	svc := spyse.NewIPService(client)

	var searchPort = "9200" // Elasticsearch port
	var ipSearchParams spyse.QueryBuilder

	//We add QueryParam to QueryBuilder, notice that if you add more QueryParam they will act with operator AND.
	ipSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().OpenPort.Name,
		Operator: svc.Params().OpenPort.Operator.Equal,
		Value:    searchPort,
	})

	//You can count result of search before search itself.
	//It's helpful to iterate over results, or you can skip SearchCount and call Search.
	countResults, err := svc.SearchCount(context.Background(), ipSearchParams.Query)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(fmt.Sprintf("IPs with %s port: %d", searchPort, countResults))

	var limit = 100
	var offset = 0
	var searchResults []spyse.IP
	var ip spyse.IP
	for ; int64(offset) < countResults; offset += limit {
		//Notice that you can obtain only first 10000 (can depend on your subscription plan) results using Search method
		searchResults, err = svc.Search(context.Background(), ipSearchParams.Query, limit, offset)
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
