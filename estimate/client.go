package estimate

import (
	"fmt"
	"github.com/Invoiced/invoiced-go"
	"github.com/Invoiced/invoiced-go/invoice"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

type Estimates []*Client

func (c *Client) Count() (int64, error) {
	endpoint := invoiced.EstimateEndpoint

	count, err := c.Api.Count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Client) Create(request *invoiced.EstimateRequest) (*Client, error) {
	endpoint := invoiced.EstimateEndpoint
	resp := new(Client)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id int64) (*Client, error) {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(id, 10)

	estimate := &Client{c.Client, new(invoiced.Estimate)}

	_, err := c.Api.Get(endpoint, estimate)

	if err != nil {
		return nil, err
	}

	return estimate, nil
}

func (c *Client) Update(request *invoiced.EstimateRequest) error {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Estimate = resp.Estimate

	return nil
}

func (c *Client) Void() (*Client, error) {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/void"
	resp := new(Client)

	err := c.Api.PostWithoutData(endpoint, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (Estimates, error) {
	endpoint := invoiced.EstimateEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

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

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (Estimates, string, error) {
	endpoint := invoiced.EstimateEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	estimates := make(Estimates, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &estimates)

	if err != nil {
		return nil, "", err
	}

	return estimates, nextEndpoint, nil
}

func (c *Client) GenerateInvoice() (*invoice.Client, error) {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/invoice"

	invResp := c.NewInvoice()

	err := c.Api.PostWithoutData(endpoint, invResp)

	if err != nil {
		return nil, err
	}

	return invResp, nil
}

func (c *Client) SendEmail(emailReq *invoiced.SendEmailRequest) error {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/emails"

	err := c.Api.Create(endpoint, emailReq, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendText(req *invoiced.SendTextMessageRequest) (invoiced.TextMessages, error) {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/text_messages"

	resp := new(invoiced.TextMessages)

	err := c.Api.Create(endpoint, req, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *Client) SendLetter() (*invoiced.Letter, error) {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/letters"

	resp := new(invoiced.Letter)

	err := c.Api.Create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListAttachments() (invoiced.Files, error) {
	endpoint := invoiced.EstimateEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/attachments"

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

func (c *Client) String() string {
	header := fmt.Sprintf("<Client id=%d at %p>", c.Id, c)

	return header + " " + "JSON: " + c.Estimate.String()
}
