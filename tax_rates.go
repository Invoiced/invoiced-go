package invoiced

const RateEndpoint = "/tax_rates"

type TaxRateRequest struct {
	Currency  *string                 `json:"currency,omitempty"`
	Id        *string                 `json:"id,omitempty"`
	Inclusive *bool                   `json:"inclusive,omitempty"`
	IsPercent *bool                   `json:"is_percent,omitempty"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	Value     *float64                `json:"value,omitempty"`
}

type TaxRate struct {
	CreatedAt int64                  `json:"created_at"`
	Currency  string                 `json:"currency"`
	Id        string                 `json:"id"`
	Inclusive bool                   `json:"inclusive"`
	IsPercent bool                   `json:"is_percent"`
	Metadata  map[string]interface{} `json:"metadata"`
	Name      string                 `json:"name"`
	Object    string                 `json:"object"`
	UpdatedAt int64                  `json:"updated_at"`
	Value     float64                `json:"value"`
}
