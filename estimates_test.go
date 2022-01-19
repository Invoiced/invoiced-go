package invoiced

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalEstimateObject(t *testing.T) {
	s := `{
    "approval": {
        "id": 250,
        "initials": "RR",
        "ip": "67.79.55.186",
        "timestamp": 1595209318,
        "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Safari/605.1.15"
    },
    "approved": "RR",
    "closed": true,
    "created_at": 1595209305,
    "currency": "usd",
    "customer": 1094399,
    "date": 1595209239,
    "deposit": 0,
    "deposit_paid": false,
    "discounts": [],
    "draft": false,
    "expiration_date": null,
    "id": 14101,
    "invoice": null,
    "metadata": {},
    "name": "Estimate",
    "notes": null,
    "number": "SDLL-00001",
    "object": "estimate",
    "payment_terms": "NET 14",
    "pdf_url": "https://tesla.sandbox.invoiced.com/estimates/krtvjafiHpGkcRGgVuFZEqAv/pdf",
    "purchase_order": null,
    "ship_to": null,
    "shipping": [],
    "status": "approved",
    "subtotal": 200,
    "taxes": [],
    "total": 200,
    "url": "https://tesla.sandbox.invoiced.com/estimates/krtvjafiHpGkcRGgVuFZEqAv"
}`

	so := new(Estimate)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
