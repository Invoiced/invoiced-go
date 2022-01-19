package invoiced

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalCustomerStatementRequest(t *testing.T) {
	s := `{"to": [{"name":"hello","email":"hello@invoiced.com"}],
  "bcc": "sales@invoiced.com",
  "subject": "Late InvoiceClient",
  "message": "Right world"
  }`

	so := new(SendEmailRequest)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
