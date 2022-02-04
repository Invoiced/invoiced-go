package payment

import (
	"reflect"
	"testing"
	"time"

	"github.com/strongdm/invoiced-go/v2"
	"github.com/strongdm/invoiced-go/v2/invdmockserver"
)

func TestPaymentCreate(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(invoiced.Payment)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.CreatedAt = time.Now().UnixNano()
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	createdPayment, err := client.Create(&invoiced.PaymentRequest{Customer: invoiced.Int64(234112)})
	if err != nil {
		t.Fatal("Error Creating payment", err)
	}

	if !reflect.DeepEqual(createdPayment, mockPaymentResponse) {
		t.Fatal("Client Was Not Created Succesfully", createdPayment, mockPaymentResponse)
	}
}

func TestPaymentCreateError(t *testing.T) {
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

	_, err = client.Create(&invoiced.PaymentRequest{Customer: invoiced.Int64(234112), Reference: invoiced.String("234")})
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
	mockPaymentResponse := new(invoiced.Payment)
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

	client := Client{invoiced.NewMockApi(key, server)}

	paymentToUpdate, err := client.Update(1523, &invoiced.PaymentRequest{Amount: invoiced.Float64(42)})

	if err != nil {
		t.Fatal("Error Updating Client", err)
	}

	if !reflect.DeepEqual(mockPaymentResponse, paymentToUpdate) {
		t.Fatal("Error Client Not Updated Properly")
	}
}

func TestPaymentUpdateError(t *testing.T) {
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

	_, err = client.Update(1234, &invoiced.PaymentRequest{Amount: invoiced.Float64(42)})

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

	server, err := invdmockserver.New(204, mockpaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.Delete(2341)

	if err != nil {
		t.Fatal("Error occurred Deleting Client")
	}
}

func TestPaymentDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.Delete(-999)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentRetrieve(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(invoiced.Payment)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.Reference = "234"
	mockPaymentResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedPayment, err := client.Retrieve(1523)
	if err != nil {
		t.Fatal("Error Creating payment", err)
	}

	if !reflect.DeepEqual(retrievedPayment, mockPaymentResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPaymentRetrieveError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockPaymentID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.Retrieve(mockPaymentID)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestPayment_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.Payment

	mockResponse := new(invoiced.Payment)
	mockResponse.Id = int64(1234)

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	result, err := client.Count()

	if result != int64(-1) || err == nil {
		t.Fatal("Unexpectedly successful")
	}
}

func TestPayment_List(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Payments
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Payment)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, nextEndpoint, err := client.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}
}

func TestPayment_ListAll(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Payments
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Payment)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.ListAll(nil, nil)
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

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.SendReceipt(1234, nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

}
