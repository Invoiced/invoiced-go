package invdendpoint

import "testing"

func TestFilter(t *testing.T) {

	f := NewFilter()
	err := f.Set("customer", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = f.Set("amount", "32311.23")
	if err != nil {
		t.Fatal(err)
	}
	err = f.Set("day", "tuesday")
	if err != nil {
		t.Fatal(err)
	}

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
