package invdapi

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestCustomerMetadata(t *testing.T) {
	conn := NewConnection("", false)
	m := make(map[string]interface{})
	m["integration_name"] = "QBO"
	mockCustomer := conn.NewCustomer()
	mockCustomer.Id = 34
	mockCustomer.Metadata = m

	b, err := json.Marshal(mockCustomer)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != `{"id":34,"metadata":{"integration_name":"QBO"}}` {
		t.Fatal("Json is wrong", "right json =>", string(b))
	}
}

func TestCustomerCreate(t *testing.T) {
	// Set up the mock customer response
	mockCustomerResponseID := int64(1523)
	mockCustomerResponse := new(Customer)
	mockCustomerResponseData := new(invdendpoint.Customer)
	mockCustomerResponse.Customer = mockCustomerResponseData
	mockCustomerResponse.Id = mockCustomerResponseID

	// Launch our mock server
	server, err := invdmockserver.New(200, mockCustomerResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	// Establish our mock connection
	key := "test api key"
	conn := mockConnection(key, server)

	customer := conn.NewCustomer()

	nowUnix := time.Now().UnixNano()
	s := strconv.FormatInt(nowUnix, 10)

	customerToCreate := customer.NewCustomer()
	customerToCreate.Name = "Test Customer Original " + s
	customerToCreate.Id = mockCustomerResponse.Id
	mockCustomerResponse.Name = customerToCreate.Name
	// mockCustomerResponse.Connection = conn

	// Make the call to create our customer
	createdCustomer, err := customer.Create(customerToCreate)
	if err != nil {
		t.Fatal("Error Creating Customer", err)
	}

	// Customer that we wanted to create should equal the customer we created
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

func TestCustomer_List(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Customers
	mockResponseId := int64(1523)
	mockNumber := "INV-3421"
	mockResponse := new(invdendpoint.Customer)
	mockResponse.Id = mockResponseId
	mockResponse.Number = mockNumber
	mockResponse.PaymentTerms = "NET15"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity := conn.NewCustomer()

	entityResp, nextEndpoint, err := entity.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entityResp[0].Customer, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestCustomer_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Customer)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewCustomer()

	retrievedPayment, err := entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.Customer, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestCustomer_GetBalance(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Balance)
	mockResponse.TotalOutstanding = 1

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewCustomer()

	retrievedItem, err := entity.GetBalance()
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if retrievedItem.TotalOutstanding != 1 {
		t.Fatal("Error messages do not match up")
	}
}

func TestCustomer_SendStatementEmail(t *testing.T) {
	key := "test api key"

	var mockEmailResponse [1]invdendpoint.EmailResponse

	mockResponse := new(invdendpoint.EmailResponse)
	mockResponse.Id = 1
	mockResponse.Message = "hello test"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockEmailResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockEmailResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewCustomer()

	sendResponse, err := subjectEntity.SendStatementEmail(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello test" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestCustomer_SendStatementText(t *testing.T) {
	key := "test api key"

	var mockTextResponse [1]invdendpoint.TextResponse

	mockResponse := new(invdendpoint.TextResponse)
	mockResponse.Id = "abcdef"
	mockResponse.Message = "hello text"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockTextResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockTextResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewCustomer()

	sendResponse, err := subjectEntity.SendStatementText(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello text" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestCustomer_SendStatementLetter(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.LetterResponse)
	mockResponse.Id = "abcdef"
	mockResponse.State = "queued"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewCustomer()

	sendResponse, err := subjectEntity.SendStatementLetter(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse.State != "queued" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestCustomer_CreateContact(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Contact)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.CreateContact(conn.NewContact())
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_RetrieveContact(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Contact)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.RetrieveContact(1234)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_UpdateContact(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Contact)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example 2"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity := defaultEntity.NewContact()
	subjectEntity.Id = int64(1234)
	subjectEntity, err = defaultEntity.UpdateContact(subjectEntity)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example 2" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_ListAllContacts(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Contacts
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Contact)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "Mock Contact"
	mockResponse.Address1 = "23 Wayne street"
	mockResponse.City = "Austin"
	mockResponse.Country = "USA"
	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.ListAllContacts()
	if err != nil {
		t.Fatal("Error with contact", err)
	}

	if subjectEntity[0].Name != "Mock Contact" {
		t.Fatal("Retrieval not correct")
	}
}

func TestCustomer_DeleteContact(t *testing.T) {
	key := "api key"

	server, err := invdmockserver.New(204, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()
	contact := customer.NewContact()
	contact.Id = int64(1234)

	err = customer.DeleteContact(int64(1234))

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}

func TestCustomer_CreatePaymentSource_Card(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Card)
	mockResponse.Id = int64(1234)
	mockResponse.Last4 = "4242"
	mockResponse.Object = "card"
	mockResponse.Brand = "Visa"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	intermediate := conn.NewPaymentSource()
	subjectEntity, err := defaultEntity.CreatePaymentSource(intermediate)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity == nil {
		t.Fatal("subjectEntity does not exist", err)
	}

	if subjectEntity.Brand != "Visa" {
		t.Fatal("Did not instantiate correctly", err)
	}
}

func TestCustomer_CreatePaymentSource_Acct(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.BankAccount)
	mockResponse.Id = int64(1234)
	mockResponse.Last4 = "4242"
	mockResponse.Object = "bank_account"
	mockResponse.Verified = true

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	intermediate := conn.NewPaymentSource()
	subjectEntity, err := defaultEntity.CreatePaymentSource(intermediate)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity == nil {
		t.Fatal("subjectEntity does not exist", err)
	}

	if !subjectEntity.Verified {
		t.Fatal("Did not instantiate correctly", err)
	}
}

func TestCustomer_ListAllPaymentSources(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.PaymentSources

	mockResponseCard := new(invdendpoint.PaymentSource)
	mockResponseCard.Object = "card"

	mockResponseAcct := new(invdendpoint.PaymentSource)
	mockResponseAcct.Object = "bank_account"

	mockResponses = append(mockResponses, *mockResponseCard)
	mockResponses = append(mockResponses, *mockResponseAcct)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.ListAllPaymentSources()
	if err != nil {
		t.Fatal("Error with source: ", err)
	}

	if subjectEntity[0].Card.Object != "card" {
		t.Fatal("Error with operation")
	}

	if subjectEntity[1].BankAccount.Object != "bank_account" {
		t.Fatal("Error with operation")
	}
}

func TestCustomer_DeleteCard(t *testing.T) {
	key := "api key"

	server, err := invdmockserver.New(204, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()

	err = customer.DeleteCard(int64(1234))

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}

func TestCustomer_DeleteBankAccount(t *testing.T) {
	key := "api key"

	server, err := invdmockserver.New(204, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()
	err = customer.DeleteBankAccount(int64(1234))

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}

func TestCustomer_CreatePendingLineItem(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.PendingLineItem)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.CreatePendingLineItem(conn.NewPendingLineItem())
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_RetrievePendingLineItem(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.PendingLineItem)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.RetrievePendingLineItem(1234)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_UpdatePendingLineItem(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.PendingLineItem)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example 2"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity := defaultEntity.NewPendingLineItem()
	subjectEntity.Id = int64(1234)
	subjectEntity, err = defaultEntity.UpdatePendingLineItem(subjectEntity)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example 2" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_DeletePendingLineItem(t *testing.T) {
	key := "api key"

	server, err := invdmockserver.New(204, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	customer := conn.NewCustomer()
	contact := customer.NewPendingLineItem()
	contact.Id = int64(1234)

	err = customer.DeletePendingLineItem(int64(1234))

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}

func TestCustomer_ListAllPendingLineItems(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.PendingLineItems
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.PendingLineItem)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "Mock Pli"

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.ListAllPendingLineItems()
	if err != nil {
		t.Fatal("Error with pli", err)
	}

	if subjectEntity[0].Name != "Mock Pli" {
		t.Fatal("Retrieval not correct")
	}
}

func TestCustomer_RetrieveNotes(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Notes
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Note)
	mockResponse.Id = mockResponseId
	mockResponse.Notes = "Mock Note"

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.RetrieveNotes()
	if err != nil {
		t.Fatal("Error with note", err)
	}

	if subjectEntity[0].Notes != "Mock Note" {
		t.Fatal("Retrieval not correct")
	}
}

func TestCustomer_TriggerInvoice(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.TriggerInvoice()
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_ConsolidateInvoices(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewCustomer()
	subjectEntity, err := defaultEntity.ConsolidateInvoices()
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}
