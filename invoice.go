package invdapi

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

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
	endpoint := invdendpoint.InvoiceEndpoint

	count, apiErr := c.count(endpoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil
}

func (c *Invoice) Create(invoice *Invoice) (*Invoice, error) {
	endpoint := invdendpoint.InvoiceEndpoint
	invResp := c.NewInvoice()

	if invoice == nil {
		return nil, errors.New("invoice cannot be nil")
	}

	// safe prune invoice data for creation
	invdInvToCreate, err := SafeInvoiceForCreation(invoice.Invoice)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endpoint, invdInvToCreate, invResp)

	if apiErr != nil {
		return nil, apiErr
	}

	return invResp, nil
}

func (c *Invoice) Delete() error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Invoice) Save() error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	invResp := new(invdendpoint.Invoice)

	invDataToUpdate, err := SafeInvoiceForUpdate(c.Invoice)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, invDataToUpdate, invResp)

	if apiErr != nil {
		return apiErr
	}

	c.Invoice = invResp

	return nil
}

func (c *Invoice) Retrieve(id int64) (*Invoice, error) {
	url := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(id, 10)

	if c.IncludeUpdatedAt {
		url = addQueryParameter(url, "include", "updated_at")
	}

	custEndpoint := new(invdendpoint.Invoice)

	invoice := &Invoice{c.Connection, custEndpoint, c.IncludeUpdatedAt}

	_, apiErr := c.retrieveDataFromAPI(url, invoice)

	if apiErr != nil {
		return nil, apiErr
	}

	return invoice, nil
}

func (c *Invoice) Void() (*Invoice, error) {
	invResp := c.NewInvoice()

	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"

	apiErr := c.postWithoutData(endpoint, invResp)

	if apiErr != nil {
		return nil, apiErr
	}

	return invResp, nil
}

func (c *Invoice) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, error) {
	url := invdendpoint.InvoiceEndpoint
	url = addFilterAndSort(url, filter, sort)

	if c.IncludeUpdatedAt {
		url = addQueryParameter(url, "include", "updated_at")
	}

	return c.ListAllHelper(url, filter, sort)
}

func (c *Invoice) ListAllHelper(endpoint string, filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, error) {
	invoices := make(Invoices, 0)
NEXT:

	tmpInvoices, endpoint, apiErr := c.ListHelper(endpoint, filter, sort)

	if apiErr != nil {
		return nil, apiErr
	}

	invoices = append(invoices, tmpInvoices...)

	if endpoint != "" {
		goto NEXT
	}

	return invoices, nil
}

func (c *Invoice) ListHelper(url string, filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, string, error) {
	if len(url) == 0 {
		url = invdendpoint.InvoiceEndpoint
		url = addFilterAndSort(url, filter, sort)
		if c.IncludeUpdatedAt {
			url = addQueryParameter(url, "include", "updated_at")
		}
	}

	invoicesToReturn := make(Invoices, 0)
	invoices := make(invdendpoint.Invoices, 0)

	nextEndpoint, apiErr := c.retrieveDataFromAPI(url, &invoices)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, invoice := range invoices {
		inv := c.Connection.NewInvoice()
		invData := invoice
		inv.Invoice = &invData
		invoicesToReturn = append(invoicesToReturn, inv)

	}

	return invoicesToReturn, nextEndpoint, nil
}

func (c *Invoice) ListAllInvoicesStartDate(filter *invdendpoint.Filter, sort *invdendpoint.Sort, invoiceDate int64) (Invoices, error) {
	return c.ListAllInvoicesStartEndDate(filter, sort, invoiceDate, 0)
}

func (c *Invoice) ListAllInvoicesEndDate(filter *invdendpoint.Filter, sort *invdendpoint.Sort, invoiceDate int64) (Invoices, error) {
	return c.ListAllInvoicesStartEndDate(filter, sort, 0, invoiceDate)
}

func (c *Invoice) ListAllInvoicesStartEndDate(filter *invdendpoint.Filter, sort *invdendpoint.Sort, startDate, endDate int64) (Invoices, error) {
	url := invdendpoint.InvoiceEndpoint
	url = addFilterAndSort(url, filter, sort)


	if startDate > 0 {
		startDateString := strconv.FormatInt(startDate, 10)
		url = addQueryParameter(url, "start_date", startDateString)
	}

	if endDate > 0 {
		endDateString := strconv.FormatInt(endDate, 10)
		url = addQueryParameter(url, "end_date", endDateString)
	}

	return c.ListAllHelper(url, filter, sort)
}

