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

	if dataType == "xml" {

		_, err := xml.Marshal(body)

		if err != nil {
			return nil, err
		}

	} else if dataType == "json" {

		_, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("dataType in MockServer not recognized")

	}

	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyMarshalled []byte

		if dataType == "xml" {
			w.Header().Set("Content-Type", "application/xml")
			bodyMarshalled, _ = xml.Marshal(body)

		} else if dataType == "json" {
			w.Header().Set("Content-Type", "application/json")
			bodyMarshalled, _ = json.Marshal(body)
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
