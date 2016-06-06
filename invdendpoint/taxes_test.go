package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalTaxObject(t *testing.T) {
	s := `{
  "id": 20554,
  "amount": 3.85,
  "tax_rate": null
}`

	so := new(Tax)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
