package invdendpoint

const NoteEndpoint = "/notes"

type Notes []Note

type Note struct {
	Id        int64  `json:"id,omitempty"`         // The note’s unique ID
	Object    string `json:"object,omitempty"`     // Object type, note
	Notes     string `json:"notes,omitempty"`      // Contents of note
	Customer  int64  `json:"customer,omitempty"`   // Customer associated with note
	CreatedAt int64  `json:"created_at,omitempty"` // Timestamp when created
	User      *User  `json:"user,omitempty"`       // Object describing user who created note
}

type CreateNoteRequest struct {
	CustomerID int64  `json:"customer_id,omitempty"` // Associated customer’s unique ID
	InvoiceID  int64  `json:"invoice_id,omitempty"`  // Associated invoice’s unique ID
	Notes      string `json:"notes,omitempty"`       // Contents of Notes
}
