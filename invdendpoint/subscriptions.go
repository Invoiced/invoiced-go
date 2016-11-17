package invdendpoint

import (
	"encoding/json"
)

const SubscriptionsEndPoint = "/subscriptions"

type Subscriptions []Subscription

type Subscription struct {
	Id          int64                  `json:"id,omitempty"`           //The subscription’s unique ID
	Customer    int64                  `json:"customer,omitempty"`     //Associated Customer
	Plan        string                 `json:"plan,omitempty"`         //Plan ID
	StartDate   int64                  `json:"start_date,omitempty"`   //Timestamp subscription starts (or started)
	Quantity    int64                  `json:"quantity,omitempty"`     //Plan quantity. Defaults to 1
	Cycles      int64                  `json:"cycles,omitempty"`       //Number of billing cycles the subscription runs for, when null runs until canceled (default).
	PeriodStart int64                  `json:"period_start,omitempty"` //Start of the current billing period
	PeriodEnd   int64                  `json:"period_end,omitempty"`   //End of the current billing period
	Status      string                 `json:"status,omitempty"`       //Subscription status, one of not_started, active, past_due, finished
	Addons      []SubscriptionAddon    `json:"addons,omitempty"`       //Collection of Subscription Addons
	Discounts   []Discount             `json:"discount,omitempty"`     //Collection of Coupon IDs
	Taxes       []Tax                  `json:"taxes,omitempty"`        //Collection of Tax Rate IDs
	Url         string                 `json:"url,omitempty"`          //URL to manage the subscription in the billing portal
	CreatedAt   int64                  `json:"created_at,omitempty"`   //Timestamp when created
	MetaData    map[string]interface{} `json:"metadata,omitempty"`     //A hash of key/value pairs that can store additional information about this object.
}

func (s *Subscription) String() string {

	b, _ := json.MarshalIndent(s, "", "    ")

	return string(b)
}
