package invdendpoint

const TaskEndpoint = "/tasks"

type Tasks []Task

// Represents an task
type Task struct {
	Id                int64  `json:"id,omitempty"`                   // The note’s unique ID
	Name              string `json:"name,omitempty"`                 // Name
	Action            string `json:"action,omitempty"`               // Action type, one of phone, letter, email, review
	CustomerId        int64  `json:"customer_id,omitempty"`          // Associated customer ID
	UserId            int64  `json:"user_id,omitempty"`              // ID of user assigned to task
	DueDate           int64  `json:"due_date,omitempty"`             // Task due date
	Complete          bool   `json:"complete,omitempty"`             // Task complete?
	CompletedDate     int64  `json:"completed_date,omitempty"`       // Date task was marked complete
	CompletedByUserId int64  `json:"completed_by_user_id,omitempty"` // User ID who completed task
	ChaseStepID       int64  `json:"chase_step_id,omitempty"`        // Chasing step ID that created task
	CreatedAt         int64  `json:"created_at,omitempty"`           // Timestamp when created

}
