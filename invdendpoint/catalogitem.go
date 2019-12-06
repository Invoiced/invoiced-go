package invdendpoint

//A catalog item represents a product or service that you sell. Catalog items can be used to generate line items and can also be used as subscription addons.

type CatalogItem struct {
	Id           string                 `json:"id,omitempty"`           //The customer’s unique ID
	Object       string                 `json:"object,omitempty"`       //Contact name
	Name         string                 `json:"name,omitempty"`         //Email address
	Currency     string                 `json:"currency,omitempty"`     //When true the contact will be copied on any account communications
	UnitCost     float64                `json:"unit_cost,omitempty"`    //First address line
	Description  string                 `json:"description,omitempty"`  //Optional description
	Type         string                 `json:"service,omitempty"`      //Optional line item type. Used to group line items by type in reporting
	Taxable      bool                   `json:"taxable,omitempty"`      //Excludes amount from taxes when false
	Taxes        []Tax                  `json:"taxes,omitempty"`        //Collection of Tax Rate Objects
	AvalaraTaxCode string               `json:"avalara_tax_code,omitempty"`  //Avalara-specific tax code
	GlAccount      string				`json:"gl_account,omitempty"`     //General ledger account code
	Discountable bool                   `json:"discountable,omitempty"` //Excludes amount from discounts when false
	CreatedAt    int64                  `json:"created_at,omitempty"`   //Timestamp when created
	MetaData     map[string]interface{} `json:"metadata,omitempty"`     //A hash of key/value pairs that can store additional information about this object.
}

type CatalogItems []CatalogItem
