package invdendpoint

const CreditBalanceAdjustmentsEndpoint = "/credit_balance_adjustments"

type BalanceAdjustment struct {
	Amount    float64 `json:"amount,omitempty"`
	CreatedAt int     `json:"created_at,omitempty"`	//Timestamp when created
	Currency  string  `json:"currency,omitempty"`
	Customer  int     `json:"customer,omitempty"`
	Date      int     `json:"date,omitempty"`
	ID        int     `json:"id,omitempty"`
	Notes     string  `json:"notes,omitempty"`
	Object    string  `json:"object,omitempty"`
	UpdatedAt int     `json:"updated_at,omitempty"`
}
