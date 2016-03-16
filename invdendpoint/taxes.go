package invdendpoint

type Tax struct {
	Id      int64   `json:"id,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
	TaxRate Rate    `json:"tax_rate,omitempty"`
}
