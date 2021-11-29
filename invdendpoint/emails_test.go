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


