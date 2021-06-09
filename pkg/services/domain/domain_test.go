//nolint:dupl
package domain

import (
	"context"
	"fmt"
	"github.com/spyse-com/go-spyse/pkg"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDomainService_Details(t *testing.T) {
	spyse.setup()
	defer spyse.teardown()
	testAPIEndpoint := "/domain/test.com"

	raw, err := ioutil.ReadFile("../data/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		spyse.testMethod(t, r, http.MethodGet)
		spyse.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := spyse.testClient.Domain.Details(context.Background(), "test.com")
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestDomainService_Search(t *testing.T) {
	spyse.setup()
	defer spyse.teardown()
	testAPIEndpoint := "/domain/search"

	raw, err := ioutil.ReadFile("../data/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		spyse.testMethod(t, r, http.MethodPost)
		spyse.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})
	var params = []map[string]spyse.SearchOption{
		{
			"name": spyse.SearchOption{
				Operator: "eq",
				Value:    "spyse.com",
			},
		},
	}
	domains, err := spyse.testClient.Domain.Search(context.Background(), params, 1, 0)
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
