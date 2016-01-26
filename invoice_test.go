package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
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

	invoiceToCreate := new(invdendpoint.Invoice)

	nowUnix := time.Now().UnixNano()

	s := strconv.FormatInt(nowUnix, 10)

	invoiceToCreate.Name = "Test invoice Original " + s
	mockInvoiceResponse.Name = invoiceToCreate.Name

	server := mockServer(200, mockInvoiceResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdInvoice, apiErr := conn.CreateInvoice(invoiceToCreate)

	if apiErr != nil {
		t.Fatal("Error Creating invoice", apiErr)
	}

	if createdInvoice.Id != mockInvoiceResponseID {
		t.Fatal("invoice was not created succesfully")
	}

}

func TestInvoiceCreateError(t *testing.T) {
	key := "test api key"
	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "Name is invalid"
	mockErrorResponse.Param = "name"

	server := mockServer(400, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	invoiceToCreate := new(invdendpoint.Invoice)
	invoiceToCreate.AmountPaid = 342.234

	_, apiErr := conn.CreateInvoice(invoiceToCreate)

	if apiErr == nil {
		t.Fatal("Api should have errored out")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error messages do not match up")
	}

}

func TestInvoiceUpdate(t *testing.T) {
	key := "test api key"

	mockInvoiceResponseID := int64(1523)
	mockUpdatedTime := time.Now().UnixNano()
	mockInvoiceResponse := new(invdendpoint.Invoice)
	mockInvoiceResponse.Id = mockInvoiceResponseID
	mockInvoiceResponse.UpdatedAt = mockUpdatedTime
	mockInvoiceResponse.Name = "MOCK invoice"

	invoiceToUpdate := new(invdendpoint.Invoice)

	mockInvoiceResponse.Balance = 42.22
	invoiceToUpdate.Balance = 42.22

	server := mockServer(200, mockInvoiceResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	updatedInvoice, apiErr := conn.UpdateInvoice(mockInvoiceResponseID, invoiceToUpdate)

	if apiErr != nil {
		t.Fatal("Error Updating invoice", apiErr)
	}

	if !reflect.DeepEqual(mockInvoiceResponse, updatedInvoice) {
		t.Fatal("Error messages do not match up")
	}

}

func TestInvoiceUpdateError(t *testing.T) {
	key := "wrong api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "We could not authenticate the supplied API Key."

	invoiceID := int64(324234)
	invoiceToUpdate := new(invdendpoint.Invoice)
	invoiceToUpdate.Balance = 400.12

	server := mockServer(401, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.UpdateInvoice(invoiceID, invoiceToUpdate)

	if apiErr == nil {
		t.Fatal("Error Updating invoice", apiErr)
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestInvoiceDelete(t *testing.T) {

	key := "api key"

	mockinvoiceResponse := ""
	mockinvoiceID := int64(2341)

	server := mockServer(204, mockinvoiceResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteInvoice(mockinvoiceID)

	if apiErr != nil {
		t.Fatal("Error occured deleting invoice")
	}

}

func TesIinvoiceDeleteError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockinvoiceID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	apiErr := conn.DeleteInvoice(mockinvoiceID)

	if apiErr == nil {
		t.Fatal("Error occured deleting invoice")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

func TestInvoiceList(t *testing.T) {

	key := "test api key"

	mockInvoiceResponseID := int64(1523)
	mockInvoiceResponse := new(invdendpoint.Invoice)
	mockInvoiceResponse.Id = mockInvoiceResponseID
	mockInvoiceResponse.PaymentTerms = "NET15"

	mockInvoiceResponse.UpdatedAt = time.Now().UnixNano()

	server := mockServer(200, mockInvoiceResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	createdinvoice, apiErr := conn.ListInvoice(mockInvoiceResponseID)

	if apiErr != nil {
		t.Fatal("Error Creating invoice", apiErr)
	}

	if createdinvoice.Id != mockInvoiceResponseID {
		t.Fatal("invoice was not created succesfully")
	}

}

func TestInvoiceListError(t *testing.T) {
	key := "api key"

	mockErrorResponse := new(APIError)
	mockErrorResponse.Type = "invalid_request"
	mockErrorResponse.Message = "You do not have permission to do that"

	mockinvoiceID := int64(-999)

	server := mockServer(403, mockErrorResponse)
	defer server.Close()

	conn := mockConnection(key, server)

	_, apiErr := conn.ListInvoice(mockinvoiceID)

	if apiErr == nil {
		t.Fatal("Error occured deleting invoice")
	}

	if !reflect.DeepEqual(mockErrorResponse, apiErr) {
		t.Fatal("Error Messages Do Not Match Up")
	}

}

// func createMockInvoice(t *testing.T, offset int64, customName string, customPaymentTerms string) *invdendpoint.Invoice {
// 	conn := NewConnection(apikey)

// 	//Create invoice
// 	invoice := createMockinvoice(t, offset)

// 	invoiceToCreate := new(invdendpoint.Invoice)

// 	nowUnix := time.Now().UnixNano() + offset
// 	s := strconv.FormatInt(nowUnix, 10)

// 	invoiceToCreate.invoice = invoice.Id
// 	if customName != "" {
// 		invoiceToCreate.Name = customName
// 	} else {
// 		invoiceToCreate.Name = "MOCK INVOICE " + s
// 	}

// 	if customPaymentTerms != "" {
// 		invoiceToCreate.PaymentTerms = customPaymentTerms
// 	} else {
// 		invoiceToCreate.PaymentTerms = "MOCK INVOICE"
// 	}

// 	lineItem := invdendpoint.LineItem{}
// 	lineItem.Description = "Mock Macbook Pro " + s
// 	lineItem.Quantity = 5
// 	lineItem.UnitCost = 34.23

// 	lineItems := append([]invdendpoint.LineItem{}, lineItem)

// 	invoiceToCreate.Items = lineItems

// 	invoice, apiErr := conn.CreateInvoice(invoiceToCreate)

// 	if apiErr != nil {
// 		t.Fatal("Error Creating Mock invoice =>", apiErr)
// 	}

// 	return invoice

// }

// func deleteMockInvoice(t *testing.T, invoice invdendpoint.Invoice) {

// 	invoiceID := invoice.Id
// 	conn := NewConnection(apikey)

// 	apiErr := conn.DeleteInvoice(invoiceID)

// 	if apiErr != nil {
// 		t.Fatal("Error Deleting Mock Invoice, ID => ", invoiceID, " =>", apiErr)
// 	}

// 	invoiceID := invoice.invoice

// 	apiErr = conn.Deleteinvoice(invoiceID)

// 	if apiErr != nil {
// 		t.Fatal("Error Deleting Mock invoice, ID => ", invoiceID, " =>", apiErr)
// 	}

// }

// func updateMockInvoice(t *testing.T, invoice *invdendpoint.Invoice) {

// 	invoiceID := invoice.Id
// 	conn := NewConnection(apikey)

// 	invoice.Name = "UPDATED " + invoice.Name

// 	apiErr := conn.UpdateInvoice(invoiceID, invoice)

// 	if apiErr != nil {
// 		t.Fatal("Error Updating Mock Invoice, ID => ", invoiceID, " =>", apiErr)
// 	}

// }

// func retrieveMockInvoice(t *testing.T, invoiceID int64) *invdendpoint.Invoice {

// 	conn := NewConnection(apikey)

// 	invoice, apiErr := conn.ListInvoice(invoiceID)

// 	if apiErr != nil {
// 		t.Fatal("Error Retrieving Mock Invoice, ID => ", invoiceID, " =>", apiErr)
// 	}

// 	return invoice

// }

// func deleteAllInvoices(t *testing.T) {
// 	conn := NewConnection(apikey)

// 	var invoiceIDs []int64

// 	invoices, apiErr := conn.ListAllInvoicesAuto(nil, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error: Getting All Invoices Auto", apiErr)
// 	}

// 	for _, invoice := range *invoices {
// 		invoiceIDs = append(invoiceIDs, invoice.Id)
// 	}

// 	for _, invoiceID := range invoiceIDs {
// 		apiErr = conn.DeleteInvoice(invoiceID)

// 		if apiErr != nil {
// 			t.Fatal("Error Deleting A Invoice with ID => ", invoiceID, apiErr)
// 		}

// 	}

// }

// func TestDeleteAllInvoices(t *testing.T) {
// 	deleteAllInvoices(t)
// }

// func TestGetAllInvoicesAuto(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	invoices, apiErr := conn.ListAllInvoicesAuto(nil, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error: Getting All Invoices Auto", apiErr)
// 	}

// 	invoiceCount, apiErr := conn.CountInvoice()

// 	if apiErr != nil {
// 		t.Fatal("Error: Getting invoice Count", apiErr)
// 	}

// 	if int64(len(*invoices)) != invoiceCount {
// 		t.Fatal("Error: Number of invoices Returned Should Equal The invoice Count ")
// 	}

// }

// func TestGetAllinvoices(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	invoices, nextEndPoint, apiErr := conn.ListAllInvoices(nil, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error Getting All Invoices", apiErr)
// 	}

// 	invoiceCount, apiErr := conn.CountInvoice()

// 	if apiErr != nil {
// 		t.Fatal("Error Getting invoice Count", apiErr)
// 	}

// 	if invoiceCount > 100 && nextEndPoint == "" {
// 		t.Fatal("invoice Count Is Greater Than 100, So The Next EndPoint Should Not Be Empty")
// 	}

// 	if len(*invoices) > 100 {
// 		t.Fatal("invoice Returned Should Less Than OR Equal To 100")
// 	}

// }

// func TestListAllInvoicesFiltered(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	invoiceCollection := new(invdendpoint.Invoices)

// 	numberOfTestInvoices := 56

// 	paymentTerm := "TGAIF"

// 	for i := 0; i < numberOfTestInvoices; i++ {
// 		invoice := createMockInvoice(t, int64(i), "", paymentTerm)

// 		*invoiceCollection = append(*invoiceCollection, *invoice)
// 	}

// 	filter := invdendpoint.NewFilter()

// 	filter.Set("payment_terms", paymentTerm)

// 	invoices, apiErr := conn.ListAllInvoicesAuto(filter, nil)

// 	if apiErr != nil {
// 		t.Fatal("Error Getting All Invoices Auto ", apiErr)
// 	}

// 	if len(*invoices) != numberOfTestInvoices {
// 		t.Fatal("The Correct Amount of Invoices Were Not Filtered", len(*invoices), "Not Equal To", numberOfTestInvoices)
// 	}

// 	//Delete invoices

// 	for _, invoice := range *invoiceCollection {
// 		deleteMockInvoice(t, invoice)

// 	}

// }

// func TestGetAllInvoicesFilteredSorted(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	invoiceCollection := new(invdendpoint.Invoices)

// 	//Create invoice A

// 	invoicePaymentTerms := "TGAIF"

// 	invoiceName := "A" + strconv.FormatInt(time.Now().Unix()+int64(0), 10)

// 	invoice := createMockInvoice(t, 0, invoiceName, invoicePaymentTerms)

// 	*invoiceCollection = append(*invoiceCollection, *invoice)
// 	//Create invoice H
// 	invoiceName = "H" + strconv.FormatInt(time.Now().Unix()+int64(0), 10)

// 	invoice = createMockInvoice(t, 0, invoiceName, invoicePaymentTerms)

// 	*invoiceCollection = append(*invoiceCollection, *invoice)

// 	//Create invoice L
// 	invoiceName = "L" + strconv.FormatInt(time.Now().Unix()+int64(0), 10)

// 	invoice = createMockInvoice(t, 0, invoiceName, invoicePaymentTerms)

// 	*invoiceCollection = append(*invoiceCollection, *invoice)

// 	//Create invoice C
// 	invoiceName = "C" + strconv.FormatInt(time.Now().Unix()+int64(0), 10)

// 	invoice = createMockInvoice(t, 0, invoiceName, invoicePaymentTerms)

// 	*invoiceCollection = append(*invoiceCollection, *invoice)

// 	filter := invdendpoint.NewFilter()
// 	filter.Set("payment_terms", invoicePaymentTerms)

// 	sort := invdendpoint.NewSort()

// 	sort.Set("name", invdendpoint.DESC)

// 	invoices, _, apiErr := conn.ListAllInvoices(filter, sort)

// 	if apiErr != nil {
// 		t.Fatal("Error Getting All invoices Auto ", apiErr)
// 	}

// 	if (*invoices)[0].Name[0:1] != "L" {
// 		t.Fatal("Sort DESC Failed With 'L'")
// 	}

// 	if (*invoices)[1].Name[0:1] != "H" {
// 		t.Fatal("Sort DESC Failed With 'H'")
// 	}
// 	if (*invoices)[2].Name[0:1] != "C" {
// 		t.Fatal("Sort DESC Failed With 'C'")
// 	}

// 	if (*invoices)[3].Name[0:1] != "A" {
// 		t.Fatal("Sort DESC Failed With 'A'")
// 	}

// 	//Delete invoices

// 	for _, invoice := range *invoiceCollection {
// 		deleteMockInvoice(t, invoice)
// 	}

// }

// func TestInvoiceCRUD(t *testing.T) {

// 	//Create invoice
// 	createdInvoice := createMockInvoice(t, 0, "", "")

// 	//Update invoice

// 	oldInvoiceName := createdInvoice.Name

// 	updateMockInvoice(t, createdInvoice)

// 	//Retrieve Invoice
// 	updatedInvoice := retrieveMockInvoice(t, createdInvoice.Id)

// 	if updatedInvoice.Name != "UPDATED "+oldInvoiceName {
// 		t.Fatal("Invoice Was Not Updated correctly")
// 	}

// 	deleteMockInvoice(t, *updatedInvoice)

// }

// func TestInvoiceCount(t *testing.T) {

// 	conn := NewConnection(apikey)

// 	_, apiErr := conn.CountInvoice()

// 	if apiErr != nil {
// 		t.Fatal("apiError Should Be Empty")
// 	}

// }
