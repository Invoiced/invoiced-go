package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

func (c *Connection) ListAllInvoicesAuto(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (*invdendpoint.Invoices, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	invoices := new(invdendpoint.Invoices)

NEXT:
	tmpInvoices := new(invdendpoint.Invoices)
	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, invoices)

	if apiErr != nil {
		return nil, apiErr
	}

	*invoices = append(*invoices, *tmpInvoices...)

	if endPoint != "" {
		goto NEXT
	}

	return invoices, apiErr

}

func (c *Connection) ListAllInvoices(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (*invdendpoint.Invoices, string, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	invoices := new(invdendpoint.Invoices)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, invoices)

	if apiErr != nil {
		return nil, "", apiErr
	}

	return invoices, nextEndPoint, apiErr

}

func (c *Connection) ListInvoice(id int64) (*invdendpoint.Invoice, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), id)

	invoice := new(invdendpoint.Invoice)

	_, apiErr := c.retrieveDataFromAPI(endPoint, invoice)

	if apiErr != nil {
		return nil, apiErr
	}

	return invoice, apiErr

}

func (c *Connection) CountInvoice() (int64, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)

	count, apiErr := c.count(endPoint)

	return count, apiErr

}

func (c *Connection) CreateInvoice(invoice *invdendpoint.Invoice) (*invdendpoint.Invoice, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.InvoicesEndPoint)
	invoiceResponse := new(invdendpoint.Invoice)

	apiErr := c.create(endPoint, invoice, invoiceResponse)

	if apiErr != nil {
		return nil, apiErr
	}

	return invoiceResponse, apiErr

}

func (c *Connection) UpdateInvoice(id int64, invoice *invdendpoint.Invoice) (*invdendpoint.Invoice, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), id)
	invoiceResponse := new(invdendpoint.Invoice)

	apiErr := c.update(endPoint, invoice, invoiceResponse)

	if apiErr != nil {
		return nil, apiErr
	}

	return invoiceResponse, apiErr

}

func (c *Connection) DeleteInvoice(id int64) *APIError {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.InvoicesEndPoint), id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return apiErr

}
