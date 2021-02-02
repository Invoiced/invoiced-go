package invdendpoint

import (
	"encoding/json"
)

const CreditNoteEndpoint = "/credit_notes"

type CreditNotes []CreditNote

type CreditNote struct {
	Id            int64                  `json:"id,omitempty"`       // The invoice’s unique ID
	Object        string                 `json:"object,omitempty"`   // Object type, estimate
	Customer      int64                  `json:"customer,omitempty"` // Customer ID
	Invoice       int64                  `json:"invoice,omitempty"`  // Customer ID
	Name          string                 `json:"name,omitempty"`     // Invoice name for internal use, defaults to “Invoice”
	Number        string                 `json:"number,omitempty"`   // The reference number assigned to the invoice for use in the dashboard
	Currency      string                 `json:"currency,omitempty"` // 3-letter ISO code
	Draft         bool                   `json:"draft,omitempty"`    // When false, the invoice is considered outstanding, or when true, the invoice is a draft
	Closed        bool                   `json:"closed,omitempty"`   // When true, an invoice is closed and considered bad debt. No further payments are allowed.
	Paid          bool                   `json:"paid,omitempty"`     // When true, an invoice is closed and considered bad debt. No further payments are allowed.
	Status        string                 `json:"status,omitempty"`   // Invoice state, one of draft, not_sent, sent, viewed, past_due, pending, paid
	Date          int64                  `json:"date,omitempty"`     // Invoice date
	PurchaseOrder string                 `json:"purchase_order,omitempty"`
	Items         []LineItem             `json:"items,omitempty"`           // Collection of Line Items
	Notes         string                 `json:"notes,omitempty"`           // Additional notes displayed on invoice
	Subtotal      float64                `json:"subtotal,omitempty"`        // Subtotal
	Discounts     []Discount             `json:"discounts,omitempty"`       // Collection of Discounts
	Taxes         []Tax                  `json:"taxes,omitempty"`           // Collection of Taxes
	Balance       float64                `json:"balance,omitempty"`         // Balance owed
	Total         float64                `json:"total,omitempty"`           // Total
	Url           string                 `json:"url,omitempty"`             // URL to download the invoice as a PDF
	PdfUrl        string                 `json:"pdf_url,omitempty"`         // URL to download the invoice as a PDF
	CreatedAt     int64                  `json:"created_at,omitempty"`      // Timestamp when created
	Metadata      map[string]interface{} `json:"metadata,omitempty"`        // A hash of key/value pairs that can store additional information about this object.
	Attachments   []int64                `json:"attachments,omitempty"`     // A list of File IDs to attach to the estimate
	CalculateTax  bool                   `json:"calculate_taxes,omitempty"` // Disables tax calculation, default is true
}

func (i *CreditNote) String() string {
	b, _ := json.MarshalIndent(i, "", "    ")

	return string(b)
}
