package invdendpoint

//Represents the application of tax to an invoice or line item.
type Tax struct {
	Id      int64   `json:"id,omitempty"`       //The tax’s unique ID
	Amount  float64 `json:"amount,omitempty"`   //Tax amount
	TaxRate TaxRate `json:"tax_rate,omitempty"` //Tax Rate the tax was computed from, if any
}
