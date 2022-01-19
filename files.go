package invoiced

const FileEndpoint = "/files"

type FileRequest struct {
	Name *string `json:"name,omitempty"`
	Size *int64  `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
	Url  *string `json:"url,omitempty"`
}

type File struct {
	CreatedAt int64  `json:"created_at"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Object    string `json:"object"`
	Size      int64  `json:"size"`
	Type      string `json:"type"`
	UpdatedAt int64  `json:"updated_at"`
	Url       string `json:"url"`
}

type Files []File
