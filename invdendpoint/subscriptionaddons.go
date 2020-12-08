package invdendpoint

type SubscriptionAddon struct {
	Id        int64  `json:"id,omitempty"`           // The subscription’s unique ID
	Item      string `json:"catalog_item,omitempty"` // Item ID
	Plan      string `json:"plan,omitempty"`         // The Subscription's Plan ID
	Quantity  int64  `json:"quantity,omitempty"`     // Quantity
	CreatedAt int64  `json:"created_at,omitempty"`   // Timestamp when created
}
