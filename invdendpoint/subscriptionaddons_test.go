package invdendpoint

import (
	"encoding/json"
	"testing"
)


func TestUnMarshalSubscriptionAddonObject(t *testing.T) {
  catalogItem := `{
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

  s := `{
    "id": 3,
    "catalog_item":` + catalogItem + `,
    "plan" : "test-plan",
    "quantity": 11,
    "created_at": 1420391704
}`

	so := new(SubscriptionAddon)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
