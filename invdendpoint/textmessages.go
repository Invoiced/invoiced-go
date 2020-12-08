package invdendpoint

type TextRequest struct {
	To      []TextDetail `json:"to,omitempty"`
	Message string       `json:"message,omitempty"`
	Type    string       `json:"type,omitempty"`
	Start   int64        `json:"start,omitempty"`
	End     int64        `json:"end,omitempty"`
	Items   string       `json:"items,omitempty"`
}

type TextDetail struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type TextResponses []TextResponse

type TextResponse struct {
	Id        string `json:"id,omitempty"`
	To        string `json:"to,omitempty"`
	State     string `json:"state,omitempty"`
	Message   string `json:"message,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
