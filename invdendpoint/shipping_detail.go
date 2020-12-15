package invdendpoint

type ShippingDetail struct {
	Name        string `json:"name,omitempty"`
	AttentionTo string `json:"attention_to,omitempty"` // Used for ATTN: address line if company
	Address1    string `json:"address1,omitempty"`     // First address line
	Address2    string `json:"address2,omitempty"`     // Optional second address line
	City        string `json:"city,omitempty"`         // City
	State       string `json:"state,omitempty"`        // State or province
	PostalCode  string `json:"postal_code,omitempty"`  // Zip or postal code
	Country     string `json:"country,omitempty"`      // Two-letter ISO code
}
