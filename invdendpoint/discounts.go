package invdendpoint

// Represents the application of a discount to an invoice or line item.
type Discount struct {
	Id      int64   `json:"id,omitempty"`      // The discount’s unique ID
	Amount  float64 `json:"amount,omitempty"`  // Discount amount
	Coupon  TaxRate `json:"coupon,omitempty"`  // Coupon the discount was computed from, if any
	Expires int64   `json:"expires,omitempty"` // Time until discount expires, if any
}
