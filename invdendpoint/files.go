package invdendpoint

const FileEndpoint = "/files"

type File struct {
	Id        int64  `json:"id,omitempty"`         // The file’s unique ID
	Object    string `json:"object,omitempty"`     // file
	Name      string `json:"name,omitempty"`       // Filename
	Size      int    `json:"size,omitempty"`       // File size in bytes
	Type      string `json:"type,omitempty"`       // MIME Type
	Url       string `json:"url,omitempty"`        // File URL
	CreatedAt int64  `json:"created_at,omitempty"`	//Timestamp when created
	UpdatedAt int64  `json:"updated_at,omitempty"` // Timestamp when updated
}

type Files []File
