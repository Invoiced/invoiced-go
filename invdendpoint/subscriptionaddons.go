package invdendpoint

type SubscriptionAddon struct {
	Id          int64       `json:"id,omitempty"`           //The subscription’s unique ID
	CatalogItem CatalogItem `json:"catalog_item,omitempty"` //Catalog Item ID
	Quantity    int64       `json:"quantity,omitempty"`     //Quantity
	CreatedAt   int64       `json:"created_at,omitempty"`   //Timestamp when created
}
