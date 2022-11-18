package invoiced

import (
	"testing"
)

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

func TestMetadataFilter(t *testing.T) {
	f := NewMetadataFilter()
	err := f.Set("icp_number", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = f.Set("tps_report", "late")
	if err != nil {
		t.Fatal(err)
	}

	correctValue := "metadata%5Bicp_number%5D=1&metadata%5Btps_report%5D=late"

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

func TestFilterAndMetaFilter(t *testing.T) {
	f := NewMetadataFilter()
	err := f.Set("icp_number", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = f.Set("tps_report", "late")
	if err != nil {
		t.Fatal(err)
	}

	f1 := NewFilter()
	err = f1.Set("customer", 131)
	if err != nil {
		t.Fatal(err)
	}
	err = f1.Set("amount", "32311.23")
	if err != nil {
		t.Fatal(err)
	}

	s, err := AddFilterAndMetaFilterAndSort("",f1,f,nil)

	if err != nil {
		t.Fatal(err)
	}

	if s != "?filter%5Bamount%5D=32311.23&filter%5Bcustomer%5D=131&metadata%5Bicp_number%5D=1&metadata%5Btps_report%5D=late" {
		t.Fatal("metafiltering test failed")
	}

	s, err = AddFilterAndMetaFilterAndSort("",nil,f,nil)

	if err != nil {
		t.Fatal(err)
	}

	if s != "?metadata%5Bicp_number%5D=1&metadata%5Btps_report%5D=late" {
		t.Fatal("metafiltering test failed")
	}

	s, err = AddFilterAndMetaFilterAndSort("",f1,nil,nil)

	if err != nil {
		t.Fatal(err)
	}

	if s != "?filter%5Bamount%5D=32311.23&filter%5Bcustomer%5D=131" {
		t.Fatal("metafiltering test failed")
	}

	s, err = AddFilterAndMetaFilterAndSort("",f,nil,nil)

	if err == nil {
		t.Fatal("error should have been thrown")
	}




}