package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalSubscriptionObject(t *testing.T) {
	s := `{
    "id": 595,
    "customer": 15444,
    "plan": "starter",
    "cycles": null,
    "quantity": 1,
    "start_date": 1420391704,
    "period_start": 1446657304,
    "period_end": 1449249304,
    "status": "active",
    "addons": [
        {
            "id": 3,
            "catalog_item": "ipad-license",
            "quantity": 11,
            "created_at": 1420391704
        }
    ],
    "discounts": [],
    "taxes": [],
    "url": "https://dundermifflin.invoiced.com/subscriptions/o2mAd2wWVfYy16XZto7xHwXX",
    "created_at": 1420391704,
    "metadata": {}
}`

	so := new(Subscription)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
