package invdendpoint

const RateEndpoint = "/tax_rates"

type TaxRate struct {
	Id        string                 `json:"id,omitempty"`
	Object    string                 `json:"object,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Currency  string                 `json:"currency,omitempty"`
	Value     float64                `json:"value,omitempty"`
	IsPercent bool                   `json:"is_percent,omitempty"`
	Inclusive bool                   `json:"inclusive,omitempty"`
	CreatedAt int64                  `json:"created_at,omitempty"`	//Timestamp when created
	UpdatedAt int64                  `json:"updated_at,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}
