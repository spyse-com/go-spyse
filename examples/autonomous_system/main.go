package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spyse-com/go-spyse"
)

const BaseURL = "https://api.spyse.com/v3/data/"

func main() {
	client, _ := spyse.NewClient(BaseURL, "your_API_token", nil)
	as, err := client.AS.Details(context.Background(), 1)
	if err != nil {
		outputErr(err)
		os.Exit(1)
	}
	outputResponse(as)
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
	outputResponse(asSearch)
}

func outputResponse(resp *spyse.ASData) {
	if resp != nil && len(resp.Data.Items) > 0 {
		b, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("marshal err: %s\n", err)
			return
		}
		fmt.Printf("%s\n", string(b))
		return
	}
	fmt.Println("data not found")
}

func outputErr(err error) {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
