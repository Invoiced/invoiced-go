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

	fmt.Println("Subscription plan -> ", so.Plan)

	if so.Id != 595 {
		t.Fatal("Subscription has incorrect id")
	}

	if so.Customer != 15444 {
		t.Fatal("Subscription has incorrect type")
	}
	fmt.Println("so.plan ", so.Plan)
	if so.Plan != "starter" {
		t.Fatal("Subscription has incorrect plan2 -> " + so.Plan)
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

func TestSubscriptionPlanUnmarshall(t *testing.T) {
	s := `{
      "bill_in": "advance",
      "bill_in_advance_days": 0,
      "cancel_at_period_end": false,
      "canceled_at": null,
      "contract_period_end": 1627448399,
      "contract_period_start": 1624856400,
      "contract_renewal_cycles": null,
      "contract_renewal_mode": "auto",
      "created_at": 1624934799,
      "customer": {
        "ach_gateway_id": null,
        "address1": null,
        "address2": null,
        "attention_to": null,
        "autopay": false,
        "autopay_delay_days": -1,
        "avalara_entity_use_code": null,
        "avalara_exemption_number": null,
        "bill_to_parent": false,
        "cc_gateway_id": null,
        "chase": true,
        "chasing_cadence": null,
        "city": null,
        "consolidated": false,
        "country": "US",
        "created_at": 1624930142,
        "credit_hold": false,
        "credit_limit": null,
        "currency": "usd",
        "email": null,
        "id": 2321739,
        "language": null,
        "name": "Parag",
        "next_chase_step": null,
        "notes": null,
        "number": "acme00209",
        "owner": 831,
        "parent_customer": null,
        "payment_terms": "NET 14",
        "phone": null,
        "postal_code": null,
        "state": null,
        "tax_id": null,
        "taxable": true,
        "taxes": [],
        "type": "company",
        "object": "customer",
        "statement_pdf_url": "https://tesla.sandbox.invoiced.com/statements/ReOax2A5W6bIt8V4paAmGvEn/pdf",
        "sign_up_url": null,
        "payment_source": null,
        "sign_up_page": null,
        "metadata": {}
      },
      "cycles": 1,
      "description": null,
      "id": 62241,
      "mrr": 80900,
      "paused": false,
      "period_end": 1627448399,
      "period_start": 1624856400,
      "plan": {
        "amount": 80900,
        "catalog_item": null,
        "created_at": 1624934793,
        "currency": "usd",
        "description": null,
        "id": "model-z",
        "interval": "month",
        "interval_count": 1,
        "name": "Model Z",
        "notes": null,
        "pricing_mode": "per_unit",
        "quantity_type": "constant",
        "tiers": null,
        "object": "plan",
        "metadata": {}
      },
      "quantity": 1,
      "recurring_total": 80900,
      "renewed_last": 1624856400,
      "renews_next": 1627448400,
      "snap_to_nth_day": null,
      "start_date": 1624856400,
      "status": "active",
      "taxes": [],
      "object": "subscription",
      "url": "https://tesla.sandbox.invoiced.com/subscriptions/pE9pBoU0HmF6dyAXxyAYstOk",
      "approval": null,
      "payment_source": null,
      "ship_to": null,
      "metadata": {},
      "addons": [],
      "discounts": []
    }`

	so := new(Subscription)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}

	if so.Plan != "model-z" {
		t.Fatal("Plan id is incorrect")
	}

}
