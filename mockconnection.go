package invdapi

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
)

// mockConnection is helpful when writing tests for
// functions that use an Invoice connection to interact
// with the Invoice API. It requires an arbitrary string
// as its key parameter and an initialized http server.
//
// Example (error checking omitted):
//	key := "test api key"
// 	mockInvoiceResponse := new(invdendpoint.Invoice)
// 	mockInvoiceResponse.Id = int64(12345)
// 	server, _ := invdmockserver.New(200, mockInvoiceResponse, "json", true)
// 	conn := mockConnection(key, server)
// 	invoice := conn.NewInvoice()
//
// Make sure that if you have rules that prune `unused-packages`,
// you make an exception for this project in order to get the
// rest of the requirements for mocking this codebase.
// 	[prune]
// 		unused-packages = true
// 		[[prune.project]]
// 			name = "github.com/Invoiced/invoiced-go"
// 			unused-packages = false

func mockConnection(key string, server *httptest.Server) *Connection {
	c := new(Connection)
	c.key = key

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c.client = &http.Client{Transport: transport}
	c.baseUrl = server.URL

	return c
}
