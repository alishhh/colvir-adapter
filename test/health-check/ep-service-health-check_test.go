package health_check

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestEndpointHealthCheckService(t *testing.T) {
	var router = mux.NewRouter()

	router.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })
	router.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })
	//http2.Routes(routes, transport.Endpoints())
	srv := httptest.NewServer(router)
	defer srv.Close()

	for _, testcase := range []struct {
		method, url, bodyRequest, status, bodyResponse string
	}{
		{"GET", srv.URL + "/health/live", "", "200", ""},
		{"GET", srv.URL + "/health/ready", "", "200", ""},
		{"GET", srv.URL + "/health/live1", "", "404", ""},
		{"GET", srv.URL + "/health/ready1", "", "404", ""},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.bodyRequest))
		res, _ := http.DefaultClient.Do(req)
		bite, _ := ioutil.ReadAll(res.Body)
		body := strings.TrimSpace(string(bite))

		if testcase.bodyResponse != "" && testcase.bodyResponse != body {
			t.Errorf("%s %s %s: want %q, have %q", testcase.method, testcase.url, testcase.bodyRequest, testcase.bodyResponse, body)
		}

		statusCode := strconv.Itoa(res.StatusCode)
		if testcase.status != statusCode {
			t.Errorf("%s %s %s: status want %q, have %q", testcase.method, testcase.url, testcase.bodyRequest, testcase.status, statusCode)
		}
	}
}
