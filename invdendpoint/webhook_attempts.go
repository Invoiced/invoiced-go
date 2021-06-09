package invdendpoint

import "encoding/json"

const WebhookEndpoint = "/webhook_attempts"

type WebhookAttempt struct {
	Attempts  []WebhookAttemptStatus `json:"attempts,omitempty"`
	CreatedAt int64                  `json:"created_at,omitempty"`
	EventId   int64                  `json:"event_id,omitempty"`
	Id        int64                  `json:"id,omitempty"`
	Payload   json.RawMessage        `json:"payload,omitempty"`
}

type WebhookAttemptStatus struct {
	StatusCode int `json:"status_code,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
}
