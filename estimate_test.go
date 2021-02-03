package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
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

	conn := mockConnection(key, server)

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

	conn := mockConnection(key, server)

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

	conn := mockConnection(key, server)

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

	conn := mockConnection(key, server)
	entity := conn.NewEstimate()

	retrievedPayment, err := entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.Estimate, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestEstimate_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.Estimate

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

	conn := mockConnection(key, server)
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

func TestEstimate_Void(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Estimate)
	mockResponse.Id = int64(1234)
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"
	mockResponse.Status = "voided"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	entityToUpdate := conn.NewEstimate()

	entityToUpdate, err = entityToUpdate.Void()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Estimate) {
		t.Fatal("Error: entity not voided correctly")
	}
}

func TestEstimate_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.Estimate

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

	conn := mockConnection(key, server)
	entity := conn.NewEstimate()

	result, err := entity.Count()

	if err == nil {
		t.Fatal("Error: ", err)
	}

	if result != int64(-1) {
		t.Fatal("Unexpectedly successful")
	}
}

func TestEstimate_SendEmail(t *testing.T) {
	key := "test api key"

	var mockEmailResponse [1]invdendpoint.EmailResponse

	mockResponse := new(invdendpoint.EmailResponse)
	mockResponse.Id = 2
	mockResponse.Message = "hello test"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockEmailResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockEmailResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewEstimate()

	sendResponse, err := subjectEntity.SendEmail(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello test" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestEstimate_SendText(t *testing.T) {
	key := "test api key"

	var mockTextResponse [1]invdendpoint.TextResponse

	mockResponse := new(invdendpoint.TextResponse)
	mockResponse.Id = "abcdef"
	mockResponse.Message = "hello text"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockTextResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockTextResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewEstimate()

	sendResponse, err := subjectEntity.SendText(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello text" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestEstimate_SendLetter(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.LetterResponse)
	mockResponse.Id = "abcdef"
	mockResponse.State = "queued"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewEstimate()

	sendResponse, err := subjectEntity.SendLetter()
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse.State != "queued" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestEstimate_List(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Estimates
	mockResponseId := int64(1523)
	mockNumber := "INV-3421"
	mockResponse := new(invdendpoint.Estimate)
	mockResponse.Id = mockResponseId
	mockResponse.Number = mockNumber
	mockResponse.PaymentTerms = "NET15"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity := conn.NewEstimate()

	entityResp, nextEndpoint, err := entity.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entityResp[0].Estimate, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestEstimate_GenerateInvoice(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "entity example"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewEstimate()
	subjectEntity, err := defaultEntity.GenerateInvoice()
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Name != "entity example" {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestEstimate_ListAttachments(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Files
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.File)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity, err := conn.NewEstimate().ListAttachments()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entity[0].File, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}
