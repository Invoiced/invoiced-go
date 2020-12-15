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
	mockPaymentResponse := new(invdendpoint.Refund)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.Charge = 234112
	mockPaymentResponse.GatewayId = "234"

	mockPaymentResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	refund := conn.NewRefund()
	err = refund.Create(234112, 123.00)

	if err != nil {
		t.Fatal("Error Creating refund", err)
	}

	if refund.Id != int64(1523) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}
