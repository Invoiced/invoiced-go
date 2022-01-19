package invdapi

import "C"
import (
	"fmt"
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type CreditNote struct {
	*Connection
	*invdendpoint.CreditNote
}

type CreditNotes []*CreditNote

func (c *Connection) NewCreditNote() *CreditNote {
	creditNote := new(invdendpoint.CreditNote)
	return &CreditNote{c, creditNote}
}

func (c *CreditNote) Count() (int64, error) {
	count, err := c.count(invdendpoint.CreditNoteEndpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *CreditNote) Create(request *invdendpoint.CreditNoteRequest) (*CreditNote, error) {
	resp := new(CreditNote)

	err := c.create(invdendpoint.CreditNoteEndpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *CreditNote) Retrieve(id int64) (*CreditNote, error) {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(id, 10)

	creditNote := &CreditNote{c.Connection, new(invdendpoint.CreditNote)}

	_, err := c.retrieveDataFromAPI(endpoint, creditNote)

	if err != nil {
		return nil, err
	}

	return creditNote, nil
}

func (c *CreditNote) Update(request *invdendpoint.CreditNoteRequest) error {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(CreditNote)

	err := c.update(endpoint, request, resp)

	if err != nil {
		return err
	}

	c.CreditNote = resp.CreditNote

	return nil
}

func (c *CreditNote) Void() (*CreditNote, error) {
	resp := new(CreditNote)

	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"

	err := c.postWithoutData(endpoint, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *CreditNote) Delete() error {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *CreditNote) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (CreditNotes, error) {
	endpoint := addFilterAndSort(invdendpoint.CreditNoteEndpoint, filter, sort)

	creditNotes := make(CreditNotes, 0)

NEXT:
	tmpCreditNotes := make(CreditNotes, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpCreditNotes)

	if err != nil {
		return nil, err
	}

	creditNotes = append(creditNotes, tmpCreditNotes...)

	if endpoint != "" {
		goto NEXT
	}

	for _, creditNote := range creditNotes {
		creditNote.Connection = c.Connection
	}

	return creditNotes, nil
}

func (c *CreditNote) ListAttachments() (Files, error) {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

	files := make(Files, 0)

NEXT:
	tempFiles := make(Files, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tempFiles)

	if err != nil {
		return nil, err
	}

	files = append(files, tempFiles...)

	if endpoint != "" {
		goto NEXT
	}

	for _, creditNote := range files {
		creditNote.Connection = c.Connection
	}

	return files, nil
}

func (c *CreditNote) SendEmail(emailReq *invdendpoint.SendEmailRequest) error {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.create(endpoint, emailReq, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *CreditNote) SendText(req *invdendpoint.SendTextMessageRequest) (invdendpoint.TextMessages, error) {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"
	resp := new(invdendpoint.TextMessages)

	err := c.create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *CreditNote) SendLetter() (*invdendpoint.Letter, error) {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"
	resp := new(invdendpoint.Letter)

	err := c.create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CreditNote) String() string {
	header := fmt.Sprintf("<CreditNote id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.CreditNote.String()
}
