package customer

import (
	"github.com/Invoiced/invoiced-go/v2"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.CustomerRequest) (*invoiced.Customer, error) {
	resp := new(invoiced.Customer)
	err := c.Api.Create("/customers", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.Customer, error) {
	resp := new(invoiced.Customer)
	_, err := c.Api.Get("/customers/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) RetrieveAccountingSyncStatus(id int64) (*invoiced.AccountingSyncStatus, error) {
	resp := new(invoiced.AccountingSyncStatus)
	_, err := c.Api.Get("/customers/"+strconv.FormatInt(id, 10)+ "/accounting_sync_status", resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.CustomerRequest) (*invoiced.Customer, error) {
	endpoint := "/customers/" + strconv.FormatInt(id, 10)
	resp := new(invoiced.Customer)
	err := c.Api.Update(endpoint, request, resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/customers/" + strconv.FormatInt(id, 10))
}

func (c *Client) Count() (int64, error) {
	return c.Api.Count("/customers")
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Customers, error) {
	endpoint := invoiced.AddFilterAndSort("/customers", filter, sort)

	customers := make(invoiced.Customers, 0)

NEXT:
	tmpCustomers := make(invoiced.Customers, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpCustomers)

	if err != nil {
		return nil, err
	}

	customers = append(customers, tmpCustomers...)

	if endpoint != "" {
		goto NEXT
	}

	return customers, nil
}

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Customers, string, error) {
	endpoint := invoiced.AddFilterAndSort("/customers", filter, sort)
	customers := make(invoiced.Customers, 0)
	nextEndpoint, err := c.Api.Get(endpoint, &customers)
	return customers, nextEndpoint, err
}

func (c *Client) ListCustomerByNumber(customerNumber string) (*invoiced.Customer, error) {
	filter := invoiced.NewFilter()
	err := filter.Set("number", customerNumber)
	if err != nil {
		return nil, err
	}

	customers, err := c.ListAll(filter, nil)
	if err != nil {
		return nil, err
	}

	if len(customers) == 0 {
		return nil, nil
	}

	return customers[0], nil
}

func (c *Client) GetBalance(id int64) (*invoiced.Balance, error) {
	endpoint := "/customers/" + strconv.FormatInt(id, 10) + "/balance"
	custBalance := new(invoiced.Balance)
	_, err := c.Api.Get(endpoint, custBalance)
	return custBalance, err
}

func (c *Client) SendStatementEmail(id int64, request *invoiced.SendStatementEmailRequest) error {
	endpoint := "/customers/" + strconv.FormatInt(id, 10) + "/emails"
	return c.Api.Create(endpoint, request, nil)
}

func (c *Client) SendStatementText(id int64, request *invoiced.SendStatementTextMessageRequest) (invoiced.TextMessages, error) {
	endpoint := "/customers/" + strconv.FormatInt(id, 10) + "/text_messages"
	custStmtResp := new(invoiced.TextMessages)
	err := c.Api.Create(endpoint, request, custStmtResp)
	return *custStmtResp, err
}

func (c *Client) SendStatementLetter(id int64, request *invoiced.SendStatementLetterRequest) (*invoiced.Letter, error) {
	endpoint := "/customers/" + strconv.FormatInt(id, 10) + "/letters"
	custStmtResp := new(invoiced.Letter)
	err := c.Api.Create(endpoint, request, custStmtResp)
	return custStmtResp, err
}

func (c *Client) CreateContact(id int64, request *invoiced.ContactRequest) (*invoiced.Contact, error) {
	endpoint := "/customers/" + strconv.FormatInt(id, 10) + "/contacts"
	contResp := new(invoiced.Contact)
	err := c.Api.Create(endpoint, request, contResp)
	return contResp, err
}

func (c *Client) RetrieveContact(customerId int64, id int64) (*invoiced.Contact, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/contacts/" + strconv.FormatInt(id, 10)
	retrievedContact := new(invoiced.Contact)
	_, err := c.Api.Get(endpoint, retrievedContact)
	return retrievedContact, err
}

func (c *Client) UpdateContact(customerId int64, id int64, request *invoiced.ContactRequest) (*invoiced.Contact, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/contacts/" + strconv.FormatInt(id, 10)

	contResp := new(invoiced.Contact)

	err := c.Api.Update(endpoint, request, contResp)
	if err != nil {
		return nil, err
	}

	return contResp, nil
}

func (c *Client) ListAllContacts(customerId int64) (invoiced.Contacts, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/contacts"

	contacts := make(invoiced.Contacts, 0)

NEXT:
	tmpContacts := make(invoiced.Contacts, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpContacts)

	if err != nil {
		return nil, err
	}

	contacts = append(contacts, tmpContacts...)

	if endpoint != "" {
		goto NEXT
	}

	return contacts, nil
}

func (c *Client) DeleteContact(customerId int64, id int64) error {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/contacts/" + strconv.FormatInt(id, 10)
	return c.Api.Delete(endpoint)
}

func (c *Client) RetrieveNotes(customerId int64) (invoiced.Notes, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/notes"

	notes := make(invoiced.Notes, 0)

NEXT:
	tmpNotes := make(invoiced.Notes, 0)

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

func (c *Client) CreatePaymentSource(customerId int64, request *invoiced.PaymentSourceRequest) (*invoiced.PaymentSource, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/payment_sources"
	resp := new(invoiced.PaymentSource)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListAllPaymentSources(customerId int64) (invoiced.PaymentSources, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/payment_sources"

	sources := make(invoiced.PaymentSources, 0)

NEXT:
	tmpSources := make(invoiced.PaymentSources, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpSources)

	if err != nil {
		return nil, err
	}

	sources = append(sources, tmpSources...)

	if endpoint != "" {
		goto NEXT
	}

	return sources, nil
}

func (c *Client) DeleteCard(customerId int64, id int64) error {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/cards/" + strconv.FormatInt(id, 10)
	return c.Api.Delete(endpoint)
}

func (c *Client) DeleteBankAccount(customerId int64, id int64) error {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/bank_accounts/" + strconv.FormatInt(id, 10)
	return c.Api.Delete(endpoint)
}

func (c *Client) CreatePendingLineItem(customerId int64, request *invoiced.PendingLineItemRequest) (*invoiced.PendingLineItem, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/line_items"
	resp := new(invoiced.PendingLineItem)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RetrievePendingLineItem(customerId int64, id int64) (*invoiced.PendingLineItem, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/line_items/" + strconv.FormatInt(id, 10)
	resp := new(invoiced.PendingLineItem)

	_, err := c.Api.Get(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) UpdatePendingLineItem(customerId int64, id int64, request *invoiced.PendingLineItemRequest) (*invoiced.PendingLineItem, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/line_items/" + strconv.FormatInt(id, 10)
	resp := new(invoiced.PendingLineItem)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListAllPendingLineItems(customerId int64) (invoiced.PendingLineItems, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/line_items"

	plis := make(invoiced.PendingLineItems, 0)

NEXT:
	tmpPlis := make(invoiced.PendingLineItems, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpPlis)

	if err != nil {
		return nil, err
	}

	plis = append(plis, tmpPlis...)

	if endpoint != "" {
		goto NEXT
	}

	return plis, nil
}

func (c *Client) TriggerInvoice(customerId int64) (*invoiced.Invoice, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/invoices"

	invoice := new(invoiced.Invoice)

	err := c.Api.Create(endpoint, nil, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (c *Client) ConsolidateInvoices(customerId int64) (*invoiced.Invoice, error) {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/consolidate_invoices"

	invoice := new(invoiced.Invoice)

	err := c.Api.Create(endpoint, nil, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (c *Client) DeletePendingLineItem(customerId int64, id int64) error {
	endpoint := "/customers/" + strconv.FormatInt(customerId, 10) + "/line_items/" + strconv.FormatInt(id, 10)
	return c.Api.Delete(endpoint)
}
