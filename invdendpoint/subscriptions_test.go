package invdendpoint

import (
	"encoding/json"
	"fmt"
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
            "plan": "ipad-license",
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

	fmt.Println("Subscription plan -> ",so.Plan)

	if so.Id != 595 {
		t.Fatal("Subscription has incorrect id")
	}

	if so.Customer != 15444 {
		t.Fatal("Subscription has incorrect type")
	}
    fmt.Println("so.plan ",so.Plan)
	if so.Plan != "starter" {
		t.Fatal("Subscription has incorrect plan2 -> " +so.Plan)
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

	if so.Addons[0].Plan != "ipad-license" {
		t.Fatal("Subscription Addon Plan 0 has incorrect value")
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
}

func TestUnMarshalSubscriptionsPreview(t *testing.T) {
	data := `{
  "first_invoice": {
    "attempt_count": null,
    "autopay": null,
    "balance": 0,
    "chase": false,
    "closed": false,
    "created_at": null,
    "csv_url": null,
    "currency": "usd",
    "customer": -1,
    "date": 1571410119,
    "discounts": [],
    "draft": true,
    "due_date": null,
    "id": false,
    "items": [
      {
        "amount": 49,
        "catalog_item": null,
        "created_at": null,
        "description": "",
        "discountable": true,
        "discounts": [],
        "id": false,
        "metadata": {},
        "name": "Starter",
        "object": "line_item",
        "plan": "starter",
        "quantity": 1,
        "taxable": true,
        "taxes": [],
        "type": "plan",
        "unit_cost": 49
      }
    ],
    "metadata": {},
    "name": "Starter",
    "needs_attention": null,
    "next_chase_on": null,
    "next_payment_attempt": null,
    "notes": null,
    "number": null,
    "object": "invoice",
    "paid": false,
    "payment_plan": null,
    "payment_source": null,
    "payment_terms": null,
    "payment_url": null,
    "pdf_url": null,
    "purchase_order": null,
    "ship_to": null,
    "shipping": [],
    "status": "draft",
    "subscription": false,
    "subtotal": 49,
    "taxes": [],
    "total": 49,
    "url": null
  },
  "mrr": 49,
  "recurring_total": 50
}`

	so := new(SubscriptionPreview)

	err := json.Unmarshal([]byte(data), so)
	if err != nil {
		t.Fatal(err)
	}

	if so.MRR != 49 {
		t.Fatal("MRR does not match")
	}

	if so.FirstInvoice == nil {
		t.Fatal("First Invoice should not be fatal")
	}

	if so.RecurringTotal != 50 {
		t.Fatal("Recurring total should not ")
	}
}
