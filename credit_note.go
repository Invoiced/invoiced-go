package invoiced

import "C"
import (
	"fmt"
	"strconv"
)

type CreditNoteClient struct {
	*Client
	*CreditNote
}

type CreditNotes []*CreditNoteClient

func (c *Client) NewCreditNote() *CreditNoteClient {
	creditNote := new(CreditNote)
	return &CreditNoteClient{c, creditNote}
}

func (c *CreditNoteClient) Count() (int64, error) {
	count, err := c.Api.Count(CreditNoteEndpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *CreditNoteClient) Create(request *CreditNoteRequest) (*CreditNoteClient, error) {
	resp := new(CreditNoteClient)

	err := c.Api.Create(CreditNoteEndpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CreditNoteClient) Retrieve(id int64) (*CreditNoteClient, error) {
	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(id, 10)

	creditNote := &CreditNoteClient{c.Client, new(CreditNote)}

	_, err := c.Api.Get(endpoint, creditNote)

	if err != nil {
		return nil, err
	}

	return creditNote, nil
}

func (c *CreditNoteClient) Update(request *CreditNoteRequest) error {
	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(CreditNoteClient)

	err := c.Api.Update(endpoint, request, resp)

	if err != nil {
		return err
	}

	c.CreditNote = resp.CreditNote

	return nil
}

func (c *CreditNoteClient) Void() (*CreditNoteClient, error) {
	resp := new(CreditNoteClient)

	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"

	err := c.Api.PostWithoutData(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CreditNoteClient) Delete() error {
	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *CreditNoteClient) ListAll(filter *Filter, sort *Sort) (CreditNotes, error) {
	endpoint := AddFilterAndSort(CreditNoteEndpoint, filter, sort)

	creditNotes := make(CreditNotes, 0)

NEXT:
	tmpCreditNotes := make(CreditNotes, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpCreditNotes)

	if err != nil {
		return nil, err
	}

	creditNotes = append(creditNotes, tmpCreditNotes...)

	if endpoint != "" {
		goto NEXT
	}

	return creditNotes, nil
}

func (c *CreditNoteClient) ListAttachments() (Files, error) {
	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

	files := make(Files, 0)

NEXT:
	tempFiles := make(Files, 0)

	endpoint, err := c.Api.Get(endpoint, &tempFiles)

	if err != nil {
		return nil, err
	}

	files = append(files, tempFiles...)

	if endpoint != "" {
		goto NEXT
	}

	return files, nil
}

func (c *CreditNoteClient) SendEmail(emailReq *SendEmailRequest) error {
	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.Api.Create(endpoint, emailReq, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *CreditNoteClient) SendText(req *SendTextMessageRequest) (TextMessages, error) {
	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"
	resp := new(TextMessages)

	err := c.Api.Create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *CreditNoteClient) SendLetter() (*Letter, error) {
	endpoint := CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"
	resp := new(Letter)

	err := c.Api.Create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CreditNoteClient) String() string {
	header := fmt.Sprintf("<CreditNoteClient id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.CreditNote.String()
}
