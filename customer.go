package invdapi

import (
	"errors"
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
	endPoint := c.MakeEndPointURL(invdendpoint.CustomersEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Customer) Create(customer *Customer) (*Customer, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CustomersEndPoint)
	custResp := new(Customer)

	apiErr := c.create(endPoint, customer, custResp)

	if apiErr != nil {
		return nil, apiErr
	}

	custResp.Connection = c.Connection

	return custResp, nil

}

func (c *Customer) Delete() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Customer) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id)
	custResp := new(Customer)
	apiErr := c.update(endPoint, c, custResp)

	if apiErr != nil {
		return apiErr
	}

	c.Customer = custResp.Customer

	return nil

}

func (c *Customer) Retrieve(id int64) (*Customer, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), id)

	custEndPoint := new(invdendpoint.Customer)

	customer := &Customer{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, customer)

	if apiErr != nil {
		return nil, apiErr
	}

	return customer, nil

}

func (c *Customer) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Customers, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CustomersEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	customers := make(Customers, 0)

NEXT:
	tmpCustomers := make(Customers, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpCustomers)

	if apiErr != nil {
		return nil, apiErr
	}

	customers = append(customers, tmpCustomers...)

	if endPoint != "" {
		goto NEXT
	}

	for _, customer := range customers {
		customer.Connection = c.Connection

	}

	return customers, nil

}

func (c *Customer) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Customers, string, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CustomersEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	customers := make(Customers, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &customers)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, customer := range customers {
		customer.Connection = c.Connection

	}

	return customers, nextEndPoint, nil

}

func (c *Customer) ListCustomersByName(customerName string) (Customers, error) {

	filter := invdendpoint.NewFilter()

	err := filter.Set("name", customerName)

	if err != nil {
		return nil, err
	}

	customers, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	return customers, nil

}

func (c *Customer) ListCustomerByNumber(customerNumber string) (*Customer, error) {

	filter := invdendpoint.NewFilter()
	err := filter.Set("number", customerNumber)

	if err != nil {
		return nil, err
	}

	customers, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(customers) == 0 {
		return nil, nil
	}

	return customers[0], nil

}

func (c *Customer) GetBalance() (*invdendpoint.CustomerBalance, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/balance"

	custBalance := new(invdendpoint.CustomerBalance)

	_, apiErr := c.retrieveDataFromAPI(endPoint, custBalance)

	if apiErr != nil {
		return nil, apiErr
	}

	return custBalance, nil
}

func (c *Customer) SendStatement(custStmtReq *invdendpoint.EmailResponse) (*invdendpoint.EmailResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/emails"

	custStmtResp := new(invdendpoint.EmailResponses)
	err := c.create(endPoint, custStmtReq, custStmtResp)

	if err != nil {
		return nil, err
	}

	return custStmtResp, nil
}

func (c *Customer) CreateContact(contact *invdendpoint.Contact) (*invdendpoint.Contact, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/contacts"

	createdContact := new(invdendpoint.Contact)

	err := c.create(endPoint, contact, createdContact)

	if err != nil {
		return nil, err
	}

	return createdContact, nil

}

func (c *Customer) RetrieveContact(contactID int64) (*invdendpoint.Contact, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/contacts" + strconv.FormatInt(contactID, 10)

	retrievedContact := new(invdendpoint.Contact)

	_, err := c.retrieveDataFromAPI(endPoint, retrievedContact)

	if err != nil {
		return nil, err
	}

	return retrievedContact, nil

}

func (c *Customer) UpdateContact(contactToUpdate *invdendpoint.Contact) (*invdendpoint.Contact, error) {

	if contactToUpdate.Id <= 0 {
		return nil, errors.New("Need to supply a contact id in order to update a contact")
	}

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/contacts" + strconv.FormatInt(contactToUpdate.Id, 10)

	contResp := new(invdendpoint.Contact)
	err := c.update(endPoint, contactToUpdate, contResp)

	if err != nil {
		return nil, err
	}

	return contResp, nil

}

func (c *Customer) ListAllContacts() (invdendpoint.Contacts, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/contacts"

	contacts := make(invdendpoint.Contacts, 0)

NEXT:
	tmpContacts := make(invdendpoint.Contacts, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpContacts)

	if apiErr != nil {
		return nil, apiErr
	}

	contacts = append(contacts, tmpContacts...)

	if endPoint != "" {
		goto NEXT
	}

	return contacts, nil

}

func (c *Customer) DeleteContact(contactID int64) error {

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/contacts" + strconv.FormatInt(contactID, 10)

	err := c.delete(endPoint)

	if err != nil {
		return err
	}

	return nil

}

func (c *Customer) CreatePendingLineItem(pendingLineItem *invdendpoint.PendingLineItem) (*invdendpoint.PendingLineItem, error) {

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/line_items"

	pendingLineItemResp := new(invdendpoint.PendingLineItem)

	err := c.create(endPoint, pendingLineItem, pendingLineItemResp)

	if err != nil {
		return nil, err
	}

	return pendingLineItemResp, nil

}

func (c *Customer) RetrievePendingLineItem(id int64) (*invdendpoint.PendingLineItem, error) {

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/line_items" + strconv.FormatInt(id, 10)

	retrievedPendingLineItem := new(invdendpoint.PendingLineItem)

	_, err := c.retrieveDataFromAPI(endPoint, retrievedPendingLineItem)

	if err != nil {
		return nil, err
	}

	return retrievedPendingLineItem, nil

}

func (c *Customer) UpdatePendingLineItem(pendingLineItem *invdendpoint.PendingLineItem) (*invdendpoint.PendingLineItem, error) {

	if pendingLineItem.Id <= 0 {
		return nil, errors.New("Need to supply a pending line item id in order to update a pending line item")
	}

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/line_items" + strconv.FormatInt(pendingLineItem.Id, 10)

	pendingLineItemResp := new(invdendpoint.PendingLineItem)
	err := c.update(endPoint, pendingLineItem, pendingLineItemResp)

	if err != nil {
		return nil, err
	}

	return pendingLineItemResp, nil

}

func (c *Customer) TriggerInvoice() (*Invoice, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id) + "/invoices"

	invoice := new(Invoice)

	err := c.create(endPoint, nil, invoice)

	if err != nil {
		return nil, err
	}

	invoice.Connection = c.Connection

	return invoice, nil

}

func (c *Customer) DeletePendingLineItem(id int64) error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CustomersEndPoint), c.Id)

	err := c.delete(endPoint)

	if err != nil {
		return err
	}

	return nil

}
