package invdapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func mockConnection(key string, server *httptest.Server) *Connection {
	c := new(Connection)
	c.key = key

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c.client = &http.Client{Transport: transport}
	c.url = server.URL

	return c
}

func mockServer(code int, body interface{}) *httptest.Server {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)

		bodyMarshalled, err := json.Marshal(body)

		if err != nil {
			panic(err)
		}

		fmt.Fprintln(w, string(bodyMarshalled))
	}))

	return server
}

func TestMockServer(t *testing.T) {

	customer := new(invdendpoint.Customer)
	customer.Name = "Parag Patel"

	server := mockServer(200, customer)
	defer server.Close()

	conn := mockConnection("whatever", server)

	customerResp, apiErr := conn.CreateCustomer(customer)

	if apiErr != nil {
		t.Fatal(apiErr)
	}

	if !reflect.DeepEqual(customerResp, customer) {
		t.Fatal("Returned Customer Is Not Equal")
	}
}
