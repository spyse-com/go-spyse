package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/spyse-com/go-spyse/pkg"
)

func main() {

	accessToken := flag.String("access_token", "", "API personal access token")
	flag.Parse()

	client, err := spyse.NewClient(*accessToken, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	svc := spyse.NewHistoryService(client)

	var domain = "google.com"
	var limit = 100
	var offset = 0
	//DNS A, AAAA, MX, NS, TXT, CNAME records type available
	dnsHistoryA, err := svc.DNS(context.Background(), domain, spyse.DNSTYPEA, limit, offset)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("A records:")
	for _, record := range dnsHistoryA {
		fmt.Println(record.Value)
	}

	dnsHistoryNS, err := svc.DNS(context.Background(), domain, spyse.DNSTYPENS, limit, offset)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("NS records:")
	for _, record := range dnsHistoryNS {
		fmt.Println(record.Value)
	}

	dnsHistoryMX, err := svc.DNS(context.Background(), domain, spyse.DNSTYPEMX, limit, offset)
	if err != nil {
		log.Fatal(err.Error())

	}
	fmt.Println("MX records:")
	for _, record := range dnsHistoryMX {
		fmt.Println(record.Value)
	}

}
