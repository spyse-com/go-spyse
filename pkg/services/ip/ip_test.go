//nolint:dupl
package ip

//func TestIPService_Details(t *testing.T) {
//	spyse.setup()
//	defer spyse.teardown()
//	testAPIEndpoint := "/ip/8.8.8.8"
//
//	raw, err := ioutil.ReadFile("../data/ip_details.json")
//	if err != nil {
//		t.Error(err.Error())
//	}
//	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
//		spyse.testMethod(t, r, http.MethodGet)
//		spyse.testRequestURL(t, r, testAPIEndpoint)
//		fmt.Fprint(w, string(raw))
//	})
//
//	ips, err := spyse.testClient.IP.Details(context.Background(), "8.8.8.8")
//	if ips == nil {
//		t.Error("Expected IP struct. IP struct is nil")
//	}
//	if err != nil {
//		t.Errorf("Error given: %s", err)
//	}
//}
//
//func TestIPService_Search(t *testing.T) {
//	spyse.setup()
//	defer spyse.teardown()
//	testAPIEndpoint := "/ip/search"
//
//	raw, err := ioutil.ReadFile("../data/ip_details.json")
//	if err != nil {
//		t.Error(err.Error())
//	}
//	spyse.testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
//		spyse.testMethod(t, r, http.MethodPost)
//		spyse.testRequestURL(t, r, testAPIEndpoint)
//		fmt.Fprint(w, string(raw))
//	})
//	var params = []map[string]spyse.SearchOption{
//		{
//			"cidr": spyse.SearchOption{
//				Operator: "eq",
//				Value:    "8.8.8.8/32",
//			},
//		},
//	}
//	ips, err := spyse.testClient.IP.Search(context.Background(), params, 1, 0)
//	if ips == nil {
//		t.Error("Expected IP struct. IP struct is nil")
//	}
//	if err != nil {
//		t.Errorf("Error given: %s", err)
//	}
//}
