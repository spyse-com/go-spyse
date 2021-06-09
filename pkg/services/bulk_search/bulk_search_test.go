package bulk_search

import (
	"context"
	"fmt"
	"github.com/spyse-com/go-spyse/pkg"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBulkSearchService_Domain(t *testing.T) {
	spyse.setup()
	defer spyse.teardown()
	testAPIEndpoint := "/bulk-search/domain"

	raw, err := ioutil.ReadFile("../data/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		spyse.testMethod(t, r, http.MethodPost)
		spyse.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := spyse.testClient.BulkSearch.Domain(
		context.Background(),
		[]string{"google.com"},
	)
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestBulkSearchService_IP(t *testing.T) {
	spyse.setup()
	defer spyse.teardown()
	testAPIEndpoint := "/bulk-search/ip"

	raw, err := ioutil.ReadFile("../data/ip_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		spyse.testMethod(t, r, http.MethodPost)
		spyse.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := spyse.testClient.BulkSearch.IP(
		context.Background(),
		[]string{"8.8.8.8"},
	)
	if domains == nil {
		t.Error("Expected IP struct. IP struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
