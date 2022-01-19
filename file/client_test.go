package file

import (
	"github.com/Invoiced/invoiced-go"
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestFile_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.File)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entity := client.NewFile()

	createdEntity, err := entity.Create(&invoiced.FileRequest{Name: invoiced.String("nomenclature")})
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.File, mockResponse) {
		t.Fatal("entity was not created", createdEntity.File, mockResponse)
	}
}

func TestFile_Delete(t *testing.T) {
	key := "api key"

	mockResponse := ""
	mockResponseId := int64(1234)

	server, err := invdmockserver.New(204, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entity := client.NewFile()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}
}

func TestFile_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.File)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)
	entity := client.NewFile()

	retrievedPayment, err := entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.File, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
