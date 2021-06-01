package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestIPService_Details(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/ip/8.8.8.8"

	raw, err := ioutil.ReadFile("../mocks/ip_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	ips, err := testClient.IP.Details(context.Background(), "8.8.8.8")
	if ips == nil {
		t.Error("Expected IP struct. IP struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIPService_Search(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/ip/search"

	raw, err := ioutil.ReadFile("../mocks/ip_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})
	var filters = []map[string]SearchParameter{
		{
			"cidr": SearchParameter{
				Operator: "eq",
				Value:    "8.8.8.8/32",
			},
		},
	}
	ips, err := testClient.IP.Search(context.Background(), filters, 1, 0)
	if ips == nil {
		t.Error("Expected IP struct. IP struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
