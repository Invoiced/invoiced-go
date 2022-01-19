package invoiced

import "testing"

func TestNewAPIError(t *testing.T) {
	error := NewAPIError("", "", "")

	if error == nil {
		t.Fatal("Error did not initialize")
	}
}
