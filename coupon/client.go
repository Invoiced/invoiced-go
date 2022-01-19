package coupon

import "github.com/Invoiced/invoiced-go"

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.CouponRequest) (*Client, error) {
	resp := new(Client)

	err := c.Api.Create("/coupons", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id string) (*invoiced.Coupon, error) {
	resp := new(invoiced.Coupon)
	_, err := c.Api.Get("/coupons" + "/" + id, resp)
	return resp, err
}

func (c *Client) Update(request *invoiced.CouponRequest) error {
	endpoint := "/coupons" + "/" + c.Id
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete() error {
	endpoint := "/coupons" + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
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
