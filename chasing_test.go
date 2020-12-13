package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestChasingCadence_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.ChasingCadence

	mockResponse := new(invdendpoint.ChasingCadence)
	mockResponse.Id = int64(123)
	mockResponse.Name = "standard"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewChasingCadence()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(result[0].ChasingCadence, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
