package invdendpoint

type SubscriptionAddonRequest struct {
	Amount    *float64 `json:"amount,omitempty"`
	CreatedAt *int64   `json:"created_at,omitempty"`
	Id        *int64   `json:"id,omitempty"`
	Plan      *string  `json:"plan,omitempty"`
	Quantity  *float64 `json:"quantity,omitempty"`
}

type SubscriptionAddon struct {
	Amount    float64 `json:"amount"`
	CreatedAt int64   `json:"created_at"`
	Id        int64   `json:"id"`
	Plan      string  `json:"plan"`
	Quantity  float64 `json:"quantity"`
}
