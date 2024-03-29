package invoiced

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalItemObject(t *testing.T) {
	s := `{
  "id": "delivery",
  "object": "catalog_item",
  "name": "Delivery",
  "currency": "usd",
  "unit_cost": 100,
  "description": null,
  "type": "service",
  "taxes": [],
  "discountable": true,
  "taxable": true,
  "unit_cost": 10,
  "created_at": 1477327516,
  "metadata": {}
}`

	so := new(Item)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
