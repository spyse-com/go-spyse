package spyse

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHistoryService_DNS(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewHistoryService(m.Client)

	testAPIEndpoint := "/history/dns/A/0.07.aa0d.ip4.static.sl-reverse.com"

	raw, err := ioutil.ReadFile("../data/dns_history.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodGet)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	ips, err := svc.DNS(context.Background(), "0.07.aa0d.ip4.static.sl-reverse.com", "A", 1, 0)
	if ips == nil {
		t.Error("Expected DNS history struct. DNS history struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestHistoryService_Whois(t *testing.T) {
	m := setup()
	defer m.teardown()

	svc := NewHistoryService(m.Client)

	testAPIEndpoint := "/history/domain-whois/101reasonsfilm.com"

	raw, err := ioutil.ReadFile("../data/whois_history.json")
	if err != nil {
		t.Error(err.Error())
	}
	m.TestMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		m.testMethod(t, r, http.MethodGet)
		m.testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	ips, err := svc.Whois(context.Background(), "101reasonsfilm.com", 1, 0)
	if ips == nil {
		t.Error("Expected WHOIS history struct. WHOIS history struct is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
