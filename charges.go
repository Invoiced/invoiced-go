package invoiced

type ChargeRequest struct {
	Customer          *int64                `json:"customer,omitempty"`
	Method            *string               `json:"method,omitempty"`
	Currency          *string               `json:"currency,omitempty"`
	Amount            *float64              `json:"amount,omitempty"`
	InvoicedToken     *string               `json:"invoiced_token,omitempty"`
	GatewayToken      *string               `json:"gateway_token,omitempty"`
	PaymentSourceType *string               `json:"payment_source_type,omitempty"`
	PaymentSourceId   *int64                `json:"payment_source_id,omitempty"`
	VaultMethod       *bool                 `json:"vault_method,omitempty"`
	MakeDefault       *bool                 `json:"make_default,omitempty"`
	ReceiptEmail      *string               `json:"receipt_email,omitempty"`
	AppliedTo         []*PaymentItemRequest `json:"applied_to,omitempty"`
}

type Charge struct {
	Id             int64          `json:"id"`
	Object         string         `json:"object"`
	Customer       int64          `json:"customer"`
	Status         string         `json:"status"`
	Gateway        string         `json:"gateway"`
	GatewayId      string         `json:"gateway_id"`
	PaymentSource  *PaymentSource `json:"payment_source"`
	Currency       string         `json:"currency"`
	Amount         float64        `json:"amount"`
	FailureMessage string         `json:"failure_message"`
	AmountRefunded float64        `json:"amount_refunded"`
	Refunded       bool           `json:"refunded"`
	Refunds        []Refund       `json:"refunds"`
	Disputed       bool           `json:"disputed"`
	CreatedAt      int64          `json:"created_at"`
	UpdatedAt      int64          `json:"updated_at"`
}
