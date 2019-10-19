package invdendpoint

import (
	"encoding/json"
)

const SubscriptionsEndPoint = "/subscriptions"

type Subscriptions []Subscription

type Subscription struct {
	Id                    int64                  `json:"id,omitempty"`                      //The subscription’s unique ID
	Object                string                 `json:"object,omitempty"`                  //Object type, subscription
	Customer              int64                  `json:"customer,omitempty"`                //Associated Customer
	Plan                  string                 `json:"plan,omitempty"`                    //Plan ID
	StartDate             int64                  `json:"start_date,omitempty"`              //Timestamp subscription starts (or started)
	BillIn                string                 `json:"bill_in,omitempty"`                 //advance or arrears. Defaults to advance
	Quantity              int64                  `json:"quantity,omitempty"`                //Plan quantity. Defaults to 1
	Addons                []SubscriptionAddon    `json:"addons,omitempty"`                  //Collection of Subscription Addons
	Discounts             []Discount             `json:"discount,omitempty"`                //Collection of Coupon IDs
	Cycles                int64                  `json:"cycles,omitempty"`                  //Number of billing cycles the subscription runs for, when null runs until canceled (default).
	PeriodStart           int64                  `json:"period_start,omitempty"`            //Start of the current billing period
	PeriodEnd             int64                  `json:"period_end,omitempty"`              //End of the current billing period
	CancelAtPeriodEnd     bool                   `json:"cancel_at_period_end"`              //When true the subscription will be canceled at the end of the current billing period
	CanceledAt            int64                  `json:"cancel_at"`                         //Timestamp the subscription was canceled
	Status                string                 `json:"status,omitempty"`                  //Subscription status, one of not_started, active, past_due, finished
	Paused                bool                   `json:"paused,omitempty"`                  //When true, subscription is paused
	SnapToNthDay          int                    `json:"snap_to_nth_day,omitempty"`         //Snap billing cycles to a specific day of the month (also known as calendar billing), off by default
	ContractPeriodStart   int64                  `json:"contract_period_start,omitempty"`   //Start of current contract period
	ContractPeriodEnd     int64                  `json:"contract_period_end,omitempty"`     //End of current contract period
	ContractRenewalCycles int                    `json:"contract_renewal_cycles,omitempty"` //Number of billing cycles in next contract period
	ContractRenewalMode   string                 `json:"contract_renewal_mode,omitempty"`   //auto, manual, renew_once, or none. Defaults to none
	Taxes                 []Tax                  `json:"taxes,omitempty"`                   //Collection of Tax Rate ID
	RecurringTotal        float64                `json:"recurring_total,omitempty"`         //Total recurring amount (includes taxes)
	Mrr                   float64                `json:"MRR,omitempty"`                     //Monthly Recurring Revenue (MRR)`
	Url                   string                 `json:"url,omitempty"`                     //URL to manage the subscription in the billing portal
	CreatedAt             int64                  `json:"created_at,omitempty"`              //Timestamp when created
	MetaData              map[string]interface{} `json:"metadata,omitempty"`                //A hash of key/value pairs that can store additional information about this object.
	Prorate               bool                   `json:"prorate,omitempty"`					//Prorate changes to plan, quantities, or addons, defaults to true
	ProrationDate        int64                    `json:"prorartiondate,omitempty"`			//Timestamp when the proration happened, defaults to now

}

func (s *Subscription) String() string {
	b, _ := json.MarshalIndent(s, "", "    ")
	return string(b)
}
