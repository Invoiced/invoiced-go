package estimate

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.EstimateRequest) (*invoiced.Estimate, error) {
	resp := new(invoiced.Estimate)
	err := c.Api.Create("/estimates", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.Estimate, error) {
	resp := new(invoiced.Estimate)
	_, err := c.Api.Get("/estimates/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.EstimateRequest) (*invoiced.Estimate, error) {
	resp := new(invoiced.Estimate)
	err := c.Api.Update("/estimates/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Void(id int64) (*invoiced.Estimate, error) {
	resp := new(invoiced.Estimate)
	err := c.Api.PostWithoutData("/estimates/"+strconv.FormatInt(id, 10)+"/void", resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/estimates/" + strconv.FormatInt(id, 10))
}

func (c *Client) Count() (int64, error) {
	return c.Api.Count("/estimates")
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Estimates, error) {
	endpoint := invoiced.AddFilterAndSort("/estimates", filter, sort)

	estimates := make(invoiced.Estimates, 0)

NEXT:
	tmpInvoices := make(invoiced.Estimates, 0)

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

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Estimates, string, error) {
	endpoint := invoiced.AddFilterAndSort("/estimates", filter, sort)

	estimates := make(invoiced.Estimates, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &estimates)

	if err != nil {
		return nil, "", err
	}

	return estimates, nextEndpoint, nil
}

func (c *Client) GenerateInvoice(id int64) (*invoiced.Invoice, error) {
	endpoint := "/estimates/" + strconv.FormatInt(id, 10) + "/invoice"
	resp := new(invoiced.Invoice)
	err := c.Api.PostWithoutData(endpoint, resp)
	return resp, err
}

func (c *Client) SendEmail(id int64, request *invoiced.SendEmailRequest) error {
	endpoint := "/estimates/" + strconv.FormatInt(id, 10) + "/emails"

	err := c.Api.Create(endpoint, request, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendText(id int64, request *invoiced.SendTextMessageRequest) (invoiced.TextMessages, error) {
	endpoint := "/estimates/" + strconv.FormatInt(id, 10) + "/text_messages"

	resp := new(invoiced.TextMessages)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return *resp, nil
}

func (c *Client) SendLetter(id int64) (*invoiced.Letter, error) {
	endpoint := "/estimates/" + strconv.FormatInt(id, 10) + "/letters"

	resp := new(invoiced.Letter)

	err := c.Api.Create(endpoint, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListAttachments(id int64) (invoiced.Files, error) {
	endpoint := "/estimates/" + strconv.FormatInt(id, 10) + "/attachments"

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
