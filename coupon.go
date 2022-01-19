package invoiced

type CouponClient struct {
	*Client
	*Coupon
}

type Coupons []*CouponClient

func (c *Client) NewCoupon() *CouponClient {
	coupon := new(Coupon)
	return &CouponClient{c, coupon}
}

func (c *CouponClient) Create(request *CouponRequest) (*CouponClient, error) {
	resp := new(CouponClient)

	err := c.Api.Create(CouponEndpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CouponClient) Retrieve(id string) (*CouponClient, error) {
	endpoint := CouponEndpoint + "/" + id

	couponEndpoint := new(Coupon)

	coupon := &CouponClient{c.Client, couponEndpoint}

	_, err := c.Api.Get(endpoint, coupon)
	if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (c *CouponClient) Update(request *CouponRequest) error {
	endpoint := CouponEndpoint + "/" + c.Id
	resp := new(CouponClient)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Coupon = resp.Coupon

	return nil
}

func (c *CouponClient) Delete() error {
	endpoint := CouponEndpoint + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *CouponClient) ListAll(filter *Filter, sort *Sort) (Coupons, error) {
	endpoint := AddFilterAndSort(CouponEndpoint, filter, sort)

	coupons := make(Coupons, 0)

NEXT:
	tmpCoupons := make(Coupons, 0)

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
