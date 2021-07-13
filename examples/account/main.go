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
}
