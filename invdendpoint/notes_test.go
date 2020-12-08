package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalNoteObject(t *testing.T) {
	s := `{
  "created_at": 1571338027,
  "customer": 15444,
  "id": 501,
  "notes": "Customer called 10/1, clarified account terms.",
  "object": "note",
  "user": {
    "created_at": 1563810757,
    "email": "invoiced@example.com",
    "first_name": "John",
    "id": 1946,
    "last_name": "Smith",
    "registered": true,
    "two_factor_enabled": true
  }
}`

	so := new(Note)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
