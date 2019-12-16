package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
	"reflect"
	"testing"
	"time"
)

func TestCoupon_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Coupon)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewCoupon()

	requestEntity := entity.NewCoupon()

	requestEntity.Id = "example"
	requestEntity.Name = "nomenclature"

	createdEntity, err := entity.Create(requestEntity)

	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.Coupon, mockResponse) {
		t.Fatal("entity was not created", createdEntity.Coupon, mockResponse)
	}

}

func TestCoupon_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.Coupon)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)

	entityToUpdate := conn.NewCoupon()

	entityToUpdate.Name = "new-name"

	err = entityToUpdate.Save()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.Coupon) {
		t.Fatal("Error: entity not updated correctly")
	}

}

func TestCoupon_Delete(t *testing.T) {

	key := "api key"

	mockResponse := ""
	mockResponseId := "example"

	server, err := invdmockserver.New(204, mockResponse, "json", true)

	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := MockConnection(key, server)

	entity := conn.NewCoupon()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}

}

func TestCoupon_Retrieve(t *testing.T) {

	key := "test api key"

	mockResponse := new(invdendpoint.Coupon)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewCoupon()

	retrievedTransaction, err := entity.Retrieve("example")

	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedTransaction.Coupon, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}

func TestCoupon_ListAll(t *testing.T) {

	key := "test api key"

	var mockListResponse [1] invdendpoint.Coupon

	mockResponse := new(invdendpoint.Coupon)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := MockConnection(key, server)
	entity := conn.NewCoupon()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)

	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].Coupon, mockResponse) {
		t.Fatal("Error messages do not match up")
	}

}