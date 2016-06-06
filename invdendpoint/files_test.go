package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalFileObject(t *testing.T) {
	s := `{
  "id": 13,
  "object": "file",
  "name": "logo-invoice.png",
  "size": 6936,
  "type": "image/png",
  "url": "https://invoiced.com/img/logo-invoice.png",
  "created_at": 1464625855
}`

	so := new(File)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
