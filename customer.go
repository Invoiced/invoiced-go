package invdapi

import (
	"errors"
	"strconv"
	"github.com/Invoiced/invoiced-go/invdendpoint"
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

func (c *Connection) NewContact() *invdendpoint.Contact {
	return &invdendpoint.Contact{}
}

func (c *Connection) NewPaymentSource() *invdendpoint.PaymentSource {
	return &invdendpoint.PaymentSource{}
}

func (c *Connection) NewPendingLineItem() *invdendpoint.PendingLineItem {
	return &invdendpoint.PendingLineItem{}
}

func (c *Customer) Count() (int64, error) {
	count, apiErr := c.count(invdendpoint.CustomerEndpoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil
}

func (c *Customer) Create(customer *Customer) (*Customer, error) {
	custResp := new(Customer)

	if customer == nil {
		return nil, errors.New("Customer is nil")
	}

	custDataToCreate, err := SafeCustomerForCreation(customer.Customer)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(invdendpoint.CustomerEndpoint, custDataToCreate, custResp)

	if apiErr != nil {
		return nil, apiErr
	}

	custResp.Connection = c.Connection

	return custResp, nil
}

func (c *Customer) Delete() error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Customer) Save() error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	custDataToUpdate, err := SafeCustomerForUpdate(c.Customer)
	if err != nil {
		return nil
	}

	custResp := new(Customer)

	apiErr := c.update(endpoint, custDataToUpdate, custResp)

	if apiErr != nil {
		return apiErr
	}

	c.Customer = custResp.Customer

	return nil
}

func (c *Customer) FreeUpdate(customerData interface{}) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	custResp := new(Customer)

	apiErr := c.update(endpoint, customerData, custResp)

	if apiErr != nil {
		return apiErr
	}

	c.Customer = custResp.Customer

	return nil
}

func (c *Customer) Retrieve(id int64) (*Customer, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(id, 10)

	custEndpoint := new(invdendpoint.Customer)

	customer := &Customer{c.Connection, custEndpoint}

	_, apiErr := c.retrieveDataFromAPI(endpoint, customer)

	if apiErr != nil {
		return nil, apiErr
	}

	return customer, nil
}

func (c *Customer) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Customers, error) {
	endpoint := invdendpoint.CustomerEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)


	customers := make(Customers, 0)

NEXT:
	tmpCustomers := make(Customers, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpCustomers)

	if apiErr != nil {
		return nil, apiErr
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

	nextEndpoint, apiErr := c.retrieveDataFromAPI(endpoint, &customers)

	if apiErr != nil {
		return nil, "", apiErr
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

	customers, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(customers) == 0 {
		return nil, nil
	}

	return customers[0], nil
}

func (c *Customer) GetBalance() (*invdendpoint.Balance, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/balance"

	custBalance := new(invdendpoint.Balance)

	_, apiErr := c.retrieveDataFromAPI(endpoint, custBalance)

	if apiErr != nil {
		return nil, apiErr
	}

	return custBalance, nil
}

func (c *Customer) SendStatementEmail(custStmtReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	custStmtResp := new(invdendpoint.EmailResponses)
	err := c.create(endpoint, custStmtReq, custStmtResp)
	if err != nil {
		return nil, err
	}

	return *custStmtResp, nil
}

func (c *Customer) SendStatementText(custStmtReq *invdendpoint.TextRequest) (invdendpoint.TextResponses, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	custStmtResp := new(invdendpoint.TextResponses)
	err := c.create(endpoint, custStmtReq, custStmtResp)
	if err != nil {
		return nil, err
	}

	return *custStmtResp, nil
}

func (c *Customer) SendStatementLetter(custStmtReq *invdendpoint.LetterRequest) (*invdendpoint.LetterResponse, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	custStmtResp := new(invdendpoint.LetterResponse)
	err := c.create(endpoint, custStmtReq, custStmtResp)
	if err != nil {
		return nil, err
	}

	return custStmtResp, nil
}

func (c *Customer) CreateContact(contact *invdendpoint.Contact) (*invdendpoint.Contact, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts"

	contactDataToCreate, err := SafeContactForCreation(contact)
	if err != nil {
		return nil, err
	}

	contResp := new(invdendpoint.Contact)

	err = c.create(endpoint, contactDataToCreate, contResp)

	if err != nil {
		return nil, err
	}

	return contResp, nil
}

func (c *Customer) RetrieveContact(contactID int64) (*invdendpoint.Contact, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts/" + strconv.FormatInt(contactID, 10)

	retrievedContact := new(invdendpoint.Contact)

	_, err := c.retrieveDataFromAPI(endpoint, retrievedContact)
	if err != nil {
		return nil, err
	}

	return retrievedContact, nil
}

func (c *Customer) UpdateContact(contactToUpdate *invdendpoint.Contact) (*invdendpoint.Contact, error) {
	if contactToUpdate.Id <= 0 {
		return nil, errors.New("Need to supply a contact id in order to update a contact")
	}

	contactDataToUpdate, err := SafeContactForUpdate(contactToUpdate)
	if err != nil {
		return nil, err
	}

	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts/" + strconv.FormatInt(contactToUpdate.Id, 10)

	contResp := new(invdendpoint.Contact)

	err = c.update(endpoint, contactDataToUpdate, contResp)

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

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpContacts)

	if apiErr != nil {
		return nil, apiErr
	}

	contacts = append(contacts, tmpContacts...)

	if endpoint != "" {
		goto NEXT
	}

	return contacts, nil
}

