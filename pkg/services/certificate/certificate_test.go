package certificate

//func TestCertificateService_Details(t *testing.T) {
//	spyse.setup()
//	defer spyse.teardown()
//	testAPIEndpoint := "/certificate/5c157070be587becb7856643c9be75ab31726a0328a88377b0093c908a53abf5"
//
//	raw, err := ioutil.ReadFile("../data/certificate_details.json")
//	if err != nil {
//		t.Error(err.Error())
//	}
//	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
//		spyse.testMethod(t, r, http.MethodGet)
//		spyse.testRequestURL(t, r, testAPIEndpoint)
//		fmt.Fprint(w, string(raw))
//	})
//
//	certificates, err := spyse.testClient.Certificate.Details(
//		context.Background(),
//		"5c157070be587becb7856643c9be75ab31726a0328a88377b0093c908a53abf5",
//	)
//	if certificates == nil {
//		t.Error("Expected Certificate struct. Certificate struct is nil")
//	}
//	if err != nil {
//		t.Errorf("Error given: %s", err)
//	}
//}
