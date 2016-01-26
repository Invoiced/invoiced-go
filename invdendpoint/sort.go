package invdendpoint

import "net/url"
import "strings"
import "sort"

type SortOrder int

const (
	ASC SortOrder = iota
	DESC
)

func (s SortOrder) String() string {
	if s == ASC {
		return "ASC"
	} else if s == DESC {
		return "DESC"
	}

	return ""

}

type Sort struct {
	orders map[string]SortOrder
}

func NewSort() *Sort {
	s := new(Sort)
	s.orders = make(map[string]SortOrder)

	return s
}

func (s *Sort) Set(column string, order SortOrder) {
	s.orders[column] = order
}

func (s *Sort) String() string {

	uValues := url.Values{}
	orderString := ""
	orderedKeys := []string{}

	for column, _ := range s.orders {
		orderedKeys = append(orderedKeys, column)
	}

	sort.Strings(orderedKeys)

	for _, column := range orderedKeys {

		orderString += column + " " + s.orders[column].String() + ","
	}
	orderString = strings.TrimRight(orderString, ",")
	if orderString == "" {
		return ""
	}

	uValues.Set("sort", orderString)

	return uValues.Encode()

}
