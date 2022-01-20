package estimate

import (
	"github.com/Invoiced/invoiced-go/v2"
	"reflect"
	"testing"
	"time"
	"github.com/Invoiced/invoiced-go/v2/invdmockserver"
)

func TestEstimate_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	createdEntity, err := client.Create(&invoiced.EstimateRequest{Name: invoiced.String("nomenclature")})
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity, mockResponse) {
		t.Fatal("entity was not created", createdEntity, mockResponse)
	}
}

func TestEstimate_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	entityToUpdate, err := client.Update(1234, &invoiced.EstimateRequest{Name: invoiced.String("new-name")})

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate) {
		t.Fatal("Error: entity not updated correctly")
	}
}

func TestEstimate_Delete(t *testing.T) {
	key := "api key"

	mockResponse := ""

	server, err := invdmockserver.New(204, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.Delete(1234)

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}
}

func TestEstimate_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	retrievedPayment, err := client.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestEstimate_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.Estimate

	mockResponse := new(invoiced.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

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

func TestEstimate_Void(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"
	mockResponse.Status = "voided"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	entityToUpdate, err := client.Void(1234)

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate) {
		t.Fatal("Error: entity not voided correctly")
	}
}

func TestEstimate_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.Estimate

	mockResponse := new(invoiced.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	result, err := client.Count()

	if err == nil {
		t.Fatal("Error: ", err)
	}

	if result != int64(-1) {
		t.Fatal("Unexpectedly successful")
	}
}

func TestEstimate_SendEmail(t *testing.T) {
	key := "test api key"

	server, err := invdmockserver.New(200, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.SendEmail(1234, nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}
}

func TestEstimate_List(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Estimates
	mockResponseId := int64(1523)
	mockNumber := "INV-3421"
	mockResponse := new(invoiced.Estimate)
	mockResponse.Id = mockResponseId
	mockResponse.Number = mockNumber
	mockResponse.PaymentTerms = "NET15"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	entityResp, nextEndpoint, err := client.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entityResp[0], mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestEstimate_GenerateInvoice(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.GenerateInvoice(1234)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestEstimate_ListAttachments(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Files
	mockResponseId := int64(1523)
	mockResponse := new(invoiced.File)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	entity, err := client.ListAttachments(1523)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entity[0], mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}
