package invdendpoint

const PlansEndPoint = "/plans/"

type Plan struct {
	Plan PlanData `json:"plan,omitempty"`
}

type Plans struct {
	Plans []PlanData `json:"plans,omitempty"`
}

type PlanData struct {
	CreatedAt               int64  `json:"created_at,omitempty"`
	UpdatedAt               int64  `json:"updated_at,omitempty"`
	Id                      int64  `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	Theme                   int64  `json:"theme,omitempty"`
	Chase                   bool   `json:"chase,omitempty"`
	Amount                  bool   `json:"amount,omitempty"`
	Interval                string `json:"interval,omitempty"`
	IntervalCount           int64  `json:"interval_count,omitempty"`
	Description             string `json:"description,omitempty"`
	Type                    string `json:"type,omitempty"`
	Note                    string `json:"note,omitempty"`
	Terms                   string `json:"terms,omitempty"`
	SendInvoiceAfterRenewal string `json:"send_invoice_after_renewal,omitempty"`
	DisabledPaymentMethods  string `json:"disabled_payment_methods,omitempty"`
}
