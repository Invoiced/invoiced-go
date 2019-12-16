package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalPliObject(t *testing.T) {
	s := `{
  "amount": 10,
  "catalog_item": "delivery",
  "customer": 15444,
  "description": null,
  "discountable": true,
  "discounts": [],
  "id": 8,
  "metadata": {},
  "name": "Delivery",
  "object": "line_item",
  "quantity": 1,
  "taxable": true,
  "taxes": [],
  "type": "service",
  "unit_cost": 10
}`

	so := new(PendingLineItem)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
