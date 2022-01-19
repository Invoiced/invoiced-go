package invoiced

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
)

func Bool(v bool) *bool {
	return &v
}

func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

func Float64(v float64) *float64 {
	return &v
}

func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

func Int64(v int64) *int64 {
	return &v
}

func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

func String(v string) *string {
	return &v
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// CreateFormFile is a convenience wrapper around CreatePart. It creates
// a new form-data header with the provided field name and file name.
func (c *Client) CreateFormFile(w *multipart.Writer, fieldname, filename string, fileType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", fileType)
	return w.CreatePart(h)
}
