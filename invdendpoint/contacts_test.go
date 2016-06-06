package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalCustomerContactObject(t *testing.T) {
	s := `{
  "id": 10403,
  "name": "Nancy Talty",
  "email": "nancy.talty@example.com",
  "primary": true,
  "address1": null,
  "address2": null,
  "city": null,
  "state": null,
  "postal_code": null,
  "country": null,
  "created_at": 1463510889
}`

	so := new(Contact)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
