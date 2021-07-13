//nolint:dupl
package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEmailService_Details(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewEmailService(m.Client)

	testAPIEndpoint := "/email/test@gmail.com"

	raw, err := ioutil.ReadFile("../data/email_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodGet)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	emails, err := svc.Details(context.Background(), "test@gmail.com")
	if emails == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestEmailService_Search(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewEmailService(m.Client)

	testAPIEndpoint := "/email/search"

	raw, err := ioutil.ReadFile("../data/email_details.json")
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
		Name:     svc.Params().Email.Name,
		Operator: svc.Params().Email.Operator.Equal,
		Value:    "test@gmail.com",
	})

	emails, err := svc.Search(context.Background(), params.Query, 1, 0)
	if emails == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
