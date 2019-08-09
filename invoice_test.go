package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"strconv"
	"testing"
	"time"
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

	conn := MockConnection(key, server)

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

	conn := MockConnection(key, server)

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

	conn := MockConnection(key, server)

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

	conn := MockConnection(key, server)

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

	conn := MockConnection(key, server)

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

	conn := MockConnection(key, server)

	invoice := conn.NewInvoice()

	err = invoice.Delete()

	if err == nil {
		t.Fatal("Error Occured Deleting Invoice")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
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

	conn := MockConnection(key, server)

	invoice := conn.NewInvoice()

	invoiceResp, err := invoice.ListInvoiceByNumber(mockInvoiceNumber)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(invoiceResp.Invoice, mockInvoiceResponse) {
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

	conn := MockConnection(key, server)

	invoice := conn.NewInvoice()

	_, err = invoice.ListInvoiceByNumber(mockInvoiceNumber)

	if err == nil {
		t.Fatal("Error occured deleting invoice")
	}

	if !reflect.DeepEqual(mockErrorResponse.Error(), err.Error()) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}
