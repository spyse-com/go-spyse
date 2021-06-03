package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCertificateService_Details(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/certificate/5c157070be587becb7856643c9be75ab31726a0328a88377b0093c908a53abf5"

	raw, err := ioutil.ReadFile("../data/certificate_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	certificates, err := testClient.Certificate.Details(
		context.Background(),
		"5c157070be587becb7856643c9be75ab31726a0328a88377b0093c908a53abf5",
	)
	if certificates == nil {
		t.Error("Expected Certificate struct. Certificate struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
