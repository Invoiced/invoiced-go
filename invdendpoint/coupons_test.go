package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalCouponObject(t *testing.T) {
	s := `{
	"created_at": 1565619435,
	"currency": null,
	"duration": 0,
	"exclusive": false,
	"expiration_date": null,
	"id": "example",
	"is_percent": true,
	"max_redemptions": 0,
	"metadata": {},
	"name": "Example",
	"object": "coupon",
	"value": 1
}`

	so := new(Coupon)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
