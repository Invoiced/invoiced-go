package note

import (
	"github.com/Invoiced/invoiced-go/v2"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.NoteRequest) (*invoiced.Note, error) {
	resp := new(invoiced.Note)
	err := c.Api.Create("/roles", request, resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.NoteRequest) (*invoiced.Note, error) {
	resp := new(invoiced.Note)
	err := c.Api.Update("/notes/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/notes/" + strconv.FormatInt(id, 10))
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Notes, error) {
	endpoint := invoiced.AddFilterAndSort("/notes", filter, sort)

	notes := make(invoiced.Notes, 0)

NEXT:
	tmpNotes := make(invoiced.Notes, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpNotes)

	if err != nil {
		return nil, err
	}

	notes = append(notes, tmpNotes...)

	if endpointTmp != "" {
		goto NEXT
	}

	return notes, nil
}
