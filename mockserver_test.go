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
	customerData := new(invdendpoint.Customer)
	customerToCreate := new(Customer)
	customerToCreate.Customer = customerData
	customerToCreate.Name = "John Doe"

	server := mockServer(200, customerToCreate)
	defer server.Close()

	conn := mockConnection("whatever", server)
	customerToCreate.Connection = conn

	customer := conn.NewCustomer()

	customerResp, apiErr := customer.Create(customerToCreate)

	if apiErr != nil {
		t.Fatal(apiErr)
	}

	if !reflect.DeepEqual(customerResp, customerToCreate) {
		t.Fatal("Returned Customer Is Not Equal", customerToCreate.Customer, customerResp.Customer)
	}
}
