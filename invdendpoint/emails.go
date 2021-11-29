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
