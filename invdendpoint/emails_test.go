package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalCustomerStatementRequest(t *testing.T) {
	s := `{"to": [{"name":"hello","email":"hello@invoiced.com"}],
  "bcc": "sales@invoiced.com",
  "subject": "Late Invoice",
  "message": "Right world"
  }`

	so := new(EmailRequest)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnMarshalCustomerStatementsResponse(t *testing.T) {
	s := `[
  {
    "id": 231,
    "state": "sent",
    "reject_reason": null,
    "email": "client@example.com",
    "template": "statement_email",
    "subject": "Statement from Dunder Mifflin, Inc.",
    "message": "Dear Client, we have attached your latest account statement. Thank you!",
    "opens": 0,
    "opens_detail": [],
    "clicks": 0,
    "clicks_detail": [],
    "created_at": 1436890047
  }
]`

	so := new(EmailResponses)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
