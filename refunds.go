package invoiced

type RefundRequest struct {
	Amount *float64 `json:"amount,omitempty"`
}

type Refund struct {
	Amount         float64 `json:"amount,omitempty"`
	Charge         int64   `json:"charge,omitempty"`
	CreatedAt      int64   `json:"created_at,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	FailureMessage string  `json:"failure_message,omitempty"`
	Gateway        string  `json:"gateway,omitempty"`
	GatewayId      string  `json:"gateway_id,omitempty"`
	Id             int64   `json:"id,omitempty"`
	Object         string  `json:"object,omitempty"`
	Status         string  `json:"status,omitempty"`
	UpdatedAt      int64   `json:"updated_at,omitempty"`
}
