package invoiced

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalBalanceAdjustmentJSON(t *testing.T) {
	s := `{
  "amount": 50,
  "created_at": 1607550710,
  "currency": "usd",
  "customer": 78,
  "date": 1607550710,
  "id": 717,
  "notes": null,
  "object": "credit_balance_adjustment"
}`

	so := new(CreditBalanceAdjustment)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
