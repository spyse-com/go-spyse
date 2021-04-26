package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEmailService_Details(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/email/test@gmail.com"

	raw, err := ioutil.ReadFile("../mocks/email_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, string(raw))
	})

	autonomousSystem, err := testClient.Email.Details(context.Background(), "test@gmail.com")
	if autonomousSystem == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestEmailService_Search(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/email/search"

	raw, err := ioutil.ReadFile("../mocks/email_details.json")
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
			"email": Filter{
				Operator: "eq",
				Value:    "test@gmail.com",
			},
		},
	}
	autonomousSystem, err := testClient.Email.Search(context.Background(), filters, 1, 0)
	if autonomousSystem == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
