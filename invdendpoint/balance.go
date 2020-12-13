package invdendpoint

type Balance struct {
	AvailableCredits int64            `json:"available_credits,omitempty"`
	PastDue          bool             `json:"past_due,omitempty"`
	Histories        BalanceHistories `json:"history,omitempty"`
	TotalOutstanding int64            `json:"total_outstanding,omitempty"`
}

type BalanceHistories []BalanceHistory

type BalanceHistory struct {
	Timestamp int64 `json:"timestamp,omitempty"`
	Balance   int64 `json:"balance,omitempty"`
}
