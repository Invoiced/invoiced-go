package credit_note

import "C"
import (
	"fmt"
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

type CreditNotes []*Client

func (c *Client) Count() (int64, error) {
	count, err := c.Api.Count(invoiced.CreditNoteEndpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Client) Create(request *invoiced.CreditNoteRequest) (*Client, error) {
	resp := new(Client)

	err := c.Api.Create(invoiced.CreditNoteEndpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id int64) (*Client, error) {
	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(id, 10)

	creditNote := &Client{c.Client, new(invoiced.CreditNote)}

	_, err := c.Api.Get(endpoint, creditNote)

	if err != nil {
		return nil, err
	}

	return creditNote, nil
}

func (c *Client) Update(request *invoiced.CreditNoteRequest) error {
	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)

	if err != nil {
		return err
	}

	c.CreditNote = resp.CreditNote

	return nil
}

func (c *Client) Void() (*Client, error) {
	resp := new(Client)

	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"

	err := c.Api.PostWithoutData(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (CreditNotes, error) {
	endpoint := invoiced.AddFilterAndSort(invoiced.CreditNoteEndpoint, filter, sort)

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

func (c *Client) ListAttachments() (invoiced.Files, error) {
	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

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

func (c *Client) SendEmail(emailReq *invoiced.SendEmailRequest) error {
	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.Api.Create(endpoint, emailReq, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendText(req *invoiced.SendTextMessageRequest) (invoiced.TextMessages, error) {
	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"
	resp := new(invoiced.TextMessages)

	err := c.Api.Create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *Client) SendLetter() (*invoiced.Letter, error) {
	endpoint := invoiced.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"
	resp := new(invoiced.Letter)

	err := c.Api.Create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) String() string {
	header := fmt.Sprintf("<Client id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.CreditNote.String()
}
