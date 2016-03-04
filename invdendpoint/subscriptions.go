package invdendpoint

const SubscriptionsEndPoint = "/subscriptions"

type Subscription struct {
	Id        int64  `json:"id,omitempty"`
	Customer  int64  `json:"customer,omitempty"`
	Plan      string `json:"plan,omitempty"`
	StartDate int64  `json:"start_date,omitempty"`

	Quantity int64 `json:"quantity,omitempty"`
	Cycles   int64 `json:"cycles,omitempty"`

	RenewsNext  int64 `json:"renews_next,omitempty"`
	RenewedLast int64 `json:"renewed_last,omitempty"`

	Status string `json:"status,omitempty"`

	Addons    []SubscriptionAddon `json:"addons,omitempty"`
	Discounts []Discount          `json:"discount,omitempty"`
	Taxes     []Tax               `json:"taxes,omitempty"`
	Url       string              `json:"url,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}
