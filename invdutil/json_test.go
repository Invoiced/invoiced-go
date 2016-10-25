package invdutil

import (
	"testing"
)

func TestJsonEqual(t *testing.T) {

	j1 := "{\"customer\":15444,\"draft\":false,\"closed\":false,\"paid\":false,\"sent\":false,\"chase\":false,\"payment_terms\":\"NET 14\",\"items\":[{\"name\":\"Copy paper, Case\",\"quantity\":1.0,\"unit_cost\":45.0,\"discountable\":false,\"taxable\":false},{\"catalog_item\":\"delivery\",\"quantity\":1.0,\"discountable\":false,\"taxable\":false}],\"taxes\":[{\"amount\":3.85}]}"

	j2 := "{\"paid\":false,                        \"customer\":15444,\"taxes\":[                {\"amount\":3.85}],\"draft\":false,\"closed\":false,\"sent\":false,\"chase\":false,\"payment_terms\":\"NET 14\",\"items\":[{\"name\":\"Copy paper, Case\",\"quantity\":1.0,\"taxable\":false,\"unit_cost\":45.0,\"discountable\":false},{\"catalog_item\":\"delivery\",\"quantity\":1.0,\"discountable\":false,\"taxable\":false}]}"

	equal, err := JsonEqual(j1, j2)

	if err != nil {
		t.Fatal(err)
	}

	if !equal {
		t.Fatal("Json is not equal")

	}

}
