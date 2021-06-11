package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBulkSearchService_Domain(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewBulkService(m.Client)

	testAPIEndpoint := "/bulk-search/domain"

	raw, err := ioutil.ReadFile("../data/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodPost)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := svc.Domain(
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
	m := setup()
	defer m.teardown()

	svc := NewBulkService(m.Client)

	testAPIEndpoint := "/bulk-search/ip"

	raw, err := ioutil.ReadFile("../data/ip_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodPost)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := svc.IP(
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
