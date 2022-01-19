package invoiced

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalPaymentObject(t *testing.T) {
	s := `{
    "id": 20939,
    "customer": 15460,
    "invoice": 44648,
    "date": 1410843600,
    "type": "payment",
    "method": "check",
    "status": "succeeded",
    "gateway": null,
    "gateway_id": null,
    "currency": "usd",
    "amount": 800,
    "fee": 0,
    "notes": null,
    "parent_payment": null,
    "pdf_url": "https://dundermifflin.invoiced.com/payments/IZmXbVOPyvfD3GPBmyd6FwXY/pdf",
    "created_at": 1415228628,
    "metadata": {}
}`

	so := new(Payment)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 20939 {
		t.Fatal("Client has incorrect periodstart")
	}

	if so.Customer != 15460 {
		t.Fatal("Client has incorrect periodstart")
	}

	if so.Date != 1410843600 {
		t.Fatal("Client has incorrect invoice")
	}

	if so.Currency != "usd" {
		t.Fatal("Client has incorrect currency")
	}

	if so.Amount != 800 {
		t.Fatal("Client has incorrect amount")
	}

	if so.PdfUrl != "https://dundermifflin.invoiced.com/payments/IZmXbVOPyvfD3GPBmyd6FwXY/pdf" {
		t.Fatal("Client has incorrect pdf")
	}

	if so.CreatedAt != 1415228628 {
		t.Fatal("Client has incorrect createdAt")
	}
}
