package as

//func TestASService_Details(t *testing.T) {
//	spyse.setup()
//	defer spyse.teardown()
//	testAPIEndpoint := "/as/1"
//
//	raw, err := ioutil.ReadFile("../data/as_details.json")
//	if err != nil {
//		t.Error(err.Error())
//	}
//	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
//		spyse.testMethod(t, r, http.MethodGet)
//		spyse.testRequestURL(t, r, testAPIEndpoint)
//		fmt.Fprint(w, string(raw))
//	})
//
//	autonomousSystem, err := spyse.testClient.AS.Details(context.Background(), "1")
//	if autonomousSystem == nil {
//		t.Error("Expected AS struct. AS struct is nil")
//	}
//	if err != nil {
//		t.Errorf("Error given: %s", err)
//	}
//}
