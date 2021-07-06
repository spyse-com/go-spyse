package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	spyse "github.com/spyse-com/go-spyse/pkg"
)

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	var apiBaseUrl = "https://api.spyse.com/v4/data/"

	client, _ := spyse.NewClient(apiBaseUrl, *accessToken, nil)
	asService := spyse.NewAccountService(client)

	account, err := asService.Quota(context.Background())
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Customer account quotas:")
	println(fmt.Sprintf("start_at: %s", account.StartAt))
	println(fmt.Sprintf("end_at: %s", account.EndAt))
	println(fmt.Sprintf("api_requests_remaining: %d", account.APIRequestsRemaining))
	println(fmt.Sprintf("api_requests_limit: %d", account.APIRequestsLimit))
	println(fmt.Sprintf("downloads_limit_remaining: %d", account.DownloadsLimitRemaining))
	println(fmt.Sprintf("downloads_limit: %d", account.DownloadsLimit))
	println()
}
