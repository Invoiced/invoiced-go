package invdendpoint

import (
	"encoding/json"
	"strconv"
)

const InvoiceEndpoint = "/invoices"

type Invoices []Invoice

type Invoice struct {
	Id                     int64                  `json:"id,omitempty"`     // The invoice’s unique ID
	Object                 string                 `json:"object,omitempty"` // Object type, invoice
	Customer               int64                  `json:"-"`
	CustomerFull           *Customer              `json:"-"`
	CustomerRaw            json.RawMessage        `json:"customer,omitempty"`
	Name                   string                 `json:"name,omitempty"`                 // Invoice name for internal use, defaults to “Invoice”
	Number                 string                 `json:"number,omitempty"`               // The reference number assigned to the invoice for use in the dashboard
	AutoPay                bool                   `json:"autopay,omitempty"`              // Invoice collection mode, auto or manual
	Currency               string                 `json:"currency,omitempty"`             // 3-letter ISO code
	Draft                  bool                   `json:"draft,omitempty"`                          // When false, the invoice is considered outstanding, or when true, the invoice is a draft
	Closed                 bool                   `json:"closed,omitempty"`                         // When true, an invoice is closed and considered bad debt. No further payments are allowed.
	Paid                   bool                   `json:"paid,omitempty"`                 // Indicates whether an invoice has been paid in full
	Status                 string                 `json:"status,omitempty"`               // Invoice state, one of draft, not_sent, sent, viewed, past_due, pending, paid
	AttemptCount           int64                  `json:"attempt_count,omitempty"`        //# of payment attempts
	NextPaymentAttempt     int64                  `json:"next_payment_attempt,omitempty"` // Next scheduled charge attempt, when in automatic collection
	Subscription           int64                  `json:"subscription,omitempty"`         // Subscription ID if invoice came from subscription
	Date                   int64                  `json:"date,omitempty"`                 // Invoice date
	DueDate                int64                  `json:"due_date,omitempty"`             // Date payment is due by
	PaymentTerms           string                 `json:"payment_terms,omitempty"`        // Payment terms for the invoice, i.e. “NET 30”
	PurchaseOrder          string                 `json:"purchase_order,omitempty"`
	Items                  []LineItem             `json:"items,omitempty"`           // Collection of Line Items
	Notes                  string                 `json:"notes,omitempty"`           // Additional notes displayed on invoice
	Subtotal               float64                `json:"subtotal,omitempty"`        // Subtotal
	Discounts              []Discount             `json:"discounts,omitempty"`       // Collection of Discounts
	Taxes                  []Tax                  `json:"taxes,omitempty"`           // Collection of Taxes
	Total                  float64                `json:"total,omitempty"`           // Total
	Balance                float64                `json:"balance,omitempty"`         // Balance owed
	Url                    string                 `json:"url,omitempty"`             // URL to view the invoice in the billing portal
	PaymentUrl             string                 `json:"payment_url,omitempty"`     // URL for the invoice payment page
	PdfUrl                 string                 `json:"pdf_url,omitempty"`         // URL to download the invoice as a PDF
	CreatedAt              int64                  `json:"created_at,omitempty"`      // Timestamp when created
	Metadata               map[string]interface{} `json:"metadata,omitempty"`        // A hash of key/value pairs that can store additional information about this object.
	CalculateTaxes         bool                   `json:"calculate_taxes,omitempty"` // Flag to indicate whether taxes should be calculated on the invoice
	ShipTo                 *ShippingDetail        `json:"ship_to,omitempty"`
	Attachments            []int64                `json:"attachments,omitempty"`
	DisabledPaymentMethods []string               `json:"disabled_payment_methods,omitempty"`
	Sent 				   bool 				  `json:"sent,omitempty"`
}

func (i *Invoice) TotalTaxAmount() float64 {
	totalTax := 0.0

	for _, lineItem := range i.Items {
		lineItemTaxes := lineItem.Taxes

		for _, lineItemTax := range lineItemTaxes {
			totalTax += lineItemTax.Amount
		}
	}

	for _, invoiceTax := range i.Taxes {
		totalTax += invoiceTax.Amount
	}

	return totalTax
}

func (i *Invoice) TotalDiscountAmount() float64 {
	totalDiscount := 0.0

	for _, lineItem := range i.Items {
		lineItemDiscounts := lineItem.Discounts

		for _, lineItemDiscount := range lineItemDiscounts {
			totalDiscount += lineItemDiscount.Amount
		}
	}

	for _, invoiceDiscount := range i.Discounts {
		totalDiscount += invoiceDiscount.Amount
	}

	return totalDiscount
}

func (i *Invoice) UnmarshalJSON(data []byte) error {
	type invoice2 Invoice
	if err := json.Unmarshal(data, (*invoice2)(i)); err != nil {
		return err
	}

	rj := i.CustomerRaw

	i.Customer, _ = strconv.ParseInt(string(rj), 10, 64)
	customer := new(Customer)

	err := json.Unmarshal(rj, customer)

	if err == nil {
		i.CustomerFull = customer
		i.Customer = customer.Id
	}

	return nil
}

func (i *Invoice) MarshalJSON() ([]byte, error) {
	type invoice2 Invoice
	i2 := (*invoice2)(i)

	if i2.Customer > 0 {
		i2.CustomerRaw = []byte(strconv.FormatInt(i2.Customer, 10))
	}

	return json.Marshal(i2)
}

func (i *Invoice) String() string {
	b, _ := json.MarshalIndent(i, "", "    ")

	return string(b)
}
