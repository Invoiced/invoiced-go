package invdendpoint

const PaymentEndpoint = "/payments"

type Payments []Payment

type Payment struct {
	Id        int64                `json:"id,omitempty"`       // The payment’s unique ID
	Object    string               `json:"object,omitempty"`   // Object type, payment
	Customer  int64                `json:"customer,omitempty"` // Customer ID, required if invoice ID is not supplied
	Date      int64                `json:"date,omitempty"`     // Payment date, defaults to current timestamp
	Method    string               `json:"method,omitempty"`   // Payment instrument used to facilitate payment, defaults to other
	Matched   bool               `json:"matched,omitempty"`
	Voided    bool               `json:"voided,omitempty"`
	Status    string        `json:"status,omitempty"`   // Payment status
	Currency  string        `json:"currency,omitempty"` // 3-letter ISO code
	Amount    float64       `json:"amount,omitempty"`   // Payment amount
	Balance   float64       `json:"balance,omitempty"`
	Reference string        `json:"reference,omitempty"`
	Source    string        `json:"source,omitempty"`
	Notes     string        `json:"notes,omitempty"` // Internal notes
	Charge    *Charge       `json:"charge,omitempty"`
	PdfUrl    string        `json:"pdf_url,omitempty"`    // URL to download the invoice as a PDF
	CreatedAt int64         `json:"created_at,omitempty"` // Timestamp when created
	AppliedTo []PaymentItem `json:"applied_to,omitempty"`
}

type PaymentItem struct {
	Type         string  `json:"type,omitempty"`
	Invoice      int64   `json:"invoice,omitempty"`
	CreditNote   int64   `json:"credit_note,omitempty"`
	Estimate     int64   `json:"estimate,omitempty"`
	DocumentType int64   `json:"document_type,omitempty"`
	Amount       float64 `json:"amount,omitempty"`
}
