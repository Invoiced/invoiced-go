package customer

import (
	"encoding/json"
	"github.com/Invoiced/invoiced-go"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestCustomerMetadata(t *testing.T) {
	m := make(map[string]interface{})
	m["integration_name"] = "QBO"
	mockCustomer := &invoiced.CustomerRequest{
		Metadata: &m,
	}

	b, err := json.Marshal(mockCustomer)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != `{"metadata":{"integration_name":"QBO"}}` {
		t.Fatal("Json is wrong", "right json =>", string(b))
	}
}

func TestCustomerCreate(t *testing.T) {
	// Set up the mock customer response
	mockCustomerResponseID := int64(1523)
	mockCustomerResponse := new(invoiced.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID

	// Launch our mock server
	server, err := invdmockserver.New(200, mockCustomerResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	// Establish our mock connection
	key := "test api key"
	client := Client{invoiced.NewMockApi(key, server)}

	nowUnix := time.Now().UnixNano()
	s := strconv.FormatInt(nowUnix, 10)

	mockCustomerResponse.Name = "Test Api Original " + s

	// Make the call to create our customer
	_, err = client.Create(&invoiced.CustomerRequest{Name: invoiced.String("Test Api Original " + s)})
	if err != nil {
		t.Fatal("Error Creating Api", err)
	}
}

func TestCustomerCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server, err := invdmockserver.New(400, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.Create(&invoiced.CustomerRequest{Email: invoiced.String("example@example.com")})

	if err == nil {
		t.Fatal("Api should have errored out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error messages do not match up", mockErrorResponse, ",", err)
	}
}

func TestCustomerUpdate(t *testing.T) {
	key := "test api key"

	mockCustomerResponseID := int64(1523)
	mockCreatedTime := time.Now().UnixNano()
	mockName := "MOCK CUSTOMER"
	mockCustomerResponse := new(invoiced.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID
	mockCustomerResponse.Name = mockName
	mockCustomerResponse.CreatedAt = mockCreatedTime

	server, err := invdmockserver.New(200, mockCustomerResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	addressToUpdate := "7500 Rialto BLVD"
	mockCustomerResponse.Address1 = addressToUpdate

	resp, err := client.Update(mockCustomerResponseID, &invoiced.CustomerRequest{
		Name:     invoiced.String("MOCK CUSTOMER"),
		Address1: invoiced.String(addressToUpdate),
	})

	if err != nil {
		t.Fatal("Error Updating Api", err)
	}

	if !reflect.DeepEqual(mockCustomerResponse, resp) {
		t.Fatal("Updated Customers Do Not Match Up")
	}
}

func TestCustomerUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	server, err := invdmockserver.New(401, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.Update(3411111, &invoiced.CustomerRequest{
		Name: invoiced.String("Parag Patel"),
		City: invoiced.String("Austin"),
	})

	if err == nil {
		t.Fatal("Error Updating Api => ", err)
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

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.Delete(mockCustomerID)

	if err != nil {
		t.Fatal("Error occurred deleting customer")
	}
}

func TestCustomerDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockCustomerID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", false)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.Delete(mockCustomerID)

	if err == nil {
		t.Fatal("Error Should Have Been Raised")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestCustomerList(t *testing.T) {
	key := "test api key"

	var mockCustomersResponse invoiced.Customers
	mockCustomerResponseID := int64(1523)
	mockCustomerResponse := new(invoiced.Customer)
	mockCustomerResponse.Id = mockCustomerResponseID
	mockCustomerResponse.Name = "Mock Api"
	mockCustomerResponse.Address1 = "23 Wayne street"
	mockCustomerResponse.City = "Austin"
	mockCustomerResponse.Country = "USA"
	mockCustomerResponse.CreatedAt = time.Now().UnixNano()
	mockCustomerResponse.Number = "CUST-21312"

	mockCustomersResponse = append(mockCustomersResponse, mockCustomerResponse)

	server, err := invdmockserver.New(200, mockCustomersResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedCustomer, err := client.ListCustomerByNumber("CUST-21312")
	if err != nil {
		t.Fatal("Error Creating Api", err)
	}

	if !reflect.DeepEqual(retrievedCustomer, mockCustomerResponse) {
		t.Fatal("Retrieved Api does not match the mock customer retrievedCustomer => ", retrievedCustomer, ", mockCustomer => ", mockCustomerResponse)
	}
}

func TestCustomerListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockCustomerNumber := "CUST-33442"

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.ListCustomerByNumber(mockCustomerNumber)

	if err == nil {
		t.Fatal("Error occured deleting customer")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestCustomer_List(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Customers
	mockResponseId := int64(1523)
	mockNumber := "INV-3421"
	mockResponse := new(invoiced.Customer)
	mockResponse.Id = mockResponseId
	mockResponse.Number = mockNumber
	mockResponse.PaymentTerms = "NET15"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	entityResp, nextEndpoint, err := client.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entityResp[0], mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestCustomer_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Customer)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedPayment, err := client.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestCustomer_GetBalance(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Balance)
	mockResponse.TotalOutstanding = 1

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedItem, err := client.GetBalance(1234)
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if retrievedItem.TotalOutstanding != 1 {
		t.Fatal("Error messages do not match up")
	}
}

func TestCustomer_SendStatementEmail(t *testing.T) {
	key := "test api key"

	server, err := invdmockserver.New(200, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.SendStatementEmail(1234, nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

}

func TestCustomer_SendStatementText(t *testing.T) {
	key := "test api key"

	var mockTextResponse [1]invoiced.TextMessage

	mockResponse := new(invoiced.TextMessage)
	mockResponse.Id = "abcdef"
	mockResponse.Message = "hello text"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockTextResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockTextResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	sendResponse, err := client.SendStatementText(1234, nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello text" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestCustomer_SendStatementLetter(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Letter)
	mockResponse.Id = "abcdef"
	mockResponse.State = "queued"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	sendResponse, err := client.SendStatementLetter(1234, nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse.State != "queued" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestCustomer_CreateContact(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Contact)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.CreateContact(1234, &invoiced.ContactRequest{Name: invoiced.String("entity example")})
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_RetrieveContact(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Contact)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.RetrieveContact(1234, 456)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_UpdateContact(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Contact)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example 2"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.UpdateContact(1234, 456, &invoiced.ContactRequest{Name: invoiced.String("entity example 2")})

	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example 2" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_ListAllContacts(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Contacts
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Contact)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "Mock Contact"
	mockResponse.Address1 = invoiced.String("23 Wayne street")
	mockResponse.City = invoiced.String("Austin")
	mockResponse.Country = invoiced.String("USA")
	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.ListAllContacts(1234)
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

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.DeleteContact(1234, 456)

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}

func TestCustomer_CreatePaymentSource_Card(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Card)
	mockResponse.Id = int64(1234)
	mockResponse.Last4 = "4242"
	mockResponse.Object = "card"
	mockResponse.Brand = "Visa"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.CreatePaymentSource(1234, &invoiced.PaymentSourceRequest{})
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

	mockResponse := new(invoiced.BankAccount)
	mockResponse.Id = int64(1234)
	mockResponse.Last4 = "4242"
	mockResponse.Object = "bank_account"
	mockResponse.Verified = true

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.CreatePaymentSource(1234, &invoiced.PaymentSourceRequest{})
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

	var mockResponses invoiced.PaymentSources

	mockResponseCard := new(invoiced.PaymentSource)
	mockResponseCard.Object = "card"

	mockResponseAcct := new(invoiced.PaymentSource)
	mockResponseAcct.Object = "bank_account"

	mockResponses = append(mockResponses, *mockResponseCard)
	mockResponses = append(mockResponses, *mockResponseAcct)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.ListAllPaymentSources(1234)
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

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.DeleteCard(1234, 456)

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

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.DeleteBankAccount(1234, 456)

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}

func TestCustomer_CreatePendingLineItem(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.PendingLineItem)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.CreatePendingLineItem(1234, &invoiced.PendingLineItemRequest{})
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_RetrievePendingLineItem(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.PendingLineItem)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.RetrievePendingLineItem(1234, 456)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_UpdatePendingLineItem(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.PendingLineItem)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example 2"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.UpdatePendingLineItem(1234, 456, &invoiced.PendingLineItemRequest{})

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

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.DeletePendingLineItem(1234, 456)

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}

func TestCustomer_ListAllPendingLineItems(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.PendingLineItems
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.PendingLineItem)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "Mock Pli"

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.ListAllPendingLineItems(1234)
	if err != nil {
		t.Fatal("Error with pli", err)
	}

	if subjectEntity[0].Name != "Mock Pli" {
		t.Fatal("Retrieval not correct")
	}
}

func TestCustomer_RetrieveNotes(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Notes
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Note)
	mockResponse.Id = mockResponseId
	mockResponse.Notes = "Mock NoteClient"

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.RetrieveNotes(1234)
	if err != nil {
		t.Fatal("Error with note", err)
	}

	if subjectEntity[0].Notes != "Mock NoteClient" {
		t.Fatal("Retrieval not correct")
	}
}

func TestCustomer_TriggerInvoice(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.TriggerInvoice(1234)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestCustomer_ConsolidateInvoices(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.ConsolidateInvoices(1234)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}
