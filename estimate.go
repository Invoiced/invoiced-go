package invdapi

import (
	"errors"
	"fmt"
	"strconv"

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
	endpoint := invdendpoint.EstimateEndpoint

	count, apiErr := c.count(endpoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil
}

func (c *Estimate) Create(estimate *Estimate) (*Estimate, error) {
	endpoint := invdendpoint.EstimateEndpoint

	estResp := new(Estimate)

	if estimate == nil {
		return nil, errors.New("invoice cannot be nil")
	}

	// safe prune invoice data for creation
	invdEstToCreate, err := SafeEstimateForCreation(estimate.Estimate)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endpoint, invdEstToCreate, estResp)

	if apiErr != nil {
		return nil, apiErr
	}

	estResp.Connection = c.Connection

	return estResp, nil
}

func (c *Estimate) Delete() error {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Estimate) Void() (*Estimate, error) {
	estResp := new(Estimate)

	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"

	apiErr := c.postWithoutData(endpoint, estResp)

	if apiErr != nil {
		return nil, apiErr
	}

	estResp.Connection = c.Connection

	return estResp, nil
}

func (c *Estimate) Save() error {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	estResp := new(Estimate)

	invdEstToUpdate, err := SafeEstimateForUpdate(c.Estimate)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, invdEstToUpdate, estResp)

	if apiErr != nil {
		return apiErr
	}

	c.Estimate = estResp.Estimate

	return nil
}

func (c *Estimate) Retrieve(id int64) (*Estimate, error) {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(id, 10)

	custEndpoint := new(invdendpoint.Estimate)

	estimate := &Estimate{c.Connection, custEndpoint}

	_, apiErr := c.retrieveDataFromAPI(endpoint, estimate)

	if apiErr != nil {
		return nil, apiErr
	}

	return estimate, nil
}

func (c *Estimate) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Estimates, error) {
	endpoint := invdendpoint.EstimateEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	estimates := make(Estimates, 0)

NEXT:
	tmpInvoices := make(Estimates, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpInvoices)

	if apiErr != nil {
		return nil, apiErr
	}

	estimates = append(estimates, tmpInvoices...)

	if endpoint != "" {
		goto NEXT
	}

	for _, estimate := range estimates {
		estimate.Connection = c.Connection
	}

	return estimates, nil
}

func (c *Estimate) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Estimates, string, error) {
	endpoint := invdendpoint.EstimateEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	estimates := make(Estimates, 0)

	nextEndpoint, apiErr := c.retrieveDataFromAPI(endpoint, &estimates)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, estimate := range estimates {
		estimate.Connection = c.Connection
	}

	return estimates, nextEndpoint, nil
}

func (c *Estimate) GenerateInvoice() (*Invoice, error) {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/invoice"

	invResp := c.NewInvoice()

	apiErr := c.postWithoutData(endpoint, invResp)

	if apiErr != nil {
		return nil, apiErr
	}

	return invResp, nil
}

func (c *Estimate) SendEmail(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endpoint, emailReq, emailResp)
	if err != nil {
		return nil, err
	}

	return *emailResp, nil
}

func (c *Estimate) SendText(req *invdendpoint.TextRequest) (invdendpoint.TextResponses, error) {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	resp := new(invdendpoint.TextResponses)

	err := c.create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *Estimate) SendLetter() (*invdendpoint.LetterResponse, error) {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	resp := new(invdendpoint.LetterResponse)

	err := c.create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Estimate) ListAttachments() (Files, error) {
	endpoint :=  invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

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

	for _, estimate := range files {
		estimate.Connection = c.Connection
	}

	return files, nil
}

func (c *Estimate) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Estimate.String()
}

// SafeEstimateForCreation prunes estimate data for just fields that can be used for creation of a invoice
func SafeEstimateForCreation(estimate *invdendpoint.Estimate) (*invdendpoint.Estimate, error) {
	if estimate == nil {
		return nil, errors.New("Estimate is nil")
	}

	estData := new(invdendpoint.Estimate)
	estData.Customer = estimate.Customer
	estData.Invoice = estimate.Invoice
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
	estData.Metadata = estimate.Metadata
	estData.Attachments = estimate.Attachments
	estData.DisabledPaymentMethods = estimate.DisabledPaymentMethods
	estData.CalculateTax = estimate.CalculateTax

	return estData, nil
}

// SafeInvoiceForCreation prunes invoice data for just fields that can be used for creation of a invoice
func SafeEstimateForUpdate(estimate *invdendpoint.Estimate) (*invdendpoint.Estimate, error) {
	if estimate == nil {
		return nil, errors.New("Estimate is nil")
	}

	estData := new(invdendpoint.Estimate)
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
	estData.Metadata = estimate.Metadata
	estData.Attachments = estimate.Attachments
	estData.DisabledPaymentMethods = estimate.DisabledPaymentMethods
	estData.CalculateTax = estimate.CalculateTax

	return estData, nil
}
