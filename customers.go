package invoiced

import (
	"encoding/json"
)

type CustomerRequest struct {
	Address1               *string                 `json:"address1,omitempty"`
	Address2               *string                 `json:"address2,omitempty"`
	AttentionTo            *string                 `json:"attention_to,omitempty"`
	AutoPay                *bool                   `json:"autopay,omitempty"`
	AutoPayDelays          *int64                  `json:"autopay_delay_days,omitempty"`
	AvalaraEntityUseCode   *string                 `json:"avalara_entity_use_code,omitempty"`
	AvalaraExemptionNumber *string                 `json:"avalara_exemption_number,omitempty"`
	BillToParent           *bool                   `json:"bill_to_parent,omitempty"`
	Chase                  *bool                   `json:"boolean,omitempty"`
	ChasingCadence         *int64                  `json:"chasing_cadence,omitempty"`
	City                   *string                 `json:"city,omitempty"`
	Country                *string                 `json:"country,omitempty"`
	CreatedAt              *int64                  `json:"created_at,omitempty"`
	CreditHold             *bool                   `json:"credit_hold,omitempty"`
	CreditLimit            *float64                `json:"credit_limit,omitempty"`
	Currency               *string                 `json:"currency,omitempty"`
	DisabledPaymentMethods []*string               `json:"disabled_payment_methods,omitempty"`
	Email                  *string                 `json:"email,omitempty"`
	Id                     *int64                  `json:"id,omitempty"`
	Language               *string                 `json:"language,omitempty"`
	Metadata               *map[string]interface{} `json:"metadata,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
	NextChaseStep          *int64                  `json:"next_chase_step,omitempty"`
	Notes                  *string                 `json:"notes,omitempty"`
	Number                 *string                 `json:"number,omitempty"`
	Object                 *string                 `json:"object,omitempty"`
	Owner                  *int64                  `json:"owner,omitempty"`
	ParentCustomer         *int64                  `json:"parent_customer,omitempty"`
	PaymentSource          *PaymentSource          `json:"payment_source,omitempty"`
	PaymentTerms           *string                 `json:"payment_terms,omitempty"`
	Phone                  *string                 `json:"phone,omitempty"`
	PostalCode             *string                 `json:"postal_code,omitempty"`
	SignUpPage             *int64                  `json:"sign_up_page,omitempty"`
	SignUpUrl              *string                 `json:"sign_up_url,omitempty"`
	State                  *string                 `json:"state,omitempty"`
	StatementPdfUrl        *string                 `json:"statement_pdf_url,omitempty"`
	TaxId                  *string                 `json:"taxid,omitempty"`
	Taxable                *bool                   `json:"taxable,omitempty"`
	Taxes                  []*TaxRate              `json:"taxes,omitempty"`
	Type                   *string                 `json:"type,omitempty"`
	UpdatedAt              *int64                  `json:"updated_at,omitempty"`
}

type Customers []*Customer

type Customer struct {
	Address1               string                 `json:"address1"`
	Address2               string                 `json:"address2"`
	AttentionTo            string                 `json:"attention_to"`
	AutoPay                bool                   `json:"autopay"`
	AutoPayDelays          int64                  `json:"autopay_delay_days"`
	AvalaraEntityUseCode   string                 `json:"avalara_entity_use_code"`
	AvalaraExemptionNumber string                 `json:"avalara_exemption_number"`
	BillToParent           bool                   `json:"bill_to_parent"`
	Chase                  bool                   `json:"boolean"`
	ChasingCadence         int64                  `json:"chasing_cadence"`
	City                   string                 `json:"city"`
	Country                string                 `json:"country"`
	CreatedAt              int64                  `json:"created_at"`
	CreditHold             bool                   `json:"credit_hold"`
	CreditLimit            float64                `json:"credit_limit"`
	Currency               string                 `json:"currency"`
	DisabledPaymentMethods []string               `json:"disabled_payment_methods"`
	Email                  string                 `json:"email"`
	Id                     int64                  `json:"id"`
	Language               string                 `json:"language"`
	Metadata               map[string]interface{} `json:"metadata"`
	Name                   string                 `json:"name"`
	NextChaseStep          int64                  `json:"next_chase_step"`
	Notes                  string                 `json:"notes"`
	Number                 string                 `json:"number"`
	Object                 string                 `json:"object"`
	Owner                  int64                  `json:"owner"`
	ParentCustomer         int64                  `json:"parent_customer"`
	PaymentSource          *PaymentSource         `json:"payment_source"`
	PaymentTerms           string                 `json:"payment_terms"`
	Phone                  string                 `json:"phone"`
	PostalCode             string                 `json:"postal_code"`
	SignUpPage             int64                  `json:"sign_up_page"`
	SignUpUrl              string                 `json:"sign_up_url"`
	State                  string                 `json:"state"`
	StatementPdfUrl        string                 `json:"statement_pdf_url"`
	TaxId                  string                 `json:"taxid"`
	Taxable                bool                   `json:"taxable"`
	Taxes                  []TaxRate              `json:"taxes"`
	Type                   string                 `json:"type"`
	UpdatedAt              int64                  `json:"updated_at"`
}

func (c *Customer) String() string {
	b, _ := json.MarshalIndent(c, "", "    ")

	return string(b)
}
