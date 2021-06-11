package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestASService_Details(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewASService(m.Client)

	testAPIEndpoint := "/as/1"
	raw, err := ioutil.ReadFile("../data/as_details.json")
	if err != nil {
		t.Error(err.Error())
	}

	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodGet)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	autonomousSystem, err := svc.Details(context.Background(), "1")
	if autonomousSystem == nil {
		t.Error("Expected AS struct. AS struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
