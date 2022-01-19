package plan

import (
	"github.com/Invoiced/invoiced-go"
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestPlan_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(Client)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entity := client.NewPlan()

	createdEntity, err := entity.Create(&invoiced.PlanRequest{Id: invoiced.String("example"), Name: invoiced.String("nomenclature")})
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.Plan, mockResponse) {
		t.Fatal("entity was not created", createdEntity.Plan, mockResponse)
	}
}

func TestPlan_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(Client)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entityToUpdate := client.NewPlan()
	err = entityToUpdate.Update(&invoiced.PlanRequest{Name: invoiced.String("new-name")})

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Plan) {
		t.Fatal("Error: entity not updated correctly")
	}
}

func TestPlan_Delete(t *testing.T) {
	key := "api key"

	mockResponse := ""
	mockResponseId := "example"

	server, err := invdmockserver.New(204, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := invoiced.NewMockApi(key, server)

	entity := client.NewPlan()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}
}

func TestPlan_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(Client)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)
	entity := client.NewPlan()

	retrievedPayment, err := entity.Retrieve("example")
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.Plan, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestPlan_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]Client

	mockResponse := new(Client)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := invoiced.NewMockApi(key, server)
	entity := client.NewPlan()

	filter := invoiced.NewFilter()
	sorter := invoiced.NewSort()

	result, err := entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].Plan, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
