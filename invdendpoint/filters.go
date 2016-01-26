package invdendpoint

import "net/url"
import "strconv"
import "errors"
import "sort"

type Filter struct {
	params map[string]string
}

func NewFilter() *Filter {
	f := new(Filter)
	f.params = make(map[string]string)

	return f
}

//Can only set Numberic Types and Strings
func (f *Filter) Set(key string, value interface{}) error {
	switch v := value.(type) {
	case string:
		f.params[key] = v
	case int:
		f.params[key] = strconv.Itoa(v)
	case int32:
		f.params[key] = strconv.FormatInt(int64(v), 10)
	case int64:
		f.params[key] = strconv.FormatInt(v, 10)
	case float32:
		f.params[key] = strconv.FormatFloat(float64(v), 'f', 2, 64)
	case float64:
		f.params[key] = strconv.FormatFloat(float64(v), 'f', 2, 64)
	default:
		return errors.New("Filter can only accept numeric (int32,int64,float32,float64) or string values")
	}

	return nil

}

func (f *Filter) Get(key string) string {
	v, ok := f.params[key]

	if !ok {
		return ""
	}

	return v

}

func (f *Filter) String() string {
	uValues := url.Values{}
	orderedKeys := []string{}

	for key, _ := range f.params {
		orderedKeys = append(orderedKeys, key)
	}

	sort.Strings(orderedKeys)

	for _, key := range orderedKeys {
		mapkey := "filter[" + key + "]"
		uValues.Set(mapkey, f.params[key])
	}

	return uValues.Encode()

}
