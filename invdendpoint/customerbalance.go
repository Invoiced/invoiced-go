package invdendpoint

type CustomerBalance struct {
	AvailableCredits int64                    `json:"available_credits,omitempty"`
	PastDue          bool                     `json:"past_due,omitempty"`
	Histories        CustomerBalanceHistories `json:"history,omitempty"`
	TotalOutstanding int64                    `json:"total_outstanding,omitempty"`
}

type CustomerBalanceHistories []CustomerBalanceHistory

type CustomerBalanceHistory struct {
	Timestamp int64 `json:"timestamp,omitempty"`
	Balance   int64 `json:"balance,omitempty"`
}
