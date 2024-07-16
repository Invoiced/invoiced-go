package invoiced

type Import struct {
	ID            int64    `json:"id"`
	CreatedAt     int64    `json:"created_at"`
	UpdatedAt     int64    `json:"updated_at"`
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Status        string   `json:"status"`
	Position      int      `json:"position"`
	NumImported   int      `json:"num_imported"`
	NumUpdated    int      `json:"num_updated"`
	NumFailed     int      `json:"num_failed"`
	TotalRecords  int      `json:"total_records"`
	Message       string   `json:"message"`
	FailureDetail []string `json:"failure_detail"`
	User          int64    `json:"user"`
	UserID        int64    `json:"user_id"`
	Object        string   `json:"object"`
}
