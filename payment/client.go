package payment

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Count() (int64, error) {
	endpoint := invoiced.PaymentEndpoint

	count, err := c.Api.Count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Client) Create(request *invoiced.PaymentRequest) (*invoiced.Payment, error) {
	endpoint := invoiced.PaymentEndpoint
	resp := c.NewPayment()

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id int64) (*invoiced.Payment, error) {
	endpoint := invoiced.PaymentEndpoint + "/" + strconv.FormatInt(id, 10)
	payment := &invoiced.Payment{c.Client, new(invoiced.Payment)}

	_, err := c.Api.Get(endpoint, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (c *Client) Update(request *invoiced.PaymentRequest) error {
	endpoint := invoiced.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := c.NewPayment()

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Client = resp.PaymentClient

	return nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Payments, error) {
	endpoint := invoiced.PaymentEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	payments := make(invoiced.Payments, 0)
	paymentsToReturn := make(invoiced.Payments, 0)

NEXT:
	tmpPayments := make(invoiced.Payments, 0)

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

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Payments, string, error) {
	endpoint := invoiced.PaymentEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	payments := make(invoiced.Payments, 0)
	paymentsToReturn := make(invoiced.Payments, 0)

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

func (c *Client) SendReceipt(request *invoiced.SendEmailRequest) error {
	endpoint := invoiced.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.Api.Create(endpoint, request, nil)
	if err != nil {
		return err
	}

	return nil
}
