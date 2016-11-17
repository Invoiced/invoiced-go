package invdendpoint

import (
	"strings"
)

type Expand struct {
	params []string
}

func NewExpand() *Expand {
	f := new(Expand)
	f.params = make([]string, 0)

	return f
}

func (e *Expand) Set(key string) {
	e.params = append(e.params, key)

}

func (e *Expand) String() string {
	s := ""
	for _, values := range e.params {
		s += values + ","
	}

	s = strings.TrimRight(s, ",")

	return s

}
