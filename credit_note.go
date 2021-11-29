package invdapi

import "C"
import (
	"errors"
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
	count, apiErr := c.count(invdendpoint.CreditNoteEndpoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil
}

func (c *CreditNote) Create(creditNote *CreditNote) (*CreditNote, error) {
	cnResp := new(CreditNote)

	if creditNote == nil {
		return nil, errors.New("credit note cannot be nil")
	}

	// safe prune invoice data for creation
	invdCNToCreate, err := SafeCreditNoteForCreation(creditNote.CreditNote)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(invdendpoint.CreditNoteEndpoint, invdCNToCreate, cnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	cnResp.Connection = c.Connection

	return cnResp, nil
}

func (c *CreditNote) Delete() error {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *CreditNote) Void() (*CreditNote, error) {
	cnResp := new(CreditNote)

	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"

	apiErr := c.postWithoutData(endpoint, cnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	cnResp.Connection = c.Connection

	return cnResp, nil
}

func (c *CreditNote) Save() error {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	cnResp := new(CreditNote)

	invdCnToUpdate, err := SafeCreditNoteForUpdate(c.CreditNote)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, invdCnToUpdate, cnResp)

	if apiErr != nil {
		return apiErr
	}

	c.CreditNote = cnResp.CreditNote

	return nil
}

func (c *CreditNote) Retrieve(id int64) (*CreditNote, error) {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(id, 10)

	custEndpoint := new(invdendpoint.CreditNote)

	creditNote := &CreditNote{c.Connection, custEndpoint}

	_, apiErr := c.retrieveDataFromAPI(endpoint, creditNote)

	if apiErr != nil {
		return nil, apiErr
	}

	return creditNote, nil
}

func (c *CreditNote) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (CreditNotes, error) {
	endpoint := addFilterAndSort(invdendpoint.CreditNoteEndpoint, filter, sort)

	creditNotes := make(CreditNotes, 0)

NEXT:
	tmpCreditNotes := make(CreditNotes, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpCreditNotes)

	if apiErr != nil {
		return nil, apiErr
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

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tempFiles)

	if apiErr != nil {
		return nil, apiErr
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

func (c *CreditNote) SendEmail(emailReq *invdendpoint.EmailRequest)  error {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"


	err := c.create(endpoint, emailReq, nil)
	if err != nil {
		return err
	}

	return  nil
}

func (c *CreditNote) SendText(req *invdendpoint.TextRequest) (invdendpoint.TextResponses, error) {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	resp := new(invdendpoint.TextResponses)

	err := c.create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *CreditNote) SendLetter() (*invdendpoint.LetterResponse, error) {
	endpoint := invdendpoint.CreditNoteEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	resp := new(invdendpoint.LetterResponse)

	err := c.create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CreditNote) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.CreditNote.String()
}

// SafeEstimateForCreation prunes credit note data for just fields that can be used for creation of a credit note
func SafeCreditNoteForCreation(creditNote *invdendpoint.CreditNote) (*invdendpoint.CreditNote, error) {
	if creditNote == nil {
		return nil, errors.New("CreditNote is nil")
	}

	cnData := new(invdendpoint.CreditNote)
	cnData.Customer = creditNote.Customer
	cnData.Invoice = creditNote.Invoice
	cnData.Name = creditNote.Name
	cnData.Number = creditNote.Number
	cnData.Currency = creditNote.Currency
	cnData.Date = creditNote.Date
	cnData.Draft = creditNote.Draft
	cnData.Closed = creditNote.Closed
	cnData.Items = creditNote.Items
	cnData.Notes = creditNote.Notes
	cnData.Discounts = creditNote.Discounts
	cnData.Taxes = creditNote.Taxes
	cnData.Metadata = creditNote.Metadata
	cnData.Attachments = creditNote.Attachments
	cnData.CalculateTax = creditNote.CalculateTax

	return cnData, nil
}

// SafeCreditNoteForUpdate prunes creditnote data for just fields that can be used for updating a credit note
func SafeCreditNoteForUpdate(creditNote *invdendpoint.CreditNote) (*invdendpoint.CreditNote, error) {
	if creditNote == nil {
		return nil, errors.New("CreditNote is nil")
	}

	cnData := new(invdendpoint.CreditNote)
	cnData.Name = creditNote.Name
	cnData.Number = creditNote.Number
	cnData.Currency = creditNote.Currency
	cnData.Date = creditNote.Date
	cnData.Draft = creditNote.Draft
	cnData.Closed = creditNote.Closed
	cnData.Items = creditNote.Items
	cnData.Notes = creditNote.Notes
	cnData.Discounts = creditNote.Discounts
	cnData.Taxes = creditNote.Taxes
	cnData.Metadata = creditNote.Metadata
	cnData.Attachments = creditNote.Attachments
	cnData.CalculateTax = creditNote.CalculateTax

	return cnData, nil
}
