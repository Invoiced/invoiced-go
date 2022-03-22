package payment

import (
	"github.com/Invoiced/invoiced-go/v2"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.PaymentRequest) (*invoiced.Payment, error) {
	resp := new(invoiced.Payment)
	err := c.Api.Create("/payments", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.Payment, error) {
	resp := new(invoiced.Payment)
	_, err := c.Api.Get("/payments/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.PaymentRequest) (*invoiced.Payment, error) {
	resp := new(invoiced.Payment)
	err := c.Api.Update("/payments/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/payments/" + strconv.FormatInt(id, 10))
}

func (c *Client) Count() (int64, error) {
	return c.Api.Count("/payments")
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Payments, error) {
	endpoint := invoiced.AddFilterAndSort("/payments", filter, sort)

	payments := make(invoiced.Payments, 0)

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

	return payments, nil
}

func (c *Client) ListAllStartEndDate(filter *invoiced.Filter, sort *invoiced.Sort,startDate,endDate int64) (invoiced.Payments, error) {
	endpoint := "/payments"

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	if startDate > 0 {
		startDateString := strconv.FormatInt(startDate, 10)
		endpoint = invoiced.AddQueryParameter(endpoint, "start_date", startDateString)
	}

	if endDate > 0 {
		endDateString := strconv.FormatInt(endDate, 10)
		endpoint = invoiced.AddQueryParameter(endpoint, "end_date", endDateString)
	}

	payments := make(invoiced.Payments, 0)

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

	return payments, nil
}

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Payments, string, error) {
	endpoint := invoiced.AddFilterAndSort("/payments", filter, sort)
	payments := make(invoiced.Payments, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &payments)
	if err != nil {
		return nil, "", err
	}

	return payments, nextEndpoint, nil
}

func (c *Client) SendReceipt(id int64, request *invoiced.SendEmailRequest) error {
	endpoint := "/payments/" + strconv.FormatInt(id, 10) + "/emails"

	return c.Api.Create(endpoint, request, nil)
}
