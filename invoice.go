package invdapi

import (
	"fmt"
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Invoice struct {
	*Connection
	*invdendpoint.Invoice
}

type Invoices []*Invoice

func (c *Connection) NewInvoice() *Invoice {
	invoice := new(invdendpoint.Invoice)
	return &Invoice{c, invoice}
}

func (c *Invoice) Count() (int64, error) {
	endpoint := invdendpoint.InvoiceEndpoint

	count, err := c.count(endpoint)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Invoice) Create(request *invdendpoint.InvoiceRequest) (*Invoice, error) {
	endpoint := invdendpoint.InvoiceEndpoint
	resp := c.NewInvoice()

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Invoice) Retrieve(id int64) (*Invoice, error) {
	url := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(id, 10)

	invoice := &Invoice{c.Connection, new(invdendpoint.Invoice)}

	_, err := c.retrieveDataFromAPI(url, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (c *Invoice) Update(request *invdendpoint.InvoiceRequest) error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(invdendpoint.Invoice)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Invoice = resp

	return nil
}

func (c *Invoice) Void() (*Invoice, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"
	resp := c.NewInvoice()

	err := c.postWithoutData(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Invoice) Delete() error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Invoice) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, error) {
	url := invdendpoint.InvoiceEndpoint
	url = addFilterAndSort(url, filter, sort)

	return c.ListAllHelper(url, filter, sort)
}

func (c *Invoice) ListAllHelper(endpoint string, filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, error) {
	invoices := make(Invoices, 0)
NEXT:

	tmpInvoices, endpoint, err := c.ListHelper(endpoint, filter, sort)

	if err != nil {
		return nil, err
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
	}

	invoicesToReturn := make(Invoices, 0)
	invoices := make(invdendpoint.Invoices, 0)

	nextEndpoint, err := c.retrieveDataFromAPI(url, &invoices)
	if err != nil {
		return nil, "", err
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

func (c *Invoice) SendEmail(request *invdendpoint.SendEmailRequest) error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.create(endpoint, request, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Invoice) SendText(request *invdendpoint.SendTextMessageRequest) (invdendpoint.TextMessages, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"
	resp := new(invdendpoint.TextMessages)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *Invoice) SendLetter() (*invdendpoint.Letter, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"
	resp := new(invdendpoint.Letter)

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

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tempFiles)

	if err != nil {
		return nil, err
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

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpNotes)

	if err != nil {
		return nil, err
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

func (c *Invoice) CreatePaymentPlan(request *invdendpoint.PaymentPlanRequest) (*invdendpoint.PaymentPlan, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_plan"
	resp := new(invdendpoint.PaymentPlan)

	err := c.create(endpoint, request, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Invoice) RetrievePaymentPlan() (*invdendpoint.PaymentPlan, error) {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_plan"
	resp := new(invdendpoint.PaymentPlan)

	_, err := c.retrieveDataFromAPI(endpoint, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Invoice) CancelPaymentPlan() error {
	endpoint := invdendpoint.InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Invoice) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Invoice.String()
}