func (c *Invoice) ListAllInvoicesUpdatedDate(filter *invdendpoint.Filter, sort *invdendpoint.Sort, invoiceDate int64) (Invoices, error) {
	url := invdendpoint.InvoiceEndpoint
	url = addFilterAndSort(url, filter, sort)

	if invoiceDate > 0 {
		updatedAfterString := strconv.FormatInt(invoiceDate, 10)
		url = addQueryParameter(url, "updated_after", updatedAfterString)
	}

	return c.ListAllHelper(url, filter, sort)
}

func (c *Invoice) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, string, error) {
	return c.ListHelper("", filter, sort)
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
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endpoint, emailReq, emailResp)
	if err != nil {
		return nil, err
	}

	return *emailResp, nil
}

func (c *Invoice) SendText(req *invdendpoint.TextRequest) (invdendpoint.TextResponses, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	resp := new(invdendpoint.TextResponses)

	err := c.create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *Invoice) SendLetter() (*invdendpoint.LetterResponse, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	resp := new(invdendpoint.LetterResponse)

	err := c.create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Invoice) Pay() error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/pay"
	invoice := new(invdendpoint.Invoice)
	err := c.create(endpoint, nil, invoice)
	if err != nil {
		return nil
	}

	c.Invoice = invoice

	return nil
}

func (c *Invoice) ListAttachments() (Files, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

	files := make(Files, 0)

NEXT:
	tempFiles := make(Files, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tempFiles)

	if apiErr != nil {
		return nil, apiErr
	}

	files = append(files, tempFiles...)

	if endpoint != "" {
		goto NEXT
	}

	for _, invoice := range files {
		invoice.Connection = c.Connection
	}

	return files, nil
}

func (c *Invoice) RetrieveNotes() (Notes, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/notes"

	notes := make(Notes, 0)

NEXT:
	tmpNotes := make(Notes, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpNotes)

	if apiErr != nil {
		return nil, apiErr
	}

	notes = append(notes, tmpNotes...)

	if endpoint != "" {
		goto NEXT
	}

	for _, invoice := range notes {
		invoice.Connection = c.Connection
	}

	return notes, nil
}

func (c *Invoice) CreatePaymentPlan(paymentPlanRequest *invdendpoint.PaymentPlanRequest) (*invdendpoint.PaymentPlan, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_plan"

	if paymentPlanRequest == nil {
		return nil, errors.New("paymentPlanRequest cannot be nil")
	}

	paymentPlanResp := new(invdendpoint.PaymentPlan)

	apiErr := c.create(endpoint, paymentPlanRequest, paymentPlanResp)

	if apiErr != nil {
		return nil, apiErr
	}

	return paymentPlanResp, nil
}

func (c *Invoice) RetrievePaymentPlan() (*invdendpoint.PaymentPlan, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_plan"

	paymentPlanResp := new(invdendpoint.PaymentPlan)

	_, apiErr := c.retrieveDataFromAPI(endpoint, paymentPlanResp)

	if apiErr != nil {
		return nil, apiErr
	}

	return paymentPlanResp, nil
}

func (c *Invoice) CancelPaymentPlan() error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Invoice) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Invoice.String()
}

// SafeInvoiceForCreation prunes invoice data for just fields that can be used for creation of a invoice
func SafeInvoiceForCreation(inv *invdendpoint.Invoice) (*invdendpoint.Invoice, error) {
	if inv == nil {
		return nil, errors.New("Invoice is nil or Invoice.Invoice is nil")
	}

	invData := new(invdendpoint.Invoice)
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
	invData.Metadata = inv.Metadata
	invData.Attachments = inv.Attachments
	invData.DisabledPaymentMethods = inv.DisabledPaymentMethods
	invData.Taxes = inv.Taxes
	invData.AutoPay = inv.AutoPay
	invData.ShipTo = inv.ShipTo
	invData.PurchaseOrder = inv.PurchaseOrder
	invData.Sent = inv.Sent

	return invData, nil
}

// SafeInvoiceForCreation prunes invoice data for just fields that can be used for creation of a invoice
func SafeInvoiceForUpdate(inv *invdendpoint.Invoice) (*invdendpoint.Invoice, error) {
	if inv == nil {
		return nil, errors.New("Invoice is nil or Invoice.Invoice is nil")
	}

	invData := new(invdendpoint.Invoice)
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
	invData.Metadata = inv.Metadata
	invData.Attachments = inv.Attachments
	invData.DisabledPaymentMethods = inv.DisabledPaymentMethods
	invData.Taxes = inv.Taxes
	invData.AutoPay = inv.AutoPay
	invData.ShipTo = inv.ShipTo
	invData.PurchaseOrder = inv.PurchaseOrder
	invData.Sent = inv.Sent

	return invData, nil
}
