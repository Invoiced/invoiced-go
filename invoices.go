package invoiced

import (
	"encoding/json"
	"strconv"
)

type InvoiceRequest struct {
	Attachments            []*int64                `json:"attachments,omitempty"`
	AutoPay                *bool                   `json:"autopay,omitempty"`
	CalculateTaxes         *bool                   `json:"calculate_taxes,omitempty"`
	Closed                 *bool                   `json:"closed,omitempty"`
	Currency               *string                 `json:"currency,omitempty"`
	Customer               *int64                  `json:"customer"`
	Date                   *int64                  `json:"date,omitempty"`
	DisabledPaymentMethods []*string               `json:"disabled_payment_methods,omitempty"`
	Discounts              []*DiscountRequest      `json:"discounts,omitempty"`
	Draft                  *bool                   `json:"draft,omitempty"`
	DueDate                *int64                  `json:"due_date,omitempty"`
	Items                  []*LineItemRequest      `json:"items,omitempty"`
	Metadata               *map[string]interface{} `json:"metadata,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
	NextPaymentAttempt     *int64                  `json:"next_payment_attempt,omitempty"`
	Notes                  *string                 `json:"notes,omitempty"`
	Number                 *string                 `json:"number,omitempty"`
	PaymentTerms           *string                 `json:"payment_terms,omitempty"`
	PurchaseOrder          *string                 `json:"purchase_order,omitempty"`
	Sent                   *bool                   `json:"sent,omitempty"`
	ShipTo                 *ShippingDetailRequest  `json:"ship_to,omitempty"`
	Taxes                  []*TaxRequest           `json:"taxes,omitempty"`
}

type Invoice struct {
	Attachments            []int64                `json:"attachments"`
	AttemptCount           int64                  `json:"attempt_count"`
	AutoPay                bool                   `json:"autopay"`
	Balance                float64                `json:"balance"`
	Closed                 bool                   `json:"closed"`
	CreatedAt              int64                  `json:"created_at"`
	Currency               string                 `json:"currency"`
	Customer               int64                  `json:"-"`
	CustomerFull           *Customer              `json:"-"`
	CustomerRaw            json.RawMessage        `json:"customer"`
	Date                   int64                  `json:"date"`
	DisabledPaymentMethods []string               `json:"disabled_payment_methods"`
	Discounts              []Discount             `json:"discounts"`
	Draft                  bool                   `json:"draft"`
	DueDate                int64                  `json:"due_date"`
	Id                     int64                  `json:"id"`
	Items                  []LineItem             `json:"items"`
	Metadata               map[string]interface{} `json:"metadata"`
	Name                   string                 `json:"name"`
	NextPaymentAttempt     int64                  `json:"next_payment_attempt"`
	Notes                  string                 `json:"notes"`
	Number                 string                 `json:"number"`
	Object                 string                 `json:"object"`
	Paid                   bool                   `json:"paid"`
	PaymentPlan            int64                  `json:"payment_plan"`
	PaymentTerms           string                 `json:"payment_terms"`
	PaymentUrl             string                 `json:"payment_url"`
	PdfUrl                 string                 `json:"pdf_url"`
	PurchaseOrder          string                 `json:"purchase_order"`
	Sent                   bool                   `json:"sent"`
	ShipTo                 *ShippingDetail        `json:"ship_to"`
	Status                 string                 `json:"status"`
	Subscription           int64                  `json:"subscription"`
	Subtotal               float64                `json:"subtotal"`
	Taxes                  []Tax                  `json:"taxes"`
	Total                  float64                `json:"total"`
	UpdatedAt              int64                  `json:"updated_at"`
	Url                    string                 `json:"url"`
}

type Invoices []*Invoice

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
