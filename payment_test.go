package invoiced

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestPaymentCreate(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(PaymentClient)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.CreatedAt = time.Now().UnixNano()
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := NewMockApi(key, server)

	payment := client.NewPayment()

	createdPayment, err := payment.Create(&PaymentRequest{Customer: Int64(234112)})
	if err != nil {
		t.Fatal("Error Creating payment", err)
	}

	if !reflect.DeepEqual(createdPayment.Payment, mockPaymentResponse) {
		t.Fatal("PaymentClient Was Not Created Succesfully", createdPayment.Payment, mockPaymentResponse)
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

	client := NewMockApi(key, server)
	payment := client.NewPayment()

	_, err = payment.Create(&PaymentRequest{Customer: Int64(234112), Reference: String("234")})
	if err == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentUpdate(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(PaymentClient)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.CreatedAt = time.Now().UnixNano()
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"
	mockPaymentResponse.Amount = 42

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi(key, server)

	paymentToUpdate := client.NewPayment()

	err = paymentToUpdate.Update(&PaymentRequest{Amount: Float64(42)})

	if err != nil {
		t.Fatal("Error Updating PaymentClient", err)
	}

	if !reflect.DeepEqual(mockPaymentResponse, paymentToUpdate.Payment) {
		t.Fatal("Error PaymentClient Not Updated Properly")
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

	client := NewMockApi(key, server)
	subcriptionToUpdate := client.NewPayment()

	err = subcriptionToUpdate.Update(&PaymentRequest{Amount: Float64(42)})

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

	client := NewMockApi(key, server)

	payment := client.NewPayment()

	payment.Id = mockpaymentID

	err = payment.Delete()

	if err != nil {
		t.Fatal("Error Occured Deleting PaymentClient")
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

	client := NewMockApi(key, server)

	payment := client.NewPayment()

	payment.Id = mockPaymentID

	err = payment.Delete()

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentRetrieve(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(PaymentClient)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"

	mockPaymentResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi(key, server)
	payment := client.NewPayment()

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

	client := NewMockApi(key, server)
	payment := client.NewPayment()

	_, err = payment.Retrieve(mockPaymentID)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPayment_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]PaymentClient

	mockResponse := new(PaymentClient)
	mockResponse.Id = int64(1234)

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi(key, server)
	entity := client.NewPayment()

	result, err := entity.Count()

	if result != int64(-1) || err == nil {
		t.Fatal("Unexpectedly successful")
	}
}

func TestPayment_List(t *testing.T) {
	key := "test api key"

	var mockResponses Payments
	mockResponseId := int64(1523)
	mockResponse := new(PaymentClient)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := NewMockApi(key, server)

	payment := client.NewPayment()

	_, nextEndpoint, err := payment.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}
}

func TestPayment_ListAll(t *testing.T) {
	key := "test api key"

	var mockResponses Payments
	mockResponseId := int64(1523)
	mockResponse := new(PaymentClient)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := NewMockApi(key, server)

	payment := client.NewPayment()

	_, err = payment.ListAll(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPayment_SendReceipt(t *testing.T) {
	key := "test api key"

	server, err := invdmockserver.New(200, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi(key, server)

	subjectEntity := client.NewPayment()

	err = subjectEntity.SendReceipt(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

}
