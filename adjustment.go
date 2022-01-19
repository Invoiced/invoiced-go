package invoiced

const CreditBalanceAdjustmentsEndpoint = "/credit_balance_adjustments"

type BalanceAdjustmentRequest struct {
	Amount   *float64 `json:"amount,omitempty"`
	Currency *string  `json:"currency,omitempty"`
	Customer *int64   `json:"customer,omitempty"`
	Date     *int64   `json:"date,omitempty"`
	Notes    *string  `json:"notes,omitempty"`
}

type BalanceAdjustment struct {
	Amount    float64 `json:"amount,omitempty"`
	CreatedAt int64   `json:"created_at,omitempty"`
	Currency  string  `json:"currency,omitempty"`
	Customer  int64   `json:"customer,omitempty"`
	Date      int64   `json:"date,omitempty"`
	ID        int64   `json:"id,omitempty"`
	Notes     string  `json:"notes,omitempty"`
	Object    string  `json:"object,omitempty"`
	UpdatedAt int64   `json:"updated_at,omitempty"`
}
