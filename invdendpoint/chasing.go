package invdendpoint

const ChasingCadenceEndpoint = "/chasing_cadences"

type ChasingCadence struct {
	AssignmentConditions string        `json:"assignment_conditions,omitempty"`
	AssignmentMode       string        `json:"assignment_mode,omitempty"`
	CreatedAt            int64         `json:"created_at,omitempty"`	//Timestamp when created
	UpdatedAt            int64         `json:"updated_at,omitempty"`
	Frequency            string        `json:"frequency,omitempty"`
	Id                   int64         `json:"id,omitempty"`
	LastRun              int64         `json:"last_run,omitempty"`
	MinBalance           float64       `json:"min_balance,omitempty"`
	Name                 string        `json:"name,omitempty"`
	NextRun              int64         `json:"nextrun,omitempty"`
	NumCustomers         int64         `json:"num_customers,omitempty"`
	Object               string        `json:"object,omitempty"`
	Paused               bool          `json:"paused,omitempty"`
	RunDate              int64         `json:"run_date,omitempty"`
	Steps                []ChasingStep `json:"steps,omitempty"`
	TimeOfDay            int           `json:"time_of_day,omitempty"`
}

type ChasingStep struct {
	Action          string `json:"action,omitempty"`
	AssignedUserId  int64  `json:"assigned_user_id,omitempty"`
	CreatedAt       int64  `json:"created_at,omitempty"`	//Timestamp when created
	EmailTemplateId string `json:"email_template_id,omitempty"`
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Schedule        string `json:"schedule,omitempty"`
	SmsTemplateId   string `json:"sms_template_id"`
	UpdatedAt       int64  `json:"updated_at,omitempty"`
}
