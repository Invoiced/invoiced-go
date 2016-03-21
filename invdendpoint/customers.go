package invdendpoint

const CustomersEndPoint = "/customers"

type Customers []Customer

type Customer struct {
	Id              int64          `json:"id,omitempty"`
	Name            string         `json:"name,omitempty"`
	Number          string         `json:"number,omitempty"`
	Type            string         `json:"type,omitempty"`
	Email           string         `json:"email,omitempty"`
	CollectionMode  string         `json:"collection_mode,omitempty"`
	PaymentTerms    string         `json:"payment_terms,omitempty"`
	PaymentSource   *PaymentSource `json:"payment_source,omitempty"`
	AttentionTo     string         `json:"attention_to,omitempty"`
	Address1        string         `json:"address1,omitempty"`
	Address2        string         `json:"address2,omitempty"`
	City            string         `json:"city,omitempty"`
	State           string         `json:"state,omitempty"`
	PostalCode      string         `json:"postal_code,omitempty"`
	Country         string         `json:"country,omitempty"`
	TaxID           string         `json:"taxid,omitempty"`
	Phone           string         `json:"phone,omitempty"`
	OtherPhone      string         `json:"other_phone,omitempty"`
	Notes           string         `json:"notes,omitempty"`
	StatementPdfUrl string         `json:"statement_pdf_url,omitempty"`
	CreatedAt       int64          `json:"created_at,omitempty"`
	UpdatedAt       int64          `json:"updated_at,omitempty"`
	MetaData        interface{}    `json:"metadata,omitempty"`
}

type PaymentSource struct {
	Id       int64  `json:"id,omitempty"`
	Brand    string `json:"brand,omitempty"`
	Last4    int64  `json:"last4,omitempty"`
	ExpMonth int64  `json:"exp_month,omitempty"`
	ExpYear  int64  `json:"exp_year,omitempty"`
	Funding  string `json:"funding,omitempty"`
}
