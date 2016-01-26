package invdendpoint

type Tax struct {
	Id     int64   `json:"id,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	//TaxRate int64   `json:"tax_rate,omitempty"`
}
