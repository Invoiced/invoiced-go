package invdapi

import (
	"reflect"
	"testing"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestMockConnection(t *testing.T) {
	customerData := new(invdendpoint.Customer)
	customerToCreate := new(Customer)
	customerToCreate.Customer = customerData
	customerToCreate.Name = "John Doe"

	server, err := invdmockserver.New(200, customerToCreate, "json", true)
	if err != nil {
		t.Fatal(err)
	}
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
