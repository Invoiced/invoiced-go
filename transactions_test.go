package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestTransactionCreate(t *testing.T) {
	key := "test api key"

	mockTransactionResponseID := int64(1523)
	mockTransactionResponse := new(invdendpoint.Transaction)
	mockTransactionResponse.Id = mockTransactionResponseID
	mockTransactionResponse.CreatedAt = time.Now().UnixNano()
	mockTransactionResponse.Customer = 234112
	mockTransactionResponse.GatewayId = "234"

	server, err := invdmockserver.New(200, mockTransactionResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	transaction := conn.NewTransaction()

	transactionToCreate := transaction.NewTransaction()

	transactionToCreate.Customer = 234112
	transactionToCreate.Gateway = "dell"

	createdTransaction, err := transaction.Create(transactionToCreate)

	if err != nil {
		t.Fatal("Error Creating transaction", err)
	}

	if !reflect.DeepEqual(createdTransaction.Transaction, mockTransactionResponse) {
		t.Fatal("Transaction Was Not Created Succesfully", createdTransaction.Transaction, mockTransactionResponse)
	}

}

func TestTransactionCreateError(t *testing.T) {
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

	conn := MockConnection(key, server)
	transaction := conn.NewTransaction()
	transactionToCreate := transaction.NewTransaction()
	transactionToCreate.Customer = 234112
	transactionToCreate.GatewayId = "234"

	_, apiErr := transaction.Create(transactionToCreate)

	if apiErr == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), apiErr.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestTransactionUpdate(t *testing.T) {
	key := "test api key"

	mockTransactionResponseID := int64(1523)
	mockTransactionResponse := new(invdendpoint.Transaction)
	mockTransactionResponse.Id = mockTransactionResponseID
	mockTransactionResponse.CreatedAt = time.Now().UnixNano()
	mockTransactionResponse.Customer = 234112
	mockTransactionResponse.GatewayId = "234"

	server, err := invdmockserver.New(200, mockTransactionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)

	transactionToUpdate := conn.NewTransaction()

	mockTransactionResponse.Amount = 42
	transactionToUpdate.Amount = 42

	err = transactionToUpdate.Save()

	if err != nil {
		t.Fatal("Error Updating Transaction", err)
	}

	if !reflect.DeepEqual(mockTransactionResponse, transactionToUpdate.Transaction) {
		t.Fatal("Error Transaction Not Updated Properly")
	}

}

func TestTransactionUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	server, err := invdmockserver.New(401, mockErrorResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)
	subcriptionToUpdate := conn.NewTransaction()

	subcriptionToUpdate.Amount = 42

	err = subcriptionToUpdate.Save()

	if err == nil {
		t.Fatal("Error Updating transaction", err)
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestTransactionDelete(t *testing.T) {

	key := "api key"

	mocktransactionResponse := ""
	mocktransactionID := int64(2341)

	server, err := invdmockserver.New(204, mocktransactionResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	transaction := conn.NewTransaction()

	transaction.Id = mocktransactionID

	err = transaction.Delete()

	if err != nil {
		t.Fatal("Error Occured Deleting Transaction")
	}

}

func TestTransactionDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockTransactionID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	transaction := conn.NewTransaction()

	transaction.Id = mockTransactionID

	err = transaction.Delete()

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestTransactionRetrieve(t *testing.T) {

	key := "test api key"

	mockTransactionResponseID := int64(1523)
	mockTransactionResponse := new(invdendpoint.Transaction)
	mockTransactionResponse.Id = mockTransactionResponseID
	mockTransactionResponse.Customer = 234112
	mockTransactionResponse.GatewayId = "234"

	mockTransactionResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockTransactionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	transaction := conn.NewTransaction()

	retrievedTransaction, err := transaction.Retrieve(mockTransactionResponseID)

	if err != nil {
		t.Fatal("Error Creating transaction", err)
	}

	if !reflect.DeepEqual(retrievedTransaction.Transaction, mockTransactionResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestTransactionRetrieveError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockTransactionID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	transaction := conn.NewTransaction()

	_, err = transaction.Retrieve(mockTransactionID)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}
