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

	m := new(customerMetaData)
	m.IntegrationName = "QBO"
	mockCustomer := new(invdendpoint.Customer)
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
	key := "test api key"

	mockCustomerResponseID := int64(1523)
	mockCustomerResponse := new(invdendpoint.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID

	customerToCreate := new(invdendpoint.Customer)

	nowUnix := time.Now().UnixNano()

	s := strconv.FormatInt(nowUnix, 10)

	customerToCreate.Name = "Test Customer Original " + s
	mockCustomerResponse.Name = customerToCreate.Name

	server := mockServer(200, mockCustomerResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdCustomer, apiErr := conn.CreateCustomer(customerToCreate)

	if apiErr != nil {
		t.Fatal("Error Creating Customer", apiErr)
	}

	if createdCustomer.Id != mockCustomerResponseID {
		t.Fatal("Customer was not created succesfully")
	}

}

func TestCustomerCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server := mockServer(400, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	customerToCreate := new(invdendpoint.Customer)
	customerToCreate.Email = "example@example.com"

	_, apiErr := conn.CreateCustomer(customerToCreate)

	if apiErr == nil {
		t.Fatal("Api should have errored out")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error messages do not match up")
	}

}

func TestCustomerUpdate(t *testing.T) {
	key := "test api key"

	mockCustomerResponseID := int64(1523)
	mockUpdatedTime := time.Now().UnixNano()
	mockCustomerResponse := new(invdendpoint.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID
	mockCustomerResponse.UpdatedAt = mockUpdatedTime
	mockCustomerResponse.Name = "MOCK CUSTOMER"

	customerToUpdate := new(invdendpoint.Customer)

	addressToUpdate := "7500 Rialto BLVD"

	mockCustomerResponse.Address1 = addressToUpdate
	customerToUpdate.Address1 = addressToUpdate

	server := mockServer(200, mockCustomerResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	updatedCustomer, apiErr := conn.UpdateCustomer(mockCustomerResponseID, customerToUpdate)

	if apiErr != nil {
		t.Fatal("Error Updating Customer", apiErr)
	}

	if !reflect.DeepEqual(mockCustomerResponse, updatedCustomer) {
		t.Fatal("Error messages do not match up")
	}

}

func TestCustomerUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	customerID := int64(324234)
	customerToUpdate := new(invdendpoint.Customer)
	customerToUpdate.Address1 = "7500 Rialto BLVD"

	server := mockServer(401, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.UpdateCustomer(customerID, customerToUpdate)

	if apiErr == nil {
		t.Fatal("Error Updating Customer", apiErr)
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestCustomerDelete(t *testing.T) {

	key := "api key"

	mockCustomerResponse := ""
	mockCustomerID := int64(2341)

	server := mockServer(204, mockCustomerResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteCustomer(mockCustomerID)

	if apiErr != nil {
		t.Fatal("Error occured deleting customer")
	}

}

func TestCustomerDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockCustomerID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteCustomer(mockCustomerID)

	if apiErr == nil {
		t.Fatal("Error occured deleting customer")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestCustomerList(t *testing.T) {

	key := "test api key"

	mockCustomerResponseID := int64(1523)
	mockCustomerResponse := new(invdendpoint.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID
	mockCustomerResponse.Name = "Mock Customer"
	mockCustomerResponse.Address1 = "23 Wayne street"
	mockCustomerResponse.City = "Austin"
	mockCustomerResponse.Country = "USA"
	mockCustomerResponse.UpdatedAt = time.Now().UnixNano()

	server := mockServer(200, mockCustomerResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdCustomer, apiErr := conn.ListCustomer(mockCustomerResponseID)

	if apiErr != nil {
		t.Fatal("Error Creating Customer", apiErr)
	}

	if createdCustomer.Id != mockCustomerResponseID {
		t.Fatal("Customer was not created succesfully")
	}

}

func TestCustomerListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockCustomerID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.ListCustomer(mockCustomerID)

	if apiErr == nil {
		t.Fatal("Error occured deleting customer")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

// func TestCustomerCRUD(t *testing.T) {
// 	key := "test api key"
// 	//Create Customer

// 	mockCustomerResponse := new(invdendpoint.Customer)
// 	mockCustomerResponse.Id = 1

// 	customerToCreate := new(invdendpoint.Customer)

// 	nowUnix := time.Now().UnixNano()

// 	s := strconv.FormatInt(nowUnix, 10)

// 	customerToCreate.Name = "Test Customer Original " + s
// 	mockCustomerResponse.Name = customerToCreate.Name

// 	server := mockServer(200, mockCustomerResponse)

// 	conn := mockConnection(key, server)

// 	createdCustomer, apiErr := conn.CreateCustomer(customerToCreate)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Customer", apiErr)
// 	}

// 	if createdCustomer.Id <= 0 {
// 		t.Fatal("Customer was not created succesfully")
// 	}

// 	server.Close()
// 	//Update Customer

// 	createdCustomer.Name = "Test Customer Updated " + s

// 	mockCustomerResponse.Name = customerToCreate.Name

// 	server = mockServer(200, mockCustomerResponse)
// 	conn = mockConnection(key, server)

// 	apiErr = conn.UpdateCustomer(createdCustomer.Id, createdCustomer)

// 	if apiErr != nil {
// 		t.Fatal("Error Updating Customer => ", apiErr)
// 	}

// 	//Retrieve Customer

// 	retrievedCustomer, apiErr := conn.ListCustomer(createdCustomer.Id)

// 	if apiErr != nil {
// 		t.Fatal("Could Not Retrieve Customer")
// 	}

// 	if retrievedCustomer.Id != createdCustomer.Id {
// 		t.Fatal("Customer Did Not Retrieve Correctly")
// 	}

// 	if retrievedCustomer.Name != createdCustomer.Name {
// 		t.Fatal("Customer Did Not Update Correctly")
// 	}

// 	//Delete Customer

// 	apiErr = conn.DeleteCustomer(retrievedCustomer.Id)

// 	if apiErr != nil {
// 		t.Fatal("Error Deleting Customer => ", apiErr)
// 	}

// 	//Make Sure Customer is Deleted

// 	_, apiErr = conn.ListCustomer(retrievedCustomer.Id)

// 	if apiErr == nil {
// 		t.Fatal("Customer ", retrievedCustomer.Id, "Should Be Deleted!")
// 	}

// }

// func createMockCustomer(t *testing.T, offset int64) *invdendpoint.Customer {
// 	conn := NewConnection(apikey)

// 	//Create Customer
// 	customerToCreate := new(invdendpoint.Customer)

// 	nowUnix := time.Now().UnixNano() + offset

// 	s := strconv.FormatInt(nowUnix, 10)

// 	customerToCreate.Name = "Mock-Customer-" + s
// 	customerToCreate.City = "Mock-Customer"

// 	customer, apiErr := conn.CreateCustomer(customerToCreate)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Mock Customer =>", apiErr)
// 	}

// 	return customer

// }

// func deleteMockCustomer(t *testing.T, customer invdendpoint.Customer) {

// 	id := customer.Id
// 	conn := NewConnection(apikey)

// 	apiErr := conn.DeleteCustomer(id)

// 	if apiErr != nil {
// 		t.Fatal("Error Deleting Mock Customer, ID => ", id, " =>", apiErr)
// 	}

// }

// func deleteAllCustomers(t *testing.T) {
// 	conn := NewConnection(apikey)

// 	var customerIDs []int64

// 	customers, apiErr := conn.ListAllCustomersAuto(nil, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error: Getting All Customers Auto", apiErr)
// 	}

// 	for _, customer := range *customers {
// 		customerIDs = append(customerIDs, customer.Id)
// 	}

// 	for _, customerID := range customerIDs {
// 		apiErr = conn.DeleteCustomer(customerID)

// 		if apiErr != nil {
// 			t.Fatal("Error Deleting A Customer with ID => ", customerID, apiErr)
// 		}

// 	}

// }

// func TestDeleteAllCustomers(t *testing.T) {
// 	deleteAllCustomers(t)
// }

// func TestGetAllCustomersAuto(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	customers, apiErr := conn.ListAllCustomersAuto(nil, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error: Getting All Customers Auto", apiErr)
// 	}

// 	customerCount, apiErr := conn.CountCustomer()

// 	if apiErr != nil {
// 		t.Fatal("Error: Getting Customer Count", apiErr)
// 	}

// 	if int64(len(*customers)) != customerCount {
// 		t.Fatal("Error: Number of Customers Returned Should Equal The Customer Count, len(*customers) => ", len(*customers), ", customerCount => ", customerCount)
// 	}

// }

// func TestListCustomersByName(t *testing.T) {
// 	conn := NewConnection(apikey)
// 	customerName := "Cool Cars"
// 	customers, apiErr := conn.ListCustomersByName(customerName)
// 	if apiErr != nil {
// 		t.Fatal(apiErr)
// 	}

// 	log.Println(customers)

// }

// func TestListCustomersByNumber(t *testing.T) {
// 	conn := NewConnection(apikey)
// 	customerNumber := "QBO-"
// 	customer, apiErr := conn.ListCustomerByNumber(customerNumber)
// 	if apiErr != nil {
// 		t.Fatal(apiErr)
// 	}

// 	log.Println(customer)

// }

// func TestGetAllCustomers(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	customers, nextEndPoint, apiErr := conn.ListAllCustomers(nil, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error Getting All Customers", apiErr)
// 	}

// 	customerCount, apiErr := conn.CountCustomer()

// 	if apiErr != nil {
// 		t.Fatal("Error Getting Customer Count", apiErr)
// 	}

// 	if customerCount > 100 && nextEndPoint == "" {
// 		t.Fatal("Customer Count Is Greater Than 100, So The Next EndPoint Should Not Be Empty")
// 	}

// 	if len(*customers) > 100 {
// 		t.Fatal("Customer Returned Should Less Than OR Equal To 100")
// 	}

// }

// func TestListAllCustomersFiltered(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	var customerIDs []int64

// 	numberOfTestCustomers := 56

// 	for i := 0; i < numberOfTestCustomers; i++ {
// 		customer := new(invdendpoint.Customer)
// 		s := strconv.FormatInt(time.Now().Unix()+int64(i), 10)
// 		customer.Name = "Test GetALLCustomersFiltered " + s
// 		customer.City = "TestGetAllCustomersFiltered"
// 		customerResponse, apiErr := conn.CreateCustomer(customer)

// 		if apiErr != nil {
// 			t.Fatal("Error Creating Customer", apiErr)
// 		}

// 		customerIDs = append(customerIDs, customerResponse.Id)

// 	}

// 	filter := invdendpoint.NewFilter()
// 	filter.Set("city", "TestGetAllCustomersFiltered")

// 	customers, apiErr := conn.ListAllCustomersAuto(filter, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error Getting All Customers Auto ", apiErr)
// 	}

// 	if len(*customers) != numberOfTestCustomers {
// 		t.Fatal("The Correct Amount of Customers Were Not Filtered", "")
// 	}

// 	//Delete Customers

// 	for _, customerID := range customerIDs {
// 		apiErr = conn.DeleteCustomer(customerID)

// 		if apiErr != nil {
// 			t.Fatal("Error Deleting A Customer with ID => ", customerID, apiErr)
// 		}

// 	}

// }

// func TestGetAllCustomersFilteredSorted(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	var customerIDs []int64

// 	//Create Customer A
// 	customer := new(invdendpoint.Customer)
// 	customer.Name = "A" + strconv.FormatInt(time.Now().Unix()+int64(0), 10)
// 	customer.City = "TestGetAllCustomersFilteredSorted"
// 	customerResponse, apiErr := conn.CreateCustomer(customer)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Customer", apiErr)
// 	}
// 	customerIDs = append(customerIDs, customerResponse.Id)

// 	//Create Customer H
// 	customer = new(invdendpoint.Customer)
// 	customer.Name = "H" + strconv.FormatInt(time.Now().Unix()+int64(1), 10)
// 	customer.City = "TestGetAllCustomersFilteredSorted"
// 	customerResponse, apiErr = conn.CreateCustomer(customer)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Customer", apiErr)
// 	}
// 	customerIDs = append(customerIDs, customerResponse.Id)

// 	//Create Customer L
// 	customer = new(invdendpoint.Customer)
// 	customer.Name = "L" + strconv.FormatInt(time.Now().Unix()+int64(2), 10)
// 	customer.City = "TestGetAllCustomersFilteredSorted"
// 	customerResponse, apiErr = conn.CreateCustomer(customer)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Customer", apiErr)
// 	}
// 	customerIDs = append(customerIDs, customerResponse.Id)

// 	//Create Customer C
// 	customer = new(invdendpoint.Customer)
// 	customer.Name = "C" + strconv.FormatInt(time.Now().Unix()+int64(3), 10)
// 	customer.City = "TestGetAllCustomersFilteredSorted"
// 	customerResponse, apiErr = conn.CreateCustomer(customer)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Customer", apiErr)
// 	}
// 	customerIDs = append(customerIDs, customerResponse.Id)

// 	filter := invdendpoint.NewFilter()
// 	filter.Set("city", "TestGetAllCustomersFilteredSorted")

// 	sort := invdendpoint.NewSort()

// 	sort.Set("name", invdendpoint.DESC)

// 	customers, _, apiErr := conn.ListAllCustomers(filter, sort)

// 	if apiErr != nil {
// 		t.Fatal("Error Getting All Customers Auto ", apiErr)
// 	}

// 	if (*customers)[0].Name[0:1] != "L" {
// 		t.Fatal("Sort DESC Failed With 'L'")
// 	}

// 	if (*customers)[1].Name[0:1] != "H" {
// 		t.Fatal("Sort DESC Failed With 'H'")
// 	}
// 	if (*customers)[2].Name[0:1] != "C" {
// 		t.Fatal("Sort DESC Failed With 'C'")
// 	}

// 	if (*customers)[3].Name[0:1] != "A" {
// 		t.Fatal("Sort DESC Failed With 'A'")
// 	}

// 	//Delete Customers

// 	for _, customerID := range customerIDs {
// 		apiErr = conn.DeleteCustomer(customerID)

// 		if apiErr != nil {
// 			t.Fatal("Error Deleting A Customer with ID => ", customerID, apiErr)
// 		}

// 	}

// }

// func TestCustomerCount(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	_, apiErr := conn.CountCustomer()

// 	if apiErr != nil {
// 		t.Fatal("apiError Should Be Empty")
// 	}

// }
