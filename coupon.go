package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Coupon struct {
	*Connection
	*invdendpoint.Coupon
}

type Coupons []*Coupon

func (c *Connection) NewCoupon() *Coupon {
	coupon := new(invdendpoint.Coupon)
	return &Coupon{c, coupon}
}

func (c *Coupon) Create(request *invdendpoint.CouponRequest) (*Coupon, error) {
	resp := new(Coupon)

	err := c.create(invdendpoint.CouponEndpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Coupon) Retrieve(id string) (*Coupon, error) {
	endpoint := invdendpoint.CouponEndpoint + "/" + id

	couponEndpoint := new(invdendpoint.Coupon)

	coupon := &Coupon{c.Connection, couponEndpoint}

	_, err := c.retrieveDataFromAPI(endpoint, coupon)
	if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (c *Coupon) Update(request *invdendpoint.CouponRequest) error {
	endpoint := invdendpoint.CouponEndpoint + "/" + c.Id
	resp := new(Coupon)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Coupon = resp.Coupon

	return nil
}

func (c *Coupon) Delete() error {
	endpoint := invdendpoint.CouponEndpoint + "/" + c.Id

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Coupon) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Coupons, error) {
	endpoint := addFilterAndSort(invdendpoint.CouponEndpoint, filter, sort)

	coupons := make(Coupons, 0)

NEXT:
	tmpCoupons := make(Coupons, 0)

	endpointTmp, err := c.retrieveDataFromAPI(endpoint, &tmpCoupons)

	if err != nil {
		return nil, err
	}

	coupons = append(coupons, tmpCoupons...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, coupon := range coupons {
		coupon.Connection = c.Connection
	}

	return coupons, nil
}
