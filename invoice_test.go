package invdapi

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestInvoiceCreate(t *testing.T) {
	key := "test api key"

	mockInvoiceResponseID := int64(1523)
	mockInvoiceResponse := new(invdendpoint.Invoice)
	mockInvoiceResponse.Id = mockInvoiceResponseID

	nowUnix := time.Now().UnixNano()

	s := strconv.FormatInt(nowUnix, 10)

	server, err := invdmockserver.New(200, mockInvoiceResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	invoice := conn.NewInvoice()

	invoiceToCreate := invoice.NewInvoice()

	invoiceToCreate.Name = "Test invoice Original " + s
	mockInvoiceResponse.Name = invoiceToCreate.Name

	createdInvoice, err := invoice.Create(invoiceToCreate)
	if err != nil {
		t.Fatal("Error Creating invoice", err)
	}

	if !reflect.DeepEqual(createdInvoice.Invoice, mockInvoiceResponse) {
		t.Fatal("Invoice Not Created Succesfully")
	}
}

func TestInvoiceCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server, err := invdmockserver.New(400, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoice := conn.NewInvoice()

	invoiceToCreate := invoice.NewInvoice()
	invoiceToCreate.Total = 342.234

	_, err = invoice.Create(invoiceToCreate)

	if err == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceUpdate(t *testing.T) {
	key := "test api key"

	mockInvoiceResponseID := int64(1523)
	mockUpdatedTime := time.Now().UnixNano()
	mockInvoiceResponse := new(invdendpoint.Invoice)
	mockInvoiceResponse.Id = mockInvoiceResponseID
	mockInvoiceResponse.CreatedAt = mockUpdatedTime
	mockInvoiceResponse.Name = "MOCK invoice"

	mockInvoiceResponse.Balance = 42.22

	server, err := invdmockserver.New(200, mockInvoiceResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoiceToUpdate := conn.NewInvoice()
	invoiceToUpdate.Balance = 42.22

	err = invoiceToUpdate.Save()

	if err != nil {
		t.Fatal("Error Updating Invoice", err)
	}

	if !reflect.DeepEqual(mockInvoiceResponse, invoiceToUpdate.Invoice) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	server, err := invdmockserver.New(401, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoiceToUpdate := conn.NewInvoice()

	invoiceToUpdate.Balance = 400.12

	err = invoiceToUpdate.Save()

	if err == nil {
		t.Fatal("Error Updating invoice", err)
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceDelete(t *testing.T) {
	key := "api key"

	mockinvoiceResponse := ""
	mockinvoiceID := int64(2341)

	server, err := invdmockserver.New(204, mockinvoiceResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoice := conn.NewInvoice()

	invoice.Id = mockinvoiceID

	err = invoice.Delete()

	if err != nil {
		t.Fatal("Error Occured Deleting Invoice")
	}
}

func TestInvoiceDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You Do Not Have Permission To Do That"

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal()
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoice := conn.NewInvoice()

	err = invoice.Delete()

	if err == nil {
		t.Fatal("Error Occured Deleting Invoice")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceListAllByNumber(t *testing.T) {
	key := "test api key"

	var mockInvoicesResponse invdendpoint.Invoices
	mockInvoiceResponseID := int64(1523)
	mockInvoiceNumber := "INV-3421"
	mockInvoiceResponse := new(invdendpoint.Invoice)
	mockInvoiceResponse.Id = mockInvoiceResponseID
	mockInvoiceResponse.Number = mockInvoiceNumber
	mockInvoiceResponse.PaymentTerms = "NET15"

	mockInvoiceResponse.CreatedAt = time.Now().UnixNano()

	mockInvoicesResponse = append(mockInvoicesResponse, *mockInvoiceResponse)

	server, err := invdmockserver.New(200, mockInvoicesResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoice := conn.NewInvoice()

	invoiceResp, err := invoice.ListInvoiceByNumber(mockInvoiceNumber)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(invoiceResp.Invoice, mockInvoiceResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceList(t *testing.T) {
	key := "test api key"

	var mockInvoicesResponse invdendpoint.Invoices
	mockInvoiceResponseID := int64(1523)
	mockInvoiceNumber := "INV-3421"
	mockInvoiceResponse := new(invdendpoint.Invoice)
	mockInvoiceResponse.Id = mockInvoiceResponseID
	mockInvoiceResponse.Number = mockInvoiceNumber
	mockInvoiceResponse.PaymentTerms = "NET15"

	mockInvoiceResponse.CreatedAt = time.Now().UnixNano()

	mockInvoicesResponse = append(mockInvoicesResponse, *mockInvoiceResponse)

	server, err := invdmockserver.New(200, mockInvoicesResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoice := conn.NewInvoice()

	invoiceResp, nextEndpoint, err := invoice.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(invoiceResp[0].Invoice, mockInvoiceResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockInvoiceNumber := "INV-32421"

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	invoice := conn.NewInvoice()

	_, err = invoice.ListInvoiceByNumber(mockInvoiceNumber)

	if err == nil {
		t.Fatal("Error occured deleting invoice")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoice_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.Invoice

	mockResponse := new(invdendpoint.Invoice)
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
	entity := conn.NewInvoice()

	result, err := entity.Count()

	if err == nil {
		t.Fatal("Error: ", err)
	}

	if result != int64(-1) {
		t.Fatal("Unexpectedly successful")
	}
}

func TestInvoice_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewInvoice()

	retrievedEntity, err := entity.Retrieve(int64(1234))
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedEntity.Invoice, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestInvoice_Void(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Invoice)
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

	entityToUpdate := conn.NewInvoice()

	entityToUpdate, err = entityToUpdate.Void()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Invoice) {
		t.Fatal("Error: entity not voided correctly")
	}
}

func TestInvoice_SendEmail(t *testing.T) {
	key := "test api key"

	var mockEmailResponse [1]invdendpoint.EmailResponse

	mockResponse := new(invdendpoint.EmailResponse)
	mockResponse.Id = 3
	mockResponse.Message = "hello test"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockEmailResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockEmailResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewInvoice()

	sendResponse, err := subjectEntity.SendEmail(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello test" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestInvoice_SendText(t *testing.T) {
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

	subjectEntity := conn.NewInvoice()

	sendResponse, err := subjectEntity.SendText(nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello text" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestInvoice_SendLetter(t *testing.T) {
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

	subjectEntity := conn.NewInvoice()

	sendResponse, err := subjectEntity.SendLetter()
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse.State != "queued" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestInvoice_Pay(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Invoice)
	mockResponse.Id = int64(1234)
	mockResponse.Balance = float64(0)

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity := conn.NewInvoice()

	err = subjectEntity.Pay()

	if err != nil {
		t.Fatal("Error with pay", err)
	}

	if subjectEntity.Balance != float64(0) {
		t.Fatal("Error: pay not completed correctly")
	}
}

func TestInvoice_ListAttachments(t *testing.T) {
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

	entity := conn.NewInvoice()
	entity.Id = 2
	attachments, err := entity.ListAttachments()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(attachments[0].File, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoice_RetrieveNotes(t *testing.T) {
	key := "test api key"

	var mockResponses invdendpoint.Notes
	mockResponseId := int64(1523)
	mockResponse := new(invdendpoint.Note)
	mockResponse.Id = mockResponseId

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, *mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity := conn.NewInvoice()
	entity.Id = 2
	notes, err := entity.RetrieveNotes()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(notes[0].Note, mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoice_CreatePaymentPlan(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.PaymentPlan)
	mockResponse.Id = int64(1234)

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	subjectEntity, err := conn.NewInvoice().CreatePaymentPlan(conn.NewPaymentPlanRequest())
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Id != int64(1234) {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestInvoice_RetrievePaymentPlan(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.PaymentPlan)
	mockResponse.Id = int64(1234)

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	defaultEntity := conn.NewInvoice()
	subjectEntity, err := defaultEntity.RetrievePaymentPlan()
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Id != int64(1234) {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestInvoice_CancelPaymentPlan(t *testing.T) {
	key := "api key"

	server, err := invdmockserver.New(204, nil, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	err = conn.NewInvoice().CancelPaymentPlan()

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}
