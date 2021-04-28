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

	apiErr := c.create(invdendpoint.ChargeEndpoint, chargeRequest, payment)

	if apiErr != nil {
		return nil, apiErr
	}

	payment.Connection = c.Connection

	return payment, nil
}
