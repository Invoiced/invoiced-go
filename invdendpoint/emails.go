package invdendpoint

type EmailRequest struct {
	To       []EmailDetail `json:"to,omitempty"`
	Bcc      string        `json:"bcc,omitempty"`
	Subject  string        `json:"subject,omitempty"`
	Message  string        `json:"message,omitempty"`
	Template string        `json:"template,omitempty"`
	Type     string        `json:"type,omitempty"`
	Start    int64         `json:"start,omitempty"`
	End      int64         `json:"end,omitempty"`
	Items    string        `json:"items,omitempty"`
}

type EmailDetail struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type EmailResponses []EmailResponse

func (e EmailResponses) Error() string {
	panic("implement me")
}

type EmailResponse struct {
	Id           int64  `json:"id,omitempty"`
	State        string `json:"state,omitempty"`
	RejectReason string `json:"reject_reason,omitempty"`
	Email        string `json:"email,omitempty"`
	Template     string `json:"template,omitempty"`
	Subject      string `json:"subject,omitempty"`
	Message      string `json:"message,omitempty"`
	Opens        int64  `json:"opens,omitempty"`
	Clicks       int64  `json:"clicks,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
}
