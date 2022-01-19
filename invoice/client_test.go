package invoice

import (
	"github.com/Invoiced/invoiced-go"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestInvoiceCreate(t *testing.T) {
	key := "test api key"

	mockInvoiceResponse := new(invoiced.Invoice)
	mockInvoiceResponse.Id = 1523

	nowUnix := time.Now().UnixNano()

	s := strconv.FormatInt(nowUnix, 10)

	server, err := invdmockserver.New(200, mockInvoiceResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	_, err = client.Create(&invoiced.InvoiceRequest{Name: invoiced.String("Test invoice Original " + s)})
	if err != nil {
		t.Fatal("Error Creating invoice", err)
	}
}

func TestInvoiceCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server, err := invdmockserver.New(400, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, err = client.Create(&invoiced.InvoiceRequest{Closed: invoiced.Bool(false)})

	if err == nil {
		t.Fatal("Api Should Have Errored Out")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceUpdate(t *testing.T) {
	key := "test api key"

	mockUpdatedTime := time.Now().UnixNano()
	mockInvoiceResponse := new(invoiced.Invoice)
	mockInvoiceResponse.Id = 1523
	mockInvoiceResponse.CreatedAt = mockUpdatedTime
	mockInvoiceResponse.Name = "MOCK invoice"

	mockInvoiceResponse.Balance = 42.22

	server, err := invdmockserver.New(200, mockInvoiceResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, err = client.Update(1523, &invoiced.InvoiceRequest{Name: invoiced.String("Test")})
	if err != nil {
		t.Fatal("Error Updating Client", err)
	}
}

func TestInvoiceUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	server, err := invdmockserver.New(401, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, err = client.Update(1234, &invoiced.InvoiceRequest{Closed: invoiced.Bool(false)})
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

	server, err := invdmockserver.New(204, mockinvoiceResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.Delete(2341)

	if err != nil {
		t.Fatal("Error Occured Deleting Client")
	}
}

func TestInvoiceDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You Do Not Have Permission To Do That"

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal()
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	err = client.Delete(1234)

	if err == nil {
		t.Fatal("Error Occured Deleting Client")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoiceListAllByNumber(t *testing.T) {
	key := "test api key"

	var mockInvoicesResponse invoiced.Invoices
	mockInvoiceNumber := "INV-3421"
	mockInvoiceResponse := new(invoiced.Invoice)
	mockInvoiceResponse.Id = 1523
	mockInvoiceResponse.Number = mockInvoiceNumber
	mockInvoiceResponse.PaymentTerms = "NET15"
	mockInvoiceResponse.CreatedAt = time.Now().UnixNano()

	mockInvoicesResponse = append(mockInvoicesResponse, mockInvoiceResponse)

	server, err := invdmockserver.New(200, mockInvoicesResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, err = client.ListInvoiceByNumber(mockInvoiceNumber)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInvoiceList(t *testing.T) {
	key := "test api key"

	var mockInvoicesResponse invoiced.Invoices
	mockInvoiceNumber := "INV-3421"
	mockInvoiceResponse := new(invoiced.Invoice)
	mockInvoiceResponse.Id = 1523
	mockInvoiceResponse.Number = mockInvoiceNumber
	mockInvoiceResponse.PaymentTerms = "NET15"

	mockInvoiceResponse.CreatedAt = time.Now().UnixNano()

	mockInvoicesResponse = append(mockInvoicesResponse, mockInvoiceResponse)

	server, err := invdmockserver.New(200, mockInvoicesResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, nextEndpoint, err := client.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if nextEndpoint != "" {
		t.Fatal(err)
	}
}

func TestInvoiceListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(invoiced.APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockInvoiceNumber := "INV-32421"

	server, err := invdmockserver.New(403, mockErrorResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, err = client.ListInvoiceByNumber(mockInvoiceNumber)

	if err == nil {
		t.Fatal("Error occured deleting invoice")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoice_Count_Error(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.Invoice

	mockResponse := new(invoiced.Invoice)
	mockResponse.Id = 1234
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

	if result != -1 {
		t.Fatal("Unexpectedly successful")
	}
}

func TestInvoice_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Invoice)
	mockResponse.Id = 1234
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, err = client.Retrieve(1234)
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}
}

func TestInvoice_Void(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Invoice)
	mockResponse.Id = 1234
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"
	mockResponse.Status = "voided"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	_, err = client.Void(1234)
	if err != nil {
		t.Fatal("Error updating entity", err)
	}
}

func TestInvoice_SendEmail(t *testing.T) {
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

func TestInvoice_SendText(t *testing.T) {
	key := "test api key"

	var mockTextResponse [1]invoiced.TextMessage

	mockResponse := new(invoiced.TextMessage)
	mockResponse.Id = "abcdef"
	mockResponse.Message = "hello text"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockTextResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockTextResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	sendResponse, err := client.SendText(1234, nil)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse[0].Message != "hello text" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestInvoice_SendLetter(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Letter)
	mockResponse.Id = "abcdef"
	mockResponse.State = "queued"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	sendResponse, err := client.SendLetter(1234)
	if err != nil {
		t.Fatal("Error with send", err)
	}

	if sendResponse.State != "queued" {
		t.Fatal("Error: send not completed correctly")
	}
}

func TestInvoice_Pay(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.Invoice)
	mockResponse.Id = 1234
	mockResponse.Balance = 0

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	invoice, err := client.Pay(1234)
	if err != nil {
		t.Fatal("Error with pay", err)
	}

	if invoice.Balance != 0 {
		t.Fatal("Error: pay not completed correctly")
	}
}

func TestInvoice_ListAttachments(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Files
	mockResponse := new(invoiced.File)
	mockResponse.Id = 1523
	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	attachments, err := client.ListAttachments(1234)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(attachments[0], mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoice_RetrieveNotes(t *testing.T) {
	key := "test api key"

	var mockResponses invoiced.Notes
	mockResponse := new(invoiced.Note)
	mockResponse.Id = 1523

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockResponses = append(mockResponses, mockResponse)

	server, err := invdmockserver.New(200, mockResponses, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	notes, err := client.RetrieveNotes(1523)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(notes[0], mockResponse) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestInvoice_CreatePaymentPlan(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.PaymentPlan)
	mockResponse.Id = 1234

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.CreatePaymentPlan(1234, &invoiced.PaymentPlanRequest{})
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Id != 1234 {
		t.Fatal("Error: operation not completed correctly")
	}
}

func TestInvoice_RetrievePaymentPlan(t *testing.T) {
	key := "test api key"

	mockResponse := new(invoiced.PaymentPlan)
	mockResponse.Id = 1234

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	subjectEntity, err := client.RetrievePaymentPlan(1234)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if subjectEntity.Id != 1234 {
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

	client := Client{invoiced.NewMockApi(key, server)}

	err = client.CancelPaymentPlan(1234)

	if err != nil {
		t.Fatal("Error occurred during deletion")
	}
}
