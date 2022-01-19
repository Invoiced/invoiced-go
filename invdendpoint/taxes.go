package invdendpoint

type TaxRequest struct {
	Amount  *float64 `json:"amount,omitempty"`
	TaxRate *TaxRate `json:"tax_rate,omitempty"`
}

type Tax struct {
	Amount  float64 `json:"amount"`
	Id      int64   `json:"id"`
	TaxRate TaxRate `json:"tax_rate"`
}
