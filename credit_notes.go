package invoiced

import (
	"encoding/json"
)

const CreditNoteEndpoint = "/credit_notes"

type CreditNoteRequest struct {
	Attachments   []*int64                `json:"attachments,omitempty"`
	CalculateTax  *bool                   `json:"calculate_taxes,omitempty"`
	Closed        *bool                   `json:"closed,omitempty"`
	Currency      *string                 `json:"currency,omitempty"`
	Customer      *int64                  `json:"customer,omitempty"`
	Date          *int64                  `json:"date,omitempty"`
	Discounts     []*DiscountRequest      `json:"discounts,omitempty"`
	Draft         *bool                   `json:"draft,omitempty"`
	Items         []*LineItemRequest      `json:"items,omitempty"`
	Metadata      *map[string]interface{} `json:"metadata,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	Notes         *string                 `json:"notes,omitempty"`
	Number        *string                 `json:"number,omitempty"`
	Paid          *bool                   `json:"paid,omitempty"`
	PurchaseOrder *string                 `json:"purchase_order,omitempty"`
	Taxes         []*TaxRequest           `json:"taxes,omitempty"`
}

type CreditNotes []CreditNote

type CreditNote struct {
	Attachments   []int64                `json:"attachments"`
	Balance       float64                `json:"balance"`
	Closed        bool                   `json:"closed"`
	CreatedAt     int64                  `json:"created_at"`
	Currency      string                 `json:"currency"`
	Customer      int64                  `json:"customer"`
	Date          int64                  `json:"date"`
	Discounts     []Discount             `json:"discounts"`
	Draft         bool                   `json:"draft"`
	Id            int64                  `json:"id"`
	Invoice       int64                  `json:"invoice"`
	Items         []LineItem             `json:"items"`
	Metadata      map[string]interface{} `json:"metadata"`
	Name          string                 `json:"name"`
	Notes         string                 `json:"notes"`
	Number        string                 `json:"number"`
	Object        string                 `json:"object"`
	Paid          bool                   `json:"paid"`
	PdfUrl        string                 `json:"pdf_url"`
	PurchaseOrder string                 `json:"purchase_order"`
	Status        string                 `json:"status"`
	Subtotal      float64                `json:"subtotal"`
	Taxes         []Tax                  `json:"taxes"`
	Total         float64                `json:"total"`
	UpdatedAt     int64                  `json:"updated_at"`
	Url           string                 `json:"url"`
}

func (i *CreditNote) String() string {
	b, _ := json.MarshalIndent(i, "", "    ")

	return string(b)
}
