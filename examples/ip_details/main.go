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

	svc := spyse.NewIPService(client)

	var IPv4 = "91.210.36.26"
	details, err := svc.Details(context.Background(), IPv4)
	if err != nil {
		log.Fatal(err.Error())
	}
	if details != nil {
		printIpInfo(details)
	} else {
		log.Println(fmt.Sprintf("IP %s not found", IPv4))
	}

}

func printIpInfo(details *spyse.IP) {
	examplesToPrint := 3
	fmt.Println("Details about " + details.IPAddress)
	fmt.Println(fmt.Sprintf("Updated at: %s", details.UpdatedAt))
	fmt.Println(fmt.Sprintf("Country: %s", details.GEOInfo.Country))
	fmt.Println(fmt.Sprintf("Abuse confidence score: %d", details.Abuses.Score))
	fmt.Println(fmt.Sprintf("Number of abuses: %d", details.Abuses.ReportsNum))
	fmt.Println(fmt.Sprintf("CIDR: %s", details.CIDR))
	fmt.Println(fmt.Sprintf("Open ports (%d)", len(details.Ports)))
	if len(details.Ports) >= examplesToPrint {
		for i := 0; i < examplesToPrint; i++ {
			fmt.Println(fmt.Sprintf("  - %d", details.Ports[i].Port))
		}
		fmt.Println(fmt.Sprintf("    --- and %d more ---", len(details.Ports)-examplesToPrint))
	}
	fmt.Println(fmt.Sprintf("Affected by %d CVEs", len(details.CVEList)))
	fmt.Println(fmt.Sprintf("PTR record: %s", details.PtrRecord.Value))
	fmt.Println(fmt.Sprintf("Technologies detected (%d)", len(details.Technologies)))
	if len(details.Technologies) >= examplesToPrint {
		for i := 0; i < 12; i++ {
			fmt.Println(fmt.Sprintf("  - %s, version %s on port %d",
				details.Technologies[i].Name,
				details.Technologies[i].Version,
				details.Technologies[i].Port,
			))
		}
		fmt.Println(fmt.Sprintf("      --- and %d more ---", len(details.Technologies)-examplesToPrint))
	}
	fmt.Println()

}
