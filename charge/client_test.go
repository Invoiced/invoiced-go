package charge

import (
	"github.com/Invoiced/invoiced-go"
	"github.com/Invoiced/invoiced-go/payment"
	"testing"
	"time"

	"github.com/Invoiced/invoiced-go/invdmockserver"
)

func TestCreateCharge(t *testing.T) {
	key := "test api key"

	mockChargeResponseID := int64(1523)
	mockChargeResponse := new(payment.Client)
	mockChargeResponse.Id = mockChargeResponseID
	mockChargeResponse.Customer = 234112
	mockChargeResponse.Reference = "234"

	mockChargeResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockChargeResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	client := Client{invoiced.NewMockApi(key, server)}

	chargeRequest := new(invoiced.ChargeRequest)

	charge, err := client.Create(chargeRequest)
	if err != nil {
		t.Fatal("Error Creating charge", err)
	}

	if charge.Id != int64(1523) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}

func TestCreateRefund(t *testing.T) {
	key := "test api key"

	mockPaymentResponseID := int64(1523)
	mockPaymentResponse := new(invoiced.Refund)
	mockPaymentResponse.Id = mockPaymentResponseID
	mockPaymentResponse.Charge = 234112
	mockPaymentResponse.GatewayId = "234"

	mockPaymentResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockPaymentResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := Client{invoiced.NewMockApi(key, server)}
	refund, err := client.Refund(234112, &invoiced.RefundRequest{invoiced.Float64(123.00)})

	if err != nil {
		t.Fatal("Error Creating refund", err)
	}

	if refund.Id != int64(1523) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}
