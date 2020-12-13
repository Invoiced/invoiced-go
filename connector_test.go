package invdapi

import (
	"strconv"
	"testing"
)

import "github.com/Invoiced/invoiced-go/invdendpoint"

func TestParseRawRelation(t *testing.T) {
	s := "          rel=\"       self     \"                           "
	parsed := parseRelValue(s)

	if parsed != "self" {
		t.Fatal("Error: Parsing of relation is not self")
	}
}

func TestParseRawURL(t *testing.T) {
	s := "     <  https://api.invoiced.com/invoices?page=1  > "
	parsed := parseLinkUrl(s)

	if parsed != "https://api.invoiced.com/invoices?page=1" {
		t.Fatal("Error: Parsing of URL", " parsed => ", parsed)
	}
}

func TestAddFilterSortToEndpointWithBothValues(t *testing.T) {
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

	endpoint := "https://www.do.com"

	value := addFilterAndSort(endpoint, f, s)

	correctValue := "https://www.do.com?filter%5Baddress%5D=121123&filter%5Bid%5D=121123&sort=age+DESC%2Cname+ASC"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}

	// endpoint2 := "https://www.do.com?"
}

func TestAddFilterSortToEndpointWithOnlySort(t *testing.T) {
	s := invdendpoint.NewSort()
	s.Set("name", invdendpoint.ASC)
	s.Set("age", invdendpoint.DESC)

	endpoint := "https://www.do.com"

	value := addFilterAndSort(endpoint, nil, s)

	correctValue := "https://www.do.com?sort=age+DESC%2Cname+ASC"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}
}

func TestAddFilterSortToEndpointWithOnlyFilter(t *testing.T) {
	f := invdendpoint.NewFilter()

	err := f.Set("id", 121123)
	if err != nil {
		t.Fatal(err)
	}

	err = f.Set("address", 121123)

	if err != nil {
		t.Fatal(err)
	}

	endpoint := "https://www.do.com"

	value := addFilterAndSort(endpoint, f, nil)

	correctValue := "https://www.do.com?filter%5Baddress%5D=121123&filter%5Bid%5D=121123"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}
}

func TestAddFilterSortToEndpointWithNothing(t *testing.T) {
	endpoint := "https://www.do.com"

	value := addFilterAndSort(endpoint, nil, nil)

	correctValue := "https://www.do.com"

	if value != correctValue {
		t.Fatal("Error: resulting URL is incorrect it should be ", correctValue, " but instead got ", value)
	}
}

func TestMakeEndpointSingular(t *testing.T) {
	endpoint := "https://www.do.com/customer"

	singularEndpoint := endpoint + "/" + strconv.FormatInt(5, 10)

	correctSingularEndpoint := "https://www.do.com/customer/5"

	if singularEndpoint != correctSingularEndpoint {
		t.Fatal("Expect =>", singularEndpoint, " Got =>", correctSingularEndpoint)
	}
}
