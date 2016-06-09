package invdendpoint

type PaymentMethod struct {
	Ach          bool `json:"ach,omitempty"`
	Bitcoin      bool `json:"bitcoin,omitempty"`
	Cash         bool `json:"cash,omitempty"`
	Check        bool `json:"check,omitempty"`
	CreditCard   bool `json:"credit_card,omitempty"`
	Other        bool `json:"other,omitempty"`
	Paypal       bool `json:"paypal,omitempty"`
	WireTransfer bool `json:"wire_transfer,omitempty"`
}
