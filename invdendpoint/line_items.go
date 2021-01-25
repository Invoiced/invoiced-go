package invdendpoint

type LineItem struct {
	Id           int64                  `json:"id,omitempty"`           // The line item’s unique ID
	Item         string                 `json:"catalog_item,omitempty"` // Optional Item ID. Fills the line item with the name and pricing of the Item.
	Type         string                 `json:"type,omitempty"`         // Optional line item type. Used to group line items by type in reporting
	Name         string                 `json:"name,omitempty"`         // Title
	Description  string                 `json:"description,omitempty"`  // Optional description
	Quantity     float64                `json:"quantity,omitempty"`               // Quantity
	UnitCost     float64                `json:"unit_cost,omitempty"`    // Unit cost or rate
	Amount       float64                `json:"amount,omitempty"`       // Computed from quantity x unit_cost
	Discountable bool                   `json:"discountable,omitempty"` // Excludes amount from invoice discounts when false
	Discounts    []Discount             `json:"discounts,omitempty"`    // Line item Discounts
	Taxable      bool                   `json:"taxable,omitempty"`      // Excludes amount from invoice taxes when false
	Taxes        []Tax                  `json:"taxes,omitempty"`        // Line item Taxes
	Plan         string                 `json:"plan,omitempty"`         // Plan ID, only present when type is plan
	PeriodStart  int64                  `json:"period_start,omitempty"`
	PeriodEnd    int64                  `json:"period_end,omitempty"`
	Prorated     bool                   `json:"prorated,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"` // A hash of key/value pairs that can store additional information about this object.
}

type LineItemPreview struct {
	Item         string                 `json:"catalog_item,omitempty"` // Optional Item ID. Fills the line item with the name and pricing of the Item.
	Type         string                 `json:"type,omitempty"`         // Optional line item type. Used to group line items by type in reporting
	Name         string                 `json:"name,omitempty"`         // Title
	Description  string                 `json:"description,omitempty"`  // Optional description
	Quantity     float64                `json:"quantity,omitempty"`     // Quantity
	UnitCost     float64                `json:"unit_cost,omitempty"`    // Unit cost or rate
	Amount       float64                `json:"amount,omitempty"`       // Computed from quantity x unit_cost
	Discountable bool                   `json:"discountable,omitempty"` // Excludes amount from invoice discounts when false
	Discounts    []Discount             `json:"discounts,omitempty"`    // Line item Discounts
	Taxable      bool                   `json:"taxable,omitempty"`      // Excludes amount from invoice taxes when false
	Taxes        []Tax                  `json:"taxes,omitempty"`        // Line item Taxes
	Plan         string                 `json:"plan,omitempty"`         // Plan ID, only present when type is plan
	PeriodStart  int64                  `json:"period_start,omitempty"`
	PeriodEnd    int64                  `json:"period_end,omitempty"`
	Prorated     bool                   `json:"prorated,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"` // A hash of key/value pairs that can store additional information about this object.
}
