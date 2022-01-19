package invdapi

import (
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

	count, err := c.count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Estimate) Create(request *invdendpoint.EstimateRequest) (*Estimate, error) {
	endpoint := invdendpoint.EstimateEndpoint
	resp := new(Estimate)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Estimate) Retrieve(id int64) (*Estimate, error) {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(id, 10)

	estimate := &Estimate{c.Connection, new(invdendpoint.Estimate)}

	_, err := c.retrieveDataFromAPI(endpoint, estimate)

	if err != nil {
		return nil, err
	}

	return estimate, nil
}

func (c *Estimate) Update(request *invdendpoint.EstimateRequest) error {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Estimate)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Estimate = resp.Estimate

	return nil
}

func (c *Estimate) Void() (*Estimate, error) {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"
	resp := new(Estimate)

	err := c.postWithoutData(endpoint, resp)

	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Estimate) Delete() error {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Estimate) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Estimates, error) {
	endpoint := invdendpoint.EstimateEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	estimates := make(Estimates, 0)

NEXT:
	tmpInvoices := make(Estimates, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpInvoices)

	if err != nil {
		return nil, err
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

	nextEndpoint, err := c.retrieveDataFromAPI(endpoint, &estimates)

	if err != nil {
		return nil, "", err
	}

	for _, estimate := range estimates {
		estimate.Connection = c.Connection
	}

	return estimates, nextEndpoint, nil
}

func (c *Estimate) GenerateInvoice() (*Invoice, error) {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/invoice"

	invResp := c.NewInvoice()

	err := c.postWithoutData(endpoint, invResp)

	if err != nil {
		return nil, err
	}

	return invResp, nil
}

func (c *Estimate) SendEmail(emailReq *invdendpoint.SendEmailRequest) error {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.create(endpoint, emailReq, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Estimate) SendText(req *invdendpoint.SendTextMessageRequest) (invdendpoint.TextMessages, error) {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	resp := new(invdendpoint.TextMessages)

	err := c.create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *Estimate) SendLetter() (*invdendpoint.Letter, error) {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	resp := new(invdendpoint.Letter)

	err := c.create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Estimate) ListAttachments() (Files, error) {
	endpoint := invdendpoint.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

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

	for _, estimate := range files {
		estimate.Connection = c.Connection
	}

	return files, nil
}

func (c *Estimate) String() string {
	header := fmt.Sprintf("<Invoice id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Estimate.String()
}
