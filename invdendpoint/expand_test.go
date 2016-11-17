package invdendpoint

import (
	"testing"
)

func TestExpand(t *testing.T) {

	e := NewExpand()
	e.Set("items.catalog_item")
	e.Set("customer")
	if e.String() != "items.catalog_item,customer" {
		t.Fatal("Expanded values do not match")
	}

}

func TestEmptyExpand(t *testing.T) {

	e := NewExpand()

	if e.String() != "" {
		t.Fatal("Expand should be the empty string")
	}

}
