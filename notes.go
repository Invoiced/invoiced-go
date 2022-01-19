package invoiced

type NoteRequest struct {
	Customer *int64  `json:"customer_id,omitempty"`
	Invoice  *int64  `json:"invoice_id,omitempty"`
	Notes    *string `json:"notes,omitempty"`
}

type Note struct {
	CreatedAt int64  `json:"created_at"`
	Customer  int64  `json:"customer"`
	Id        int64  `json:"id"`
	Notes     string `json:"notes"`
	Object    string `json:"object"`
	UpdatedAt int64  `json:"updated_at"`
	User      *User  `json:"user"`
}

type Notes []*Note
