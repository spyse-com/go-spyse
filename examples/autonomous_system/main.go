package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	spyse "github.com/spyse-com/go-spyse/pkg"
)

const BaseURL = "https://api.spyse.com/v4/data/"

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	client, _ := spyse.NewClient(BaseURL, *accessToken, nil)
	asResponse, err := client.AS.Details(context.Background(), 1)
	if err != nil {
		outputErr(err)
		os.Exit(1)
	}
	println(asResponse.Number)
	var params = []map[string]spyse.SearchParameter{
		{
			"asn": spyse.SearchParameter{
				Operator: "eq",
				Value:    "1",
			},
		},
	}
	asSearchResponse, err := client.AS.Search(context.Background(), params, 1, 0)
	if err != nil {
		outputErr(err)
		os.Exit(1)
	}
	println(asSearchResponse[0].Number)
}

func outputErr(err error) {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
