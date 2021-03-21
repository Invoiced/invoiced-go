package invdendpoint

const NotificationEndpoint = "/notifications"

type NotificationResponses []NotificationResponse

type NotificationRequest struct {
	Enabled         bool                `json:"enabled,omitempty"`
	Event  			string  			`json:"event,omitempty"`
	Medium          string              `json:"medium,omitempty"`
	Role            string              `json:"role,omitempty"`
	UserId           int64              `json:"user_id,omitempty"`
}

type NotificationResponse struct {
	Conditions       string  `json:"conditions,omitempty"`
	Enabled          bool	 `json:"enabled,omitempty"`
	Event  			 string  `json:"event,omitempty"`
	Id               int64   `json:"id,omitempty"`
	MatchMode        string  `json:"match_mode,omitempty"`
	Medium           string  `json:"medium,omitempty"`
	UserId           int64   `json:"user_id,omitempty"`
}

type Notification struct {
	Conditions       string  `json:"conditions,omitempty"`
	Enabled          bool	 `json:"enabled,omitempty"`
	Event  			 string  `json:"event,omitempty"`
	Id               int64   `json:"id,omitempty"`
	MatchMode        string  `json:"match_mode,omitempty"`
	Medium           string  `json:"medium,omitempty"`
	UserId           int64   `json:"user_id,omitempty"`
}
