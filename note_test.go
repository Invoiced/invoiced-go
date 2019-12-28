package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"testing"
	"time"
)

func TestNote_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Note)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Notes = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewNote()

	request := invdendpoint.CreateNoteRequest{CustomerID:int64(1234)}

	createdEntity, err := entity.Create(request)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.Note, mockResponse) {
		t.Fatal("entity was not created", createdEntity.Note, mockResponse)
	}

}

func TestNote_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Note)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Notes = "new-notes"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)

	entityToUpdate := conn.NewNote()

	entityToUpdate.Notes = "new-notes"

	err = entityToUpdate.Save()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Note) {
		t.Fatal("Error: entity not updated correctly")
	}

}

func TestNote_Delete(t *testing.T) {

	key := "api key"

	mockResponse := ""
	mockResponseId := int64(1234)

	server, err := invdmockserver.New(204, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewNote()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}

}

func TestNote_ListAll(t *testing.T) {

	key := "test api key"

	var mockListResponse [1] invdendpoint.Note

	mockResponse := new(invdendpoint.Note)
	mockResponse.Id = int64(1234)
	mockResponse.Notes = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewNote()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)

	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].Note, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}