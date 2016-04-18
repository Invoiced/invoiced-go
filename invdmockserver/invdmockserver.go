package invdmockserver

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func New(code int, body interface{}, dataType string, ssl bool) (*httptest.Server, error) {

	var bodyMarshalled []byte
	var err error

	if dataType == "xml" {

		bodyMarshalled, err = xml.Marshal(body)

		if err != nil {
			return nil, err
		}

	} else if dataType == "json" {

		bodyMarshalled, err = json.Marshal(body)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("dataType in MockServer not recognized")

	}

	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if dataType == "xml" {
			w.Header().Set("Content-Type", "application/xml")
		} else if dataType == "json" {
			w.Header().Set("Content-Type", "application/json")
		}

		w.WriteHeader(code)

		fmt.Fprintln(w, string(bodyMarshalled))
	})

	var server *httptest.Server

	if ssl {
		server = httptest.NewTLSServer(f)
	} else {
		server = httptest.NewServer(f)
	}

	return server, nil
}
