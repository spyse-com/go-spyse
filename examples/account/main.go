package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	spyse "github.com/spyse-com/go-spyse/pkg"
)

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()
	client, err := spyse.NewClient(*accessToken, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	asService := spyse.NewAccountService(client)

	account, err := asService.Quota(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Customer account quotas:")
	fmt.Println(fmt.Sprintf("Subscription period start at: %s", account.StartAt))
	fmt.Println(fmt.Sprintf("Subscription period end at: %s", account.EndAt))
	fmt.Println(fmt.Sprintf("API requests remaining: %d", account.APIRequestsRemaining))
	fmt.Println(fmt.Sprintf("API requests limit: %d", account.APIRequestsLimit))
	fmt.Println(fmt.Sprintf("Downloads remaining: %d", account.DownloadsLimitRemaining))
	fmt.Println(fmt.Sprintf("Downloads limit: %d", account.DownloadsLimit))
	fmt.Println(fmt.Sprintf("On demand scan remaining: %d", account.OnDemandScansRemaining))
	fmt.Println(fmt.Sprintf("On demand scan limit: %d", account.OnDemandScansLimit))
	fmt.Println(fmt.Sprintf("Is scroll search enabled: %t", account.IsScrollSearchEnabled))
	fmt.Println(fmt.Sprintf("Search params limit: %d", account.SearchParamsLimit))
	fmt.Println(fmt.Sprintf("Requests rate limit per second: %d", account.RequestsRateLimit))
}
