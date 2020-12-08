package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalSmsRequest(t *testing.T) {
	s := `{
  "to": [{"phone": "2345678900", "name": "Test McGee"}],
  "message": "test, hello, hi",
  "type": "open_item",
  "start": 1234567890,
  "end": 1234567891,
  "items": "past_due"
  }`

	so := new(TextRequest)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnMarshalSmsResponse(t *testing.T) {
	s := `[{
    "created_at": 1571086718,
    "id": "c05c9cae8c5799da1e5723a0fff355b3",
    "message": "Acme Inc.: You have a new statement https://acme.invoiced.com/statements/5X5g7Sb46KIR9IzxjjEjdnI9",
    "state": "sent",
    "to": "+12345678900"
}]`

	so := new(TextResponses)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
