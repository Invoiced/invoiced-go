package invdendpoint

import (
	"encoding/json"
)

const EstimateEndpoint = "/estimates"

type Estimates []Estimate

type Estimate struct {
	Id                     int64                  `json:"id,omitempty"`              // The invoice’s unique ID
	Object                 string                 `json:"object,omitempty"`          // Object type, estimate
	Customer               int64                  `json:"customer,omitempty"`        // Customer ID
	Invoice                int64                  `json:"invoice,omitempty"`         // Customer ID
	Name                   string                 `json:"name,omitempty"`            // Invoice name for internal use, defaults to “Invoice”
	Number                 string                 `json:"number,omitempty"`          // The reference number assigned to the invoice for use in the dashboard
	Currency               string                 `json:"currency,omitempty"`        // 3-letter ISO code
	Draft                  bool                   `json:"draft,omitempty"`           // When false, the invoice is considered outstanding, or when true, the invoice is a draft
	Closed                 bool                   `json:"closed,omitempty"`          // When true, an invoice is closed and considered bad debt. No further payments are allowed.
	Approved               string                 `json:"approved,omitempty"`        // When true, an invoice is closed and considered bad debt. No further payments are allowed.
	Status                 string                 `json:"status,omitempty"`          // Invoice state, one of draft, not_sent, sent, viewed, past_due, pending, paid
	Date                   int64                  `json:"date,omitempty"`            // Invoice date
	ExpirationDate         int64                  `json:"expiration_date,omitempty"` // Estimate expiration date
	PaymentTerms           string                 `json:"payment_terms,omitempty"`   // Payment terms for the invoice, i.e. “NET 30”
	PurchaseOrder          string                 `json:"purchase_order,omitempty"`
	Items                  []LineItem             `json:"items,omitempty"`                    // Collection of Line Items
	Notes                  string                 `json:"notes,omitempty"`                    // Additional notes displayed on invoice
	Subtotal               float64                `json:"subtotal,omitempty"`                 // Subtotal
	Discounts              []Discount             `json:"discounts,omitempty"`                // Collection of Discounts
	Taxes                  []Tax                  `json:"taxes,omitempty"`                    // Collection of Taxes
	ShipTo                 string                 `json:"ship_to,omitempty"`                  // Shipipng Detail object
	Total                  float64                `json:"total,omitempty"`                    // Total
	Deposit                float64                `json:"deposit,omitempty"`                  // Deposit
	DepositPaid            bool                   `json:"deposit_paid,omoitempty"`            // Deposit Paid
	Url                    string                 `json:"url,omitempty"`                      // URL to download the invoice as a PDF
	PdfUrl                 string                 `json:"pdf_url,omitempty"`                  // URL to download the invoice as a PDF
	CreatedAt              int64                  `json:"created_at,omitempty"`               // Timestamp when created
	Metadata               map[string]interface{} `json:"metadata,omitempty"`                 // A hash of key/value pairs that can store additional information about this object.
	Attachments            []int64                `json:"attachments,omitempty"`              // A list of File IDs to attach to the estimate
	DisabledPaymentMethods []string               `json:"disabled_payment_methods,omitempty"` // List of payment methods to disable for this estimate, i.e. ["credit_card", "wire_transfer"].
	CalculateTax           bool                   `json:"calculate_taxes,omitempty"`          // Disables tax calculation, default is true
}

func (i *Estimate) String() string {
	b, _ := json.MarshalIndent(i, "", "    ")

	return string(b)
}
