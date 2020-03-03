package invdendpoint

import (
	"encoding/json"
	"github.com/Invoiced/invoiced-go/invdutil"
	"strings"
	"testing"
)

func TestUnMarshalEventObject(t *testing.T) {
	s := `{
    "id": 1228003,
    "type": "transaction.created",
    "data": {
        "object": {
            "amount": 55,
            "created_at": 1451500772,
            "currency": "usd",
            "customer": 15455,
            "date": 1451500771,
            "fee": 0,
            "gateway": null,
            "gateway_id": null,
            "id": 212047,
            "invoice": 196539,
            "metadata": [],
            "method": "other",
            "notes": null,
            "parent_transaction": null,
            "status": "succeeded",
            "theme": null,
            "type": "payment",
            "pdf_url": "https:\/\/dundermifflin.invoiced.com\/payments\/59FHO96idoXFeiBDu1y5Zggg\/pdf",
            "metadata": {}
        }
    }
}`

	so := new(Event)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 1228003 {
		t.Fatal("Event id is incorrect")
	}

	if so.Type != "transaction.created" {
		t.Fatal("Event type is incorrect")
	}

	object := `{
	       "object": {
	           "amount": 55,
	           "created_at": 1451500772,
	           "currency": "usd",
	           "customer": 15455,
	           "date": 1451500771,
	           "fee": 0,
	           "gateway": null,
	           "gateway_id": null,
	           "id": 212047,
	           "invoice": 196539,
	           "metadata": [],
	           "method": "other",
	           "notes": null,
	           "parent_transaction": null,
	           "status": "succeeded",
	           "theme": null,
	           "type": "payment",
	           "pdf_url": "https:\/\/dundermifflin.invoiced.com\/payments\/59FHO96idoXFeiBDu1y5Zggg\/pdf",
	           "metadata": {}
	       }
	   }`

	equal, err := invdutil.JsonEqual(object, string(so.Data))

	if err != nil {
		t.Fatal(err)
	}

	if !equal {
		t.Fatal("Event object is incorrect")
	}

}

func TestUnMarshalEventObject2(t *testing.T) {
	s := `{
    "id": 5986597,
    "timestamp": 1583095640,
    "type": "invoice.created",
    "data": {
        "object": {
            "attempt_count": 0,
            "autopay": true,
            "balance": 2341,
            "chase": false,
            "closed": false,
            "created_at": 1583095640,
            "currency": "usd",
            "customer": {
                "address1": null,
                "address2": null,
                "attention_to": null,
                "autopay": true,
                "autopay_delay_days": -1,
                "avalara_entity_use_code": null,
                "avalara_exemption_number": null,
                "bill_to_parent": false,
                "chase": true,
                "chasing_cadence": null,
                "city": null,
                "consolidated": false,
                "country": "US",
                "created_at": 1582820561,
                "credit_hold": false,
                "credit_limit": null,
                "email": null,
                "id": 725981,
                "language": null,
                "name": "Professor Xavier",
                "next_chase_step": null,
                "notes": null,
                "number": "CUST-00001",
                "owner": null,
                "parent_customer": null,
                "payment_terms": null,
                "phone": null,
                "postal_code": null,
                "state": null,
                "tax_id": null,
                "taxable": true,
                "taxes": [],
                "type": "company",
                "object": "customer",
                "statement_pdf_url": "https:\/\/tesla198.sandbox.invoiced.com\/statements\/0gf8EYB8McG65lDB9l53vziu\/pdf",
                "sign_up_url": null,
                "payment_source": {
                    "bank_name": "Invoiced Test Bank",
                    "chargeable": true,
                    "country": "US",
                    "created_at": 1582831581,
                    "currency": "usd",
                    "failure_reason": null,
                    "gateway": "lawpay",
                    "gateway_customer": null,
                    "gateway_id": "QtZucfjRSsKWvuL01zuicQ",
                    "id": 696,
                    "last4": "6789",
                    "merchant_account": 324,
                    "receipt_email": null,
                    "routing_number": "110000000",
                    "verified": true,
                    "object": "bank_account"
                },
                "sign_up_page": null,
                "metadata": {}
            },
            "date": 1583095541,
            "draft": false,
            "due_date": null,
            "id": 2759436,
            "name": "Invoice",
            "needs_attention": false,
            "next_chase_on": null,
            "next_payment_attempt": 1583095541,
            "notes": null,
            "number": "INV-00001",
            "paid": false,
            "payment_plan": null,
            "payment_terms": "AutoPay",
            "purchase_order": null,
            "status": "not_sent",
            "subscription": null,
            "subtotal": 2341,
            "total": 2341,
            "object": "invoice",
            "url": "https:\/\/tesla198.sandbox.invoiced.com\/invoices\/qB0bqF4G7z3edBX097yylfuc",
            "pdf_url": "https:\/\/tesla198.sandbox.invoiced.com\/invoices\/qB0bqF4G7z3edBX097yylfuc\/pdf",
            "csv_url": "https:\/\/tesla198.sandbox.invoiced.com\/invoices\/qB0bqF4G7z3edBX097yylfuc\/csv",
            "payment_url": null,
            "ship_to": null,
            "payment_source": null,
            "metadata": {},
            "items": [
                {
                    "amount": 2341,
                    "catalog_item": null,
                    "created_at": 1583095640,
                    "description": "",
                    "discountable": true,
                    "id": 26354944,
                    "name": "test",
                    "quantity": 1,
                    "taxable": true,
                    "type": null,
                    "unit_cost": 2341,
                    "object": "line_item",
                    "metadata": {},
                    "discounts": [],
                    "taxes": []
                }
            ],
            "discounts": [],
            "taxes": [],
            "shipping": []
        }
    }
}`

	so := new(Event)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		panic(err)
	}

	ie, err := so.ParseInvoiceEvent()

	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(ie)

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(b),"725981") {
		t.Fatal("Customer id was not set")
	}

}