package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"testing"
	"time"
)

func TestCreditNote_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.CreditNote)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewCreditNote()

	requestEntity := entity.NewCreditNote()

	requestEntity.Id = int64(1234)
	requestEntity.Name = "nomenclature"

	createdEntity, err := entity.Create(requestEntity)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.CreditNote, mockResponse) {
		t.Fatal("entity was not created", createdEntity.CreditNote, mockResponse)
	}

}

func TestCreditNote_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.CreditNote)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)

	entityToUpdate := conn.NewCreditNote()

	entityToUpdate.Name = "new-name"

	err = entityToUpdate.Save()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.CreditNote) {
		t.Fatal("Error: entity not updated correctly")
	}

}

func TestCreditNote_Delete(t *testing.T) {

	key := "api key"

	mockResponse := ""
	mockResponseId := int64(1234)

	server, err := invdmockserver.New(204, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewCreditNote()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}

}

func TestCreditNote_Retrieve(t *testing.T) {

	key := "test api key"

	mockResponse := new(invdendpoint.CreditNote)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewCreditNote()

	retrievedTransaction, err := entity.Retrieve(int64(1234))

	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedTransaction.CreditNote, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}

func TestCreditNote_ListAll(t *testing.T) {

	key := "test api key"

	var mockListResponse [1] invdendpoint.CreditNote

	mockResponse := new(invdendpoint.CreditNote)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewCreditNote()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)

	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].CreditNote, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}