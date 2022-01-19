package invoiced

import (
	"testing"
)

func TestExclude(t *testing.T) {
	e := NewExclude()
	e.Set("items.catalog_item")
	e.Set("customer")
	if e.String() != "items.catalog_item,customer" {
		t.Fatal("Expanded values do not match")
	}
}

func TestEmptyExclude(t *testing.T) {
	e := NewExclude()

	if e.String() != "" {
		t.Fatal("Expand should be the empty string")
	}
}
