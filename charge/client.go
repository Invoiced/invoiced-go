package charge

import (
	"strconv"

	"github.com/strongdm/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.ChargeRequest) (*invoiced.Charge, error) {
	resp := new(invoiced.Charge)
	err := c.Api.Create("/charges", request, resp)
	return resp, err
}

func (c *Client) Refund(chargeId int64, request *invoiced.RefundRequest) (*invoiced.Refund, error) {
	refund := new(invoiced.Refund)
	err := c.Api.Create("/charges/"+strconv.FormatInt(chargeId, 10)+"/refunds", request, refund)
	return refund, err
}
