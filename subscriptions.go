package invoiced

import (
	"encoding/json"
	"strconv"
	"strings"
)

type SubscriptionRequest struct {
	Addons                []*SubscriptionAddonRequest `json:"addons,omitempty"`
	Amount                *float64                    `json:"amount,omitempty"`
	BillIn                *string                     `json:"bill_in,omitempty"`
	BillInAdvanceDays     *int64                      `json:"bill_in_advance_days,omitempty"`
	CancelAtPeriodEnd     *bool                       `json:"cancel_at_period_end,omitempty"`
	ContractPeriodEnd     *int64                      `json:"contract_period_end,omitempty"`
	ContractPeriodStart   *int64                      `json:"contract_period_start,omitempty"`
	ContractRenewalCycles *int64                      `json:"contract_renewal_cycles,omitempty"`
	ContractRenewalMode   *string                     `json:"contract_renewal_mode,omitempty"`
	Customer              *int64                      `json:"customer,omitempty"`
	Cycles                *int64                      `json:"cycles,omitempty"`
	Discounts             []*DiscountRequest          `json:"discount,omitempty"`
	Metadata              *map[string]interface{}     `json:"metadata,omitempty"`
	Paused                *bool                       `json:"paused,omitempty"`
	PeriodEnd             *int64                      `json:"period_end,omitempty"`
	Plan                  *string                     `json:"plan,omitempty"`
	Prorate               *bool                       `json:"prorate,omitempty"`
	ProrationDate         *int64                      `json:"proration_date,omitempty"`
	Quantity              *float64                    `json:"quantity,omitempty"`
	ShipTo                *ShippingDetail             `json:"ship_to,omitempty"`
	SnapToNthDay          *int64                      `json:"snap_to_nth_day,omitempty"`
	StartDate             *int64                      `json:"start_date,omitempty"`
	Taxes                 []*TaxRequest               `json:"taxes,omitempty"`
}

type Subscription struct {
	Addons                []SubscriptionAddon    `json:"addons"`
	Amount                float64                `json:"amount"`
	BillIn                string                 `json:"bill_in"`
	BillInAdvanceDays     int64                  `json:"bill_in_advance_days"`
	CancelAtPeriodEnd     bool                   `json:"cancel_at_period_end"`
	CanceledAt            int64                  `json:"cancel_at"`
	ContractPeriodEnd     int64                  `json:"contract_period_end"`
	ContractPeriodStart   int64                  `json:"contract_period_start"`
	ContractRenewalCycles int64                  `json:"contract_renewal_cycles"`
	ContractRenewalMode   string                 `json:"contract_renewal_mode"`
	CreatedAt             int64                  `json:"created_at"`
	Customer              int64                  `json:"-"`
	CustomerFull          *Customer              `json:"-"`
	CustomerRaw           json.RawMessage        `json:"customer"`
	Cycles                int64                  `json:"cycles"`
	Discounts             []Discount             `json:"discount"`
	Id                    int64                  `json:"id"`
	Metadata              map[string]interface{} `json:"metadata"`
	Mrr                   float64                `json:"MRR"`
	Object                string                 `json:"object"`
	Paused                bool                   `json:"paused"`
	PeriodEnd             int64                  `json:"period_end"`
	PeriodStart           int64                  `json:"period_start"`
	Plan                  string                 `json:"-"`
	PlanFull              *Plan                  `json:"-"`
	PlanRaw               json.RawMessage        `json:"plan"`
	Prorate               bool                   `json:"prorate"`
	Quantity              float64                `json:"quantity"`
	RecurringTotal        float64                `json:"recurring_total"`
	RenewsNext            int64                  `json:"renews_next"`
	ShipTo                *ShippingDetail        `json:"ship_to"`
	StartDate             int64                  `json:"start_date"`
	Status                string                 `json:"status"`
	Taxes                 []Tax                  `json:"taxes"`
	UpdatedAt             int64                  `json:"updated_at"`
	Url                   string                 `json:"url"`
}

type Subscriptions []*Subscription

func (s *Subscription) String() string {
	b, _ := json.MarshalIndent(s, "", "    ")
	return string(b)
}

type SubscriptionPreviewRequest struct {
	Addons           []*SubscriptionAddonRequest `json:"addons,omitempty"`
	Customer         *int64                      `json:"customer,omitempty"`
	Discounts        []*DiscountRequest          `json:"discounts,omitempty"`
	PendingLineItems []*PendingLineItemRequest   `json:"pending_line_item,omitempty"`
	Plan             *string                     `json:"plan,omitempty"`
	Quantity         *float64                    `json:"quantity,omitempty"`
	Taxes            []*TaxRequest               `json:"Taxes,omitempty"`
}

type SubscriptionPreview struct {
	FirstInvoice   *SubscriptionPreviewInvoice `json:"first_invoice"`
	MRR            float64                     `json:"mrr"`
	RecurringTotal float64                     `json:"recurring_total"`
}

type SubscriptionPreviewInvoice struct {
	AttemptCount       int64                  `json:"attempt_count"`
	AutoPay            bool                   `json:"autopay"`
	Balance            float64                `json:"balance"`
	Closed             bool                   `json:"closed"`
	CreatedAt          int64                  `json:"created_at"`
	Currency           string                 `json:"currency"`
	Customer           int64                  `json:"customer"`
	Date               int64                  `json:"date"`
	Discounts          []Discount             `json:"discounts"`
	Draft              bool                   `json:"draft"`
	DueDate            int64                  `json:"due_date"`
	Items              []LineItemPreview      `json:"items"`
	Metadata           map[string]interface{} `json:"metadata"`
	Name               string                 `json:"name"`
	NextPaymentAttempt int64                  `json:"next_payment_attempt"`
	Notes              string                 `json:"notes"`
	Number             string                 `json:"number"`
	Paid               bool                   `json:"paid"`
	PaymentTerms       string                 `json:"payment_terms"`
	PaymentUrl         string                 `json:"payment_url"`
	PdfUrl             string                 `json:"pdf_url"`
	Status             string                 `json:"status"`
	Subtotal           float64                `json:"subtotal"`
	Taxes              []Tax                  `json:"taxes"`
	Total              float64                `json:"total"`
	UpdatedAt          int64                  `json:"updated_at"`
	Url                string                 `json:"url"`
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

	aj := i.PlanRaw

	i.Plan = string(aj)
	i.Plan = strings.Trim(i.Plan, "\"")

	plan := new(Plan)

	err = json.Unmarshal(aj, plan)

	if err == nil {
		i.PlanFull = plan
		i.Plan = plan.Id
	}

	return nil
}

func (i *Subscription) MarshalJSON() ([]byte, error) {
	type subscription2 Subscription
	i2 := (*subscription2)(i)

	if i2.Customer > 0 {
		i2.CustomerRaw = []byte(strconv.FormatInt(i2.Customer, 10))
	}

	if len(i2.Plan) > 0 {
		b, err := json.Marshal(i2.Plan)

		if err != nil {
			return nil, err
		}

		i2.PlanRaw = b
	}

	return json.Marshal(i2)
}
