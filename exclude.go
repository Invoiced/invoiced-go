package invoiced

import (
	"strings"
)

type Exclude struct {
	params []string
}

func NewExclude() *Exclude {
	f := new(Exclude)
	f.params = make([]string, 0)

	return f
}

func (e *Exclude) Set(key string) {
	e.params = append(e.params, key)
}

func (e *Exclude) String() string {
	s := ""
	for _, values := range e.params {
		s += values + ","
	}

	s = strings.TrimRight(s, ",")

	return s
}
