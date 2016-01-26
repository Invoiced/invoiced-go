package invdendpoint

type SubscriptionAddon struct {
	Id          int64  `json:"id,omitempty"`
	CatalogItem string `json:"catalog_item,omitempty"`
	Quantity    int64  `json:"quantity,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}
