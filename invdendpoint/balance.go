package invdendpoint

type Balance struct {
	Currency         string           `json:"currency"`
	AvailableCredits float64          `json:"available_credits"`
	PastDue          bool             `json:"past_due"`
	Histories        BalanceHistories `json:"history"`
	TotalOutstanding float64          `json:"total_outstanding"`
	DueNow           float64          `json:"due_now"`
}

type BalanceHistories []BalanceHistory

type BalanceHistory struct {
	Currency  string  `json:"currency"`
	Timestamp int64   `json:"timestamp"`
	Balance   float64 `json:"balance"`
}
