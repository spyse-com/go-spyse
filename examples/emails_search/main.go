package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spyse-com/go-spyse/pkg"
)

func main() {
	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	client, err := spyse.NewClient(*accessToken, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	var emailDomain = "tesla.com"
	svc := spyse.NewEmailService(client)

	var emailsSearchParams spyse.QueryBuilder

	emailsSearchParams.AppendParam(spyse.QueryParam{
		Name:     svc.Params().Email.Name,
		Operator: svc.Params().Email.Operator.EndsWith,
		Value:    "@" + emailDomain,
	})

	countResults, err := svc.SearchCount(context.Background(), emailsSearchParams.Query)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Emails found:", countResults)

	var offset = 0
	var limit = 100
	emails, err := svc.Search(context.Background(), emailsSearchParams.Query, limit, offset)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Emails list:")
	for _, email := range emails {
		fmt.Println(email.Email)
		for _, source := range email.Sources {
			fmt.Println(" found on:", source.Target)
		}
	}
}
