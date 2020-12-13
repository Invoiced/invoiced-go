package invdapi

import (
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Refund struct {
	*Connection
	*invdendpoint.Refund
}

func (c *Connection) NewRefund() *Refund {
	return &Refund{c, new(invdendpoint.Refund)}
}

func (c *Refund) Create(chargeId int64, amount float64) error {
	endpoint := invdendpoint.ChargeEndpoint + "/" + strconv.FormatInt(chargeId, 10) + "/refunds"
	c.Refund = new(invdendpoint.Refund)
	err := c.create(endpoint, nil, c.Refund)
	if err != nil {
		return nil
	}

	return nil
}
