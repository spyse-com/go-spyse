package main

import (
	"context"
	"flag"
	"fmt"
	spyse "github.com/spyse-com/go-spyse/pkg"
	"os"
)

const BaseURL = "https://api.spyse.com/v3/data/"

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	client, _ := spyse.NewClient(BaseURL, *accessToken, nil)
	as, err := client.AS.Details(context.Background(), 1)
	if err != nil {
		outputErr(err)
		os.Exit(1)
	}
	println(as.ASN)
	var filters = []map[string]spyse.Filter{
		{
			"as_num": spyse.Filter{
				Operator: "eq",
				Value:    "1",
			},
		},
	}
	asSearch, err := client.AS.Search(context.Background(), filters, 1, 0)
	if err != nil {
		outputErr(err)
		os.Exit(1)
	}
	println(asSearch[0].ASN)
}

func outputErr(err error) {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
