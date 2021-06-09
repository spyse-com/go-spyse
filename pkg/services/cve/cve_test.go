package cve

//func TestCVEService_Details(t *testing.T) {
//	spyse.setup()
//	defer spyse.teardown()
//	testAPIEndpoint := "/cve/CVE-2004-2343"
//
//	raw, err := ioutil.ReadFile("../data/cve_details.json")
//	if err != nil {
//		t.Error(err.Error())
//	}
//	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
//		spyse.testMethod(t, r, http.MethodGet)
//		spyse.testRequestURL(t, r, testAPIEndpoint)
//		fmt.Fprint(w, string(raw))
//	})
//
//	cves, err := spyse.testClient.CVE.Details(context.Background(), "CVE-2004-2343")
//	if cves == nil {
//		t.Error("Expected CVE struct. CVE struct is nil")
//	}
//	if err != nil {
//		t.Errorf("Error given: %s", err)
//	}
//}
