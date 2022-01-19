package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestItem_Create(t *testing.T) {
	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invdendpoint.Item)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	requestEntity := conn.NewItem()

	requestEntity, err = requestEntity.Create(&invdendpoint.ItemRequest{Name: String("delivery"), Type: String("service")})

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(requestEntity.Item, mockResponse) {
		t.Fatal("entity was not created", requestEntity.Item, mockResponse)
	}
}

func TestItem_Save(t *testing.T) {
	key := "test api key"

	mockResponseId := "delivery"
	mockResponse := new(invdendpoint.Item)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	entityToUpdate := conn.NewItem()

	err = entityToUpdate.Update(&invdendpoint.ItemRequest{Name: String("new-name")})

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Item) {
		t.Fatal("Error: entity not updated correctly")
	}
}

func TestItem_Delete(t *testing.T) {
	key := "api key"

	mockResponse := ""
	mockResponseId := "example"

	server, err := invdmockserver.New(204, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity := conn.NewItem()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error Occured Deleting Payment")
	}
}

func TestItem_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invdendpoint.Item)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewItem()

	retrievedPayment, err := entity.Retrieve(mockResponseId)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.Item, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestItem_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.Item

	mockResponse := new(invdendpoint.Item)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewItem()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(result[0].Item, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
