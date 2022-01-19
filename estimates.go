package invoiced

import (
	"encoding/json"
)

const EstimateEndpoint = "/estimates"

type EstimateRequest struct {
	Approved               *string                 `json:"approved,omitempty"`
	Attachments            []*int64                `json:"attachments,omitempty"`
	CalculateTax           *bool                   `json:"calculate_taxes,omitempty"`
	Closed                 *bool                   `json:"closed,omitempty"`
	Currency               *string                 `json:"currency,omitempty"`
	Customer               *int64                  `json:"customer,omitempty"`
	Date                   *int64                  `json:"date,omitempty"`
	Deposit                *float64                `json:"deposit,omitempty"`
	DepositPaid            *bool                   `json:"deposit_paid,omitempty"`
	DisabledPaymentMethods []*string               `json:"disabled_payment_methods,omitempty"`
	Discounts              []*DiscountRequest      `json:"discounts,omitempty"`
	Draft                  *bool                   `json:"draft,omitempty"`
	ExpirationDate         *int64                  `json:"expiration_date,omitempty"`
	Items                  []*LineItemRequest      `json:"items,omitempty"`
	Metadata               *map[string]interface{} `json:"metadata,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
	Notes                  *string                 `json:"notes,omitempty"`
	Number                 *string                 `json:"number,omitempty"`
	PaymentTerms           *string                 `json:"payment_terms,omitempty"`
	PurchaseOrder          *string                 `json:"purchase_order,omitempty"`
	ShipTo                 *string                 `json:"ship_to,omitempty"`
	Taxes                  []*TaxRequest           `json:"taxes,omitempty"`
	UpdatedAt              *int64                  `json:"updated_at,omitempty"`
}

type Estimates []Estimate

type Estimate struct {
	Approved               string                 `json:"approved"`
	Attachments            []int64                `json:"attachments"`
	Closed                 bool                   `json:"closed"`
	CreatedAt              int64                  `json:"created_at"`
	Currency               string                 `json:"currency"`
	Customer               int64                  `json:"customer"`
	Date                   int64                  `json:"date"`
	Deposit                float64                `json:"deposit"`
	DepositPaid            bool                   `json:"deposit_paid"`
	DisabledPaymentMethods []string               `json:"disabled_payment_methods"`
	Discounts              []Discount             `json:"discounts"`
	Draft                  bool                   `json:"draft"`
	ExpirationDate         int64                  `json:"expiration_date"`
	Id                     int64                  `json:"id"`
	Invoice                int64                  `json:"invoice"`
	Items                  []LineItem             `json:"items"`
	Metadata               map[string]interface{} `json:"metadata"`
	Name                   string                 `json:"name"`
	Notes                  string                 `json:"notes"`
	Number                 string                 `json:"number"`
	Object                 string                 `json:"object"`
	PaymentTerms           string                 `json:"payment_terms"`
	PdfUrl                 string                 `json:"pdf_url"`
	PurchaseOrder          string                 `json:"purchase_order"`
	ShipTo                 string                 `json:"ship_to"`
	Status                 string                 `json:"status"`
	Subtotal               float64                `json:"subtotal"`
	Taxes                  []Tax                  `json:"taxes"`
	Total                  float64                `json:"total"`
	UpdatedAt              int64                  `json:"updated_at"`
	Url                    string                 `json:"url"`
}

func (i *Estimate) String() string {
	b, _ := json.MarshalIndent(i, "", "    ")

	return string(b)
}
