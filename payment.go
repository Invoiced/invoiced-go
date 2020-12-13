package invdapi

import (
	"errors"
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Payment struct {
	*Connection
	*invdendpoint.Payment
}

type Payments []*Payment

func (c *Connection) NewPayment() *Payment {
	return &Payment{c, new(invdendpoint.Payment)}
}

func (c *Payment) Count() (int64, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.PaymentEndpoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil
}

func (c *Payment) Create(payment *Payment) (*Payment, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.PaymentEndpoint)
	txnResp := new(Payment)

	if payment == nil {
		return nil, errors.New("payment cannot be nil")
	}

	// safe prune invoice data for creation
	invdTransDataToCreate, err := SafePaymentForCreation(payment.Payment)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdTransDataToCreate, txnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	txnResp.Connection = c.Connection

	return txnResp, nil
}

func (c *Payment) Delete() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.PaymentEndpoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Payment) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.PaymentEndpoint), c.Id)
	txnResp := new(Payment)

	// safe prune invoice data for updating
	invdTransDataToUpdate, err := SafePaymentForUpdate(c.Payment)
	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, invdTransDataToUpdate, txnResp)

	if apiErr != nil {
		return apiErr
	}

	c.Payment = txnResp.Payment

	return nil
}

func (c *Payment) Retrieve(id int64) (*Payment, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.PaymentEndpoint), id)

	custEndPoint := new(invdendpoint.Payment)

	payment := &Payment{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, payment)

	if apiErr != nil {
		return nil, apiErr
	}

	return payment, nil
}

func (c *Payment) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Payments, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.PaymentEndpoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	payments := make(Payments, 0)

NEXT:
	tmpPayments := make(Payments, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpPayments)

	if apiErr != nil {
		return nil, apiErr
	}

	payments = append(payments, tmpPayments...)

	if endPoint != "" {
		goto NEXT
	}

	for _, payment := range payments {
		payment.Connection = c.Connection
	}

	return payments, nil
}

func (c *Payment) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Payments, string, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.PaymentEndpoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	payments := make(Payments, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &payments)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, payment := range payments {
		payment.Connection = c.Connection
	}

	return payments, nextEndPoint, nil
}

func (c *Payment) ListSuccessfulByInvoiceID(invoiceID int64) (Payments, error) {
	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	filter := invdendpoint.NewFilter()
	err := filter.Set("invoice", invoiceIDString)
	if err != nil {
		return nil, err
	}

	err = filter.Set("status", "succeeded")

	if err != nil {
		return nil, err
	}

	payments, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(payments) == 0 {
		return nil, nil
	}

	return payments, nil
}

func (c *Payment) ListSuccessfulChargesByInvoiceID(invoiceID int64) (Payments, error) {
	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	filter := invdendpoint.NewFilter()
	err := filter.Set("invoice", invoiceIDString)
	if err != nil {
		return nil, err
	}
	err = filter.Set("status", "succeeded")
	if err != nil {
		return nil, err
	}
	err = filter.Set("type", "charge")
	if err != nil {
		return nil, err
	}

	payments, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(payments) == 0 {
		return nil, nil
	}

	return payments, nil
}

func (c *Payment) ListSuccessfulRefundsByInvoiceID(invoiceID int64) (Payments, error) {
	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	filter := invdendpoint.NewFilter()
	err := filter.Set("invoice", invoiceIDString)
	if err != nil {
		return nil, err
	}

	err = filter.Set("status", "succeeded")

	if err != nil {
		return nil, err
	}

	err = filter.Set("type", "refund")

	if err != nil {
		return nil, err
	}

	payments, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(payments) == 0 {
		return nil, nil
	}

	return payments, nil
}

func (c *Payment) ListSuccessfulPaymentsByInvoiceID(invoiceID int64) (Payments, error) {
	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	filter := invdendpoint.NewFilter()
	err := filter.Set("invoice", invoiceIDString)
	if err != nil {
		return nil, err
	}
	err = filter.Set("status", "succeeded")

	if err != nil {
		return nil, err
	}

	err = filter.Set("type", "payment")

	if err != nil {
		return nil, err
	}

	payments, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(payments) == 0 {
		return nil, nil
	}

	return payments, nil
}

func (c *Payment) ListSuccessfulChargesAndPaymentsByInvoiceID(invoiceID int64) (Payments, error) {
	charges, err := c.ListSuccessfulChargesByInvoiceID(invoiceID)
	if err != nil {
		return nil, err
	}

	payments, err := c.ListSuccessfulPaymentsByInvoiceID(invoiceID)
	if err != nil {
		return nil, err
	}

	chargesPayments := append(charges, payments...)

	return chargesPayments, nil
}

func (c *Payment) SendReceipt(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endPoint, emailReq, emailResp)
	if err != nil {
		return nil, err
	}

	return *emailResp, nil
}

// SafePaymentForCreation prunes payment data for just fields that can be used for creation of a payment
func SafePaymentForCreation(payment *invdendpoint.Payment) (*invdendpoint.Payment, error) {
	if payment == nil {
		return nil, errors.New("Payment is nil")
	}

	transData := new(invdendpoint.Payment)
	transData.Customer = payment.Customer
	transData.Invoice = payment.Invoice
	transData.Date = payment.Date
	transData.CreditNote = payment.CreditNote
	transData.Type = payment.Type
	transData.Method = payment.Method
	transData.Status = payment.Status
	transData.Gateway = payment.Gateway
	transData.GatewayId = payment.GatewayId
	transData.Currency = payment.Currency
	transData.Amount = payment.Amount
	transData.Notes = payment.Notes
	transData.Metadata = payment.Metadata

	return transData, nil
}

// SafePaymentForUpdate prunes payment data for just fields that can be used for creation of a transactiobn
func SafePaymentForUpdate(payment *invdendpoint.Payment) (*invdendpoint.Payment, error) {
	if payment == nil {
		return nil, errors.New("Payment is nil")
	}

	transData := new(invdendpoint.Payment)

	transData.Date = payment.Date
	transData.Method = payment.Method
	transData.Status = payment.Status
	transData.Gateway = payment.Gateway
	transData.GatewayId = payment.GatewayId
	transData.Currency = payment.Currency
	transData.Amount = payment.Amount
	transData.Notes = payment.Notes
	transData.Metadata = payment.Metadata

	return transData, nil
}
