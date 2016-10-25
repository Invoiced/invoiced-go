package invdendpoint

import "encoding/json"

// type Invoices struct {
// 	Invoices []Invoice `json:"invoices"`
// }

const InvoicesEndPoint = "/invoices"

type Invoices []Invoice

//Todo Add Tags Object and Attachements Objects
type Invoice struct {
	Id             int64  `json:"id,omitempty"`              //The invoice’s unique ID
	Customer       int64  `json:"customer,omitempty"`        //Customer ID
	Name           string `json:"name,omitempty"`            //Invoice name for internal use, defaults to “Invoice”
	Number         string `json:"number,omitempty"`          //The reference number assigned to the invoice for use in the dashboard
	CollectionMode string `json:"collection_mode,omitempty"` //Invoice collection mode, auto or manual
	Currency       string `json:"currency,omitempty"`        //3-letter ISO code
	Draft          bool   `json:"draft,omitempty"`           //When false, the invoice is considered outstanding, or when true, the invoice is a draft
	Closed         bool   `json:"closed,omitempty"`          //When true, an invoice is closed and considered bad debt. No further payments are allowed.

	Paid               bool   `json:"paid,omitempty"`                 //Indicates whether an invoice has been paid in full
	Status             string `json:"status,omitempty"`               //Invoice state, one of draft, not_sent, sent, viewed, past_due, pending, paid
	Chase              bool   `json:"chase,omitempty"`                //Whether chasing is enabled for the invoice
	NextChaseOn        int64  `json:"next_chase_on,omitempty"`        //Next scheduled chase
	AttemptCount       int64  `json:"attempt_count,omitempty"`        //# of payment attempts
	NextPaymentAttempt int64  `json:"next_payment_attempt,omitempty"` //Next scheduled charge attempt, when in automatic collection

	Subscription int64  `json:"subscription,omitempty"`  //Subscription ID if invoice came from subscription
	Date         int64  `json:"date,omitempty"`          //Invoice date
	DueDate      int64  `json:"due_date,omitempty"`      //Date payment is due by
	PaymentTerms string `json:"payment_terms,omitempty"` //Payment terms for the invoice, i.e. “NET 30”

	Items     []LineItem `json:"items,omitempty"`     //Collection of Line Items
	Notes     string     `json:"notes,omitempty"`     //Additional notes displayed on invoice
	Subtotal  float64    `json:"subtotal,omitempty"`  //Subtotal
	Discounts []Discount `json:"discounts,omitempty"` //Collection of Discounts
	Taxes     []Tax      `json:"taxes,omitempty"`     //Collection of Taxes

	Total   float64 `json:"total,omitempty"`   //Total
	Balance float64 `json:"balance,omitempty"` //Balance owed

	Tags       []string               `json:"tags,omitempty"`        //Invoice tags
	Url        string                 `json:"url,omitempty"`         //URL to view the invoice in the billing portal
	PaymentUrl string                 `json:"payment_url,omitempty"` //URL for the invoice payment page
	PdfUrl     string                 `json:"pdf_url,omitempty"`     //URL to download the invoice as a PDF
	CreatedAt  int64                  `json:"created_at,omitempty"`  //Timestamp when created
	MetaData   map[string]interface{} `json:"metadata,omitempty"`    //A hash of key/value pairs that can store additional information about this object.

	// add disabled payment methods
}

func (i *Invoice) String() string {

	b, _ := json.MarshalIndent(i, "", "    ")

	return string(b)
}
