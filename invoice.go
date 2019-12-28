package invdapi

import (
	"errors"
	"fmt"
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

const defaultExpandInvoice = "items.catalog_item"

type Invoice struct {
	*Connection
	*invdendpoint.Invoice
	IncludeUpdatedAt bool
}

type Invoices []*Invoice

func (c *Connection) NewInvoice() *Invoice {
	invoice := new(invdendpoint.Invoice)
	return &Invoice{c, invoice, false}

}

func (c *Connection) NewPaymentPlanRequest() *invdendpoint.PaymentPlanRequest {
	return &invdendpoint.PaymentPlanRequest{}
}

func (c *Invoice) Count() (int64, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.InvoicesEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Invoice) Create(invoice *Invoice) (*Invoice, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.InvoicesEndPoint)
	invResp := new(Invoice)

	if invoice == nil {
		return nil, errors.New("invoice cannot be nil")
	}

	//safe prune invoice data for creation
	invdInvToCreate,err := SafeInvoiceForCreation(invoice.Invoice)

	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdInvToCreate, invResp)

	if apiErr != nil {
		return nil, apiErr
	}

	invResp.Connection = c.Connection

	return invResp, nil

}

func (c *Invoice) Delete() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Invoice) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id)

	invResp := new(Invoice)

	invDataToUpdate, err := SafeInvoiceForUpdate(c.Invoice)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, invDataToUpdate, invResp)

	if apiErr != nil {
		return apiErr
	}

	c.Invoice = invResp.Invoice

	return nil

}

func (c *Invoice) Retrieve(id int64) (*Invoice, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), id)

	if c.IncludeUpdatedAt {
		endPoint = addIncludeToEndPoint(endPoint, "updated_at")
	}

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	custEndPoint := new(invdendpoint.Invoice)

	invoice := &Invoice{c.Connection, custEndPoint, c.IncludeUpdatedAt}

	_, apiErr := c.retrieveDataFromAPI(endPoint, invoice)

	if apiErr != nil {
		return nil, apiErr
	}

	return invoice, nil

}

func (c *Invoice) Void() (*Invoice, error) {

	invResp := new(Invoice)

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/void"

	apiErr := c.postWithoutData(endPoint,invResp)

	if apiErr != nil {
		return nil,apiErr
	}

	invResp.Connection = c.Connection

	return invResp,nil

}

func (c *Invoice) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.InvoicesEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	if c.IncludeUpdatedAt {
		endPoint = addIncludeToEndPoint(endPoint, "updated_at")
	}

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	invoices := make(Invoices, 0)

NEXT:
	tmpInvoices := make(Invoices, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpInvoices)

	if apiErr != nil {
		return nil, apiErr
	}

	invoices = append(invoices, tmpInvoices...)

	if endPoint != "" {
		goto NEXT
	}

	for _, invoice := range invoices {
		invoice.Connection = c.Connection

	}

	return invoices, nil

}

func (c *Invoice) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, string, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.InvoicesEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)
	if c.IncludeUpdatedAt {
		endPoint = addIncludeToEndPoint(endPoint, "updated_at")
	}

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	invoices := make(Invoices, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &invoices)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, invoice := range invoices {
		invoice.Connection = c.Connection

	}

	return invoices, nextEndPoint, nil

}

func (c *Invoice) ListInvoiceByNumber(invoiceNumber string) (*Invoice, error) {

	filter := invdendpoint.NewFilter()
	err := filter.Set("number", invoiceNumber)

	if err != nil {
		return nil, err
	}

	invoices, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(invoices) == 0 {
		return nil, nil
	}

	return invoices[0], nil

}

func (c *Invoice) SendEmail(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endPoint, emailReq, emailResp)

	if err != nil {
		return nil, err
	}

	return *emailResp, nil

}

func (c *Invoice) SendText(req *invdendpoint.TextRequest) (invdendpoint.TextResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/text_messages"

	resp := new(invdendpoint.TextResponses)

	err := c.create(endPoint, req, resp)

	if err != nil {
		return nil, err
	}

	return *resp, nil

}

