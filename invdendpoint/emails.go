package invdendpoint

type SendStatementEmailRequest struct {
	Bcc      *string           `json:"bcc,omitempty"`
	End      *int64            `json:"end,omitempty"`
	Items    *string           `json:"items,omitempty"`
	Message  *string           `json:"message,omitempty"`
	Start    *int64            `json:"start,omitempty"`
	Subject  *string           `json:"subject,omitempty"`
	Template *string           `json:"template,omitempty"`
	To       []*EmailRecipient `json:"to,omitempty"`
	Type     *string           `json:"type,omitempty"`
}

type SendEmailRequest struct {
	Bcc      *string           `json:"bcc,omitempty"`
	Message  *string           `json:"message,omitempty"`
	Subject  *string           `json:"subject,omitempty"`
	Template *string           `json:"template,omitempty"`
	To       []*EmailRecipient `json:"to,omitempty"`
}

type EmailRecipient struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}
