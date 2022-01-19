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

	server, err := invdmockserver.New(200, customerToCreate, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection("whatever", server)
	customerToCreate.Connection = conn

	customer := conn.NewCustomer()

	customerResp, err := customer.Create(&invdendpoint.CustomerRequest{Name: String("John Doe")})

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(customerResp, customerToCreate) {
		t.Fatal("Returned Customer Is Not Equal", customerToCreate.Customer, customerResp.Customer)
	}
}
