package invdendpoint

const FilesEndPoint = "/files"

//Represents an external file.
type File struct {
	Id        int64  `json:"id"`                   //The file’s unique ID
	Object    string `json:"object"`               //file
	Name      string `json:"name,omitempty`        //Filename
	Size      int    `json:"size,omitempty`        //File size in bytes
	Type      string `json:"type,omitempty"`       //MIME Type
	Url       string `json:"url,omitempty"`        //File URL
	CreatedAt int64  `json:"created_at,omitempty"` //Timestamp when created
}

type Files []File
