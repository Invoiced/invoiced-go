package invdendpoint

type Balance struct {
	Currency         string           `json:"currency,omitempty"`
	AvailableCredits float64          `json:"available_credits,omitempty"`
	PastDue          bool             `json:"past_due,omitempty"`
	Histories        BalanceHistories `json:"history,omitempty"`
	TotalOutstanding float64          `json:"total_outstanding,omitempty"`
	DueNow           float64          `json:"due_now,omitempty"`
}

type BalanceHistories []BalanceHistory

type BalanceHistory struct {
	Currency  string `json:"currency,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Balance   int64  `json:"balance,omitempty"`
}
