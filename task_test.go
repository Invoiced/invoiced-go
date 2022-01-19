package invoiced

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestTask_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(Task)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := NewMockApi(key, server)

	entity := client.NewTask()
	createdEntity, err := entity.Create(&TaskRequest{Name: String("nomenclature")})
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.Task, mockResponse) {
		t.Fatal("entity was not created", createdEntity.Task, mockResponse)
	}
}

func TestTask_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(Task)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi(key, server)

	entityToUpdate := client.NewTask()
	err = entityToUpdate.Update(&TaskRequest{Name: String("new-name")})

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Task) {
		t.Fatal("Error: entity not updated correctly")
	}
}

func TestTask_Delete(t *testing.T) {
	key := "api key"

	mockResponse := ""
	mockResponseId := int64(1234)

	server, err := invdmockserver.New(204, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := NewMockApi(key, server)

	entity := client.NewTask()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}
}

func TestTask_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(Task)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi(key, server)
	entity := client.NewTask()

	retrievedPayment, err := entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.Task, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestTask_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]Task

	mockResponse := new(Task)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := NewMockApi(key, server)
	entity := client.NewTask()

	filter := NewFilter()
	sorter := NewSort()

	result, err := entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].Task, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
