package invoiced

type ChasingCadence struct {
	AssignmentConditions *string       `json:"assignment_conditions"`
	AssignmentMode       string        `json:"assignment_mode"`
	CreatedAt            int64         `json:"created_at"`
	UpdatedAt            int64         `json:"updated_at"`
	Frequency            string        `json:"frequency"`
	Id                   int64         `json:"id"`
	LastRun              *int64        `json:"last_run"`
	MinBalance           *float64      `json:"min_balance"`
	Name                 string        `json:"name"`
	NextRun              *int64        `json:"nextrun"`
	NumCustomers         int64         `json:"num_customers"`
	Object               string        `json:"object"`
	Paused               bool          `json:"paused"`
	RunDate              int64         `json:"run_date"`
	Steps                []ChasingStep `json:"steps"`
	TimeOfDay            int64         `json:"time_of_day"`
}

type ChasingStep struct {
	Action          string  `json:"action"`
	AssignedUserId  *int64  `json:"assigned_user_id"`
	CreatedAt       int64   `json:"created_at"`
	EmailTemplateId *string `json:"email_template_id"`
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Schedule        string  `json:"schedule"`
	SmsTemplateId   *string `json:"sms_template_id"`
	UpdatedAt       int64   `json:"updated_at"`
}

type ChasingCadences []*ChasingCadence
