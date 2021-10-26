package invdendpoint

const PlanEndpoint = "/plans"

type Plan struct {
	Id            string                 `json:"id,omitempty"`
	Object        string                 `json:"object,omitempty"`
	Item          string                 `json:"catalog_item,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Currency      string                 `json:"currency,omitempty"`
	Amount        float64                `json:"amount,omitempty"`
	PricingMode   string                 `json:"pricing_mode,omitempty"`
	QuantityType  string                 `json:"quantity_type,omitempty"`
	Interval      string                 `json:"interval,omitempty"`
	IntervalCount float64                `json:"interval_count,omitempty"`
	Tiers         []Tier                 `json:"tier,omitempty"`
	CreatedAt     int64                  `json:"created_at,omitempty"`	//Timestamp when created
	UpdatedAt     int64                  `json:"updated_at,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	NumberOfSubscriptions *int           `json:"num_subscriptions,omitempty"`
}

type Tier struct {
	MaxQty   float64 `json:"max_qty,omitempty"`
	UnitCost float64 `json:"unit_cost,omitempty"`
	MinQty   float64 `json:"min_qty,omitempty"`
}
