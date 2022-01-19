package invoiced

const NotificationEndpoint = "/notifications"

type Notifications []Notification

type NotificationRequest struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Event   *string `json:"event,omitempty"`
	Medium  *string `json:"medium,omitempty"`
	Role    *string `json:"role,omitempty"`
	User    *int64  `json:"user_id,omitempty"`
}

type Notification struct {
	Conditions string `json:"conditions"`
	Enabled    bool   `json:"enabled"`
	Event      string `json:"event"`
	Id         int64  `json:"id"`
	MatchMode  string `json:"match_mode"`
	Medium     string `json:"medium"`
	User       int64  `json:"user_id"`
}
