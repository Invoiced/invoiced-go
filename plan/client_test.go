package plan

import (
	"reflect"
	"testing"
	"time"

	"github.com/strongdm/invoiced-go/v2"
	"github.com/strongdm/invoiced-go/v2/invdmockserver"
)

func TestPlan_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Plan)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	createdEntity, err := client.Create(&invoiced.PlanRequest{Id: invoiced.String("example"), Name: invoiced.String("nomenclature")})
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity, mockResponse) {
		t.Fatal("entity was not created", createdEntity, mockResponse)
	}
}

func TestPlan_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Plan)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	entityToUpdate, err := client.Update("example", &invoiced.PlanRequest{Name: invoiced.String("new-name")})

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate) {
		t.Fatal("Error: entity not updated correctly")
	}
}

func TestPlan_Delete(t *testing.T) {
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
		t.Fatal("Error occurred deleting entity")
	}
}

func TestPlan_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Plan)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedPayment, err := client.Retrieve("example")
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestPlan_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]*invoiced.Plan

	mockResponse := new(invoiced.Plan)
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
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0], mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
