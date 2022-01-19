package subscription

import (
	"github.com/Invoiced/invoiced-go"
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestSubscriptionCreate(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invoiced.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.CreatedAt = time.Now().UnixNano()
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = "234"

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	createdSubscription, err := client.Create(&invoiced.SubscriptionRequest{Customer: invoiced.Int64(234112), Plan: invoiced.String("234")})
	if err != nil {
		t.Fatal("Error Creating subscription", err)
	}

	if !reflect.DeepEqual(createdSubscription, mockSubscriptionResponse) {
		t.Fatal("Client Was Not Created Succesfully", createdSubscription, mockSubscriptionResponse)
	}
}

func TestSubscriptionCreateError(t *testing.T) {
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

	_, err = client.Create(&invoiced.SubscriptionRequest{Customer: invoiced.Int64(234112), Plan: invoiced.String("234")})
	if err == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscriptionUpdate(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invoiced.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.CreatedAt = time.Now().UnixNano()
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = "234"
	mockSubscriptionResponse.Cycles = 42

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subscriptionToUpdate, err := client.Update(1523, &invoiced.SubscriptionRequest{Cycles: invoiced.Int64(42)})

	if err != nil {
		t.Fatal("Error Updating Client", err)
	}

	if !reflect.DeepEqual(mockSubscriptionResponse, subscriptionToUpdate) {
		t.Fatal("Error Client Not Updated Properly")
	}
}

func TestSubscriptionUpdateError(t *testing.T) {
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

	_, err = client.Update(1234, &invoiced.SubscriptionRequest{Cycles: invoiced.Int64(42)})
	if err == nil {
		t.Fatal("Error Updating subscription", err)
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscriptionDelete(t *testing.T) {
	key := "api key"

	mocksubscriptionResponse := ""

	server, err := invdmockserver.New(204, mocksubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	err = client.Cancel(2341)

	if err != nil {
		t.Fatal("Error Occurred Canceling Client")
	}
}

func TestSubscriptionDeleteError(t *testing.T) {
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
	err = client.Cancel(-999)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscription_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.Subscription

	mockResponse := new(invoiced.Subscription)
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

func TestSubscriptionRetrieve(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invoiced.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = "234"

	mockSubscriptionResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedSubscription, err := client.Retrieve(mockSubscriptionResponseID)
	if err != nil {
		t.Fatal("Error Creating subscription", err)
	}

	if !reflect.DeepEqual(retrievedSubscription, mockSubscriptionResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscriptionRetrieveError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockSubscriptionID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.Retrieve(mockSubscriptionID)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscription_List(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Subscriptions
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Subscription)
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

func TestSubscription_ListAll(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Subscriptions
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Subscription)
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

func TestSubscription_Preview(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponse := new(invoiced.SubscriptionPreview)
	mockSubscriptionResponse.FirstInvoice = nil
	mockSubscriptionResponse.MRR = float64(123.34)

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.Preview(&invoiced.SubscriptionPreviewRequest{})
	if err != nil {
		t.Fatal("Error Creating subscription", err)
	}
}
