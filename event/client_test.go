package event

import (
	"github.com/Invoiced/invoiced-go/v2"
	"testing"
	"github.com/Invoiced/invoiced-go/v2/invdmockserver"
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

	client := Client{invoiced.NewMockApi(key, server)}

	filter := invoiced.NewFilter()
	sorter := invoiced.NewSort()

	_, err = client.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}
}

func TestEvent_List(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Events
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.Event)
	mockResponse.Id = mockResponseId

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

func TestEvent_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Event)
	mockResponse.Id = int64(1234)

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}
}
