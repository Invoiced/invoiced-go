package invdendpoint

import (
	"encoding/json"
	"strconv"
)

const PaymentEndpoint = "/payments"

type Payments []Payment

type Payment struct {
	Id        int64         `json:"id,omitempty"`       // The payment’s unique ID
	Object    string        `json:"object,omitempty"`   // Object type, payment
	Customer               int64                  `json:"-"`
	CustomerFull           *Customer              `json:"-"`
	CustomerRaw             json.RawMessage        `json:"customer,omitempty"`
	Date      int64         `json:"date,omitempty"`     // Payment date, defaults to current timestamp
	Method    string        `json:"method,omitempty"`   // Payment instrument used to facilitate payment, defaults to other
	Matched   bool          `json:"matched,omitempty"`
	Voided    bool          `json:"voided,omitempty"`
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
