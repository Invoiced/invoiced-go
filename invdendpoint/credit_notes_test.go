package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalCreditNoteObject(t *testing.T) {
	s := `{
  "balance": 51.15,
  "closed": false,
  "created_at": 1415229884,
  "currency": "usd",
  "customer": 15444,
  "date": 1416290400,
  "discounts": [],
  "draft": false,
  "id": 2048,
  "invoice": 46225,
  "items": [
    {
      "amount": 45,
      "catalog_item": null,
      "description": null,
      "discountable": true,
      "discounts": [],
      "id": 7,
      "metadata": {},
      "name": "Copy Paper, Case",
      "object": "line_item",
      "quantity": 1,
      "taxable": true,
      "taxes": [],
      "type": "product",
      "unit_cost": 45
    },
    {
      "amount": 10,
      "catalog_item": "delivery",
      "description": null,
      "discountable": true,
      "discounts": [],
      "id": 8,
      "metadata": {},
      "name": "Delivery",
      "object": "line_item",
      "quantity": 1,
      "taxable": true,
      "taxes": [],
      "type": "service",
      "unit_cost": 10
    }
  ],
  "metadata": {},
  "name": null,
  "notes": null,
  "number": "CN-0016",
  "object": "credit_note",
  "paid": false,
  "pdf_url": "https://dundermifflin.invoiced.com/credit_notes/IZmXbVOPyvfD3GPBmyd6FwXY/pdf",
  "status": "open",
  "subtotal": 55,
  "taxes": [
    {
      "amount": 3.85,
      "id": 20554,
      "object": "tax",
      "tax_rate": null
    }
  ],
  "total": 51.15,
  "url": "https://dundermifflin.invoiced.com/credit_notes/IZmXbVOPyvfD3GPBmyd6FwXY"
}
`

	so := new(CreditNote)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
