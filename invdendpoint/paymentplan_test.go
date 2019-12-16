package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalPaymentPlan(t *testing.T) {
	s := `{
  "approval": {
    "id": 12,
    "ip": "217.15.151.36",
    "timestamp": 1479827803,
    "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:50.0) Gecko/20100101 Firefox/50.0"
  },
  "created_at": 1479827791,
  "id": 6,
  "installments": [
    {
      "amount": 500,
      "balance": 500,
      "date": 1480572000,
      "id": 23
    },
    {
      "amount": 500,
      "balance": 500,
      "date": 1481176800,
      "id": 24
    },
    {
      "amount": 500,
      "balance": 500,
      "date": 1481781600,
      "id": 25
    },
    {
      "amount": 500,
      "balance": 500,
      "date": 1482386400,
      "id": 26
    }
  ],
  "object": "payment_plan",
  "status": "active"
}`

	so := new(PaymentPlan)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 6 {
		t.Fatal("Id does not match")
	}

	if so.Installments[0].Amount != 500 {
		t.Fatal("Id does not match")
	}

	if so.Approval.Ip !="217.15.151.36" {
		t.Fatal("ip address is incorrect")
	}

}
