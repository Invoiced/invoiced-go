package invdapi

import (
	"errors"
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

func (c *Note) Create(createNoteRequest invdendpoint.CreateNoteRequest) (*Note, error) {
	endpoint := invdendpoint.NoteEndpoint

	noteResp := new(Note)

	apiErr := c.create(endpoint, createNoteRequest, noteResp)

	if apiErr != nil {
		return nil, apiErr
	}

	noteResp.Connection = c.Connection

	return noteResp, nil
}

func (c *Note) Save() error {
	endpoint := invdendpoint.NoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	noteResp := new(Note)

	noteDataToUpdate, err := SafeNoteForUpdating(c.Note)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, noteDataToUpdate, noteResp)

	if apiErr != nil {
		return apiErr
	}

	c.Note = noteResp.Note

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

	endpointTmp, apiErr := c.retrieveDataFromAPI(endpoint, &tmpNotes)

	if apiErr != nil {
		return nil, apiErr
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

// SafeCustomerForCreation prunes note data for just fields that can be used for creation of a note
func SafeNoteForUpdating(note *invdendpoint.Note) (*invdendpoint.Note, error) {
	if note == nil {
		return nil, errors.New("task is nil")
	}

	noteData := new(invdendpoint.Note)
	noteData.Notes = note.Notes

	return noteData, nil
}
