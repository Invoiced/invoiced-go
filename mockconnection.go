package invdapi

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
)

func mockConnection(key string, server *httptest.Server) *Connection {
	c := new(Connection)
	c.key = key

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c.client = &http.Client{Transport: transport}
	c.url = server.URL

	return c
}
