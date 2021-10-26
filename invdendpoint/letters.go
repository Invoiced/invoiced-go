package invdendpoint

type LetterRequest struct {
	Type  string `json:"type,omitempty"`
	Start int64  `json:"start,omitempty"`
	End   int64  `json:"end,omitempty"`
	Items string `json:"items,omitempty"`
}

type LetterResponse struct {
	Id                   string `json:"id,omitempty"`
	To                   string `json:"to,omitempty"`
	State                string `json:"state,omitempty"`
	NumPages             int64  `json:"num_pages,omitempty"`
	ExpectedDeliveryDate int64  `json:"expected_delivery_date,omitempty"`
	CreatedAt            int64  `json:"created_at,omitempty"`	//Timestamp when created
	UpdatedAt            int64  `json:"updated_at,omitempty"`
}
