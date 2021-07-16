//nolint:dupl
package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDomainService_Details(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewDomainService(m.Client)

	testAPIEndpoint := "/domain/test.com"

	raw, err := ioutil.ReadFile("../data/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodGet)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	domains, err := svc.Details(context.Background(), "test.com")
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestDomainService_Search(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewDomainService(m.Client)

	testAPIEndpoint := "/domain/search"

	raw, err := ioutil.ReadFile("../data/domain_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodPost)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})
	var params QueryBuilder

	params.AppendParam(QueryParam{
		Name:     svc.Params().Name.Name,
		Operator: svc.Params().Name.Operator.Equal,
		Value:    "spyse.com",
	})

	domains, err := svc.Search(context.Background(), params.Query, 1, 0)
	if domains == nil {
		t.Error("Expected Domain struct. Domain struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
