package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
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

	conn := mockConnection(key, server)

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

	conn := mockConnection(key, server)

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

func TestCreditNote_Void(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.CreditNote)
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

	entityToUpdate := conn.NewCreditNote()

	entityToUpdate, err = entityToUpdate.Void()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.CreditNote) {
		t.Fatal("Error: entity not voided correctly")
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

	conn := mockConnection(key, server)

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

	conn := mockConnection(key, server)
	entity := conn.NewCreditNote()

	retrievedPayment, err := entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.CreditNote, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestCreditNote_CountErr(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.CreditNote

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

	conn := mockConnection(key, server)
	entity := conn.NewCreditNote()

	result, err := entity.Count()

	if err == nil {
		t.Fatal("Error: ", err)
	}

	if result != int64(-1) {
		t.Fatal("Unexpectedly successful")
	}
}

func TestCreditNote_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.CreditNote

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

	conn := mockConnection(key, server)
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

func TestCreditNote_ListAttachments(t *testing.T) {
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

	entity, err := conn.NewCreditNote().ListAttachments()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(entity[0].File, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestCreditNote_SendEmail(t *testing.T) {
	key := "test api key"

	var mockEmailResponse [1]invdendpoint.EmailResponse

	mockResponse := new(invdendpoint.EmailResponse)
	mockResponse.Id = 1
	mockResponse.Message = "hello test"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockEmailResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockEmailResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewCreditNote()

	sendResponse, err := subjectEntity.SendEmail(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello test" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestCreditNote_SendText(t *testing.T) {
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

	subjectEntity := conn.NewCreditNote()

	sendResponse, err := subjectEntity.SendText(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello text" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestCreditNote_SendLetter(t *testing.T) {
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

	subjectEntity := conn.NewCreditNote()

	sendResponse, err := subjectEntity.SendLetter()
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse.State != "queued" {
		t.Fatal("Error: send not completed correctly")
	}
}
