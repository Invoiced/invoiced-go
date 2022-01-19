package invdendpoint

const PlanEndpoint = "/plans"

type PlanRequest struct {
	Amount        *float64                `json:"amount,omitempty"`
	Currency      *string                 `json:"currency,omitempty"`
	Id            *string                 `json:"id,omitempty"`
	Interval      *string                 `json:"interval,omitempty"`
	IntervalCount *float64                `json:"interval_count,omitempty"`
	Item          *string                 `json:"catalog_item,omitempty"`
	Metadata      *map[string]interface{} `json:"metadata,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	PricingMode   *string                 `json:"pricing_mode,omitempty"`
	QuantityType  *string                 `json:"quantity_type,omitempty"`
	Tiers         []*TierRequest          `json:"tier,omitempty"`
}

type TierRequest struct {
	MaxQty   *float64 `json:"max_qty,omitempty"`
	MinQty   *float64 `json:"min_qty,omitempty"`
	UnitCost *float64 `json:"unit_cost,omitempty"`
}

type Plan struct {
	Amount                float64                `json:"amount"`
	CreatedAt             int64                  `json:"created_at"`
	Currency              string                 `json:"currency"`
	Id                    string                 `json:"id"`
	Interval              string                 `json:"interval"`
	IntervalCount         float64                `json:"interval_count"`
	Item                  string                 `json:"catalog_item"`
	Metadata              map[string]interface{} `json:"metadata"`
	Name                  string                 `json:"name"`
	NumberOfSubscriptions *int64                 `json:"num_subscriptions"`
	Object                string                 `json:"object"`
	PricingMode           string                 `json:"pricing_mode"`
	QuantityType          string                 `json:"quantity_type"`
	Tiers                 []Tier                 `json:"tier"`
	UpdatedAt             int64                  `json:"updated_at"`
}

type Tier struct {
	MaxQty   float64 `json:"max_qty"`
	MinQty   float64 `json:"min_qty"`
	UnitCost float64 `json:"unit_cost"`
}
