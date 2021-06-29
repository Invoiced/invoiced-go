package invdendpoint

import (
	"encoding/json"
	"strconv"
)

const SubscriptionEndpoint = "/subscriptions"

type Subscriptions []Subscription

type Subscription struct {
	Id                    int64                  `json:"id,omitempty"`         // The subscription’s unique ID
	Object                string                 `json:"object,omitempty"`     // Object type, subscription
	Customer               int64                  `json:"-"`
	CustomerFull           *Customer              `json:"-"`
	CustomerRaw             json.RawMessage        `json:"customer,omitempty"`
	Plan                  string                 `json:"plan,omitempty"`       // Plan ID
	StartDate             int64                  `json:"start_date,omitempty"` // Timestamp subscription starts (or started)
	BillIn                string                 `json:"bill_in,omitempty"`    // advance or arrears. Defaults to advance
	BillInAdvanceDays     int64                  `json:"bill_in_advance_days,omitempty"`
	Quantity              int64                  `json:"quantity,omitempty"`                // Plan quantity. Defaults to 1
	Cycles                int64                  `json:"cycles,omitempty"`                  // Number of billing cycles the subscription runs for, when null runs until canceled (default).
	PeriodStart           int64                  `json:"period_start,omitempty"`            // Start of the current billing period
	PeriodEnd             int64                  `json:"period_end,omitempty"`              // End of the current billing period
	SnapToNthDay          int                    `json:"snap_to_nth_day,omitempty"`         // Snap billing cycles to a specific day of the month (also known as calendar billing), off by default
	CancelAtPeriodEnd     bool                   `json:"cancel_at_period_end,omitempty"`    // When true the subscription will be canceled at the end of the current billing period
	CanceledAt            int64                  `json:"cancel_at,omitempty"`               // Timestamp the subscription was canceled
	Prorate               bool                   `json:"prorate,omitempty"`                 // Prorate changes to plan, quantities, or addons, defaults to true
	ProrationDate         int64                  `json:"proration_date,omitempty"`          // Timestamp when the proration happened, defaults to now
	Paused                bool                   `json:"paused"`                            // When true, subscription is paused
	ContractPeriodStart   int64                  `json:"contract_period_start,omitempty"`   // Start of current contract period
	ContractPeriodEnd     int64                  `json:"contract_period_end,omitempty"`     // End of current contract period
	ContractRenewalCycles int                    `json:"contract_renewal_cycles,omitempty"` // Number of billing cycles in next contract period
	ContractRenewalMode   string                 `json:"contract_renewal_mode,omitempty"`   // auto, manual, renew_once, or none. Defaults to none
	Status                string                 `json:"status,omitempty"`                  // Subscription status, one of not_started, active, past_due, finished
	RecurringTotal        float64                `json:"recurring_total,omitempty"`         // Total recurring amount (includes taxes)
	Mrr                   float64                `json:"MRR,omitempty"`                     // Monthly Recurring Revenue (MRR)`
	Addons                []SubscriptionAddon    `json:"addons,omitempty"`                  // Collection of Subscription Addons
	Discounts             []Discount             `json:"discount,omitempty"`                // Collection of Coupon IDs
	Taxes                 []Tax                  `json:"taxes,omitempty"`                   // Collection of Tax Rate ID
	ShipTo                *ShippingDetail        `json:"ship_to,omitempty"`
	Url                   string                 `json:"url,omitempty"`        // URL to manage the subscription in the billing portal
	CreatedAt             int64                  `json:"created_at,omitempty"` // Timestamp when created
	Metadata              map[string]interface{} `json:"metadata,omitempty"`   // A hash of key/value pairs that can store additional information about this object.
}

func (s *Subscription) String() string {
	b, _ := json.MarshalIndent(s, "", "    ")
	return string(b)
}

type SubscriptionPreviewRequest struct {
	Customer         int64               `json:"customer,omitempty"`
	Plan             string              `json:"plan,omitempty"`
	Quantity         int                 `json:"quantity,omitempty"`
	Addons           []SubscriptionAddon `json:"addons,omitempty"`
	Discounts        []Discount          `json:"discounts,omitempty"`
	Taxes            []Tax               `json:"Taxes,omitempty"`
	PendingLineItems []PendingLineItem   `json:"pending_line_item,omitempty"`
}

type SubscriptionPreview struct {
	FirstInvoice   *SubscriptionPreviewInvoice `json:"first_invoice,omitempty"`
	MRR            float64                     `json:"mrr,omitempty"`
	RecurringTotal float64                     `json:"recurring_total,omitempty"`
}

type SubscriptionPreviewInvoice struct {
	Customer           int64                  `json:"customer,omitempty"`             // Customer ID
	Name               string                 `json:"name,omitempty"`                 // Invoice name for internal use, defaults to “Invoice”
	Number             string                 `json:"number,omitempty"`               // The reference number assigned to the invoice for use in the dashboard
	AutoPay            bool                   `json:"autopay,omitempty"`              // Invoice collection mode, auto or manual
	Currency           string                 `json:"currency,omitempty"`             // 3-letter ISO code
	Draft              bool                   `json:"draft,omitempty"`                // When false, the invoice is considered outstanding, or when true, the invoice is a draft
	Closed             bool                   `json:"closed,omitempty"`               // When true, an invoice is closed and considered bad debt. No further payments are allowed.
	Paid               bool                   `json:"paid,omitempty"`                 // Indicates whether an invoice has been paid in full
	Status             string                 `json:"status,omitempty"`               // Invoice state, one of draft, not_sent, sent, viewed, past_due, pending, paid
	AttemptCount       int64                  `json:"attempt_count,omitempty"`        //# of payment attempts
	NextPaymentAttempt int64                  `json:"next_payment_attempt,omitempty"` // Next scheduled charge attempt, when in automatic collection
	Date               int64                  `json:"date,omitempty"`                 // Invoice date
	DueDate            int64                  `json:"due_date,omitempty"`             // Date payment is due by
	PaymentTerms       string                 `json:"payment_terms,omitempty"`        // Payment terms for the invoice, i.e. “NET 30”
	Items              []LineItemPreview      `json:"items,omitempty"`                // Collection of Line Items
	Notes              string                 `json:"notes,omitempty"`                // Additional notes displayed on invoice
	Subtotal           float64                `json:"subtotal,omitempty"`             // Subtotal
	Discounts          []Discount             `json:"discounts,omitempty"`            // Collection of Discounts
	Taxes              []Tax                  `json:"taxes,omitempty"`                // Collection of Taxes
	Total              float64                `json:"total,omitempty"`                // Total
	Balance            float64                `json:"balance,omitempty"`              // Balance owed
	Url                string                 `json:"url,omitempty"`                  // URL to view the invoice in the billing portal
	PaymentUrl         string                 `json:"payment_url,omitempty"`          // URL for the invoice payment page
	PdfUrl             string                 `json:"pdf_url,omitempty"`              // URL to download the invoice as a PDF
	CreatedAt          int64                  `json:"created_at,omitempty"`           // Timestamp when created
	Metadata           map[string]interface{} `json:"metadata,omitempty"`             // A hash of key/value pairs that can store additional information about this object.
}

func (i *Subscription) UnmarshalJSON(data []byte) error {
	type subscription2 Subscription

	if err := json.Unmarshal(data, (*subscription2)(i)); err != nil {
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

func (i *Subscription) MarshalJSON() ([]byte, error) {
	type subscription2 Subscription
	i2 := (*subscription2)(i)

	if i2.Customer > 0 {
		i2.CustomerRaw = []byte(strconv.FormatInt(i2.Customer, 10))
	}

	return json.Marshal(i2)
}