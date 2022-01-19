package invdendpoint

type LineItemRequest struct {
	Amount       *float64                `json:"amount,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	Discountable *bool                   `json:"discountable,omitempty"`
	Discounts    []*DiscountRequest      `json:"discounts,omitempty"`
	Item         *string                 `json:"catalog_item,omitempty"`
	Metadata     *map[string]interface{} `json:"metadata,omitempty"`
	Name         *string                 `json:"name,omitempty"`
	PeriodEnd    *int64                  `json:"period_end,omitempty"`
	PeriodStart  *int64                  `json:"period_start,omitempty"`
	Plan         *string                 `json:"plan,omitempty"`
	Prorated     *bool                   `json:"prorated,omitempty"`
	Quantity     *float64                `json:"quantity,omitempty"`
	Taxable      *bool                   `json:"taxable,omitempty"`
	Taxes        []*TaxRequest           `json:"taxes,omitempty"`
	Type         *string                 `json:"type,omitempty"`
	UnitCost     *float64                `json:"unit_cost,omitempty"`
}

type LineItem struct {
	Amount       float64                `json:"amount"`
	Description  string                 `json:"description"`
	Discountable bool                   `json:"discountable"`
	Discounts    []Discount             `json:"discounts"`
	Id           int64                  `json:"id"`
	Item         string                 `json:"catalog_item"`
	Metadata     map[string]interface{} `json:"metadata"`
	Name         string                 `json:"name"`
	PeriodEnd    int64                  `json:"period_end"`
	PeriodStart  int64                  `json:"period_start"`
	Plan         string                 `json:"plan"`
	Prorated     bool                   `json:"prorated"`
	Quantity     float64                `json:"quantity"`
	Taxable      bool                   `json:"taxable"`
	Taxes        []Tax                  `json:"taxes"`
	Type         string                 `json:"type"`
	UnitCost     float64                `json:"unit_cost"`
}

type LineItemPreview struct {
	Amount       float64                `json:"amount"`
	Description  string                 `json:"description"`
	Discountable bool                   `json:"discountable"`
	Discounts    []Discount             `json:"discounts"`
	Item         string                 `json:"catalog_item"`
	Metadata     map[string]interface{} `json:"metadata"`
	Name         string                 `json:"name"`
	PeriodEnd    int64                  `json:"period_end"`
	PeriodStart  int64                  `json:"period_start"`
	Plan         string                 `json:"plan"`
	Prorated     bool                   `json:"prorated"`
	Quantity     float64                `json:"quantity"`
	Taxable      bool                   `json:"taxable"`
	Taxes        []Tax                  `json:"taxes"`
	Type         string                 `json:"type"`
	UnitCost     float64                `json:"unit_cost"`
}
