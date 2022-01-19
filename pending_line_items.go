package invoiced

type PendingLineItems []PendingLineItem

type PendingLineItemRequest struct {
	Description  *string                 `json:"description,omitempty"`
	Discountable *bool                   `json:"discountable,omitempty"`
	Discounts    []*DiscountRequest      `json:"discounts,omitempty"`
	Item         *string                 `json:"catalog_item,omitempty"`
	Metadata     *map[string]interface{} `json:"metadata,omitempty"`
	Name         *string                 `json:"name,omitempty"`
	Quantity     *float64                `json:"quantity,omitempty"`
	Taxable      *bool                   `json:"taxable,omitempty"`
	Taxes        []*TaxRequest           `json:"taxes,omitempty"`
	Type         *string                 `json:"type,omitempty"`
	UnitCost     *float64                `json:"unit_cost,omitempty"`
}

type PendingLineItem struct {
	Description  string                 `json:"description"`
	Discountable bool                   `json:"discountable"`
	Discounts    []Discount             `json:"discounts"`
	Id           int64                  `json:"id"`
	Item         string                 `json:"catalog_item"`
	Metadata     map[string]interface{} `json:"metadata"`
	Name         string                 `json:"name"`
	Quantity     float64                `json:"quantity"`
	Taxable      bool                   `json:"taxable"`
	Taxes        []Tax                  `json:"taxes"`
	Type         string                 `json:"type"`
	UnitCost     float64                `json:"unit_cost"`
}
