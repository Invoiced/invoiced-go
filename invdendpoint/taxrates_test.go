package invdendpoint

import (
	"encoding/json"
	"testing"
)

func testUnmarshalTaxRateObject(t *testing.T) {
	s := `{
  "created_at": 1477418268,
  "currency": null,
  "id": "vat",
  "inclusive": false,
  "is_percent": true,
  "metadata": {},
  "name": "VAT",
  "object": "tax_rate",
  "value": 5
}`

	so := new(CatalogItem)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
