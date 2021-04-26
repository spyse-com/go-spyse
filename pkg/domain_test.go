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
	testAPIEdpoint := "/domain/test.com"

	raw, err := ioutil.ReadFile("../mocks/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, string(raw))
	})

	autonomousSystem, err := testClient.Domain.Details(context.Background(), "test.com")
	if autonomousSystem == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestDomainService_Search(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/domain/search"

	raw, err := ioutil.ReadFile("../mocks/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, string(raw))
	})
	var filters = []map[string]Filter{
		{
			"name": Filter{
				Operator: "eq",
				Value:    "spyse.com",
			},
		},
	}
	autonomousSystem, err := testClient.Domain.Search(context.Background(), filters, 1, 0)
	if autonomousSystem == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
