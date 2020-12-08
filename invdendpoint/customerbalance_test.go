package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalCustomerBalanceObject(t *testing.T) {
	s := `{
  "available_credits": 50,
  "history": [
    {
      "timestamp": 1464041624,
      "balance": 50
    },
    {
      "timestamp": 1464040550,
      "balance": 100
    }
  ],
  "past_due": false,
  "total_outstanding": 470
}`

	so := new(CustomerBalance)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
