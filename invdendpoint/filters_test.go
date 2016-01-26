package invdendpoint

import "testing"

func TestFilter(t *testing.T) {

	f := NewFilter()
	f.Set("customer", 1)
	f.Set("amount", "32311.23")
	f.Set("day", "tuesday")

	correctValue := "filter%5Bamount%5D=32311.23&filter%5Bcustomer%5D=1&filter%5Bday%5D=tuesday"

	for i := 0; i < 1000; i++ {
		tmp := f.String()
		if tmp != correctValue {
			t.Fatal("Expected => ", correctValue, ", Got => ", tmp)
		}
	}

}

func TestEmptyFilter(t *testing.T) {

	f := NewFilter()

	if f.String() != "" {
		t.Fatal("URL String is not equal")
	}

}
