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

	if so.Id != 8 {
		t.Fatal("Item 1 has incorrect id")
	}

	if so.Type != "service" {
		t.Fatal("Item 1 has incorrect type")
	}

	if so.Name != "Delivery" {
		t.Fatal("Item 0 has incorrect name")
	}

	if so.Quantity != 1.0 {
		t.Fatal("Item 1 has incorrect quantity")
	}

	if so.UnitCost != 10 {
		t.Fatal("Item 1 has incorrect unit cost")
	}

	if so.Amount != 10 {
		t.Fatal("Item 1 has incorrect amount")
	}

	if !so.Taxable {
		t.Fatal("Item 1 has incorrect taxable")
	}

}