func (c *Invoice) SendLetter() (*invdendpoint.LetterResponse, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/letters"

	resp := new(invdendpoint.LetterResponse)

	err := c.create(endPoint, nil, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (c *Invoice) Pay() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/pay"
	invoice := new(invdendpoint.Invoice)
	err := c.create(endPoint, nil, invoice)

	if err != nil {
		return nil
	}

	c.Invoice = invoice

	return nil

}

func (c *Invoice) ListAttachments() (Files, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/attachments"

	files := make(Files, 0)

NEXT:
	tempFiles := make(Files, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tempFiles)

	if apiErr != nil {
		return nil, apiErr
	}

	files = append(files, tempFiles...)

	if endPoint != "" {
		goto NEXT
	}

	for _, invoice := range files {
		invoice.Connection = c.Connection
	}

	return files, nil

}

func (c *Invoice) RetrieveNotes() (Notes, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/notes"

	notes := make(Notes, 0)

NEXT:
	tmpNotes := make(Notes, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpNotes)

	if apiErr != nil {
		return nil, apiErr
	}

	notes = append(notes, tmpNotes...)

	if endPoint != "" {
		goto NEXT
	}

	for _, invoice := range notes {
		invoice.Connection = c.Connection
	}

	return notes, nil

}


func (c *Invoice) CreatePaymentPlan(paymentPlanRequest *invdendpoint.PaymentPlanRequest) (*invdendpoint.PaymentPlan, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/payment_plan"

	if paymentPlanRequest == nil {
		return nil, errors.New("paymentPlanRequest cannot be nil")
	}

	paymentPlanResp := new(invdendpoint.PaymentPlan)

	apiErr := c.create(endPoint, paymentPlanRequest, paymentPlanResp)

	if apiErr != nil {
		return nil, apiErr
	}


	return paymentPlanResp, nil

}


func (c *Invoice) RetrievePaymentPlan() (*invdendpoint.PaymentPlan, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/payment_plan"

	paymentPlanResp := new(invdendpoint.PaymentPlan)

	_, apiErr := c.retrieveDataFromAPI(endPoint, paymentPlanResp)

	if apiErr != nil {
		return nil, apiErr
	}

	return paymentPlanResp, nil

}

func (c *Invoice) CancelPaymentPlan() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Invoice) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Invoice.String()
}

//SafeInvoiceForCreation prunes invoice data for just fields that can be used for creation of a invoice
func SafeInvoiceForCreation(inv *invdendpoint.Invoice) (*invdendpoint.Invoice, error) {
	if inv == nil  {
		return nil, errors.New("Invoice is nil or Invoice.Invoice is nil")
	}

	invData :=new(invdendpoint.Invoice)
	invData.Customer = inv.Customer
	invData.Name = inv.Name
	invData.Number = inv.Number
	invData.Currency = inv.Currency
	invData.PaymentTerms = inv.PaymentTerms
	invData.Date = inv.Date
	invData.DueDate = inv.DueDate
	invData.Draft = inv.Draft
	invData.Closed = inv.Closed
	invData.Items = inv.Items
	invData.Notes = inv.Notes
	invData.Discounts = inv.Discounts
	invData.MetaData = inv.MetaData
	invData.Attachments = inv.Attachments
	invData.DisabledPaymentMethods = inv.DisabledPaymentMethods


	return invData,nil
}

//SafeInvoiceForCreation prunes invoice data for just fields that can be used for creation of a invoice
func SafeInvoiceForUpdate(inv *invdendpoint.Invoice) (*invdendpoint.Invoice, error) {
	if inv == nil  {
		return nil, errors.New("Invoice is nil or Invoice.Invoice is nil")
	}

	invData :=new(invdendpoint.Invoice)
	invData.Name = inv.Name
	invData.Number = inv.Number
	invData.Currency = inv.Currency
	invData.PaymentTerms = inv.PaymentTerms
	invData.Date = inv.Date
	invData.DueDate = inv.DueDate
	invData.Draft = inv.Draft
	invData.Sent = inv.Sent
	invData.Closed = inv.Closed
	invData.Items = inv.Items
	invData.Notes = inv.Notes
	invData.Discounts = inv.Discounts
	invData.MetaData = inv.MetaData
	invData.Attachments = inv.Attachments
	invData.DisabledPaymentMethods = inv.DisabledPaymentMethods


	return invData,nil
}