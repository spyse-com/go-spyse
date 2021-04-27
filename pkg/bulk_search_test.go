package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBulkSearchService_Domain(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/bulk-search/domain"

	raw, err := ioutil.ReadFile("../mocks/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := testClient.BulkSearch.Domain(
		context.Background(),
		[]string{"google.com"},
		1, 0,
	)
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestBulkSearchService_IP(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/bulk-search/ip"

	raw, err := ioutil.ReadFile("../mocks/ip_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := testClient.BulkSearch.IP(
		context.Background(),
		[]string{"8.8.8.8"},
		1, 0,
	)
	if domains == nil {
		t.Error("Expected IP struct. IP struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
