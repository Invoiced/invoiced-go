package invdendpoint

type LineItem struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Quantity    int64      `json:"quantity"`
	UnitCost    float64    `json:"unit_cost"`
	Amount      float64    `json:"amount"`
	Discounts   []Discount `json:"discounts,omitempty"`
	Taxes       []Tax      `json:"taxes,omitempty"`
	Shipping    []Rate     `json:"shipping,omitempty"`
}
