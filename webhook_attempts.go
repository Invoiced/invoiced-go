package invoiced

import "encoding/json"

type WebhookAttempt struct {
	Attempts  []WebhookAttemptStatus `json:"attempts,omitempty"`
	CreatedAt int64                  `json:"created_at,omitempty"`
	EventId   int64                  `json:"event_id,omitempty"`
	Id        int64                  `json:"id,omitempty"`
	Payload   json.RawMessage        `json:"payload,omitempty"`
	UpdatedAt int64                  `json:"updated_at,omitempty"`
}

type WebhookAttemptStatus struct {
	StatusCode int64 `json:"status_code,omitempty"`
	Timestamp  int64 `json:"timestamp,omitempty"`
}

type WebhookAttempts []*WebhookAttempt
