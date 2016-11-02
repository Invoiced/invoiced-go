package invdmockserver

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

const dir = "./resources/"

var rrActionMap *RRActionMap = nil

func GetRRActionMap() *RRActionMap {
	return rrActionMap
}

func LoadJsonMappings() error {

	rrActionMap = NewRRActionMap()

	files, _ := ioutil.ReadDir(dir)

	jsonFiles := []os.FileInfo{}
	for _, f := range files {
		if strings.Contains(f.Name(), ".json") {
			jsonFiles = append(jsonFiles, f)
		}
	}

	for _, jsonFile := range jsonFiles {

		b, err := ioutil.ReadFile(dir + jsonFile.Name())

		if err != nil {
			return err
		}

		rrActionObject := new(RRActionObject)

		err = json.Unmarshal(b, rrActionObject)

		if err != nil {
			fmt.Println(jsonFile.Name())
			fmt.Println(string(b))
			return err
		}

		rrActionMap.Put(rrActionObject)

	}

	return nil

}

func NewJsonFileServer(ssl bool) (*httptest.Server, error) {

	var server *httptest.Server

	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var bodyMarshalled []byte

		url := r.RequestURI
		method := r.Method

		w.Header().Set("Content-Type", "application/json")

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err.Error())
			return
		}

		rrActionObject, found, err := rrActionMap.Get(method, url, string(body))

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err.Error())
			return
		}

		if found {

			w.WriteHeader(rrActionObject.Response.Status)
			fmt.Fprintln(w, rrActionObject.Response.Body)

		} else {
			w.WriteHeader(504)
			fmt.Fprintln(w, `{"errorMessage":"Resource Invalid"}`)

		}

		// fmt.Fprintln(w, string(bodyMarshalled))
	})

	if ssl {
		server = httptest.NewTLSServer(f)
	} else {
		server = httptest.NewServer(f)
	}

	return server, nil

}

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
