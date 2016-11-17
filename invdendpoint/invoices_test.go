package invdendpoint

import (
	"encoding/json"
	"testing"
)

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
      "catalog_item": {
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
},
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
  "metadata": {}
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

	if so.Chase {
		t.Fatal("Chase is incorrect")
	}

	if so.CollectionMode != "manual" {
		t.Fatal("Collection Mode is incorrect")
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
