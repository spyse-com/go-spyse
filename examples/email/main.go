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
	var examplesToPrint = 10
	var emailDomain = "tesla.com"

	client, _ := spyse.NewClient(apiBaseUrl, *accessToken, nil)

	var emailsSearchParams spyse.QueryBuilder

	emailsSearchParams.AppendParam(spyse.QueryParam{
		Name:     spyse.EmailParamEmail,
		Operator: spyse.OperatorEndsWith,
		Value:    "@" + emailDomain,
	})

	countResults, err := client.Email.SearchCount(context.Background(), emailsSearchParams.Query)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	var offset = 0
	emails, err := client.Email.Search(context.Background(), emailsSearchParams.Query, examplesToPrint, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println()
	println(fmt.Sprintf("All emails seen with domain %s (%d)", emailDomain, countResults))
	println()

	if len(emails) >= examplesToPrint {
		printFormat := "  %-35v%-35v%-25v%s"
		println(fmt.Sprintf(printFormat, "Email", "Last seen", "At location", "Location type"))
		for emailIdx := 0; emailIdx < examplesToPrint; emailIdx++ {
			if len(emails[emailIdx].Sources) == 0 {
				println(fmt.Sprintf(printFormat,
					emails[emailIdx].Email,
					"n/a",
					"n/a",
					"n/a",
				))
			} else {
				for sourceIdx := 0; sourceIdx < len(emails[emailIdx].Sources); sourceIdx++ {
					println(fmt.Sprintf(printFormat,
						emails[emailIdx].Email,
						emails[emailIdx].Sources[sourceIdx].LastSeen,
						emails[emailIdx].Sources[sourceIdx].Target,
						emails[emailIdx].Sources[sourceIdx].Type,
					))
				}
			}
		}
		println(fmt.Sprintf("  --- and %d more ---", countResults-int64(examplesToPrint)))
	}
	println()
}
