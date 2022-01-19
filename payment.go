package invdapi

import (
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
	return &Payment{c, p}
}

func (c *Payment) Count() (int64, error) {
	endpoint := invdendpoint.PaymentEndpoint

	count, err := c.count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Payment) Create(request *invdendpoint.PaymentRequest) (*Payment, error) {
	endpoint := invdendpoint.PaymentEndpoint
	resp := c.NewPayment()

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Payment) Retrieve(id int64) (*Payment, error) {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(id, 10)
	payment := &Payment{c.Connection, new(invdendpoint.Payment)}

	_, err := c.retrieveDataFromAPI(endpoint, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (c *Payment) Update(request *invdendpoint.PaymentRequest) error {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := c.NewPayment()

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Payment = resp.Payment

	return nil
}

func (c *Payment) Delete() error {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Payment) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Payments, error) {
	endpoint := invdendpoint.PaymentEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	payments := make(invdendpoint.Payments, 0)
	paymentsToReturn := make(Payments, 0)

NEXT:
	tmpPayments := make(invdendpoint.Payments, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpPayments)

	if err != nil {
		return nil, err
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

	nextEndpoint, err := c.retrieveDataFromAPI(endpoint, &payments)

	if err != nil {
		return nil, "", err
	}

	for _, payment := range payments {
		inv := c.Connection.NewPayment()
		invData := payment
		inv.Payment = &invData
		paymentsToReturn = append(paymentsToReturn, inv)

	}

	return paymentsToReturn, nextEndpoint, nil
}

func (c *Payment) SendReceipt(request *invdendpoint.SendEmailRequest) error {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.create(endpoint, request, nil)
	if err != nil {
		return err
	}

	return nil
}
