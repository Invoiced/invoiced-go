package invoiced

import (
	"reflect"
	"testing"
	"github.com/Invoiced/invoiced-go/v2/invdmockserver"
)

func TestMockConnection(t *testing.T) {
	customerData := new(Customer)

	server, err := invdmockserver.New(200, customerData, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi("whatever", server)

	customerResp := new(Customer)
	name := "John Doe"
	err = client.Create("/test", &CustomerRequest{Name: &name}, customerResp)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(customerResp, customerData) {
		t.Fatal("Returned Customer Is Not Equal", customerData, customerResp)
	}
}
