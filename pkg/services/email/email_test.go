//nolint:dupl
package email

import (
	"context"
	"fmt"
	"github.com/spyse-com/go-spyse/pkg"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEmailService_Details(t *testing.T) {
	spyse.setup()
	defer spyse.teardown()
	testAPIEndpoint := "/email/test@gmail.com"

	raw, err := ioutil.ReadFile("../data/email_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		spyse.testMethod(t, r, http.MethodGet)
		spyse.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	emails, err := spyse.testClient.Email.Details(context.Background(), "test@gmail.com")
	if emails == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestEmailService_Search(t *testing.T) {
	spyse.setup()
	defer spyse.teardown()
	testAPIEndpoint := "/email/search"

	raw, err := ioutil.ReadFile("../data/email_details.json")
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
			"email": spyse.SearchOption{
				Operator: "eq",
				Value:    "test@gmail.com",
			},
		},
	}
	emails, err := spyse.testClient.Email.Search(context.Background(), params, 1, 0)
	if emails == nil {
		t.Error("Expected Email struct. Email struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
