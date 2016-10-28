package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalContactObject(t *testing.T) {
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

	if so.Id != 10403 {
		t.Fatal("Id is incorrect")
	}

	if so.Name != "Nancy Talty" {
		t.Fatal("Name is incorrect")
	}

	if !so.Primary {
		t.Fatal("Primary should be true")
	}

	if so.Email != "nancy.talty@example.com" {
		t.Fatal("Primary should be true")
	}

	if so.CreatedAt != 1463510889 {
		t.Fatal("Created At is incorrect")
	}

}
