package invdapi

import (
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdendpoint"
	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestCreateRefund(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(invdendpoint.Payment)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.Customer = 234112
	mockPaymentResponse.GatewayId = "234"

	mockPaymentResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	refund := conn.NewPayment()
	err = refund.Refund(123.00)

	if err != nil {
		t.Fatal("Error Creating refund", err)
	}

	if refund.Id != int64(1523) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}
