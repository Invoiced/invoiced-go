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
            "catalog_item": "ipad-license",
            "quantity": 11,
            "created_at": 1420391704
        }
    ],
    "discounts": [],
    "taxes": [],
    "url": "https://dundermifflin.invoiced.com/subscriptions/o2mAd2wWVfYy16XZto7xHwXX",
    "created_at": 1420391704,
    "metadata": {},
    "prorate": true,
    "contract_renewal_mode": "manual",
    "ship_to": {
        "address1": "123 Main St",
        "address2": "Ste 100",
        "attention_to": "Regina Smith",
        "city": "Austin",
        "country": "US",
        "name": "Company Name",
        "postal_code": "78730",
        "state": "TX"
    }
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

	if so.Addons[0].CatalogItem != "ipad-license" {
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

	if so.Prorate != true {
		t.Fatal("Subscription has incorrect Prorate status")
	}

	if so.ContractRenewalMode != "manual" {
		t.Fatal("Subscription Ahas incorrect ContractRenewalMode status")
	}

	if so.CreatedAt != 1420391704 {
		t.Fatal("Subscription CreatedAt is incorrect")
	}

	if so.ShipTo.Address1 != "123 Main St" {
		t.Fatal("Subscription ShipTo.Address1 is incorrect")
	}

	if so.ShipTo.Address2 != "Ste 100" {
		t.Fatal("Subscription ShipTo.Address2 is incorrect")
	}

	if so.ShipTo.AttentionTo != "Regina Smith" {
		t.Fatal("Subscription ShipTo.AttentionTo is incorrect")
	}

	if so.ShipTo.City != "Austin" {
		t.Fatal("Subscription ShipTo.City is incorrect")
	}

	if so.ShipTo.Country != "US" {
		t.Fatal("Subscription ShipTo.Country is incorrect")
	}

	if so.ShipTo.Name != "Company Name" {
		t.Fatal("Subscription ShipTo.Name is incorrect")
	}

	if so.ShipTo.PostalCode != "78730" {
		t.Fatal("Subscription ShipTo.PostalCode is incorrect")
	}

	if so.ShipTo.State != "TX" {
		t.Fatal("Subscription ShipTo.State is incorrect")
	}
}
