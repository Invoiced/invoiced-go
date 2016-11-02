package invdmockserver

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type JsonTest struct {
	Msg string `json:"Msg"`
}

type XmlTest struct {
	Msg string `xml:"Msg"`
}

func TestLoadJsonMappings(t *testing.T) {

	LoadJsonMappings()

	if GetRRActionMap() == nil {
		t.Fatal("GetRRActionMap should not be nil")
	}

}

func TestJsonFileServer(t *testing.T) {
	//references connection_rr_52.json
	LoadJsonMappings()

	server, err := NewJsonFileServer(false)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := http.Get(server.URL + "/customers/198971")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 404 {
		t.Fatal("Status code is incorrect ", resp.StatusCode)
	}

}

func TestJsonMockServer(t *testing.T) {

	j := new(JsonTest)
	j.Msg = "Hello World"
	server, err := New(200, j, "json", false)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := `{"Msg":"Hello World"}`

	if strings.TrimSpace(string(b)) != expectedResponse {
		t.Fatal("Incorrect Response From JsonMockServer, actual response => ", string(b), " ,expected resonse => ", expectedResponse)
	}

}

func TestXMLMockServer(t *testing.T) {
	j := new(XmlTest)
	j.Msg = "Hello World"
	server, err := New(200, j, "xml", false)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := `<XmlTest><Msg>Hello World</Msg></XmlTest>`

	if strings.TrimSpace(string(b)) != expectedResponse {
		t.Fatal("Incorrect Response From XMLMockServer, actual response => ", string(b), " ,expected resonse => ", expectedResponse)
	}
}

func TestRespCodeMockServer(t *testing.T) {
	j := new(XmlTest)
	j.Msg = "Hello World"
	statusCode := 400
	server, err := New(statusCode, j, "xml", false)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != statusCode {
		t.Fatal("Status Code does not match")
	}

}
