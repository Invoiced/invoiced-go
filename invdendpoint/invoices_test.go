package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalObject2(t *testing.T) {
	s := `{
    "attempt_count": 0,
    "autopay": false,
    "balance": 351.81,
    "chase": false,
    "closed": false,
    "created_at": 1565892968,
    "csv_url": "https://paragtestcorp.sandbox.invoiced.com/invoices/Do6BicUQ6iPIv3O1waf7IFti/csv",
    "currency": "usd",
    "customer": 470508,
    "date": 1532674800,
    "discounts": [],
    "draft": false,
    "due_date": 1535266800,
    "id": 2194174,
    "items": [
        {
            "amount": 325,
            "catalog_item": null,
            "created_at": 1565892968,
            "description": "Merlin 4412D: The most powerful features avail in a 12 button display phone",
            "discountable": true,
            "discounts": [],
            "id": 21954084,
            "metadata": {
                "netsuite_quantity": "1",
                "netsuite_rate": "325.00"
            },
            "name": "ACC00004",
            "object": "line_item",
            "quantity": 1,
            "taxable": true,
            "taxes": [],
            "type": null,
            "unit_cost": 325
        }
    ],
    "metadata": {
        "netsuite_invoice_id": "8840",
        "subsidiary": "Honeycomb Holdings Inc."
    },
    "name": "Invoice",
    "needs_attention": false,
    "next_chase_on": null,
    "next_payment_attempt": null,
    "notes": null,
    "number": "SOINV10000010",
    "object": "invoice",
    "paid": false,
    "payment_plan": null,
    "payment_source": null,
    "payment_terms": "Net 30",
    "payment_url": "https://paragtestcorp.sandbox.invoiced.com/invoices/Do6BicUQ6iPIv3O1waf7IFti/payment",
    "pdf_url": "https://paragtestcorp.sandbox.invoiced.com/invoices/Do6BicUQ6iPIv3O1waf7IFti/pdf",
    "purchase_order": null,
    "ship_to": null,
    "shipping": [],
    "status": "past_due",
    "subscription": null,
    "subtotal": 325,
    "taxes": [
        {
            "amount": 26.81,
            "id": 2020973,
            "object": "tax",
            "tax_rate": null
        }
    ],
    "total": 351.81,
    "url": "https://paragtestcorp.sandbox.invoiced.com/invoices/Do6BicUQ6iPIv3O1waf7IFti"
}`

	so := new(Invoice)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnMarshalInvoiceObject(t *testing.T) {
	s := `{
  "id": 46225,
  "customer": 15444,
  "name": null,
  "currency": "usd",
  "draft": false,
  "closed": false,
  "paid": false,
  "status": "not_sent",
  "chase": false,
  "next_chase_on": null,
  "collection_mode": "manual",
  "attempt_count": 0,
  "next_payment_attempt": null,
  "subscription": null,
  "number": "INV-0016",
  "date": 1416290400,
  "due_date": 1417500000,
  "payment_terms": "NET 14",
  "items": [
    {
      "id": 7,
      "catalog_item": null,
      "type": "product",
      "name": "Copy Paper, Case",
      "description": null,
      "quantity": 1,
      "unit_cost": 45,
      "amount": 45,
      "discountable": true,
      "discounts": [],
      "taxable": true,
      "taxes": [],
      "metadata": {}
    },
    {
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
    }
  ],
  "notes": null,
  "subtotal": 55,
  "discounts": [],
  "taxes": [
    {
      "id": 20554,
      "amount": 3.85,
      "tax_rate": null
    }
  ],
  "total": 51.15,
  "balance": 51.15,
  "tags": [],
  "url": "https://dundermifflin.invoiced.com/invoices/IZmXbVOPyvfD3GPBmyd6FwXY",
  "payment_url": "https://dundermifflin.invoiced.com/invoices/IZmXbVOPyvfD3GPBmyd6FwXY/payment",
  "pdf_url": "https://dundermifflin.invoiced.com/invoices/IZmXbVOPyvfD3GPBmyd6FwXY/pdf",
  "created_at": 1415229884,
  "metadata": {},
  "calculate_taxes": true
}`

	so := new(Invoice)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 46225 {
		t.Fatal("Id is incorrect")
	}

	if so.Customer != 15444 {
		t.Fatal("Customer is incorrect")
	}

	if so.Currency != "usd" {
		t.Fatal("Number is incorrect")
	}

	if so.Draft {
		t.Fatal("Draft is incorrect")
	}

	if so.Closed {
		t.Fatal("Closed is incorrect")
	}

	if so.Paid {
		t.Fatal("Paid is incorrect")
	}

	if so.Date != 1416290400 {
		t.Fatal("Date is incorrect")
	}

	if so.DueDate != 1417500000 {
		t.Fatal("Date is incorrect")
	}

	if so.PaymentTerms != "NET 14" {
		t.Fatal("Payment terms are incorrect")
	}

	if so.Items[0].Id != 7 {
		t.Fatal("Item 0 has incorrect id")
	}

	if so.Items[0].Type != "product" {
		t.Fatal("Item 0 has incorrect type")
	}

	if so.Items[0].Name != "Copy Paper, Case" {
		t.Fatal("Item 0 has incorrect name")
	}

	if so.Items[0].Quantity != 1.0 {
		t.Fatal("Item 0 has incorrect quantity")
	}

	if so.Items[0].UnitCost != 45 {
		t.Fatal("Item 0 has incorrect unit cost")
	}

	if so.Items[0].Amount != 45 {
		t.Fatal("Item 0 has incorrect amount")
	}

	if !so.Items[0].Taxable {
		t.Fatal("Item 0 should be taxable")
	}

	if so.Items[1].Id != 8 {
		t.Fatal("Item 1 has incorrect id")
	}

	if so.Items[1].Type != "service" {
		t.Fatal("Item 1 has incorrect type")
	}

	if so.Items[1].Name != "Delivery" {
		t.Fatal("Item 0 has incorrect name")
	}

	if so.Items[1].Quantity != 1.0 {
		t.Fatal("Item 1 has incorrect quantity")
	}

	if so.Items[1].UnitCost != 10 {
		t.Fatal("Item 1 has incorrect unit cost")
	}

	if so.Items[1].Amount != 10 {
		t.Fatal("Item 1 has incorrect amount")
	}

	if !so.Items[1].Taxable {
		t.Fatal("Item 1 has incorrect taxable")
	}

	if so.Subtotal != 55 {
		t.Fatal("Subtotal is incorrect")
	}

	if so.Taxes[0].Id != 20554 {
		t.Fatal("Tax id is incorrect")
	}

	if so.Taxes[0].Id != 20554 {
		t.Fatal("Tax id is incorrect")
	}

	if so.Taxes[0].Amount != 3.85 {
		t.Fatal("Tax amount is incorrect")
	}

	if so.Total != 51.15 {
		t.Fatal("Total is incorrect")
	}

	if so.Balance != 51.15 {
		t.Fatal("Total is incorrect")
	}

	if so.Url != "https://dundermifflin.invoiced.com/invoices/IZmXbVOPyvfD3GPBmyd6FwXY" {
		t.Fatal("Url is incorrect")
	}

	if so.PaymentUrl != "https://dundermifflin.invoiced.com/invoices/IZmXbVOPyvfD3GPBmyd6FwXY/payment" {
		t.Fatal("Payment Url is incorrect")
	}

	if so.PdfUrl != "https://dundermifflin.invoiced.com/invoices/IZmXbVOPyvfD3GPBmyd6FwXY/pdf" {
		t.Fatal("Pdf Url is incorrect")
	}

	if so.CreatedAt != 1415229884 {
		t.Fatal("CreatedAt is incorrect")
	}
}

func TestTotalTaxAmount(t *testing.T) {
	s := `{
      "attempt_count": 0,
      "autopay": false,
      "balance": 457.32,
      "chase": false,
      "closed": false,
      "created_at": 1583684877,
      "currency": "usd",
      "customer": 757661,
      "date": 1583684797,
      "draft": false,
      "due_date": null,
      "id": 2818436,
      "name": "Invoice",
      "needs_attention": false,
      "next_chase_on": null,
      "next_payment_attempt": null,
      "notes": null,
      "number": "INV-00004",
      "paid": false,
      "payment_plan": null,
      "payment_terms": null,
      "purchase_order": null,
      "status": "not_sent",
      "subscription": null,
      "subtotal": 412,
      "total": 457.32,
      "object": "invoice",
      "url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww",
      "pdf_url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww/pdf",
      "csv_url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww/csv",
      "payment_url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww/payment",
      "ship_to": null,
      "payment_source": null,
      "items": [
        {
          "amount": 412,
          "catalog_item": null,
          "created_at": 1583684877,
          "description": "",
          "discountable": true,
          "id": 26684187,
          "name": "Service II",
          "quantity": 1,
          "taxable": true,
          "type": null,
          "unit_cost": 412,
          "object": "line_item",
          "discounts": [],
          "taxes": [
            {
              "amount": 16.48,
              "id": 2486098,
              "object": "tax",
              "tax_rate": {
                "created_at": 1583684856,
                "currency": null,
                "id": "linetax1",
                "inclusive": false,
                "is_percent": true,
                "name": "linetax1",
                "value": 4,
                "object": "tax_rate"
              }
            }
          ]
        }
      ],
      "discounts": [],
      "taxes": [
        {
          "amount": 28.84,
          "id": 2486099,
          "object": "tax",
          "tax_rate": {
            "created_at": 1583684874,
            "currency": null,
            "id": "state_tax",
            "inclusive": false,
            "is_percent": true,
            "name": "state_tax",
            "value": 7,
            "object": "tax_rate"
          }
        }
      ],
      "shipping": []
    }`

	so := new(Invoice)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}

	totalTax := so.TotalTaxAmount()

	if totalTax != 45.32 {
		t.Fatal("Tax amount does not match")
	}
}

