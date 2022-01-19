package invoiced

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalLetterRequest(t *testing.T) {
	s := `{"type": "open_item",
  "start": 1234567890,
  "end": 1234567891,
  "items": "past_due"
  }`

	so := new(SendStatementLetterRequest)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnMarshalLetterResponse(t *testing.T) {
	s := `{
  "created_at": 1570826337,
  "expected_delivery_date": 1571776737,
  "id": "2678c1e7e6dd1011ce13fb6b76db42df",
  "num_pages": 1,
  "state": "queued",
  "to": "Acme Inc.\n5301 Southwest Pkwy\nAustin, TX 78735"
}`

	so := new(Letter)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
