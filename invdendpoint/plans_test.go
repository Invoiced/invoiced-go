package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalPlanObject(t *testing.T) {
	s := `{
  "amount": 49,
  "catalog_item": "software-subscription",
  "created_at": 1477418268,
  "currency": "usd",
  "id": "starter",
  "interval": "month",
  "interval_count": 1,
  "metadata": {},
  "name": "Starter",
  "object": "plan",
  "pricing_mode": "per_unit",
  "quantity_type": "constant",
  "tiers": null
}`

	so := new(Plan)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
