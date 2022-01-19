package invoiced

import (
	"fmt"
	"strconv"
)

type InvoiceClient struct {
	*Client
	*Invoice
}

type Invoices []*InvoiceClient

func (c *Client) NewInvoice() *InvoiceClient {
	invoice := new(Invoice)
	return &InvoiceClient{c, invoice}
}

func (c *InvoiceClient) Count() (int64, error) {
	endpoint := InvoiceEndpoint

	count, err := c.Api.Count(endpoint)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *InvoiceClient) Create(request *InvoiceRequest) (*InvoiceClient, error) {
	endpoint := InvoiceEndpoint
	resp := c.NewInvoice()

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InvoiceClient) Retrieve(id int64) (*InvoiceClient, error) {
	url := InvoiceEndpoint + "/" + strconv.FormatInt(id, 10)

	invoice := &InvoiceClient{c.Client, new(Invoice)}

	_, err := c.Api.Get(url, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (c *InvoiceClient) Update(request *InvoiceRequest) error {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Invoice)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Invoice = resp

	return nil
}

func (c *InvoiceClient) Void() (*InvoiceClient, error) {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"
	resp := c.NewInvoice()

	err := c.Api.PostWithoutData(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InvoiceClient) Delete() error {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *InvoiceClient) ListAll(filter *Filter, sort *Sort) (Invoices, error) {
	url := InvoiceEndpoint
	url = AddFilterAndSort(url, filter, sort)

	return c.ListAllHelper(url, filter, sort)
}

func (c *InvoiceClient) ListAllHelper(endpoint string, filter *Filter, sort *Sort) (Invoices, error) {
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

func (c *InvoiceClient) ListHelper(url string, filter *Filter, sort *Sort) (Invoices, string, error) {
	if len(url) == 0 {
		url = InvoiceEndpoint
		url = AddFilterAndSort(url, filter, sort)
	}

	invoicesToReturn := make(Invoices, 0)
	invoices := make(Invoices, 0)

	nextEndpoint, err := c.Api.Get(url, &invoices)
	if err != nil {
		return nil, "", err
	}

	for _, invoice := range invoices {
		inv := c.Client.NewInvoice()
		invData := invoice
		inv.Invoice = &invData
		invoicesToReturn = append(invoicesToReturn, inv)
	}

	return invoicesToReturn, nextEndpoint, nil
}

func (c *InvoiceClient) ListAllInvoicesStartDate(filter *Filter, sort *Sort, invoiceDate int64) (Invoices, error) {
	return c.ListAllInvoicesStartEndDate(filter, sort, invoiceDate, 0)
}

func (c *InvoiceClient) ListAllInvoicesEndDate(filter *Filter, sort *Sort, invoiceDate int64) (Invoices, error) {
	return c.ListAllInvoicesStartEndDate(filter, sort, 0, invoiceDate)
}

func (c *InvoiceClient) ListAllInvoicesStartEndDate(filter *Filter, sort *Sort, startDate, endDate int64) (Invoices, error) {
	url := InvoiceEndpoint
	url = AddFilterAndSort(url, filter, sort)

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

func (c *InvoiceClient) ListAllInvoicesUpdatedDate(filter *Filter, sort *Sort, invoiceDate int64) (Invoices, error) {
	url := InvoiceEndpoint
	url = AddFilterAndSort(url, filter, sort)

	if invoiceDate > 0 {
		updatedAfterString := strconv.FormatInt(invoiceDate, 10)
		url = addQueryParameter(url, "updated_after", updatedAfterString)
	}

	return c.ListAllHelper(url, filter, sort)
}

func (c *InvoiceClient) List(filter *Filter, sort *Sort) (Invoices, string, error) {
	return c.ListHelper("", filter, sort)
}

func (c *InvoiceClient) ListInvoiceByNumber(invoiceNumber string) (*InvoiceClient, error) {
	filter := NewFilter()
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

func (c *InvoiceClient) SendEmail(request *SendEmailRequest) error {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.Api.Create(endpoint, request, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *InvoiceClient) SendText(request *SendTextMessageRequest) (TextMessages, error) {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"
	resp := new(TextMessages)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *InvoiceClient) SendLetter() (*Letter, error) {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"
	resp := new(Letter)

	err := c.Api.Create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InvoiceClient) Pay() error {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/pay"
	invoice := new(Invoice)
	err := c.Api.Create(endpoint, nil, invoice)
	if err != nil {
		return nil
	}

	c.Invoice = invoice

	return nil
}

func (c *InvoiceClient) ListAttachments() (Files, error) {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

	files := make(Files, 0)

NEXT:
	tempFiles := make(Files, 0)

	endpoint, err := c.Api.Get(endpoint, &tempFiles)

	if err != nil {
		return nil, err
	}

	files = append(files, tempFiles...)

	if endpoint != "" {
		goto NEXT
	}

	return files, nil
}

func (c *InvoiceClient) RetrieveNotes() (Notes, error) {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/notes"

	notes := make(Notes, 0)

NEXT:
	tmpNotes := make(Notes, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpNotes)

	if err != nil {
		return nil, err
	}

	notes = append(notes, tmpNotes...)

	if endpoint != "" {
		goto NEXT
	}

	return notes, nil
}

func (c *InvoiceClient) CreatePaymentPlan(request *PaymentPlanRequest) (*PaymentPlan, error) {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_plan"
	resp := new(PaymentPlan)

	err := c.Api.Create(endpoint, request, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InvoiceClient) RetrievePaymentPlan() (*PaymentPlan, error) {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_plan"
	resp := new(PaymentPlan)

	_, err := c.Api.Get(endpoint, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InvoiceClient) CancelPaymentPlan() error {
	endpoint := InvoiceEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *InvoiceClient) String() string {
	header := fmt.Sprintf("<InvoiceClient id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Invoice.String()
}
