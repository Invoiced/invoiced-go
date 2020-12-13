package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"strconv"
)

func (c *Payment) Refund(refund float64) error {
	endpoint := invdendpoint.PaymentEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/refunds"
	c.Payment = new(invdendpoint.Payment)
	err := c.create(endpoint, nil, c.Payment)
	if err != nil {
		return nil
	}

	return nil
}
