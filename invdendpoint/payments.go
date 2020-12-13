package invdendpoint

const PaymentEndpoint = "/payments"

type Payments []Payment

type Payment struct {
	Id            int64                  `json:"id,omitempty"`             // The payment’s unique ID
	Object        string                 `json:"object,omitempty"`         // Object type, payment
	Customer      int64                  `json:"customer,omitempty"`       // Customer ID, required if invoice ID is not supplied
	Invoice       int64                  `json:"invoice,omitempty"`        // Invoice ID, if any
	CreditNote    int64                  `json:"credit_note,omitempty"`    // Associated Credit Note, if any
	Type          string                 `json:"type,omitempty"`           // Payment type, charge, payment, refund, or adjustment - required
	Date          int64                  `json:"date,omitempty"`           // Payment date, defaults to current timestamp
	Method        string                 `json:"method,omitempty"`         // Payment instrument used to facilitate payment, defaults to other
	Status        string                 `json:"status,omitempty"`         // Payment status, one of succeeded, pending, or failed, defaults to succeeded
	Gateway       string                 `json:"gateway,omitempty"`        // Payment gateway that processed the payment, if any
	GatewayId     string                 `json:"gateway_id,omitempty"`     // Payment ID from the payment gateway, or check # if method is check
	PaymentSource *PaymentSource         `json:"payment_source,omitempty"` // Payment source used for payment, if any
	Currency      string                 `json:"currency,omitempty"`       // 3-letter ISO code
	Amount        float64                `json:"amount,omitempty"`         // Payment amount
	Fee           float64                `json:"fee,omitempty"`            // Processing fees
	Notes         string                 `json:"notes,omitempty"`          // Internal notes
	FailureReason string                 `json:"failure_reason,omitempty"` // Failure message from the payment gateway (only available when status = failed)
	ParentPayment int64                  `json:"parent_payment,omitempty"` // ID of the original payment for refunds
	PdfUrl        string                 `json:"pdf_url,omitempty"`        // URL to download the invoice as a PDF
	CreatedAt     int64                  `json:"created_at,omitempty"`     // Timestamp when created
	Metadata      map[string]interface{} `json:"metadata,omitempty"`       // A hash of key/value pairs that can store additional information about this object.
}
