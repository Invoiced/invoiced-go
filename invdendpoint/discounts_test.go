package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalDiscountObject(t *testing.T) {
	s := `{
  "id": 20553,
  "amount": 5,
  "coupon": null,
  "expires": null
}`

	so := new(Discount)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 20553 {
		t.Fatal("Discount id has incorrect id")
	}

	if so.Amount != 5 {
		t.Fatal("Amount is incorrect")
	}

}
