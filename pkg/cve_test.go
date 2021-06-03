package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCVEService_Details(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/cve/CVE-2004-2343"

	raw, err := ioutil.ReadFile("../data/cve_details.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	cves, err := testClient.CVE.Details(context.Background(), "CVE-2004-2343")
	if cves == nil {
		t.Error("Expected CVE struct. CVE struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
