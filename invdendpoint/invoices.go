package invdendpoint

// type Invoices struct {
// 	Invoices []Invoice `json:"invoices"`
// }

const InvoicesEndPoint = "/invoices"

type Invoices []Invoice

type Invoice struct {
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	Id        int64  `json:"id,omitempty"`
	Customer  int64  `json:"customer,omitempty"`
	Name      string `json:"name,omitempty"`
	//Currency               string         `json:"currency,omitempty"`
	Draft  bool `json:"draft,omitempty"`
	Sent   bool `json:"sent,omitempty"`
	Closed bool `json:"closed,omitempty"`
	Paid   bool `json:"paid,omitempty"`
	//Status                 string         `json:"status,omitempty"`
	Chase              bool  `json:"chase,omitempty"`
	NextChaseOn        int64 `json:"next_chase_on,omitempty"`
	AutoBilled         bool  `json:"auto_billed,omitempty"`
	AttemptCount       int64 `json:"attempt_count,omitempty"`
	NextPaymentAttempt int64 `json:"next_payment_attempt,omitempty"`
	//Theme                  string         `json:"theme,omitempty"`
	DisabledPaymentMethods *PaymentMethod `json:"disabled_payment_methods,omitempty"`
	Subscription           *int64         `json:"subscription,omitempty"`
	Number                 string         `json:"number,omitempty"`
	Date                   int64          `json:"date,omitempty"`
	DueDate                int64          `json:"due_date,omitempty"`
	PaymentTerms           string         `json:"payment_terms,omitempty"`
	//PurchaseOrder          string         `json:"purchaseorder,omitempty"`
	Items     []LineItem `json:"items,omitempty"`
	Terms     string     `json:"terms,omitempty"`
	Notes     string     `json:"notes,omitempty"`
	Subtotal  float64    `json:"subtotal,omitempty"`
	Discounts []Discount `json:"discounts,omitempty"`
	Taxes     []Tax      `json:"taxes,omitempty"`

	Total          float64 `json:"total,omitempty"`
	AmountPaid     float64 `json:"amountpaid,omitempty"`
	AmountAdjusted float64 `json:"amountadjusted,omitempty"`
	Balance        float64 `json:"balance,omitempty"`
	// Url                          string      `json:"url,omitempty"`
	// PdfUrl                       string      `json:"pdf_url,omitempty"`
	// CSVUrl                       string      `json:"csv_url,omitempty"`
	LatePaymentRemindersDisabled bool        `json:"late_payment_reminders_disabled,omitempty"`
	MetaData                     interface{} `json:"metadata,omitempty"`
	//add disabled payment methods
}
