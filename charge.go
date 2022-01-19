package invdapi

import (
	"errors"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Charge struct {
	*Connection
	*invdendpoint.Charge
}

func (c *Connection) NewCharge() *Charge {
	return &Charge{c, new(invdendpoint.Charge)}
}

func (c *Charge) Create(chargeRequest *invdendpoint.ChargeRequest) (*Payment, error) {
	payment := c.NewPayment()

	if chargeRequest == nil {
		return nil, errors.New("chargeRequest cannot be nil")
	}

	err := c.create(invdendpoint.ChargeEndpoint, chargeRequest, payment)

	if err != nil {
		return nil, err
	}

	payment.Connection = c.Connection

	return payment, nil
}
