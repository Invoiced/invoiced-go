package invdendpoint

const ItemEndPoint = "/items"

//An item represents a product or service that you sell. Items can be used to generate line items and can also be used as subscription addons.

type Item struct {
	Id             string                 `json:"id,omitempty"`               //The item’s unique ID
	Object         string                 `json:"object,omitempty"`           //Object name
	Name           string                 `json:"name,omitempty"`
	Currency       string                 `json:"currency,omitempty"`
	UnitCost       float64                `json:"unit_cost,omitempty"`
	Description    string                 `json:"description,omitempty"`      //Optional description
	Type           string                 `json:"service,omitempty"`          //Optional line item type. Used to group line items by type in reporting
	Taxable        bool                   `json:"taxable,omitempty"`          //Excludes amount from taxes when false
	Taxes          []Tax                  `json:"taxes,omitempty"`            //Collection of Tax Rate Objects
	AvalaraTaxCode string                 `json:"avalara_tax_code,omitempty"` //Avalara-specific tax code
	GlAccount      string                 `json:"gl_account,omitempty"`       //General ledger account code
	Discountable   bool                   `json:"discountable,omitempty"`     //Excludes amount from discounts when false
	CreatedAt      int64                  `json:"created_at,omitempty"`       //Timestamp when created
	Metadata       map[string]interface{} `json:"metadata,omitempty"`         //A hash of key/value pairs that can store additional information about this object.
}

type Items []Item
