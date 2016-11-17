package invdapi

import (
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

func (c *Invoice) Count() (int64, error) {
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Invoice) Create(invoice *Invoice) (*Invoice, error) {
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)
	invResp := new(Invoice)

	apiErr := c.create(endPoint, invoice, invResp)

	if apiErr != nil {
		return nil, apiErr
	}

	invResp.Connection = c.Connection

	return invResp, nil

}

func (c *Invoice) Delete() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Invoice) Save() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id)
	invResp := new(Invoice)
	apiErr := c.update(endPoint, c, invResp)

	if apiErr != nil {
		return apiErr
	}

	c.Invoice = invResp.Invoice

	return nil

}

func (c *Invoice) Retrieve(id int64) (*Invoice, error) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), id)

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

func (c *Invoice) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Invoices, error) {
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)
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
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)
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
	filter.Set("number", invoiceNumber)

	invoices, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(invoices) == 0 {
		return nil, nil
	}

	return invoices[0], nil

}

func (c *Invoice) Send(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endPoint, emailReq, emailResp)

	if err != nil {
		return nil, err
	}

	return *emailResp, nil

}

func (c *Invoice) Pay() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/pay"
	invoice := new(invdendpoint.Invoice)
	err := c.create(endPoint, nil, invoice)

	if err != nil {
		return nil
	}

	c.Invoice = invoice

	return nil

}

func (c *Invoice) ListAttachements() (Files, error) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/attachments"
	files := make(Files, 0)
	err := c.create(endPoint, nil, files)

	if err != nil {
		return nil, err
	}

	return files, nil

}

func (c *Invoice) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Invoice.String()
}
