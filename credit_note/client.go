package credit_note

import "C"
import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.CreditNoteRequest) (*invoiced.CreditNote, error) {
	resp := new(invoiced.CreditNote)
	err := c.Api.Create("/credit_notes", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.CreditNote, error) {
	resp := new(invoiced.CreditNote)
	_, err := c.Api.Get("/credit_notes/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.CreditNoteRequest) (*invoiced.CreditNote, error) {
	resp := new(invoiced.CreditNote)
	err := c.Api.Update("/credit_notes/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Void(id int64) (*Client, error) {
	resp := new(Client)

	endpoint := "/credit_notes/" + strconv.FormatInt(id, 10) + "/void"

	err := c.Api.PostWithoutData(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/credit_notes/" + strconv.FormatInt(id, 10))
}

func (c *Client) Count() (int64, error) {
	return c.Api.Count("/credit_notes")
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.CreditNotes, error) {
	endpoint := invoiced.AddFilterAndSort("/credit_notes", filter, sort)

	creditNotes := make(invoiced.CreditNotes, 0)

NEXT:
	tmpCreditNotes := make(invoiced.CreditNotes, 0)

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

func (c *Client) ListAttachments(id int64) (invoiced.Files, error) {
	endpoint := "/credit_notes/" + strconv.FormatInt(id, 10) + "/attachments"

	files := make(invoiced.Files, 0)

NEXT:
	tempFiles := make(invoiced.Files, 0)

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

func (c *Client) SendEmail(id int64, request *invoiced.SendEmailRequest) error {
	return c.Api.Create("/credit_notes/" + strconv.FormatInt(id, 10) + "/emails", request, nil)
}
