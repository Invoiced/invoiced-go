package invoiced

import (
	"fmt"
	"strconv"
)

type EstimateClient struct {
	*Client
	*Estimate
}

type Estimates []*EstimateClient

func (c *Client) NewEstimate() *EstimateClient {
	estimate := new(Estimate)
	return &EstimateClient{c, estimate}
}

func (c *EstimateClient) Count() (int64, error) {
	endpoint := EstimateEndpoint

	count, err := c.Api.Count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *EstimateClient) Create(request *EstimateRequest) (*EstimateClient, error) {
	endpoint := EstimateEndpoint
	resp := new(EstimateClient)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *EstimateClient) Retrieve(id int64) (*EstimateClient, error) {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(id, 10)

	estimate := &EstimateClient{c.Client, new(Estimate)}

	_, err := c.Api.Get(endpoint, estimate)

	if err != nil {
		return nil, err
	}

	return estimate, nil
}

func (c *EstimateClient) Update(request *EstimateRequest) error {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(EstimateClient)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Estimate = resp.Estimate

	return nil
}

func (c *EstimateClient) Void() (*EstimateClient, error) {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"
	resp := new(EstimateClient)

	err := c.Api.PostWithoutData(endpoint, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *EstimateClient) Delete() error {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *EstimateClient) ListAll(filter *Filter, sort *Sort) (Estimates, error) {
	endpoint := EstimateEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

	estimates := make(Estimates, 0)

NEXT:
	tmpInvoices := make(Estimates, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpInvoices)

	if err != nil {
		return nil, err
	}

	estimates = append(estimates, tmpInvoices...)

	if endpoint != "" {
		goto NEXT
	}

	return estimates, nil
}

func (c *EstimateClient) List(filter *Filter, sort *Sort) (Estimates, string, error) {
	endpoint := EstimateEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)

	estimates := make(Estimates, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &estimates)

	if err != nil {
		return nil, "", err
	}

	return estimates, nextEndpoint, nil
}

func (c *EstimateClient) GenerateInvoice() (*InvoiceClient, error) {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/invoice"

	invResp := c.NewInvoice()

	err := c.Api.PostWithoutData(endpoint, invResp)

	if err != nil {
		return nil, err
	}

	return invResp, nil
}

func (c *EstimateClient) SendEmail(emailReq *SendEmailRequest) error {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.Api.Create(endpoint, emailReq, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *EstimateClient) SendText(req *SendTextMessageRequest) (TextMessages, error) {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	resp := new(TextMessages)

	err := c.Api.Create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *EstimateClient) SendLetter() (*Letter, error) {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	resp := new(Letter)

	err := c.Api.Create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *EstimateClient) ListAttachments() (Files, error) {
	endpoint := EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

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

func (c *EstimateClient) String() string {
	header := fmt.Sprintf("<InvoiceClient id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Estimate.String()
}
