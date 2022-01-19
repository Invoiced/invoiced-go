package invoiced

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
)

// NewMockApi is helpful when writing tests for
// functions that use an InvoiceClient connection to interact
// with the InvoiceClient API. It requires an arbitrary string
// as its Key parameter and an initialized http server.
//
// Example (error checking omitted):
//	Key := "test api Key"
// 	mockInvoiceResponse := new(invdendpoint.InvoiceClient)
// 	mockInvoiceResponse.Id = int64(12345)
// 	server, _ := invdmockserver.New(200, mockInvoiceResponse, "json", true)
// 	api := NewMockApi(Key, server)
// 	invoice := api.NewInvoice()
//
// Make sure that if you have rules that prune `unused-packages`,
// you make an exception for this project in order to get the
// rest of the requirements for mocking this codebase.
// 	[prune]
// 		unused-packages = true
// 		[[prune.project]]
// 			name = "github.com/Invoiced/invoiced-go"
// 			unused-packages = false

func NewMockApi(key string, server *httptest.Server) *Api {
	c := new(Api)
	c.Key = key

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c.client = &http.Client{Transport: transport}
	c.baseUrl = server.URL

	return c
}
