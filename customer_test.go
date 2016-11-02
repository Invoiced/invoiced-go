package invdapi

import (
	"encoding/json"
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"strconv"
	"testing"
	"time"
)

type customerMetaData map[string]interface{}

func TestCustomerMetaData(t *testing.T) {
	conn := NewConnection("", false)
	m := make(customerMetaData)
	m["integration_name"] = "QBO"
	mockCustomer := conn.NewCustomer()
	mockCustomer.Id = 34
	mockCustomer.MetaData = m

	b, err := json.Marshal(mockCustomer)

	if err != nil {
		t.Fatal(err)
	}

	if string(b) != `{"id":34,"metadata":{"integration_name":"QBO"}}` {
		t.Fatal("Json is wrong", "right json =>", string(b))
	}

}

func TestCustomerCreate(t *testing.T) {

	//Set up the mock customer response
	mockCustomerResponseID := int64(1523)
	mockCustomerResponse := new(Customer)
	mockCustomerResponseData := new(invdendpoint.Customer)
	mockCustomerResponse.Customer = mockCustomerResponseData
	mockCustomerResponse.Id = mockCustomerResponseID

	//Launch our mock server
	server, err := invdmockserver.New(200, mockCustomerResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	//Establish our mock connection
	key := "test api key"
	conn := mockConnection(key, server)

	customer := conn.NewCustomer()

	nowUnix := time.Now().UnixNano()
	s := strconv.FormatInt(nowUnix, 10)

	customerToCreate := customer.NewCustomer()
	customerToCreate.Name = "Test Customer Original " + s
	customerToCreate.Id = mockCustomerResponse.Id
	mockCustomerResponse.Name = customerToCreate.Name
	//mockCustomerResponse.Connection = conn

	//Make the call to create our customer
	createdCustomer, err := customer.Create(customerToCreate)

	if err != nil {
		t.Fatal("Error Creating Customer", err)
	}

	//Customer that we wanted to create should equal the customer we created
	if !reflect.DeepEqual(createdCustomer, customerToCreate) {
		t.Fatal(createdCustomer.Customer, customerToCreate.Customer)
	}

}

func TestCustomerCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server, err := invdmockserver.New(400, mockErrorResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	custConn := conn.NewCustomer()

	customerToCreate := custConn.NewCustomer()
	customerToCreate.Email = "example@example.com"

	_, apiErr := custConn.Create(customerToCreate)

	if apiErr == nil {
		t.Fatal("Api should have errored out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), apiErr.Error()) {
		t.Fatal("Error messages do not match up", mockErrorResponse, ",", apiErr)
	}

}

func TestCustomerUpdate(t *testing.T) {
	key := "test api key"

	mockCustomerResponseID := int64(1523)
	mockCreatedTime := time.Now().UnixNano()
	mockName := "MOCK CUSTOMER"
	mockCustomerResponse := new(invdendpoint.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID
	mockCustomerResponse.Name = mockName
	mockCustomerResponse.CreatedAt = mockCreatedTime

	server, err := invdmockserver.New(200, mockCustomerResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customerToUpdate := conn.NewCustomer()

	customerToUpdate.Id = mockCustomerResponseID
	customerToUpdate.Name = "MOCK CUSTOMER"
	addressToUpdate := "7500 Rialto BLVD"
	customerToUpdate.Address1 = addressToUpdate
	mockCustomerResponse.Address1 = addressToUpdate

	apiErr := customerToUpdate.Save()

	if apiErr != nil {
		t.Fatal("Error Updating Customer", apiErr)
	}

	if !reflect.DeepEqual(mockCustomerResponse, customerToUpdate.Customer) {
		t.Fatal("Updated Customers Do Not Match Up")
	}

}

func TestCustomerUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	server, err := invdmockserver.New(401, mockErrorResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()
	customer.Name = "Parag Patel"
	customer.Id = 3411111
	customer.City = "Austin"

	err = customer.Save()

	if err == nil {
		t.Fatal("Error Updating Customer => ", err)
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestCustomerDelete(t *testing.T) {

	key := "api key"

	mockCustomerResponse := ""
	mockCustomerID := int64(2341)

	server, err := invdmockserver.New(204, mockCustomerResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()

	customer.Id = mockCustomerID

	err = customer.Delete()

	if err != nil {
		t.Fatal("Error occured deleting customer")
	}

}

func TestCustomerDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockCustomerID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", false)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()

	customer.Id = mockCustomerID

	err = customer.Delete()

	if err == nil {
		t.Fatal("Error Should Have Been Raised")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestCustomerList(t *testing.T) {

	key := "test api key"

	var mockCustomersResponse invdendpoint.Customers
	mockCustomerResponseID := int64(1523)
	mockCustomerResponse := new(invdendpoint.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID
	mockCustomerResponse.Name = "Mock Customer"
	mockCustomerResponse.Address1 = "23 Wayne street"
	mockCustomerResponse.City = "Austin"
	mockCustomerResponse.Country = "USA"
	mockCustomerResponse.CreatedAt = time.Now().UnixNano()
	mockCustomerResponse.Number = "CUST-21312"

	mockCustomersResponse = append(mockCustomersResponse, *mockCustomerResponse)

	server, err := invdmockserver.New(200, mockCustomersResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()

	retrievedCustomer, err := customer.ListCustomerByNumber("CUST-21312")

	if err != nil {
		t.Fatal("Error Creating Customer", err)
	}

	if !reflect.DeepEqual(retrievedCustomer.Customer, mockCustomerResponse) {
		t.Fatal("Retrieved Customer does not match the mock customer retrievedCustomer => ", retrievedCustomer.Customer, ", mockCustomer => ", mockCustomerResponse)
	}

}

func TestCustomerListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockCustomerNumber := "CUST-33442"

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)

	if err != nil {

		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()

	_, err = customer.ListCustomerByNumber(mockCustomerNumber)

	if err == nil {
		t.Fatal("Error occured deleting customer")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}
