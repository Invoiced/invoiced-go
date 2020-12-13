package invdapi

import "github.com/Invoiced/invoiced-go/invdendpoint"

func (c *Payment) Refund(refund float64) error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.PaymentEndpoint), c.Id) + "/refunds"
	c.Payment = new(invdendpoint.Payment)
	err := c.create(endPoint, nil, c.Payment)
	if err != nil {
		return nil
	}

	return nil
}
