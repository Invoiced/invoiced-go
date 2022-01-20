package chasing

import (
	"github.com/Invoiced/invoiced-go/v2"
	"reflect"
	"testing"
	"time"
	"github.com/Invoiced/invoiced-go/v2/invdmockserver"
)

func TestChasingCadence_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invoiced.ChasingCadence

	mockResponse := new(invoiced.ChasingCadence)
	mockResponse.Id = int64(123)
	mockResponse.Name = "standard"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}

	filter := invoiced.NewFilter()
	sorter := invoiced.NewSort()

	result, err := client.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(result[0], mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
