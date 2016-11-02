package invdmockserver

import (
	"github.com/Invoiced/invoiced-go/invdutil"
)

type RequestObject struct {
	Method string         `json:"method,omitempty"`
	Url    string         `json:"url,omitempty"`
	Body   []BodyPatterns `json:"bodyPatterns"`
}

type ResponseObject struct {
	Status int    `json:"status,omitempty"`
	Body   string `json:"body,omitempty"`
}

type RRActionObject struct {
	Request  RequestObject  `json:"request,omitempty"`
	Response ResponseObject `json:"response,omitempty"`
}

type BodyPatterns struct {
	EqualToJson string `json:"equalToJson,omitempty"`
}

type RRActionMap struct {
	store map[string]map[string][]*RRActionObject
}

func NewRRActionMap() *RRActionMap {
	store := make(map[string]map[string][]*RRActionObject)

	rrActionMap := new(RRActionMap)

	rrActionMap.store = store

	return rrActionMap

}

func (r *RRActionMap) Put(rrActionObject *RRActionObject) error {

	method := rrActionObject.Request.Method
	url := rrActionObject.Request.Url

	_, found0 := r.store[url]

	if !found0 {
		r.store[url] = make(map[string][]*RRActionObject)
	}

	rrActionObjects, found := r.store[url][method]

	if !found {
		r.store[url][method] = append(r.store[url][method], rrActionObject)
	} else {

		for _, rrActionObjectItem := range rrActionObjects {

			jsonBody1 := "{}"
			jsonBody2 := "{}"

			if len(rrActionObjectItem.Request.Body) > 0 {
				jsonBody1 = rrActionObjectItem.Request.Body[0].EqualToJson
			}

			if len(rrActionObject.Request.Body) > 0 {
				jsonBody2 = rrActionObject.Request.Body[0].EqualToJson
			}

			equal, err := invdutil.JsonEqual(jsonBody1, jsonBody2)

			if err != nil {
				return err
			}

			if equal {
				return nil
			}

		}

		r.store[url][method] = append(r.store[url][method], rrActionObject)

	}

	return nil

}

func (r *RRActionMap) Get(method, url, body string) (*RRActionObject, bool, error) {

	if len(body) == 0 {
		body = "{}"
	}

	_, found0 := r.store[url]

	if !found0 {
		return nil, false, nil
	}

	rrActionObjects, found := r.store[url][method]

	if !found {
		return nil, false, nil
	}

	for _, rrActionObject := range rrActionObjects {
		jsonBody1 := "{}"
		if len(rrActionObject.Request.Body) > 0 {
			jsonBody1 = rrActionObject.Request.Body[0].EqualToJson
		}

		jsonBody2 := body

		equal, err := invdutil.JsonEqual(jsonBody1, jsonBody2)

		if err != nil {
			return nil, false, err
		}

		if equal {
			return rrActionObject, true, nil
		}

	}

	return nil, false, nil

}
