package invdendpoint

//Metered billing on Invoiced allows you to bill customers for charges that occur during a billing cycle outside of their ordinary subscription. These charges are called pending line items. A pending line item is a Line Item that has been attached to a customer, but not billed yet. Pending line items will be swept up by the next invoice that is triggered for the customer. This happens automatically with subscription invoices or when triggering an invoice manually.
type PendingLineItem struct {
	Id           int64                  `json:"id,omitempty"`
	CatalogItem  string                 `json:"catalog_item,omitempty"`  //Optional Catalog Item ID. Fills the line item with the name and pricing of the Catalog Item.
	Type         string                 `json:"type,omitempty"`          //Optional line item type. Used to group line items by type in reporting
	Name         string                 `json:"name,omitempty"`         //Title
	Description  string                 `json:"description,omitempty"`  //Optional description
	Quantity     float64                `json:"quantity,omitempty"`     //Quantity
	UnitCost     float64                `json:"unit_cost,omitempty"`    //Unit cost or rate
	Discountable bool                   `json:"discountable,omitempty"` //Excludes amount from invoice discounts when false, defaults to `true
	Discounts    []Discount             `json:"discounts,omitempty"`    //Line item Discounts
	Taxable      bool                   `json:"taxable,omitempty"`      //Excludes amount from invoice taxes when false, defaults to `true
	Taxes        []Tax                  `json:"taxes,omitempty"`        //Line item Taxes
	Metadata     map[string]interface{} `json:"metadata,omitempty"`      //A hash of key/value pairs that can store additional information about this object.
}
