package invoiced

import (
	"strconv"
)

type NoteClient struct {
	*Client
	*Note
}

type Notes []*NoteClient

func (c *Client) NewNote() *NoteClient {
	note := new(Note)
	return &NoteClient{c, note}
}

func (c *NoteClient) Create(request *NoteRequest) (*NoteClient, error) {
	endpoint := NoteEndpoint
	resp := new(NoteClient)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *NoteClient) Update(request *NoteRequest) error {
	endpoint := NoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(NoteClient)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Note = resp.Note

	return nil
}

func (c *NoteClient) Delete() error {
	endpoint := NoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *NoteClient) ListAll(filter *Filter, sort *Sort) (Notes, error) {
	endpoint := NoteEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

	notes := make(Notes, 0)

NEXT:
	tmpNotes := make(Notes, 0)

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
