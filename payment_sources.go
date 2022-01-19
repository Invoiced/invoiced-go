package invoiced

import (
	"encoding/json"
	"errors"
)

type PaymentSourceRequest struct {
	GatewayToken  *string `json:"gateway_token,omitempty"`
	InvoicedToken *string `json:"invoiced_token,omitempty"`
	MakeDefault   *bool   `json:"make_default,omitempty"`
	Method        *string `json:"method,omitempty"`
	Object        *string `json:"object,omitempty"`
}

type PaymentSource struct {
	*BankAccount
	*Card
	Object string `json:"object"`
}

type PaymentSources []PaymentSource

type Card struct {
	Brand           string `json:"brand"`
	Chargeable      bool   `json:"chargeable"`
	CreatedAt       int64  `json:"created_at"`
	ExpMonth        int64  `json:"exp_month"`
	ExpYear         int64  `json:"exp_year"`
	FailureReason   string `json:"failure_reason"`
	Funding         string `json:"funding"`
	Gateway         string `json:"gateway"`
	GatewayCustomer string `json:"gateway_customer"`
	GatewayId       string `json:"gateway_id"`
	Id              int64  `json:"id"`
	Last4           string `json:"last4"`
	Object          string `json:"object"`
	ReceiptEmail    string `json:"receipt_email"`
	UpdatedAt       int64  `json:"updated_at"`
}

type BankAccount struct {
	BankName        string `json:"bank_name"`
	Chargeable      bool   `json:"chargeable"`
	Country         string `json:"country"`
	CreatedAt       int64  `json:"created_at"`
	Currency        string `json:"currency"`
	FailureReason   string `json:"failure_reason"`
	Gateway         string `json:"gateway"`
	GatewayCustomer string `json:"gateway_customer"`
	GatewayId       string `json:"gateway_id"`
	Id              int64  `json:"id"`
	Last4           string `json:"last4"`
	Object          string `json:"object"`
	ReceiptEmail    string `json:"receipt_email"`
	RoutingNumber   string `json:"routing_number"`
	UpdatedAt       int64  `json:"updated_at"`
	Verified        bool   `json:"verified"`
}

func (d *PaymentSource) UnmarshalJSON(data []byte) error {
	temp := struct {
		Object string `json:"object"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Object == "card" {
		var c Card
		if err := json.Unmarshal(data, &c); err != nil {
			return err
		}
		d.Card = &c
		d.BankAccount = nil
	} else if temp.Object == "bank_account" {
		var ba BankAccount
		if err := json.Unmarshal(data, &ba); err != nil {
			return err
		}
		d.BankAccount = &ba
		d.Card = nil
	} else {
		return errors.New("Invalid object value")
	}
	return nil
}
