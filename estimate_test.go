package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"testing"
	"time"
)

func TestEstimate_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewEstimate()

	requestEntity := entity.NewEstimate()

	requestEntity.Id = int64(1234)
	requestEntity.Name = "nomenclature"

	createdEntity, err := entity.Create(requestEntity)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.Estimate, mockResponse) {
		t.Fatal("entity was not created", createdEntity.Estimate, mockResponse)
	}

}

func TestEstimate_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)

	entityToUpdate := conn.NewEstimate()

	entityToUpdate.Name = "new-name"

	err = entityToUpdate.Save()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Estimate) {
		t.Fatal("Error: entity not updated correctly")
	}

}

func TestEstimate_Delete(t *testing.T) {

	key := "api key"

	mockResponse := ""
	mockResponseId := int64(1234)

	server, err := invdmockserver.New(204, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewEstimate()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}

}

func TestEstimate_Retrieve(t *testing.T) {

	key := "test api key"

	mockResponse := new(invdendpoint.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewEstimate()

	retrievedTransaction, err := entity.Retrieve(int64(1234))

	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedTransaction.Estimate, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}

func TestEstimate_ListAll(t *testing.T) {

	key := "test api key"

	var mockListResponse [1] invdendpoint.Estimate

	mockResponse := new(invdendpoint.Estimate)
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
	entity := conn.NewEstimate()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)

	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].Estimate, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}