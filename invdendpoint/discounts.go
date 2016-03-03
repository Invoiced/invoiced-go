package invdendpoint

type Discount struct {
	Id      int64   `json:"id,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
	Coupon  Rate    `json:"coupon,omitempty"`
	Expires int64   `json:"expires,omitempty"`
}
