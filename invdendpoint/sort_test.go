package invdendpoint

import "testing"

func TestSort(t *testing.T) {
	s := NewSort()
	s.Set("name", DESC)
	s.Set("address", ASC)

	correctValue := "sort=address+ASC%2Cname+DESC"

	for i := 0; i < 1000; i++ {
		tmp := s.String()
		if tmp != correctValue {
			t.Fatal("Expected => ", correctValue, ", Got => ", tmp)
		}
	}

}

func TestEmptySort(t *testing.T) {

	f := NewSort()

	if f.String() != "" {
		t.Fatal("URL String is not equal", f.String())
	}

}
