package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalSubscriptionObject(t *testing.T) {
	s := `{
    "id": 595,
    "customer": 15444,
    "plan": "starter",
    "cycles": null,
    "quantity": 1,
    "start_date": 1420391704,
    "period_start": 1446657304,
    "period_end": 1449249304,
    "status": "active",
    "addons": [
        {
            "id": 3,
            "catalog_item": {
  "id": "ipad-license",
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
},
            "quantity": 11,
            "created_at": 1420391704
        }
    ],
    "discounts": [],
    "taxes": [],
    "url": "https://dundermifflin.invoiced.com/subscriptions/o2mAd2wWVfYy16XZto7xHwXX",
    "created_at": 1420391704,
    "metadata": {}
}`

	so := new(Subscription)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 595 {
		t.Fatal("Subscription has incorrect id")
	}

	if so.Customer != 15444 {
		t.Fatal("Subscription has incorrect type")
	}

	if so.Plan != "starter" {
		t.Fatal("Subscription has incorrect plan")
	}

	if so.Quantity != 1 {
		t.Fatal("Subscription has incorrect quantity")
	}

	if so.StartDate != 1420391704 {
		t.Fatal("Subscription has incorrect quantity")
	}

	if so.PeriodStart != 1446657304 {
		t.Fatal("Subscription has incorrect periodstart")
	}

	if so.PeriodEnd != 1449249304 {
		t.Fatal("Subscription has incorrect periodstart")
	}

	if so.Status != "active" {
		t.Fatal("Subscription has incorrect status")
	}

	if so.Addons[0].Id != 3 {
		t.Fatal("Subscription Addon 0 has incorrect status")
	}

	if so.Addons[0].CatalogItem.Id != "ipad-license" {
		t.Fatal("Subscription Addon CatalogItem 0  has incorrect status")
	}

	if so.Addons[0].Quantity != 11 {
		t.Fatal("Subscription Addon Quantity 0  has incorrect status")
	}

	if so.Addons[0].CreatedAt != 1420391704 {
		t.Fatal("Quantity Addon CreatedAT has incorrect status")
	}

	if so.Url != "https://dundermifflin.invoiced.com/subscriptions/o2mAd2wWVfYy16XZto7xHwXX" {
		t.Fatal("Url is incorrect")

	}

	if so.CreatedAt != 1420391704 {
		t.Fatal("Subscription CreatedAt is incorrect")

	}

}
