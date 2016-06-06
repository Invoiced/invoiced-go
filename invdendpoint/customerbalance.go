package invdendpoint

type CustomerBalance struct {
	AvailableCredits int64                    `json:"available_credits"`
	PastDue          bool                     `json:"past_due"`
	Histories        CustomerBalanceHistories `json:"history"`
	TotalOutstanding int64                    `json:"total_outstanding"`
}

type CustomerBalanceHistories []CustomerBalanceHistory

type CustomerBalanceHistory struct {
	Timestamp int64 `json:"timestamp"`
	Balance   int64 `json:"balance"`
}
