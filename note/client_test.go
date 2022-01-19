package note

import (
	"github.com/Invoiced/invoiced-go"
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestNote_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Note)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Notes = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entity := client.NewNote()

	createdEntity, err := entity.Create(&invoiced.NoteRequest{Customer: invoiced.Int64(1234)})
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.Note, mockResponse) {
		t.Fatal("entity was not created", createdEntity.Note, mockResponse)
	}
}

func TestNote_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Note)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Notes = "new-notes"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entityToUpdate := client.NewNote()

	err = entityToUpdate.Update(&invoiced.NoteRequest{Notes: invoiced.String("new-notes")})

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

	client := invoiced.NewMockApi(key, server)

	entity := client.NewNote()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}
}

func TestNote_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.Note

	mockResponse := new(invoiced.Note)
	mockResponse.Id = int64(1234)
	mockResponse.Notes = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)
	entity := client.NewNote()

	filter := invoiced.NewFilter()
	sorter := invoiced.NewSort()

	result, err := entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].Note, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