func (c *Customer) DeleteContact(contactID int64) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/contacts/" + strconv.FormatInt(contactID, 10)

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

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpNotes)

	if apiErr != nil {
		return nil, apiErr
	}

	notes = append(notes, tmpNotes...)

	if endpoint != "" {
		goto NEXT
	}

	return notes, nil
}

func (c *Customer) CreatePaymentSource(source *invdendpoint.PaymentSource) (*invdendpoint.PaymentSource, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_sources"

	sourceDataToCreate, err := SafeSourceForCreation(source)
	if err != nil {
		return nil, err
	}

	resp := new(invdendpoint.PaymentSource)

	err = c.create(endpoint, sourceDataToCreate, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SafeSourceForCreation prunes source object for just fields that can be used for creation
func SafeSourceForCreation(source *invdendpoint.PaymentSource) (*invdendpoint.PaymentSource, error) {
	if source == nil {
		return nil, errors.New("Source is nil")
	}

	sourceData := new(invdendpoint.PaymentSource)
	sourceData.Method = source.Method
	sourceData.MakeDefault = source.MakeDefault
	sourceData.InvoicedToken = source.InvoicedToken
	sourceData.GatewayToken = source.GatewayToken

	return sourceData, nil
}

func (c *Customer) ListAllPaymentSources() (invdendpoint.PaymentSources, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/payment_sources"

	sources := make(invdendpoint.PaymentSources, 0)

NEXT:
	tmpSources := make(invdendpoint.PaymentSources, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpSources)

	if apiErr != nil {
		return nil, apiErr
	}

	sources = append(sources, tmpSources...)

	if endpoint != "" {
		goto NEXT
	}

	return sources, nil
}

func (c *Customer) DeleteCard(cardID int64) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/cards/" + strconv.FormatInt(cardID, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Customer) DeleteBankAccount(acctID int64) error {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/bank_accounts/" + strconv.FormatInt(acctID, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Customer) CreatePendingLineItem(pendingLineItem *invdendpoint.PendingLineItem) (*invdendpoint.PendingLineItem, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items"

	pliDataToUpdate, err := SafePendingLineItemForCreation(pendingLineItem)
	if err != nil {
		return nil, err
	}

	pendingLineItemResp := new(invdendpoint.PendingLineItem)

	err = c.create(endpoint, pliDataToUpdate, pendingLineItemResp)

	if err != nil {
		return nil, err
	}

	return pendingLineItemResp, nil
}

func (c *Customer) RetrievePendingLineItem(id int64) (*invdendpoint.PendingLineItem, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items/" + strconv.FormatInt(id, 10)

	retrievedPendingLineItem := new(invdendpoint.PendingLineItem)

	_, err := c.retrieveDataFromAPI(endpoint, retrievedPendingLineItem)
	if err != nil {
		return nil, err
	}

	return retrievedPendingLineItem, nil
}

func (c *Customer) UpdatePendingLineItem(pendingLineItem *invdendpoint.PendingLineItem) (*invdendpoint.PendingLineItem, error) {
	if pendingLineItem.Id <= 0 {
		return nil, errors.New("Need to supply a pending line item id in order to update a pending line item")
	}

	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items/" + strconv.FormatInt(pendingLineItem.Id, 10)

	pliDataToUpdate, err := SafePendingLineItemForUpdate(pendingLineItem)
	if err != nil {
		return nil, err
	}

	pendingLineItemResp := new(invdendpoint.PendingLineItem)

	err = c.update(endpoint, pliDataToUpdate, pendingLineItemResp)

	if err != nil {
		return nil, err
	}

	return pendingLineItemResp, nil
}

func (c *Customer) ListAllPendingLineItems() (invdendpoint.PendingLineItems, error) {
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/line_items"

	plis := make(invdendpoint.PendingLineItems, 0)

NEXT:
	tmpPlis := make(invdendpoint.PendingLineItems, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpPlis)

	if apiErr != nil {
		return nil, apiErr
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
	endpoint := invdendpoint.CustomerEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) CreateCreditBalanceAdjustment(amount float64) (*invdendpoint.BalanceAdjustment, error) {
	endpoint := invdendpoint.CreditBalanceAdjustmentsEndpoint

	adjustmentRequest := invdendpoint.BalanceAdjustment{
		Customer: int(c.Id),
		Amount:   amount,
	}

	var response invdendpoint.BalanceAdjustment

	err := c.create(endpoint, adjustmentRequest, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafeCustomerForCreation(cust *invdendpoint.Customer) (*invdendpoint.Customer, error) {
	if cust == nil {
		return nil, errors.New("Customer is nil")
	}

	custData := new(invdendpoint.Customer)
	custData.Name = cust.Name
	custData.Number = cust.Number
	custData.Email = cust.Email
	custData.AutoPay = cust.AutoPay
	custData.AutoPayDelays = cust.AutoPayDelays
	custData.PaymentTerms = cust.PaymentTerms
	custData.AttentionTo = cust.AttentionTo
	custData.Address1 = cust.Address1
	custData.Address2 = cust.Address2
	custData.City = cust.City
	custData.State = cust.State
	custData.PostalCode = cust.PostalCode
	custData.Language = cust.Language
	custData.Currency = cust.Currency
	custData.Chase = cust.Chase
	custData.Phone = cust.Phone
	custData.CreditHold = cust.CreditHold
	custData.CreditLimit = cust.CreditLimit
	custData.Owner = cust.Owner
	custData.Taxable = cust.Taxable
	custData.Taxes = cust.Taxes
	custData.TaxId = cust.TaxId
	custData.AvalaraEntityUseCode = cust.AvalaraEntityUseCode
	custData.AvalaraExemptionNumber = cust.AvalaraExemptionNumber
	custData.Type = cust.Type
	custData.ParentCustomer = cust.ParentCustomer
	custData.Notes = cust.Notes
	custData.SignUpPage = cust.SignUpPage
	custData.Metadata = cust.Metadata
	custData.DisabledPaymentMethods = cust.DisabledPaymentMethods

	return custData, nil
}

// SafeInvoiceForCreation prunes invoice data for just fields that can be used for creation of a invoice
func SafeCustomerForUpdate(cust *invdendpoint.Customer) (*invdendpoint.Customer, error) {
	if cust == nil {
		return nil, errors.New("Customer is nil")
	}

	custData := new(invdendpoint.Customer)
	custData.Name = cust.Name
	custData.Number = cust.Number
	custData.Email = cust.Email
	custData.AutoPay = cust.AutoPay
	custData.PaymentTerms = cust.PaymentTerms
	custData.AttentionTo = cust.AttentionTo
	custData.Address1 = cust.Address1
	custData.Address2 = cust.Address2
	custData.City = cust.City
	custData.State = cust.State
	custData.PostalCode = cust.PostalCode
	custData.Country = cust.Country
	custData.Language = cust.Language
	custData.Currency = cust.Currency
	custData.Chase = cust.Chase
	custData.ChasingCadence = cust.ChasingCadence
	custData.Phone = cust.Phone
	custData.CreditHold = cust.CreditHold
	custData.CreditLimit = cust.CreditLimit
	custData.Owner = cust.Owner
	custData.Taxable = cust.Taxable
	custData.Taxes = cust.Taxes
	custData.TaxId = cust.TaxId
	custData.AvalaraEntityUseCode = cust.AvalaraEntityUseCode
	custData.AvalaraExemptionNumber = cust.AvalaraExemptionNumber
	custData.Type = cust.Type
	custData.ParentCustomer = cust.ParentCustomer
	custData.Notes = cust.Notes
	custData.SignUpPage = cust.SignUpPage
	custData.Metadata = cust.Metadata
	custData.DisabledPaymentMethods = cust.DisabledPaymentMethods

	return custData, nil
}

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafeContactForCreation(contact *invdendpoint.Contact) (*invdendpoint.Contact, error) {
	if contact == nil {
		return nil, errors.New("Contact is nil")
	}

	contData := new(invdendpoint.Contact)
	contData.Name = contact.Name
	contData.Title = contact.Title
	contData.Email = contact.Email
	contData.Phone = contact.Phone
	contData.Primary = contact.Primary
	contData.SmsEnabled = contact.SmsEnabled
	contData.Department = contact.Department
	contData.Address1 = contact.Address1
	contData.Address2 = contact.Address2
	contData.City = contact.City
	contData.State = contact.State
	contData.PostalCode = contact.PostalCode
	contData.Country = contact.Country

	return contData, nil
}

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafeContactForUpdate(contact *invdendpoint.Contact) (*invdendpoint.Contact, error) {
	if contact == nil {
		return nil, errors.New("Contact is nil")
	}

	contData := new(invdendpoint.Contact)
	contData.Name = contact.Name
	contData.Title = contact.Title
	contData.Email = contact.Email
	contData.Phone = contact.Phone
	contData.Primary = contact.Primary
	contData.SmsEnabled = contact.SmsEnabled
	contData.Department = contact.Department
	contData.Address1 = contact.Address1
	contData.Address2 = contact.Address2
	contData.City = contact.City
	contData.State = contact.State
	contData.PostalCode = contact.PostalCode
	contData.Country = contact.Country

	return contData, nil
}

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafePendingLineItemForCreation(pendingLineItem *invdendpoint.PendingLineItem) (*invdendpoint.PendingLineItem, error) {
	if pendingLineItem == nil {
		return nil, errors.New("PendingLineItem is nil")
	}

	pliData := new(invdendpoint.PendingLineItem)
	pliData.Item = pendingLineItem.Item
	pliData.Type = pendingLineItem.Type
	pliData.Name = pendingLineItem.Name
	pliData.Description = pendingLineItem.Description
	pliData.Quantity = pendingLineItem.Quantity
	pliData.UnitCost = pendingLineItem.UnitCost
	pliData.Discountable = pendingLineItem.Discountable
	pliData.Discounts = pendingLineItem.Discounts
	pliData.Taxes = pendingLineItem.Taxes
	pliData.Metadata = pendingLineItem.Metadata

	return pliData, nil
}

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafePendingLineItemForUpdate(pendingLineItem *invdendpoint.PendingLineItem) (*invdendpoint.PendingLineItem, error) {
	if pendingLineItem == nil {
		return nil, errors.New("PendingLineItem is nil")
	}

	pliData := new(invdendpoint.PendingLineItem)
	pliData.Item = pendingLineItem.Item
	pliData.Type = pendingLineItem.Type
	pliData.Name = pendingLineItem.Name
	pliData.Description = pendingLineItem.Description
	pliData.Quantity = pendingLineItem.Quantity
	pliData.UnitCost = pendingLineItem.UnitCost
	pliData.Discountable = pendingLineItem.Discountable
	pliData.Discounts = pendingLineItem.Discounts
	pliData.Taxes = pendingLineItem.Taxes
	pliData.Metadata = pendingLineItem.Metadata

	return pliData, nil
}
