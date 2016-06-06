package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalTransactionObject(t *testing.T) {
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
    "parent_transaction": null,
    "pdf_url": "https://dundermifflin.invoiced.com/payments/IZmXbVOPyvfD3GPBmyd6FwXY/pdf",
    "created_at": 1415228628,
    "metadata": {}
}`

	so := new(Transaction)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