func TestTotalDiscountAmount(t *testing.T) {
	s := `{
    "attempt_count": 0,
    "autopay": false,
    "balance": 435.22,
    "chase": false,
    "closed": false,
    "created_at": 1583684877,
    "csv_url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww/csv",
    "currency": "usd",
    "customer":78687565,
    "date": 1583684797,
    "discounts": [
        {
            "amount": 12.11,
            "coupon": {
                "created_at": 1583685424,
                "currency": null,
                "duration": 0,
                "exclusive": false,
                "expiration_date": null,
                "id": "discount2",
                "is_percent": true,
                "max_redemptions": 0,
                "name": "discount2",
                "object": "coupon",
                "value": 3
            },
            "expires": null,
            "id": 2486561,
            "object": "discount"
        }
    ],
    "draft": false,
    "due_date": null,
    "id": 2818436,
    "items": [
        {
            "amount": 412,
            "catalog_item": null,
            "created_at": 1583684877,
            "description": "",
            "discountable": true,
            "discounts": [
                {
                    "amount": 8.24,
                    "coupon": {
                        "created_at": 1583685415,
                        "currency": null,
                        "duration": 0,
                        "exclusive": false,
                        "expiration_date": null,
                        "id": "discount1",
                        "is_percent": true,
                        "max_redemptions": 0,
                        "name": "discount1",
                        "object": "coupon",
                        "value": 2
                    },
                    "expires": null,
                    "id": 2486560,
                    "object": "discount"
                }
            ],
            "id": 26684187,
            "name": "Service II",
            "object": "line_item",
            "quantity": 1,
            "taxable": true,
            "taxes": [
                {
                    "amount": 16.15,
                    "id": 2486098,
                    "object": "tax",
                    "tax_rate": {
                        "created_at": 1583684856,
                        "currency": null,
                        "id": "linetax1",
                        "inclusive": false,
                        "is_percent": true,
                        "name": "linetax1",
                        "object": "tax_rate",
                        "value": 4
                    }
                }
            ],
            "type": null,
            "unit_cost": 412
        }
    ],
    "name": "Invoice",
    "needs_attention": false,
    "next_chase_on": null,
    "next_payment_attempt": null,
    "notes": null,
    "number": "INV-00004",
    "object": "invoice",
    "paid": false,
    "payment_plan": null,
    "payment_source": null,
    "payment_terms": null,
    "payment_url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww/payment",
    "pdf_url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww/pdf",
    "purchase_order": null,
    "ship_to": null,
    "shipping": [],
    "status": "not_sent",
    "subscription": null,
    "subtotal": 412,
    "taxes": [
        {
            "amount": 27.42,
            "id": 2486099,
            "object": "tax",
            "tax_rate": {
                "created_at": 1583684874,
                "currency": null,
                "id": "state_tax",
                "inclusive": false,
                "is_percent": true,
                "name": "state_tax",
                "object": "tax_rate",
                "value": 7
            }
        }
    ],
    "total": 435.22,
    "url": "https://tesla198.sandbox.invoiced.com/invoices/sBpN7NmRSPTYZ472rzynY6Ww"
}`

	so := new(Invoice)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}

	totalDiscount := so.TotalDiscountAmount()

	if totalDiscount != 20.35 {
		t.Fatal("Total discount amount does not match")
	}
}
