package invoiced

import (
	"encoding/json"
	"strconv"
)

type PaymentRequest struct {
	Amount    *float64               `json:"amount,omitempty"`
	AppliedTo []*PaymentItemRequest  `json:"applied_to,omitempty"`
	Currency  *string                `json:"currency,omitempty"`
	Customer  *int64                 `json:"-"`
	Date      *int64                 `json:"date,omitempty"`
	Method    *string                `json:"method,omitempty"`
	Notes     *string                `json:"notes,omitempty"`
	Metadata  map[string]interface{} `json:"metadata"`
	Reference *string                `json:"reference,omitempty"`
	Source    *string                `json:"source,omitempty"`
	Voided    *bool                  `json:"voided,omitempty"`
}

type PaymentItemRequest struct {
	Amount       *float64 `json:"amount,omitempty"`
	CreditNote   *int64   `json:"credit_note,omitempty"`
	DocumentType *string  `json:"document_type,omitempty"`
	Estimate     *int64   `json:"estimate,omitempty"`
	Invoice      *int64   `json:"invoice,omitempty"`
	Type         *string  `json:"type,omitempty"`
}

type Payment struct {
	Amount       float64                `json:"amount"`
	AppliedTo    []PaymentItem          `json:"applied_to"`
	Balance      float64                `json:"balance"`
	Charge       *Charge                `json:"charge"`
	CreatedAt    int64                  `json:"created_at"`
	Currency     string                 `json:"currency"`
	Customer     int64                  `json:"-"`
	CustomerFull *Customer              `json:"-"`
	CustomerRaw  json.RawMessage        `json:"customer"`
	Date         int64                  `json:"date"`
	Id           int64                  `json:"id"`
	Matched      bool                   `json:"matched"`
	Metadata     map[string]interface{} `json:"metadata"`
	Method       string                 `json:"method"`
	Notes        string                 `json:"notes"`
	Object       string                 `json:"object"`
	PdfUrl       string                 `json:"pdf_url"`
	Reference    string                 `json:"reference"`
	Source       string                 `json:"source"`
	Status       string                 `json:"status"`
	UpdatedAt    int64                  `json:"updated_at"`
	Voided       bool                   `json:"voided"`
}

type PaymentItem struct {
	Amount       float64 `json:"amount"`
	CreditNote   int64   `json:"credit_note"`
	DocumentType string  `json:"document_type"`
	Estimate     int64   `json:"estimate"`
	Invoice      int64   `json:"invoice"`
	Type         string  `json:"type"`
}

type Payments []*Payment

func (i *Payment) UnmarshalJSON(data []byte) error {
	type payment2 Payment

	if err := json.Unmarshal(data, (*payment2)(i)); err != nil {
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

func (i *Payment) MarshalJSON() ([]byte, error) {
	type payment2 Payment
	i2 := (*payment2)(i)

	if i2.Customer > 0 {
		i2.CustomerRaw = []byte(strconv.FormatInt(i2.Customer, 10))
	}

	return json.Marshal(i2)
}

func (i *Payment) String() string {
	b, _ := json.MarshalIndent(i, "", "    ")

	return string(b)
}
