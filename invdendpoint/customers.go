package invdendpoint

import (
	"encoding/json"
)

const CustomersEndPoint = "/customers"

type Customers []Customer

type Customer struct {
	Id                     int64                  `json:"id,omitempty"` // The customer’s unique ID
	Object                 string                 `json:"object,omitempty"`
	Name                   string                 `json:"name,omitempty"`               // Customer name
	Number                 string                 `json:"number,omitempty"`             // A unique ID to help tie your customer to your external systems
	Email                  string                 `json:"email,omitempty"`              // Email address
	AutoPay                bool                   `json:"autopay,omitempty"`            // Autopay
	AutoPayDelays          int                    `json:"autopay_delay_days,omitempty"` // Number of days to delay AutoPay
	PaymentTerms           string                 `json:"payment_terms,omitempty"`      // Payment terms used for manual collection mode, i.e. “NET 30”
	StripeToken            string                 `json:"stripe_token,omitempty"`       // When provided sets the customer’s payment source to the tokenized Stripe card
	PaymentSource          *PaymentSource         `json:"payment_source,omitempty"`     // Customer’s payment source, if attached
	AttentionTo            string                 `json:"attention_to,omitempty"`       // Used for ATTN: address line if company
	Address1               string                 `json:"address1,omitempty"`           // First address line
	Address2               string                 `json:"address2,omitempty"`           // Optional second address line
	City                   string                 `json:"city,omitempty"`               // City
	State                  string                 `json:"state,omitempty"`              // State or province
	PostalCode             string                 `json:"postal_code,omitempty"`        // Zip or postal code
	Country                string                 `json:"country,omitempty"`            // Two-letter ISO code
	Language               string                 `json:"language,omitempty"`           // Two-letter ISO code
	Chase                  bool                   `json:"boolean,omitempty"`            // Chasing enabled? - defaults to true
	ChasingCadence         int64                  `json:"chasing_cadence,omitempty"`    // Cadence ID
	NextChaseStep          int64                  `json:"next_chase_step,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`                    // Phone #
	CreditHold             bool                   `json:"credit_hold,omitempty"`              // When true, customer is on credit hold
	CreditLimit            float64                `json:"credit_limit,omitempty"`             // Customer credit limit
	Owner                  int64                  `json:"owner,omitempty"`                    // Customer credit limit
	Taxable                bool                   `json:"taxable,omitempty"`                  // Customer taxable?
	Taxes                  []TaxRate              `json:"taxes,omitempty"`                    // Collection of Tax Rate IDs
	TaxId                  string                 `json:"taxid,omitempty"`                    // Tax ID to be displayed on documents
	AvalaraEntityUseCode   string                 `json:"avalara_entity_use_code,omitempty"`  // Avalara-specific entity use code
	AvalaraExemptionNumber string                 `json:"avalara_exemption_number,omitempty"` // Tax-exempt number to pass to Avalara
	Type                   string                 `json:"type,omitempty"`                     // Organization type, company or person
	ParentCustomer         int64                  `json:"parent_customer,omitempty"`          // Parent customer ID
	Notes                  string                 `json:"notes,omitempty"`                    // Private customer notes
	SignUpPage             int64                  `json:"sign_up_page,omitempty"`
	SignUpUrl              string                 `json:"sign_up_url,omitempty"`       // URL to download the latest account statement
	StatementPdfUrl        string                 `json:"statement_pdf_url,omitempty"` // URL to download the latest account statement
	CreatedAt              int64                  `json:"created_at,omitempty"`        // Timestamp when created
	Metadata               map[string]interface{} `json:"metadata,omitempty"`          // A hash of key/value pairs that can store additional information about this object.
	DisabledPaymentMethods []string               `json:"disabled_payment_methods,omitempty"`
}

func (c *Customer) String() string {
	b, _ := json.MarshalIndent(c, "", "    ")

	return string(b)
}
