package coupon

import (
	"github.com/Invoiced/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.CouponRequest) (*invoiced.Coupon, error) {
	resp := new(invoiced.Coupon)
	err := c.Api.Create("/coupons", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id string) (*invoiced.Coupon, error) {
	resp := new(invoiced.Coupon)
	_, err := c.Api.Get("/coupons/"+id, resp)
	return resp, err
}

func (c *Client) Update(id string, request *invoiced.CouponRequest) (*invoiced.Coupon, error) {
	resp := new(invoiced.Coupon)
	err := c.Api.Update("/coupons/"+id, request, resp)
	return resp, err
}

func (c *Client) Delete(id string) error {
	return c.Api.Delete("/coupons/" + id)
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Coupons, error) {
	endpoint := invoiced.AddFilterAndSort("/coupons", filter, sort)

	coupons := make(invoiced.Coupons, 0)

NEXT:
	tmpCoupons := make(invoiced.Coupons, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpCoupons)

	if err != nil {
		return nil, err
	}

	coupons = append(coupons, tmpCoupons...)

	if endpointTmp != "" {
		goto NEXT
	}

	return coupons, nil
}
