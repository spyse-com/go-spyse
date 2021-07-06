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
	asService := spyse.NewASService(client)

	var detailsAsn = "10000"
	details, err := asService.Details(context.Background(), detailsAsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(fmt.Sprintf("Details about ASN %d", details.Number))
	fmt.Println(fmt.Sprintf("  %s", details.Organization))
	fmt.Println(fmt.Sprintf("  %d IPv4 prefix(es)", len(details.IPv4Prefixes)))
	fmt.Println(fmt.Sprintf("  %d IPv6 prefix(es)", len(details.IPv6Prefixes)))

}
