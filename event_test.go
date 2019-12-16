package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"testing"
)

func TestEvent_ListAll(t *testing.T) {

	key := "test api key"

	var mockListResponse [1] invdendpoint.Event

	mockResponse := new(invdendpoint.Event)
	mockResponse.Id = int64(123)

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewEvent()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(result[0].Event, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}