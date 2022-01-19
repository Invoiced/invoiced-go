package invoiced

import (
	"strconv"
)

type PaymentClient struct {
	*Client
	*PaymentClient
}

func (c *Client) NewPayment() *PaymentClient {
	p := new(PaymentClient)
	return &PaymentClient{c, p}
}

func (c *PaymentClient) Count() (int64, error) {
	endpoint := PaymentEndpoint

	count, err := c.Api.Count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *PaymentClient) Create(request *PaymentRequest) (*Payment, error) {
	endpoint := PaymentEndpoint
	resp := c.NewPayment()

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *PaymentClient) Retrieve(id int64) (*Payment, error) {
	endpoint := PaymentEndpoint + "/" + strconv.FormatInt(id, 10)
	payment := &Payment{c.Client, new(Payment)}

	_, err := c.Api.Get(endpoint, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (c *PaymentClient) Update(request *PaymentRequest) error {
	endpoint := PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := c.NewPayment()

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.PaymentClient = resp.PaymentClient

	return nil
}

func (c *PaymentClient) Delete() error {
	endpoint := PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *PaymentClient) ListAll(filter *Filter, sort *Sort) (Payments, error) {
	endpoint := PaymentEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)

	payments := make(Payments, 0)
	paymentsToReturn := make(Payments, 0)

NEXT:
	tmpPayments := make(Payments, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpPayments)

	if err != nil {
		return nil, err
	}

	payments = append(payments, tmpPayments...)

	if endpoint != "" {
		goto NEXT
	}

	for _, payment := range payments {
		inv := c.Client.NewPayment()
		invData := payment
		inv.PaymentClient = &invData
		paymentsToReturn = append(paymentsToReturn, inv)
	}

	return paymentsToReturn, nil
}

func (c *PaymentClient) List(filter *Filter, sort *Sort) (Payments, string, error) {
	endpoint := PaymentEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)

	payments := make(Payments, 0)
	paymentsToReturn := make(Payments, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &payments)

	if err != nil {
		return nil, "", err
	}

	for _, payment := range payments {
		inv := c.Client.NewPayment()
		invData := payment
		inv.PaymentClient = &invData
		paymentsToReturn = append(paymentsToReturn, inv)

	}

	return paymentsToReturn, nextEndpoint, nil
}

func (c *PaymentClient) SendReceipt(request *SendEmailRequest) error {
	endpoint := PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.Api.Create(endpoint, request, nil)
	if err != nil {
		return err
	}

	return nil
}
