package invdapi

import (
	"errors"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

func (c *Payment) InitiateCharge(chargeRequest *invdendpoint.ChargeRequest) (*Payment, error) {
	txnResp := new(Payment)

	if chargeRequest == nil {
		return nil, errors.New("chargeRequest cannot be nil")
	}

	apiErr := c.create(invdendpoint.ChargeEndpoint, chargeRequest, txnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	txnResp.Connection = c.Connection

	return txnResp, nil
}
