package invdapi

import (
	"encoding/json"
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"reflect"
	"strconv"
	"testing"
	"time"
)

type customerMetaData struct {
	IntegrationName string `json:"integration_name,omitempty"`
}

func TestCustomerMetaData(t *testing.T) {
	conn := NewConnection("", false)
	m := new(customerMetaData)
	m.IntegrationName = "QBO"
	mockCustomer := conn.NewCustomer()
	mockCustomer.Id = 34
	mockCustomer.MetaData = m

	b, err := json.Marshal(mockCustomer)

	if err != nil {
		panic(err)
	}

	if string(b) != `{"id":34,"metadata":{"integration_name":"QBO"}}` {
		t.Fatal("Json is wrong")
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
	server := mockServer(200, mockCustomerResponse)
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
	createdCustomer, apiErr := customer.Create(customerToCreate)

	if apiErr != nil {
		t.Fatal("Error Creating Customer", apiErr)
	}

	//Customer that we wanted to create should equal the customer we created
	if !reflect.DeepEqual(createdCustomer, customerToCreate) {
		t.Fatal(createdCustomer.Customer, customerToCreate.Customer)
	}

}

func TestCustomerCreateError(t *testing.T) {
	// key := "test api key"
	// mockErrorResponse := new(APIError)
	// mockErrorResponse.Type = "invalid_request"
	// mockErrorResponse.Message = "Name is invalid"
	// mockErrorResponse.Param = "name"

	// emptyCustomer := new(Customer)

	// server := mockServer(400, mockErrorResponse)
	// defer server.Close()

	// conn := mockConnection(key, server)
	// emptyCustomer.Connection = conn

	// customerToCreate := new(invdendpoint.Customer)
	// customerToCreate.Email = "example@example.com"

	// _, apiErr := conn.CreateCustomer(customerToCreate)

	// if apiErr == nil {
	// 	t.Fatal("Api should have errored out")
	// }

	// if !reflect.DeepEqual(mockErrorResponse, apiErr) {
	// 	t.Fatal("Error messages do not match up")
	// }

}

// func TestCustomerUpdate(t *testing.T) {
// 	key := "test api key"

// 	mockCustomerResponseID := int64(1523)
// 	mockUpdatedTime := time.Now().UnixNano()
// 	mockCustomerResponse := new(invdendpoint.Customer)
// 	mockCustomerResponse.Id = mockCustomerResponseID
// 	mockCustomerResponse.UpdatedAt = mockUpdatedTime
// 	mockCustomerResponse.Name = "MOCK CUSTOMER"

// 	customerToUpdate := new(invdendpoint.Customer)

// 	addressToUpdate := "7500 Rialto BLVD"

// 	mockCustomerResponse.Address1 = addressToUpdate
// 	customerToUpdate.Address1 = addressToUpdate

// 	server := mockServer(200, mockCustomerResponse)
// 	defer server.Close()

// 	conn := mockConnection(key, server)

// 	updatedCustomer, apiErr := conn.UpdateCustomer(mockCustomerResponseID, customerToUpdate)

// 	if apiErr != nil {
// 		t.Fatal("Error Updating Customer", apiErr)
// 	}

// 	if !reflect.DeepEqual(mockCustomerResponse, updatedCustomer) {
// 		t.Fatal("Error messages do not match up")
// 	}

// }

// func TestCustomerUpdateError(t *testing.T) {
// 	key := "wrong api key"

// 	mockErrorResponse := new(APIError)
// 	mockErrorResponse.Type = "invalid_request"
// 	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

// 	customerID := int64(324234)
// 	customerToUpdate := new(invdendpoint.Customer)
// 	customerToUpdate.Address1 = "7500 Rialto BLVD"

// 	server := mockServer(401, mockErrorResponse)
// 	defer server.Close()

// 	conn := mockConnection(key, server)

// 	_, apiErr := conn.UpdateCustomer(customerID, customerToUpdate)

// 	if apiErr == nil {
// 		t.Fatal("Error Updating Customer", apiErr)
// 	}

// 	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
// 		t.Fatal("Error Messages Do Not Match Up")
// 	}

// }

// func TestCustomerDelete(t *testing.T) {

// 	key := "api key"

// 	mockCustomerResponse := ""
// 	mockCustomerID := int64(2341)

// 	server := mockServer(204, mockCustomerResponse)
// 	defer server.Close()

// 	conn := mockConnection(key, server)

// 	apiErr := conn.DeleteCustomer(mockCustomerID)

// 	if apiErr != nil {
// 		t.Fatal("Error occured deleting customer")
// 	}

// }

// func TestCustomerDeleteError(t *testing.T) {
// 	key := "api key"

// 	mockErrorResponse := new(APIError)
// 	mockErrorResponse.Type = "invalid_request"
// 	mockErrorResponse.Message = "You do not have permission to do that"

// 	mockCustomerID := int64(-999)

// 	server := mockServer(403, mockErrorResponse)
// 	defer server.Close()

// 	conn := mockConnection(key, server)

// 	apiErr := conn.DeleteCustomer(mockCustomerID)

// 	if apiErr == nil {
// 		t.Fatal("Error occured deleting customer")
// 	}

// 	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
// 		t.Fatal("Error Messages Do Not Match Up")
// 	}

// }

// func TestCustomerList(t *testing.T) {

// 	key := "test api key"

// 	mockCustomerResponseID := int64(1523)
// 	mockCustomerResponse := new(invdendpoint.Customer)
// 	mockCustomerResponse.Id = mockCustomerResponseID
// 	mockCustomerResponse.Name = "Mock Customer"
// 	mockCustomerResponse.Address1 = "23 Wayne street"
// 	mockCustomerResponse.City = "Austin"
// 	mockCustomerResponse.Country = "USA"
// 	mockCustomerResponse.UpdatedAt = time.Now().UnixNano()

// 	server := mockServer(200, mockCustomerResponse)
// 	defer server.Close()

// 	conn := mockConnection(key, server)

// 	createdCustomer, apiErr := conn.ListCustomer(mockCustomerResponseID)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Customer", apiErr)
// 	}

// 	if createdCustomer.Id != mockCustomerResponseID {
// 		t.Fatal("Customer was not created succesfully")
// 	}

// }

// func TestCustomerListError(t *testing.T) {
// 	key := "api key"

// 	mockErrorResponse := new(APIError)
// 	mockErrorResponse.Type = "invalid_request"
// 	mockErrorResponse.Message = "You do not have permission to do that"

// 	mockCustomerID := int64(-999)

// 	server := mockServer(403, mockErrorResponse)
// 	defer server.Close()

// 	conn := mockConnection(key, server)

// 	_, apiErr := conn.ListCustomer(mockCustomerID)

// 	if apiErr == nil {
// 		t.Fatal("Error occured deleting customer")
// 	}

// 	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
// 		t.Fatal("Error Messages Do Not Match Up")
// 	}

// }
