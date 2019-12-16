package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"testing"
	"time"
)

func TestCatalogItem_Create(t *testing.T) {
	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invdendpoint.CatalogItem)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewCatalogItem()

	requestEntity := entity.NewCatalogItem()

	requestEntity.Id = "example"
	requestEntity.Name = "delivery"
	requestEntity.Type = "service"

	createdEntity, err := entity.Create(requestEntity)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.CatalogItem, mockResponse) {
		t.Fatal("entity was not created", createdEntity.CatalogItem, mockResponse)
	}

}

func TestCatalogItem_Save(t *testing.T) {
	key := "test api key"

	mockResponseId := "delivery"
	mockResponse := new(invdendpoint.CatalogItem)
	mockResponse.Id = mockResponseId
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"
	mockResponse.Type = "service"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)

	entityToUpdate := conn.NewCatalogItem()

	entityToUpdate.Name = "new-name"

	err = entityToUpdate.Save()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.CatalogItem) {
		t.Fatal("Error: entity not updated correctly")
	}

}

func TestCatalogItem_Delete(t *testing.T) {

	key := "api key"

	mockResponse := ""
	mockResponseId := "example"

	server, err := invdmockserver.New(204, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewCatalogItem()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error Occured Deleting Transaction")
	}

}

func TestCatalogItem_Retrieve(t *testing.T) {

	key := "test api key"

	mockResponseId := "example"
	mockResponse := new(invdendpoint.CatalogItem)
	mockResponse.Id = mockResponseId
	mockResponse.Name = "delivery"
	mockResponse.Type = "service"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewCatalogItem()

	retrievedTransaction, err := entity.Retrieve(mockResponseId)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(retrievedTransaction.CatalogItem, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}