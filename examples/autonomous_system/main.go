package main

import (
	"context"
	"flag"
	"fmt"
	spyse "go-spyse/pkg"
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

//
//func outputResponse(resp *spyse.ASData) {
//	if resp != nil && len(resp.Data.Items) > 0 {
//		b, err := json.Marshal(resp)
//		if err != nil {
//			fmt.Printf("marshal err: %s\n", err)
//			return
//		}
//		fmt.Printf("%s\n", string(b))
//		return
//	}
//	fmt.Println("data not found")
//}

func outputErr(err error) {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
