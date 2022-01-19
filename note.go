package invdapi

import (
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Note struct {
	*Connection
	*invdendpoint.Note
}

type Notes []*Note

func (c *Connection) NewNote() *Note {
	note := new(invdendpoint.Note)
	return &Note{c, note}
}

func (c *Note) Create(request *invdendpoint.NoteRequest) (*Note, error) {
	endpoint := invdendpoint.NoteEndpoint
	resp := new(Note)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Note) Update(request *invdendpoint.NoteRequest) error {
	endpoint := invdendpoint.NoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Note)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Note = resp.Note

	return nil
}

func (c *Note) Delete() error {
	endpoint := invdendpoint.NoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Note) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Notes, error) {
	endpoint := invdendpoint.NoteEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	notes := make(Notes, 0)

NEXT:
	tmpNotes := make(Notes, 0)

	endpointTmp, err := c.retrieveDataFromAPI(endpoint, &tmpNotes)

	if err != nil {
		return nil, err
	}

	notes = append(notes, tmpNotes...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, note := range notes {
		note.Connection = c.Connection
	}

	return notes, nil
}
