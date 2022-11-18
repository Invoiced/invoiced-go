package invoiced

import (
	"errors"
	"net/url"
	"sort"
	"strconv"
)

type Filter struct {
	params   map[string]string
	metadata bool
}

func NewFilter() *Filter {
	f := new(Filter)
	f.params = make(map[string]string)
	f.metadata = false
	return f
}

func NewMetadataFilter() *Filter {
	f := new(Filter)
	f.params = make(map[string]string)
	f.metadata = true
	return f
}

// Can only set Numeric Types and Strings
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

	for key := range f.params {
		orderedKeys = append(orderedKeys, key)
	}

	sort.Strings(orderedKeys)

	if f.metadata {
		for _, key := range orderedKeys {
			mapkey := "metadata[" + key + "]"
			uValues.Set(mapkey, f.params[key])
		}
	} else {
		for _, key := range orderedKeys {
			mapkey := "filter[" + key + "]"
			uValues.Set(mapkey, f.params[key])
		}
	}

	return uValues.Encode()
}

func AddFilterAndSort(url string, filter *Filter, sort *Sort) string {
	emptyFilter := true
	emptySort := true

	if filter != nil && filter.String() != "" {
		emptyFilter = false
	}

	if sort != nil && sort.String() != "" {
		emptySort = false
	}

	if !emptyFilter && !emptySort {
		return url + "?" + filter.String() + "&" + sort.String()
	} else if !emptyFilter && emptySort {
		return url + "?" + filter.String()
	} else if emptyFilter && !emptySort {
		return url + "?" + sort.String()
	}

	return url
}

func AddFilterAndMetaFilterAndSort(url string, filter *Filter, metaFilter *Filter, sort *Sort) (string,error) {
	emptyFilter := true
	emptySort := true

	if filter != nil {
		if filter.metadata == true {
			return "",errors.New("filter parameter should not have metadata")
		}
	}


	if metaFilter != nil {
		if metaFilter.metadata == false {
			return "",errors.New("Metafilter parameter should have metadata")
		}
	}


	filterString := ""

	if (filter != nil && filter.String() != "") ||  (metaFilter != nil && metaFilter.String() != ""){
		emptyFilter = false
		if filter != nil {
			filterString = filter.String()
			if metaFilter != nil {
				filterString = filterString + "&" + metaFilter.String()
			}
		} else if metaFilter != nil {
			filterString = metaFilter.String()
		}

	}

	if sort != nil && sort.String() != "" {
		emptySort = false
	}



	if !emptyFilter && !emptySort {
		return url + "?" + filterString + "&" + sort.String(),nil
	} else if !emptyFilter && emptySort {
		return url + "?" + filterString, nil
	} else if emptyFilter && !emptySort {
		return url + "?" + filterString, nil
	}

	return url, nil
}
