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
	endPoint := c.MakeEndPointURL(invdendpoint.CouponsEndPoint)

	couponResp := new(Coupon)


	if coupon == nil {
		return nil, errors.New("coupon is nil")
	}

	//safe prune file data for creation
	invdCouponDataToCreate, err := SafeCouponForCreation(coupon.Coupon)

	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdCouponDataToCreate, couponResp)

	if apiErr != nil {
		return nil, apiErr
	}

	couponResp.Connection = c.Connection

	return couponResp, nil

}

func (c *Coupon) Save() error {
	endPoint := c.MakeEndPointURL(invdendpoint.CouponsEndPoint) + "/" + c.Id

	taxRateResp := new(Coupon)

	invdTaxRatDataToUpdate, err := SafeCouponForUpdating(c.Coupon)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, invdTaxRatDataToUpdate, taxRateResp)

	if apiErr != nil {
		return apiErr
	}

	c.Coupon = taxRateResp.Coupon

	return nil

}

func (c *Coupon) Delete() error {
	endPoint := c.MakeEndPointURL(invdendpoint.CouponsEndPoint) + "/" + c.Id

	err := c.delete(endPoint)

	if err != nil {
		return err
	}

	return nil

}

func (c *Coupon) Retrieve(id string) (*Coupon, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CouponsEndPoint) + "/" + id

	couponEndPoint := new(invdendpoint.Coupon)

	coupon := &Coupon{c.Connection, couponEndPoint}

	_, err := c.retrieveDataFromAPI(endPoint, coupon)

	if err != nil {
		return nil, err
	}

	return coupon, nil

}

func (c *Coupon) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Coupons, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CouponsEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	coupons := make(Coupons, 0)

NEXT:
	tmpCoupons := make(Coupons, 0)

	endPointTmp, apiErr := c.retrieveDataFromAPI(endPoint, &tmpCoupons)

	if apiErr != nil {
		return nil, apiErr
	}

	coupons = append(coupons, tmpCoupons...)

	if endPointTmp != "" {
		goto NEXT
	}

	for _, coupon := range coupons {
		coupon.Connection = c.Connection

	}

	return coupons, nil

}

//SafeCouponForCreation prunes coupon data for just fields that can be used for creation of a coupon
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

//SafeTaxRateForUpdating prunes coupon data for just fields that can be used for updating of a plan
func SafeCouponForUpdating(coupon *invdendpoint.Coupon) (*invdendpoint.Coupon, error) {

	if coupon == nil {
		return nil, errors.New("coupon is nil")
	}


	couponData := new(invdendpoint.Coupon)
	couponData.Name = coupon.Name
	couponData.Metadata = coupon.Metadata

	return couponData, nil
}


