package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestSubscriptionCreate(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invdendpoint.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.CreatedAt = time.Now().UnixNano()
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = "234"

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	subscription := conn.NewSubscription()

	subscriptionToCreate := subscription.NewSubscription()

	subscriptionToCreate.Customer = 234112
	subscriptionToCreate.Plan = "234"

	createdSubscription, err := subscription.Create(subscriptionToCreate)
	if err != nil {
		t.Fatal("Error Creating subscription", err)
	}

	if !reflect.DeepEqual(createdSubscription.Subscription, mockSubscriptionResponse) {
		t.Fatal("Subscription Was Not Created Succesfully", createdSubscription.Subscription, mockSubscriptionResponse)
	}
}

func TestSubscriptionCreateError(t *testing.T) {
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
	subscription := conn.NewSubscription()
	subscriptionToCreate := subscription.NewSubscription()
	subscriptionToCreate.Customer = 234112
	subscriptionToCreate.Plan = "234"

	_, apiErr := subscription.Create(subscriptionToCreate)

	if apiErr == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), apiErr.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscriptionUpdate(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invdendpoint.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.CreatedAt = time.Now().UnixNano()
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = "234"

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subscriptionToUpdate := conn.NewSubscription()

	mockSubscriptionResponse.Cycles = 42
	subscriptionToUpdate.Cycles = 42

	err = subscriptionToUpdate.Save()

	if err != nil {
		t.Fatal("Error Updating Subscription", err)
	}

	if !reflect.DeepEqual(mockSubscriptionResponse, subscriptionToUpdate.Subscription) {
		t.Fatal("Error Subscription Not Updated Properly")
	}
}

func TestSubscriptionUpdateError(t *testing.T) {
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
	subcriptionToUpdate := conn.NewSubscription()

	subcriptionToUpdate.Cycles = 42

	err = subcriptionToUpdate.Save()

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
	mocksubscriptionID := int64(2341)

	server, err := invdmockserver.New(204, mocksubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	subscription := conn.NewSubscription()

	subscription.Id = mocksubscriptionID

	err = subscription.Cancel()

	if err != nil {
		t.Fatal("Error Occured Canceling Subscription")
	}
}

func TestSubscriptionDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockSubscriptionID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	subscription := conn.NewSubscription()

	subscription.Id = mockSubscriptionID

	err = subscription.Cancel()

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscription_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.Subscription

	mockResponse := new(invdendpoint.Subscription)
	mockResponse.Id = int64(1234)

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewSubscription()

	result, err := entity.Count()

	if result != int64(-1) || err == nil {
		t.Fatal("Unexpectedly successful")
	}
}

func TestSubscriptionRetrieve(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invdendpoint.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = "234"

	mockSubscriptionResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	subscription := conn.NewSubscription()

	retrievedSubscription, err := subscription.Retrieve(mockSubscriptionResponseID)
	if err != nil {
		t.Fatal("Error Creating subscription", err)
	}

	if !reflect.DeepEqual(retrievedSubscription.Subscription, mockSubscriptionResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscriptionRetrieveError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockSubscriptionID := int64(-999)

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	subscription := conn.NewSubscription()

	_, err = subscription.Retrieve(mockSubscriptionID)

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscription_List(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Subscriptions
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Subscription)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	subscription := conn.NewSubscription()

	invoiceResp, nextEndpoint, err := subscription.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(invoiceResp[0].Subscription, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up",invoiceResp[0].Subscription,"apple",mockResponse,reflect.DeepEqual(invoiceResp[0].Subscription, mockResponse))

	}
}

func TestSubscription_ListAll(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Subscriptions
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Subscription)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	subscription := conn.NewSubscription()

	invoiceResp, err := subscription.ListAll(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(invoiceResp[0].Subscription, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestSubscription_Preview(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponse := new(invdendpoint.SubscriptionPreview)
	mockSubscriptionResponse.FirstInvoice = nil
	mockSubscriptionResponse.MRR = float64(123.34)

	server, err := invdmockserver.New(200, mockSubscriptionResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	subscription := conn.NewSubscription()

	retrievedSubscription, err := subscription.Preview(conn.NewPreviewRequest())
	if err != nil {
		t.Fatal("Error Creating subscription", err)
	}

	if !reflect.DeepEqual(retrievedSubscription, mockSubscriptionResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}
