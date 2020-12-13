package invdapi

import (
	"errors"

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

func (c *Coupon) Create(coupon *Coupon) (*Coupon, error) {
	couponResp := new(Coupon)

	if coupon == nil {
		return nil, errors.New("coupon is nil")
	}

	// safe prune file data for creation
	invdCouponDataToCreate, err := SafeCouponForCreation(coupon.Coupon)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(invdendpoint.CouponEndpoint, invdCouponDataToCreate, couponResp)

	if apiErr != nil {
		return nil, apiErr
	}

	couponResp.Connection = c.Connection

	return couponResp, nil
}

func (c *Coupon) Save() error {
	endpoint := invdendpoint.CouponEndpoint + "/" + c.Id

	taxRateResp := new(Coupon)

	invdTaxRatDataToUpdate, err := SafeCouponForUpdating(c.Coupon)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, invdTaxRatDataToUpdate, taxRateResp)

	if apiErr != nil {
		return apiErr
	}

	c.Coupon = taxRateResp.Coupon

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

func (c *Coupon) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Coupons, error) {
	endpoint := addFilterAndSort(invdendpoint.CouponEndpoint, filter, sort)

	coupons := make(Coupons, 0)

NEXT:
	tmpCoupons := make(Coupons, 0)

	endpointTmp, apiErr := c.retrieveDataFromAPI(endpoint, &tmpCoupons)

	if apiErr != nil {
		return nil, apiErr
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

// SafeCouponForCreation prunes coupon data for just fields that can be used for creation of a coupon
func SafeCouponForCreation(coupon *invdendpoint.Coupon) (*invdendpoint.Coupon, error) {
	if coupon == nil {
		return nil, errors.New("coupon is nil")
	}

	couponData := new(invdendpoint.Coupon)
	couponData.Id = coupon.Id
	couponData.Name = coupon.Name
	couponData.Currency = coupon.Currency
	couponData.Value = coupon.Value
	couponData.IsPercent = coupon.IsPercent
	couponData.Exclusive = coupon.Exclusive
	couponData.ExpirationDate = coupon.ExpirationDate
	couponData.MaxRedemptions = coupon.MaxRedemptions
	couponData.Metadata = coupon.Metadata

	return couponData, nil
}

// SafeTaxRateForUpdating prunes coupon data for just fields that can be used for updating of a plan
func SafeCouponForUpdating(coupon *invdendpoint.Coupon) (*invdendpoint.Coupon, error) {
	if coupon == nil {
		return nil, errors.New("coupon is nil")
	}

	couponData := new(invdendpoint.Coupon)
	couponData.Name = coupon.Name
	couponData.Metadata = coupon.Metadata

	return couponData, nil
}
