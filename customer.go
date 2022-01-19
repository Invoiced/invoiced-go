package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"strconv"
)

type Customer struct {
	*Connection
	*invdendpoint.Customer
}

type Customers []*Customer

func (c *Connection) NewCustomer() *Customer {
	customer := new(invdendpoint.Customer)
	return &Customer{c, customer}
}

func (c *Customer) Count() (int64, error) {
	count, err := c.count(invdendpoint.CustomerEndpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Customer) Create(request *invdendpoint.CustomerRequest) (*Customer, error) {
	resp := new(Customer)

	err := c.create(invdendpoint.CustomerEndpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Customer) Update(request *invdendpoint.CustomerRequest) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Customer)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Customer = resp.Customer

	return nil
}

func (c *Customer) Retrieve(id int64) (*Customer, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(id, 10)

	customer := &Customer{c.Connection, new(invdendpoint.Customer)}

	_, err := c.retrieveDataFromAPI(endpoint, customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) Delete() error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Customers, error) {
	endpoint := invdendpoint.CustomerEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	customers := make(Customers, 0)

NEXT:
	tmpCustomers := make(Customers, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpCustomers)

	if err != nil {
		return nil, err
	}

	customers = append(customers, tmpCustomers...)

	if endpoint != "" {
		goto NEXT
	}

	for _, customer := range customers {
		customer.Connection = c.Connection
	}

	return customers, nil
}

func (c *Customer) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Customers, string, error) {
	endpoint := invdendpoint.CustomerEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	customers := make(Customers, 0)

	nextEndpoint, err := c.retrieveDataFromAPI(endpoint, &customers)

	if err != nil {
		return nil, "", err
	}

	for _, customer := range customers {
		customer.Connection = c.Connection
	}

	return customers, nextEndpoint, nil
}

func (c *Customer) ListCustomerByNumber(customerNumber string) (*Customer, error) {
	filter := invdendpoint.NewFilter()
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

func (c *Customer) GetBalance() (*invdendpoint.Balance, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/balance"

	custBalance := new(invdendpoint.Balance)

	_, err := c.retrieveDataFromAPI(endpoint, custBalance)

	if err != nil {
		return nil, err
	}

	return custBalance, nil
}

func (c *Customer) SendStatementEmail(request *invdendpoint.SendStatementEmailRequest) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.create(endpoint, request, nil)
	if err != nil {
		return nil
	}

	return nil
}

func (c *Customer) SendStatementText(request *invdendpoint.SendStatementTextMessageRequest) (invdendpoint.TextMessages, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	custStmtResp := new(invdendpoint.TextMessages)
	err := c.create(endpoint, request, custStmtResp)
	if err != nil {
		return nil, err
	}

	return *custStmtResp, nil
}

func (c *Customer) SendStatementLetter(request *invdendpoint.SendStatementLetterRequest) (*invdendpoint.Letter, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	custStmtResp := new(invdendpoint.Letter)
	err := c.create(endpoint, request, custStmtResp)
	if err != nil {
		return nil, err
	}

	return custStmtResp, nil
}

func (c *Customer) CreateContact(request *invdendpoint.ContactRequest) (*invdendpoint.Contact, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts"

	contResp := new(invdendpoint.Contact)

	err := c.create(endpoint, request, contResp)
	if err != nil {
		return nil, err
	}

	return contResp, nil
}

func (c *Customer) RetrieveContact(id int64) (*invdendpoint.Contact, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts/" + strconv.FormatInt(id, 10)

	retrievedContact := new(invdendpoint.Contact)

	_, err := c.retrieveDataFromAPI(endpoint, retrievedContact)
	if err != nil {
		return nil, err
	}

	return retrievedContact, nil
}

func (c *Customer) UpdateContact(id int64, request *invdendpoint.ContactRequest) (*invdendpoint.Contact, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts/" + strconv.FormatInt(id, 10)

	contResp := new(invdendpoint.Contact)

	err := c.update(endpoint, request, contResp)
	if err != nil {
		return nil, err
	}

	return contResp, nil
}

func (c *Customer) ListAllContacts() (invdendpoint.Contacts, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts"

	contacts := make(invdendpoint.Contacts, 0)

NEXT:
	tmpContacts := make(invdendpoint.Contacts, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpContacts)

	if err != nil {
		return nil, err
	}

	contacts = append(contacts, tmpContacts...)

	if endpoint != "" {
		goto NEXT
	}

	return contacts, nil
}

func (c *Customer) DeleteContact(id int64) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts/" + strconv.FormatInt(id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) RetrieveNotes() (invdendpoint.Notes, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/notes"

	notes := make(invdendpoint.Notes, 0)

NEXT:
	tmpNotes := make(invdendpoint.Notes, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpNotes)

	if err != nil {
		return nil, err
	}

	notes = append(notes, tmpNotes...)

	if endpoint != "" {
		goto NEXT
	}

	return notes, nil
}

func (c *Customer) CreatePaymentSource(request *invdendpoint.PaymentSourceRequest) (*invdendpoint.PaymentSource, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_sources"
	resp := new(invdendpoint.PaymentSource)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Customer) ListAllPaymentSources() (invdendpoint.PaymentSources, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_sources"

	sources := make(invdendpoint.PaymentSources, 0)

NEXT:
	tmpSources := make(invdendpoint.PaymentSources, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpSources)

	if err != nil {
		return nil, err
	}

	sources = append(sources, tmpSources...)

	if endpoint != "" {
		goto NEXT
	}

	return sources, nil
}

func (c *Customer) DeleteCard(id int64) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/cards/" + strconv.FormatInt(id, 10)

	err := c.delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) DeleteBankAccount(id int64) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/bank_accounts/" + strconv.FormatInt(id, 10)

	err := c.delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) CreatePendingLineItem(request *invdendpoint.PendingLineItemRequest) (*invdendpoint.PendingLineItem, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items"
	resp := new(invdendpoint.PendingLineItem)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Customer) RetrievePendingLineItem(id int64) (*invdendpoint.PendingLineItem, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items/" + strconv.FormatInt(id, 10)
	resp := new(invdendpoint.PendingLineItem)

	_, err := c.retrieveDataFromAPI(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Customer) UpdatePendingLineItem(id int64, request *invdendpoint.PendingLineItemRequest) (*invdendpoint.PendingLineItem, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items/" + strconv.FormatInt(id, 10)
	resp := new(invdendpoint.PendingLineItem)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Customer) ListAllPendingLineItems() (invdendpoint.PendingLineItems, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items"

	plis := make(invdendpoint.PendingLineItems, 0)

NEXT:
	tmpPlis := make(invdendpoint.PendingLineItems, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpPlis)

	if err != nil {
		return nil, err
	}

	plis = append(plis, tmpPlis...)

	if endpoint != "" {
		goto NEXT
	}

	return plis, nil
}

func (c *Customer) TriggerInvoice() (*Invoice, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/invoices"

	invoice := c.NewInvoice()

	err := c.create(endpoint, nil, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (c *Customer) ConsolidateInvoices() (*Invoice, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/consolidate_invoices"

	invoice := c.NewInvoice()

	err := c.create(endpoint, nil, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (c *Customer) DeletePendingLineItem(id int64) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items/" + strconv.FormatInt(id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) CreateCreditBalanceAdjustment(request *invdendpoint.BalanceAdjustmentRequest) (*invdendpoint.BalanceAdjustment, error) {
	endpoint := invdendpoint.CreditBalanceAdjustmentsEndpoint
	resp := new(invdendpoint.BalanceAdjustment)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
