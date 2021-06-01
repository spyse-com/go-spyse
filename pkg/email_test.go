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
	testAPIEndpoint := "/email/test@gmail.com"

	raw, err := ioutil.ReadFile("../mocks/email_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	emails, err := testClient.Email.Details(context.Background(), "test@gmail.com")
	if emails == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestEmailService_Search(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/email/search"

	raw, err := ioutil.ReadFile("../mocks/email_details.json")
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
			"email": SearchParameter{
				Operator: "eq",
				Value:    "test@gmail.com",
			},
		},
	}
	emails, err := testClient.Email.Search(context.Background(), params, 1, 0)
	if emails == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
