package item

import (
	"github.com/Invoiced/invoiced-go"
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestItem_Create(t *testing.T) {
	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invoiced.Item)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	item, err := client.Create(&invoiced.ItemRequest{Name: invoiced.String("delivery"), Type: invoiced.String("service")})

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(item, mockResponse) {
		t.Fatal("entity was not created", item, mockResponse)
	}
}

func TestItem_Save(t *testing.T) {
	key := "test api key"

	mockResponseId := "delivery"
	mockResponse := new(invoiced.Item)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	entityToUpdate, err := client.Update("example", &invoiced.ItemRequest{Name: invoiced.String("new-name")})

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate) {
		t.Fatal("Error: entity not updated correctly")
	}
}

func TestItem_Delete(t *testing.T) {
	key := "api key"

	mockResponse := ""

	server, err := invdmockserver.New(204, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	err = client.Delete("example")

	if err != nil {
		t.Fatal("Error occurred Deleting Client")
	}
}

func TestItem_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invoiced.Item)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedPayment, err := client.Retrieve(mockResponseId)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestItem_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]*invoiced.Item
	mockResponse := new(invoiced.Item)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockListResponse[0] = mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	filter := invoiced.NewFilter()
	sorter := invoiced.NewSort()

	result, err := client.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(result[0], mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
