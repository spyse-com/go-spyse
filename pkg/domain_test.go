package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDomainService_Details(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/domain/test.com"

	raw, err := ioutil.ReadFile("../mocks/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := testClient.Domain.Details(context.Background(), "test.com")
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestDomainService_Search(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/domain/search"

	raw, err := ioutil.ReadFile("../mocks/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})
	var params = []map[string]SearchParameter{
		{
			"name": SearchParameter{
				Operator: "eq",
				Value:    "spyse.com",
			},
		},
	}
	domains, err := testClient.Domain.Search(context.Background(), params, 1, 0)
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
