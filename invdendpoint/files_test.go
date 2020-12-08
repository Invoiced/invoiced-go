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

	if so.Id != 13 {
		t.Fatal("File has incorrect id")
	}

	if so.Object != "file" {
		t.Fatal("File has incorrect object")
	}

	if so.Name != "logo-invoice.png" {
		t.Fatal("File has incorrect logo")
	}

	if so.Size != 6936 {
		t.Fatal("File has incorrect size")
	}

	if so.Type != "image/png" {
		t.Fatal("File has incorrect type")
	}

	if so.Url != "https://invoiced.com/img/logo-invoice.png" {
		t.Fatal("File url is incorrect")
	}

	if so.CreatedAt != 1464625855 {
		t.Fatal("CreatedAt is incorrect")
	}
}
