package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalLineItemObject(t *testing.T) {
	s := `{
  "id": 8,
  "catalog_item": "delivery",
  "type": "service",
  "name": "Delivery",
  "description": "",
  "quantity": 1,
  "unit_cost": 10,
  "amount": 10,
  "discountable": true,
  "discounts": [],
  "taxable": true,
  "taxes": [],
  "metadata": {}
}`

	so := new(LineItem)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
