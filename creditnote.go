package invdapi

import (
"errors"
"fmt"
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
	endPoint := c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *CreditNote) Create(creditNote *CreditNote) (*CreditNote, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint)

	cnResp := new(CreditNote)

	if creditNote == nil {
		return nil, errors.New("credit note cannot be nil")
	}

	//safe prune invoice data for creation
	invdCNToCreate,err := SafeCreditNoteForCreation(creditNote.CreditNote)

	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdCNToCreate, cnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	cnResp.Connection = c.Connection

	return cnResp, nil

}

func (c *CreditNote) Delete() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *CreditNote) Void() (*CreditNote, error) {

	cnResp := new(CreditNote)

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), c.Id) + "/void"

	apiErr := c.postWithoutData(endPoint,cnResp)

	if apiErr != nil {
		return nil,apiErr
	}

	cnResp.Connection = c.Connection

	return cnResp,nil

}

func (c *CreditNote) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), c.Id)

	cnResp := new(CreditNote)

	invdCnToUpdate, err := SafeCreditNoteForUpdate(c.CreditNote)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, invdCnToUpdate, cnResp)

	if apiErr != nil {
		return apiErr
	}

	c.CreditNote = cnResp.CreditNote

	return nil

}

func (c *CreditNote) Retrieve(id int64) (*CreditNote, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), id)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	custEndPoint := new(invdendpoint.CreditNote)

	creditNote := &CreditNote{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, creditNote)

	if apiErr != nil {
		return nil, apiErr
	}

	return creditNote, nil

}

func (c *CreditNote) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (CreditNotes, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	creditNotes := make(CreditNotes, 0)

NEXT:
	tmpCreditNotes := make(CreditNotes, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpCreditNotes)

	if apiErr != nil {
		return nil, apiErr
	}

	creditNotes = append(creditNotes, tmpCreditNotes...)

	if endPoint != "" {
		goto NEXT
	}

	for _, creditNote := range creditNotes {
		creditNote.Connection = c.Connection
	}

	return creditNotes, nil

}

func (c *CreditNote) ListAttachments() (Files, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), c.Id) + "/attachments"

	files := make(Files, 0)

NEXT:
	tempFiles := make(Files, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tempFiles)

	if apiErr != nil {
		return nil, apiErr
	}

	files = append(files, tempFiles...)

	if endPoint != "" {
		goto NEXT
	}

	for _, creditNote := range files {
		creditNote.Connection = c.Connection
	}

	return files, nil

}

func (c *CreditNote) SendEmail(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), c.Id) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endPoint, emailReq, emailResp)

	if err != nil {
		return nil, err
	}

	return *emailResp, nil

}

func (c *CreditNote) SendText(req *invdendpoint.TextRequest) (invdendpoint.TextResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), c.Id) + "/text_messages"

	resp := new(invdendpoint.TextResponses)

	err := c.create(endPoint, req, resp)

	if err != nil {
		return nil, err
	}

	return *resp, nil

}

func (c *CreditNote) SendLetter() (*invdendpoint.LetterResponse, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.CreditNotesEndPoint), c.Id) + "/letters"

	resp := new(invdendpoint.LetterResponse)

	err := c.create(endPoint, nil, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil

}



func (c *CreditNote) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.CreditNote.String()
}

//SafeEstimateForCreation prunes credit note data for just fields that can be used for creation of a credit note
func SafeCreditNoteForCreation(creditNote *invdendpoint.CreditNote) (*invdendpoint.CreditNote, error) {

	if creditNote == nil  {
		return nil, errors.New("CreditNote is nil")
	}

	cnData :=new(invdendpoint.CreditNote)
	cnData.Customer = creditNote.Customer
	cnData.Invoice= creditNote.Invoice
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
	cnData.MetaData = creditNote.MetaData
	cnData.Attachments = creditNote.Attachments
	cnData.CalculateTax = creditNote.CalculateTax

	return cnData,nil
}

//SafeCreditNoteForUpdate prunes creditnote data for just fields that can be used for updating a credit note
func SafeCreditNoteForUpdate(creditNote *invdendpoint.CreditNote) (*invdendpoint.CreditNote, error) {

	if creditNote == nil  {
		return nil, errors.New("CreditNote is nil")
	}

	cnData :=new(invdendpoint.CreditNote)
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
	cnData.MetaData = creditNote.MetaData
	cnData.Attachments = creditNote.Attachments
	cnData.CalculateTax = creditNote.CalculateTax

	return cnData,nil
}