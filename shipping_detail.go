package invoiced

type ShippingDetailRequest struct {
	Address1    *string `json:"address1,omitempty"`
	Address2    *string `json:"address2,omitempty"`
	AttentionTo *string `json:"attention_to,omitempty"`
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	Name        *string `json:"name,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty"`
	State       *string `json:"state,omitempty"`
}

type ShippingDetail struct {
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	AttentionTo string `json:"attention_to"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Name        string `json:"name"`
	PostalCode  string `json:"postal_code"`
	State       string `json:"state"`
}
