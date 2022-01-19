package invoiced

const ItemEndpoint = "/items"

type ItemRequest struct {
	AvalaraLocationCode *string                 `json:"avalara_location_code,omitempty"`
	AvalaraTaxCode      *string                 `json:"avalara_tax_code,omitempty"`
	Currency            *string                 `json:"currency,omitempty"`
	Description         *string                 `json:"description,omitempty"`
	Discountable        *bool                   `json:"discountable,omitempty"`
	GlAccount           *string                 `json:"gl_account,omitempty"`
	Id                  *string                 `json:"id,omitempty"`
	Metadata            *map[string]interface{} `json:"metadata,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	Taxable             *bool                   `json:"taxable,omitempty"`
	Taxes               []*TaxRequest           `json:"taxes,omitempty"`
	Type                *string                 `json:"service,omitempty"`
	UnitCost            *float64                `json:"unit_cost,omitempty"`
}

type Item struct {
	AvalaraLocationCode string                 `json:"avalara_location_code"`
	AvalaraTaxCode      string                 `json:"avalara_tax_code"`
	CreatedAt           int64                  `json:"created_at"`
	Currency            string                 `json:"currency"`
	Description         string                 `json:"description"`
	Discountable        bool                   `json:"discountable"`
	GlAccount           string                 `json:"gl_account"`
	Id                  string                 `json:"id"`
	Metadata            map[string]interface{} `json:"metadata"`
	Name                string                 `json:"name"`
	Object              string                 `json:"object"`
	Taxable             bool                   `json:"taxable"`
	Taxes               []Tax                  `json:"taxes"`
	Type                string                 `json:"service"`
	UnitCost            float64                `json:"unit_cost"`
	UpdatedAt           int64                  `json:"updated_at"`
}

type Items []Item
