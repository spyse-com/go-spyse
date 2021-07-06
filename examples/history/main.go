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
	var domain = "google.com"
	var examplesToPrint = 3
	var limit = 100
	var offset = 0

	println()
	println("DNS history for domain " + domain)
	println()

	client, _ := spyse.NewClient(apiBaseUrl, *accessToken, nil)

	svc := spyse.NewHistoryService(client)

	dnsHistoryA, err := svc.DNS(context.Background(), domain, spyse.DNSTYPEA, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if len(dnsHistoryA) >= examplesToPrint {
		println(fmt.Sprintf("  %-17v%-12v%s", "A record", "First seen", "Last seen"))
		for i := 0; i < examplesToPrint; i++ {
			println(fmt.Sprintf("  %-17v%-12v%s", dnsHistoryA[i].Value, dnsHistoryA[i].FirstSeen, dnsHistoryA[i].LastSeen))
		}
		println(fmt.Sprintf("  --- and %d more ---", len(dnsHistoryA)-examplesToPrint))
	}
	println()

	dnsHistoryNS, err := svc.DNS(context.Background(), domain, spyse.DNSTYPENS, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if len(dnsHistoryNS) >= examplesToPrint {
		println(fmt.Sprintf("  %-17v%-12v%s", "NS record", "First seen", "Last seen"))
		for i := 0; i < examplesToPrint; i++ {
			println(fmt.Sprintf("  %-17v%-12v%s", dnsHistoryNS[i].Value, dnsHistoryNS[i].FirstSeen, dnsHistoryNS[i].LastSeen))
		}
		println(fmt.Sprintf("  --- and %d more ---", len(dnsHistoryNS)-examplesToPrint))
	}
	println()

	dnsHistoryMX, err := svc.DNS(context.Background(), domain, spyse.DNSTYPEMX, limit, offset)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if len(dnsHistoryNS) >= examplesToPrint {
		println(fmt.Sprintf("  %-25v%-12v%s", "MX record", "First seen", "Last seen"))
		for i := 0; i < examplesToPrint; i++ {
			println(fmt.Sprintf("  %-25v%-12v%s", dnsHistoryMX[i].Value, dnsHistoryMX[i].FirstSeen, dnsHistoryMX[i].LastSeen))
		}
		println(fmt.Sprintf("  --- and %d more ---", len(dnsHistoryMX)-examplesToPrint))
	}
	println()
}
