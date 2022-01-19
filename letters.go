package invoiced

type SendStatementLetterRequest struct {
	End   *int64  `json:"end,omitempty"`
	Items *string `json:"items,omitempty"`
	Start *int64  `json:"start,omitempty"`
	Type  *string `json:"type,omitempty"`
}

type Letter struct {
	CreatedAt            int64  `json:"created_at"`
	ExpectedDeliveryDate int64  `json:"expected_delivery_date"`
	Id                   string `json:"id"`
	NumPages             int64  `json:"num_pages"`
	State                string `json:"state"`
	To                   string `json:"to"`
	UpdatedAt            int64  `json:"updated_at"`
}
