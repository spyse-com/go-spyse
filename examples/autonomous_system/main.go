package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spyse-com/go-spyse"
)

func main() {
	client, _ := spyse.NewClient(http.DefaultClient, "eb63839d-2fbb-4763-9be2-134bd65182f0")
	as, err := client.AS.Details(context.Background(), 11)
	if err != nil {
		outputErr(err)
	}
	outputResponse(as)
}

func outputResponse(resp *spyse.ASData) {
	if resp != nil && len(resp.Data.Items) > 0 {
		for _, v := range resp.Data.Items {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Printf("marshal err: %s", err)
				return
			}
			fmt.Printf("%s", string(b))
		}
		return
	}
	fmt.Println("data not found")
}

func outputErr(err error) {
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
