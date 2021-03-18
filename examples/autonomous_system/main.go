package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spyse-com/go-spyse"
)

func main() {
	client, _ := spyse.NewClient(http.DefaultClient, "your_API_token")
	as, err := client.AS().Details(context.Background(), 1)
	if err != nil {
		outputErr(err)
		os.Exit(1)
	}
	outputResponse(as)
	filters := &spyse.ASSearchFilters{
		Filters: map[string]spyse.Filter{
			"as_num": {
				Operator: "eq",
				Value:    "1",
			},
		},
	}
	asSearch, err := client.ASSearch().
		Filter(filters).
		Size(1).From(0).
		Do(context.Background())
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
