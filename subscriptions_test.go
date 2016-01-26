package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"reflect"
	"testing"
	"time"
)

func TestSubscriptionCreate(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invdendpoint.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.UpdatedAt = time.Now().UnixNano()
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = 234

	subscriptionToCreate := new(invdendpoint.Subscription)

	subscriptionToCreate.Customer = 234112
	subscriptionToCreate.Plan = 234

	server := mockServer(200, mockSubscriptionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdSubscription, apiErr := conn.CreateSubscription(subscriptionToCreate)

	if apiErr != nil {
		t.Fatal("Error Creating subscription", apiErr)
	}

	if reflect.DeepEqual(createdSubscription, subscriptionToCreate) {
		t.Fatal("Subscription Was Not Created Succesfully")
	}

}

func TestSubscriptionCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server := mockServer(400, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	subscriptionToCreate := new(invdendpoint.Subscription)
	subscriptionToCreate.Customer = 234112
	subscriptionToCreate.Plan = 234

	_, apiErr := conn.CreateSubscription(subscriptionToCreate)

	if apiErr == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestSubscriptionUpdate(t *testing.T) {
	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invdendpoint.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.UpdatedAt = time.Now().UnixNano()
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = 234

	subscriptionToUpdate := new(invdendpoint.Subscription)

	mockSubscriptionResponse.Cycles = 42
	subscriptionToUpdate.Cycles = 42

	server := mockServer(200, mockSubscriptionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	updatedSubscription, apiErr := conn.UpdateSubscription(mockSubscriptionResponseID, subscriptionToUpdate)

	if apiErr != nil {
		t.Fatal("Error Updating subscription", apiErr)
	}

	if !reflect.DeepEqual(mockSubscriptionResponse, updatedSubscription) {
		t.Fatal("Error Messages Do Not Match Up", mockSubscriptionResponse, updatedSubscription.Id)
	}

}

func TestSubscriptionUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	subscriptionID := int64(324234)
	subscriptionToUpdate := new(invdendpoint.Subscription)
	subscriptionToUpdate.Cycles = 42

	server := mockServer(401, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.UpdateSubscription(subscriptionID, subscriptionToUpdate)

	if apiErr == nil {
		t.Fatal("Error Updating subscription", apiErr)
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestSubscriptionDelete(t *testing.T) {

	key := "api key"

	mocksubscriptionResponse := ""
	mocksubscriptionID := int64(2341)

	server := mockServer(204, mocksubscriptionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteSubscription(mocksubscriptionID)

	if apiErr != nil {
		t.Fatal("Error occured deleting subscription")
	}

}

func TestSubscriptionDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mocksubscriptionID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteSubscription(mocksubscriptionID)

	if apiErr == nil {
		t.Fatal("Error occured deleting subscription")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestSubscriptionList(t *testing.T) {

	key := "test api key"

	mockSubscriptionResponseID := int64(1523)
	mockSubscriptionResponse := new(invdendpoint.Subscription)
	mockSubscriptionResponse.Id = mockSubscriptionResponseID
	mockSubscriptionResponse.Customer = 234112
	mockSubscriptionResponse.Plan = 234

	mockSubscriptionResponse.UpdatedAt = time.Now().UnixNano()

	server := mockServer(200, mockSubscriptionResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdsubscription, apiErr := conn.ListSubscription(mockSubscriptionResponseID)

	if apiErr != nil {
		t.Fatal("Error Creating subscription", apiErr)
	}

	if createdsubscription.Id != mockSubscriptionResponseID {
		t.Fatal("subscription was not created succesfully")
	}

}

func TestSubscriptionListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mocksubscriptionID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.ListSubscription(mocksubscriptionID)

	if apiErr == nil {
		t.Fatal("Error occured deleting subscription")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}
