package invoiced

const TaskEndpoint = "/tasks"

type TaskRequest struct {
	Action   *string `json:"action,omitempty"`
	Complete *bool   `json:"complete,omitempty"`
	Customer *int64  `json:"customer_id,omitempty"`
	DueDate  *int64  `json:"due_date,omitempty"`
	Name     *string `json:"name,omitempty"`
	User     *int64  `json:"user_id,omitempty"`
}

type Tasks []Task

type Task struct {
	Action          string `json:"action"`
	ChaseStep       int64  `json:"chase_step_id"`
	Complete        bool   `json:"complete"`
	CompletedByUser int64  `json:"completed_by_user_id"`
	CompletedDate   int64  `json:"completed_date"`
	CreatedAt       int64  `json:"created_at"`
	Customer        int64  `json:"customer_id"`
	DueDate         int64  `json:"due_date"`
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	UpdatedAt       int64  `json:"updated_at"`
	User            int64  `json:"user_id"`
}
