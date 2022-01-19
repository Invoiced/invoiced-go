package invoiced

type Balance struct {
	AvailableCredits float64          `json:"available_credits"`
	Currency         string           `json:"currency"`
	DueNow           float64          `json:"due_now"`
	Histories        BalanceHistories `json:"history"`
	PastDue          bool             `json:"past_due"`
	TotalOutstanding float64          `json:"total_outstanding"`
}

type BalanceHistory struct {
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
	Timestamp int64   `json:"timestamp"`
}

type BalanceHistories []*BalanceHistory
