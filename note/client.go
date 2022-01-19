package note

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

const NoteEndpoint = "/notes"

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.NoteRequest) (*Client, error) {
	endpoint := NoteEndpoint
	resp := new(Client)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Update(request *invoiced.NoteRequest) error {
	endpoint := NoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Note = resp.Note

	return nil
}

func (c *Client) Delete() error {
	endpoint := NoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Notes, error) {
	endpoint := NoteEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

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
