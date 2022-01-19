package invdendpoint

type SendStatementTextMessageRequest struct {
	To      []*TextMessageRecipient `json:"to,omitempty"`
	Message *string                 `json:"message,omitempty"`
	Type    *string                 `json:"type,omitempty"`
	Start   *int64                  `json:"start,omitempty"`
	End     *int64                  `json:"end,omitempty"`
	Items   *string                 `json:"items,omitempty"`
}

type SendTextMessageRequest struct {
	To      []*TextMessageRecipient `json:"to,omitempty"`
	Message *string                 `json:"message,omitempty"`
}

type TextMessageRecipient struct {
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

type TextMessages []TextMessage

type TextMessage struct {
	Id        string `json:"id"`
	To        string `json:"to"`
	State     string `json:"state"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
