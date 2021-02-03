package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestPaymentCreate(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(invdendpoint.Payment)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.CreatedAt = time.Now().UnixNano()
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	payment := conn.NewPayment()

	paymentToCreate := payment.NewPayment()

	paymentToCreate.Customer = 234112

	createdPayment, err := payment.Create(paymentToCreate)
	if err != nil {
		t.Fatal("Error Creating payment", err)
	}

	if !reflect.DeepEqual(createdPayment.Payment, mockPaymentResponse) {
		t.Fatal("Payment Was Not Created Succesfully", createdPayment.Payment, mockPaymentResponse)
	}
}

func TestPaymentCreateError(t *testing.T) {
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
	payment := conn.NewPayment()
	paymentToCreate := payment.NewPayment()
	paymentToCreate.Customer = 234112
	paymentToCreate.Reference = "234"

	_, apiErr := payment.Create(paymentToCreate)

	if apiErr == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), apiErr.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentUpdate(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(invdendpoint.Payment)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.CreatedAt = time.Now().UnixNano()
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	paymentToUpdate := conn.NewPayment()

	mockPaymentResponse.Amount = 42
	paymentToUpdate.Amount = 42

	err = paymentToUpdate.Save()

	if err != nil {
		t.Fatal("Error Updating Payment", err)
	}

	if !reflect.DeepEqual(mockPaymentResponse, paymentToUpdate.Payment) {
		t.Fatal("Error Payment Not Updated Properly")
	}
}

func TestPaymentUpdateError(t *testing.T) {
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
	subcriptionToUpdate := conn.NewPayment()

	subcriptionToUpdate.Amount = 42

	err = subcriptionToUpdate.Save()

	if err == nil {
		t.Fatal("Error Updating payment", err)
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentDelete(t *testing.T) {
	key := "api key"

	mockpaymentResponse := ""
	mockpaymentID := int64(2341)

	server, err := invdmockserver.New(204, mockpaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	payment := conn.NewPayment()

	payment.Id = mockpaymentID

	err = payment.Delete()

	if err != nil {
		t.Fatal("Error Occured Deleting Payment")
	}
}

func TestPaymentDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockPaymentID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	payment := conn.NewPayment()

	payment.Id = mockPaymentID

	err = payment.Delete()

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentRetrieve(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(invdendpoint.Payment)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"

	mockPaymentResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	payment := conn.NewPayment()

	retrievedPayment, err := payment.Retrieve(mockPaymentResponseID)
	if err != nil {
		t.Fatal("Error Creating payment", err)
	}

	if !reflect.DeepEqual(retrievedPayment.Payment, mockPaymentResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentRetrieveError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockPaymentID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	payment := conn.NewPayment()

	_, err = payment.Retrieve(mockPaymentID)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPayment_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.Payment

	mockResponse := new(invdendpoint.Payment)
	mockResponse.Id = int64(1234)

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewPayment()

	result, err := entity.Count()

	if result != int64(-1) || err == nil {
		t.Fatal("Unexpectedly successful")
	}
}

func TestPayment_List(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Payments
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Payment)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	payment := conn.NewPayment()

	invoiceResp, nextEndpoint, err := payment.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(invoiceResp[0].Payment, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPayment_ListAll(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Payments
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Payment)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	subscription := conn.NewPayment()

	invoiceResp, err := subscription.ListAll(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(invoiceResp[0].Payment, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPayment_SendReceipt(t *testing.T) {
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

	subjectEntity := conn.NewPayment()

	sendResponse, err := subjectEntity.SendReceipt(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello test" {
		t.Fatal("Error: send not completed correctly")
	}
}
