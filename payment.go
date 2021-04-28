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
	 p := new(invdendpoint.Payment)
	return &Payment{c,  p}
}

func (c *Payment) Count() (int64, error) {
	endpoint := invdendpoint.PaymentEndpoint

	count, apiErr := c.count(endpoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil
}

func (c *Payment) Create(payment *Payment) (*Payment, error) {
	endpoint := invdendpoint.PaymentEndpoint
	txnResp := c.NewPayment()

	if payment == nil {
		return nil, errors.New("payment cannot be nil")
	}

	// safe prune invoice data for creation
	invdTransDataToCreate, err := SafePaymentForCreation(payment.Payment)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endpoint, invdTransDataToCreate, txnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	txnResp.Connection = c.Connection

	return txnResp, nil
}

func (c *Payment) Delete() error {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Payment) Save() error {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	txnResp := c.NewPayment()

	// safe prune invoice data for updating
	invdTransDataToUpdate, err := SafePaymentForUpdate(c.Payment)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, invdTransDataToUpdate, txnResp)

	if apiErr != nil {
		return apiErr
	}

	c.Payment = txnResp.Payment

	return nil
}

func (c *Payment) Retrieve(id int64) (*Payment, error) {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(id, 10)

	custEndpoint := new(invdendpoint.Payment)

	payment := &Payment{c.Connection, custEndpoint}

	_, apiErr := c.retrieveDataFromAPI(endpoint, payment)

	if apiErr != nil {
		return nil, apiErr
	}

	return payment, nil
}

func (c *Payment) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Payments, error) {
	endpoint := invdendpoint.PaymentEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	payments := make(invdendpoint.Payments, 0)
	paymentsToReturn := make(Payments, 0)

NEXT:
	tmpPayments := make(invdendpoint.Payments, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpPayments)

	if apiErr != nil {
		return nil, apiErr
	}

	payments = append(payments, tmpPayments...)

	if endpoint != "" {
		goto NEXT
	}

	for _, payment := range payments {
		inv := c.Connection.NewPayment()
		invData := payment
		inv.Payment = &invData
		paymentsToReturn = append(paymentsToReturn, inv)
	}

	return paymentsToReturn, nil
}

func (c *Payment) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Payments, string, error) {
	endpoint := invdendpoint.PaymentEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	payments := make(invdendpoint.Payments, 0)
	paymentsToReturn := make(Payments, 0)

	nextEndpoint, apiErr := c.retrieveDataFromAPI(endpoint, &payments)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, payment := range payments {
		inv := c.Connection.NewPayment()
		invData := payment
		inv.Payment = &invData
		paymentsToReturn = append(paymentsToReturn, inv)

	}

	return paymentsToReturn, nextEndpoint, nil
}

func (c *Payment) SendReceipt(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endpoint, emailReq, emailResp)
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

	paymentData := new(invdendpoint.Payment)
	paymentData.Customer = payment.Customer
	paymentData.Date = payment.Date
	paymentData.Method = payment.Method
	paymentData.Currency = payment.Currency
	paymentData.Amount = payment.Amount
	paymentData.Notes = payment.Notes
	paymentData.Source = payment.Source
	paymentData.Reference = payment.Reference
	paymentData.AppliedTo = payment.AppliedTo

	return paymentData, nil
}

// SafePaymentForUpdate prunes payment data for just fields that can be used for creation of a payment
func SafePaymentForUpdate(payment *invdendpoint.Payment) (*invdendpoint.Payment, error) {
	if payment == nil {
		return nil, errors.New("Payment is nil")
	}

	paymentData := new(invdendpoint.Payment)

	paymentData.Customer = payment.Customer
	paymentData.Date = payment.Date
	paymentData.Method = payment.Method
	paymentData.Currency = payment.Currency
	paymentData.Amount = payment.Amount
	paymentData.Notes = payment.Notes
	paymentData.Source = payment.Source
	paymentData.Reference = payment.Reference
	paymentData.AppliedTo = payment.AppliedTo

	return paymentData, nil
}
