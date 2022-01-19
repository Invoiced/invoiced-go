package event

import (
	"github.com/Invoiced/invoiced-go"
	"testing"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestEvent_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.Event

	mockResponse := new(invoiced.Event)
	mockResponse.Id = int64(123)

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)
	entity := client.NewEvent()

	filter := invoiced.NewFilter()
	sorter := invoiced.NewSort()

	_, err = entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}
}

func TestEvent_List(t *testing.T) {
	key := "test api key"

	var mockResponses Events
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Event)
	mockResponse.Id = mockResponseId

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entity := client.NewEvent()

	_, nextEndpoint, err := entity.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}
}

func TestEvent_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Event)
	mockResponse.Id = int64(1234)

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)
	entity := client.NewEvent()

	_, err = entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}
}
