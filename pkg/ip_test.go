//nolint:dupl
package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestIPService_Details(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewIPService(m.Client)

	testAPIEndpoint := "/ip/8.8.8.8"

	raw, err := ioutil.ReadFile("../data/ip_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodGet)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	ips, err := svc.Details(context.Background(), "8.8.8.8")
	if ips == nil {
		t.Error("Expected IP struct. IP struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIPService_Search(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewIPService(m.Client)

	testAPIEndpoint := "/ip/search"

	raw, err := ioutil.ReadFile("../data/ip_details.json")
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
		Name:     svc.Params().CIDR.Name,
		Operator: svc.Params().CIDR.Operator.Equal,
		Value:    "8.8.8.8/32",
	})

	ips, err := svc.Search(context.Background(), params.Query, 1, 0)
	if ips == nil {
		t.Error("Expected IP struct. IP struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
