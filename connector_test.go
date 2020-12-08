package invdapi

import (
	"testing"
)

import "github.com/Invoiced/invoiced-go/invdendpoint"

func TestParseRawRelation(t *testing.T) {
	s := "          rel=\"       self     \"                           "
	parsed := parseRawRelation(s)

	if parsed != "self" {
		t.Fatal("Error: Parsing of relation is not self")
	}
}

func TestParseRawURL(t *testing.T) {
	s := "     <  https://api.invoiced.com/invoices?page=1  > "
	parsed := parseRawURL(s)

	if parsed != "https://api.invoiced.com/invoices?page=1" {
		t.Fatal("Error: Parsing of URL", " parsed => ", parsed)
	}
}

func TestAddFilterSortToEndPointWithBothValues(t *testing.T) {
	f := invdendpoint.NewFilter()
	err := f.Set("id", 121123)
	if err != nil {
		t.Fatal(err)
	}

	err = f.Set("address", 121123)

	if err != nil {
		t.Fatal(err)
	}

	s := invdendpoint.NewSort()
	s.Set("name", invdendpoint.ASC)
	s.Set("age", invdendpoint.DESC)

	endPoint := "https://www.do.com"

	value := addFilterSortToEndPoint(endPoint, f, s)

	correctValue := "https://www.do.com?filter%5Baddress%5D=121123&filter%5Bid%5D=121123&sort=age+DESC%2Cname+ASC"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}

	// endpoint2 := "https://www.do.com?"
}

func TestAddFilterSortToEndPointWithOnlySort(t *testing.T) {
	s := invdendpoint.NewSort()
	s.Set("name", invdendpoint.ASC)
	s.Set("age", invdendpoint.DESC)

	endPoint := "https://www.do.com"

	value := addFilterSortToEndPoint(endPoint, nil, s)

	correctValue := "https://www.do.com?sort=age+DESC%2Cname+ASC"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}
}

func TestAddFilterSortToEndPointWithOnlyFilter(t *testing.T) {
	f := invdendpoint.NewFilter()

	err := f.Set("id", 121123)
	if err != nil {
		t.Fatal(err)
	}

	err = f.Set("address", 121123)

	if err != nil {
		t.Fatal(err)
	}

	endPoint := "https://www.do.com"

	value := addFilterSortToEndPoint(endPoint, f, nil)

	correctValue := "https://www.do.com?filter%5Baddress%5D=121123&filter%5Bid%5D=121123"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}
}

func TestAddFilterSortToEndPointWithNothing(t *testing.T) {
	endPoint := "https://www.do.com"

	value := addFilterSortToEndPoint(endPoint, nil, nil)

	correctValue := "https://www.do.com"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}
}

func TestMakeEndPointSingular(t *testing.T) {
	endpoint := "https://www.do.com/customer"

	singularEndPoint := makeEndPointSingular(endpoint, 5)

	correctSingularEndPoint := "https://www.do.com/customer/5"

	if singularEndPoint != correctSingularEndPoint {
		t.Fatal("Expect =>", singularEndPoint, " Got =>", correctSingularEndPoint)
	}
}
