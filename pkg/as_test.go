package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestASService_Details(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/as/1"

	raw, err := ioutil.ReadFile("../mocks/as_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	autonomousSystem, err := testClient.AS.Details(context.Background(), "1")
	if autonomousSystem == nil {
		t.Error("Expected AS struct. AS struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
