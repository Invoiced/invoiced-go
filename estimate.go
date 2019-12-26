package invdapi

import (
"errors"
"fmt"
"github.com/Invoiced/invoiced-go/invdendpoint"
)


type Estimate struct {
	*Connection
	*invdendpoint.Estimate
}

type Estimates []*Estimate

func (c *Connection) NewEstimate() *Estimate {
	estimate := new(invdendpoint.Estimate)
	return &Estimate{c, estimate}

}

func (c *Estimate) Count() (int64, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.EstimatesEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Estimate) Create(estimate *Estimate) (*Estimate, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.EstimatesEndPoint)

	estResp := new(Estimate)

	if estimate == nil {
		return nil, errors.New("invoice cannot be nil")
	}

	//safe prune invoice data for creation
	invdEstToCreate,err := SafeEstimateForCreation(estimate.Estimate)

	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdEstToCreate, estResp)

	if apiErr != nil {
		return nil, apiErr
	}

	estResp.Connection = c.Connection

	return estResp, nil

}

func (c *Estimate) Delete() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EstimatesEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Estimate) Void() (*Estimate, error) {

	estResp := new(Estimate)

	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EstimatesEndPoint), c.Id) + "/void"

	apiErr := c.postWithoutData(endPoint,estResp)

	if apiErr != nil {
		return nil,apiErr
	}

	estResp.Connection = c.Connection

	return estResp,nil

}

func (c *Estimate) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EstimatesEndPoint), c.Id)

	estResp := new(Estimate)

	invdEstToUpdate, err := SafeEstimateForUpdate(c.Estimate)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, invdEstToUpdate, estResp)

	if apiErr != nil {
		return apiErr
	}

	c.Estimate = estResp.Estimate

	return nil

}

func (c *Estimate) Retrieve(id int64) (*Estimate, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EstimatesEndPoint), id)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	custEndPoint := new(invdendpoint.Estimate)

	estimate := &Estimate{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, estimate)

	if apiErr != nil {
		return nil, apiErr
	}

	return estimate, nil

}

func (c *Estimate) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Estimates, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.EstimatesEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	estimates := make(Estimates, 0)

NEXT:
	tmpInvoices := make(Estimates, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpInvoices)

	if apiErr != nil {
		return nil, apiErr
	}

	estimates = append(estimates, tmpInvoices...)

	if endPoint != "" {
		goto NEXT
	}

	for _, estimate := range estimates {
		estimate.Connection = c.Connection
	}

	return estimates, nil

}

func (c *Estimate) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Estimates, string, error) {

	endPoint := c.MakeEndPointURL(invdendpoint.EstimatesEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)


	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	estimates := make(Estimates, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &estimates)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, estimate := range estimates {
		estimate.Connection = c.Connection

	}

	return estimates, nextEndPoint, nil

}

func (c *Estimate) GenerateInvoice() (*Invoice, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.EstimatesEndPoint) + "/invoice"

	invResp := new(Invoice)

	apiErr := c.postWithoutData(endPoint, invResp)

	if apiErr != nil {
		return nil, apiErr
	}

	invResp.Connection = c.Connection

	return invResp, nil

}

func (c *Estimate) SendEmail(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EstimatesEndPoint), c.Id) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endPoint, emailReq, emailResp)

	if err != nil {
		return nil, err
	}

	return *emailResp, nil

}

func (c *Estimate) SendText(req *invdendpoint.TextRequest) (invdendpoint.TextResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EstimatesEndPoint), c.Id) + "/text_messages"

	resp := new(invdendpoint.TextResponses)

	err := c.create(endPoint, req, resp)

	if err != nil {
		return nil, err
	}

	return *resp, nil

}

func (c *Estimate) SendLetter(req *invdendpoint.LetterRequest) (*invdendpoint.LetterResponse, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EstimatesEndPoint), c.Id) + "/letters"

	resp := new(invdendpoint.LetterResponse)

	err := c.create(endPoint, req, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (c *Estimate) ListAttachments() (Files, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.EstimatesEndPoint) + "/attachments"

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

	for _, estimate := range files {
		estimate.Connection = c.Connection
	}

	return files, nil

}



func (c *Estimate) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Estimate.String()
}

//SafeEstimateForCreation prunes estimate data for just fields that can be used for creation of a invoice
func SafeEstimateForCreation(estimate *invdendpoint.Estimate) (*invdendpoint.Estimate, error) {

	if estimate == nil  {
		return nil, errors.New("Estimate is nil")
	}

	estData :=new(invdendpoint.Estimate)
	estData.Customer = estimate.Customer
	estData.Invoice= estimate.Invoice
	estData.Name = estimate.Name
	estData.Number = estimate.Number
	estData.Currency = estimate.Currency
	estData.Date = estimate.Date
	estData.ExpirationDate = estimate.ExpirationDate
	estData.PaymentTerms = estimate.PaymentTerms
	estData.Draft = estimate.Draft
	estData.Closed = estimate.Closed
	estData.Items = estimate.Items
	estData.Notes = estimate.Notes
	estData.Discounts = estimate.Discounts
	estData.ShipTo = estimate.ShipTo
	estData.Deposit = estimate.Deposit
	estData.DepositPaid = estimate.DepositPaid
	estData.MetaData = estimate.MetaData
	estData.Attachments = estimate.Attachments
	estData.DisabledPaymentMethods = estimate.DisabledPaymentMethods
	estData.CalculateTax = estimate.CalculateTax

	return estData,nil
}

//SafeInvoiceForCreation prunes invoice data for just fields that can be used for creation of a invoice
func SafeEstimateForUpdate(estimate *invdendpoint.Estimate) (*invdendpoint.Estimate, error) {

	if estimate == nil  {
		return nil, errors.New("Estimate is nil")
	}

	estData :=new(invdendpoint.Estimate)
	estData.Name = estimate.Name
	estData.Number = estimate.Number
	estData.Currency = estimate.Currency
	estData.Date = estimate.Date
	estData.ExpirationDate = estimate.ExpirationDate
	estData.PaymentTerms = estimate.PaymentTerms
	estData.Draft = estimate.Draft
	estData.Closed = estimate.Closed
	estData.Items = estimate.Items
	estData.Notes = estimate.Notes
	estData.Discounts = estimate.Discounts
	estData.Taxes = estimate.Taxes
	estData.ShipTo = estimate.ShipTo
	estData.Deposit = estimate.Deposit
	estData.DepositPaid = estimate.DepositPaid
	estData.MetaData = estimate.MetaData
	estData.Attachments = estimate.Attachments
	estData.DisabledPaymentMethods = estimate.DisabledPaymentMethods
	estData.CalculateTax = estimate.CalculateTax

	return estData,nil
}