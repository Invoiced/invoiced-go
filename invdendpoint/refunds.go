package invdendpoint

type RefundRequest struct {
	Amount float64 `json:"amount,omitempty"`
}

type Refund struct {
	Id             int64   `json:"id,omitempty"`              // The payment’s unique ID
	Object         string  `json:"object,omitempty"`          // Object type, payment
	Charge         int64   `json:"charge,omitempty"`          // Charge ID
	Status         string  `json:"status,omitempty"`          // Payment status, one of succeeded, pending, or failed, defaults to succeeded
	Gateway        string  `json:"gateway,omitempty"`         // Payment gateway that processed the payment, if any
	GatewayId      string  `json:"gateway_id,omitempty"`      // Payment ID from the payment gateway, or check # if method is check
	Currency       string  `json:"currency,omitempty"`        // 3-letter ISO code
	Amount         float64 `json:"amount,omitempty"`          // Payment amount
	FailureMessage string  `json:"failure_message,omitempty"` // Failure message from the payment gateway (only available when status = failed)
	CreatedAt      int64   `json:"created_at,omitempty"`      // Timestamp when created
}
