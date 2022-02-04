package credit_balance_adjustment

import (
	"strconv"

	"github.com/strongdm/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.CreditBalanceAdjustmentRequest) (*invoiced.CreditBalanceAdjustment, error) {
	resp := new(invoiced.CreditBalanceAdjustment)
	err := c.Api.Create("/credit_balance_adjustments", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.CreditBalanceAdjustment, error) {
	resp := new(invoiced.CreditBalanceAdjustment)
	_, err := c.Api.Get("/credit_balance_adjustments/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.CreditBalanceAdjustmentRequest) (*invoiced.CreditBalanceAdjustment, error) {
	resp := new(invoiced.CreditBalanceAdjustment)
	err := c.Api.Update("/credit_balance_adjustments/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/credit_balance_adjustments/" + strconv.FormatInt(id, 10))
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.CreditBalanceAdjustments, error) {
	endpoint := invoiced.AddFilterAndSort("/credit_balance_adjustments", filter, sort)

	adjustments := make(invoiced.CreditBalanceAdjustments, 0)

NEXT:
	tmpAdjustments := make(invoiced.CreditBalanceAdjustments, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpAdjustments)

	if err != nil {
		return nil, err
	}

	adjustments = append(adjustments, tmpAdjustments...)

	if endpointTmp != "" {
		goto NEXT
	}

	return adjustments, nil
}
