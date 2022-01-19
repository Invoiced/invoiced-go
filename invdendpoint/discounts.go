package invdendpoint

type DiscountRequest struct {
	Amount  *float64 `json:"amount,omitempty"`
	Coupon  *TaxRate `json:"coupon,omitempty"`
	Expires *int64   `json:"expires,omitempty"`
}

type Discount struct {
	Id      int64   `json:"id"`
	Amount  float64 `json:"amount"`
	Coupon  TaxRate `json:"coupon"`
	Expires int64   `json:"expires"`
}
