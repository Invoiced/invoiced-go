package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"reflect"
	"testing"
	"time"
)

func TestTransactionCreate(t *testing.T) {
	key := "test api key"

	mockTransactionResponseID := int64(1523)
	mockTransactionResponse := new(invdendpoint.Transaction)
	mockTransactionResponse.Id = mockTransactionResponseID
	mockTransactionResponse.UpdatedAt = time.Now().UnixNano()
	mockTransactionResponse.Type = "payment"
	mockTransactionResponse.Amount = 34.99

	transactionToCreate := new(invdendpoint.Transaction)

	transactionToCreate.Type = "payment"
	transactionToCreate.Amount = 34.99

	server := mockServer(200, mockTransactionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdTransaction, apiErr := conn.CreateTransaction(transactionToCreate)

	if apiErr != nil {
		t.Fatal("Error Creating transaction", apiErr)
	}

	if reflect.DeepEqual(createdTransaction, transactionToCreate) {
		t.Fatal("Transaction Was Not Created Succesfully")
	}

}

func TestTransactionCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server := mockServer(400, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	transactionToCreate := new(invdendpoint.Transaction)
	transactionToCreate.Type = "payment"
	transactionToCreate.Amount = 342.234

	_, apiErr := conn.CreateTransaction(transactionToCreate)

	if apiErr == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestTransactionUpdate(t *testing.T) {
	key := "test api key"

	mockTransactionResponseID := int64(1523)
	mockTransactionResponse := new(invdendpoint.Transaction)
	mockTransactionResponse.Id = mockTransactionResponseID
	mockTransactionResponse.UpdatedAt = time.Now().UnixNano()
	mockTransactionResponse.Type = "payment"
	mockTransactionResponse.Amount = 34.99

	transactionToUpdate := new(invdendpoint.Transaction)

	mockTransactionResponse.Amount = 42.22
	transactionToUpdate.Amount = 42.22

	server := mockServer(200, mockTransactionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	updatedTransaction, apiErr := conn.UpdateTransaction(mockTransactionResponseID, transactionToUpdate)

	if apiErr != nil {
		t.Fatal("Error Updating transaction", apiErr)
	}

	if !reflect.DeepEqual(mockTransactionResponse, updatedTransaction) {
		t.Fatal("Error Messages Do Not Match Up", mockTransactionResponse, updatedTransaction.Id)
	}

}

func TestTransactionUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	transactionID := int64(324234)
	transactionToUpdate := new(invdendpoint.Transaction)
	transactionToUpdate.Amount = 400.12

	server := mockServer(401, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.UpdateTransaction(transactionID, transactionToUpdate)

	if apiErr == nil {
		t.Fatal("Error Updating transaction", apiErr)
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestTransactionDelete(t *testing.T) {

	key := "api key"

	mocktransactionResponse := ""
	mocktransactionID := int64(2341)

	server := mockServer(204, mocktransactionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteTransaction(mocktransactionID)

	if apiErr != nil {
		t.Fatal("Error occured deleting transaction")
	}

}

func TestTransactionDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mocktransactionID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteTransaction(mocktransactionID)

	if apiErr == nil {
		t.Fatal("Error occured deleting transaction")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestTransactionList(t *testing.T) {

	key := "test api key"

	mockTransactionResponseID := int64(1523)
	mockTransactionResponse := new(invdendpoint.Transaction)
	mockTransactionResponse.Id = mockTransactionResponseID
	mockTransactionResponse.Amount = 23.22
	mockTransactionResponse.Gateway = "Stripe"

	mockTransactionResponse.UpdatedAt = time.Now().UnixNano()

	server := mockServer(200, mockTransactionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdtransaction, apiErr := conn.ListTransaction(mockTransactionResponseID)

	if apiErr != nil {
		t.Fatal("Error Creating transaction", apiErr)
	}

	if createdtransaction.Id != mockTransactionResponseID {
		t.Fatal("transaction was not created succesfully")
	}

}

func TestTransactionListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mocktransactionID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.ListTransaction(mocktransactionID)

	if apiErr == nil {
		t.Fatal("Error occured deleting transaction")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}
