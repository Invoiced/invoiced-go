package invdapi

import (
	"reflect"
	"testing"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestEvent_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.Event

	mockResponse := new(invdendpoint.Event)
	mockResponse.Id = int64(123)

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewEvent()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(result[0].Event, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestEvent_List(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Events
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Event)
	mockResponse.Id = mockResponseId

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity := conn.NewEvent()

	entityResp, nextEndpoint, err := entity.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entityResp[0].Event, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestEvent_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Event)
	mockResponse.Id = int64(1234)

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewEvent()

	retrievedPayment, err := entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.Event, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
