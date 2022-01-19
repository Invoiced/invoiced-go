package file

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.FileRequest) (*Client, error) {
	endpoint := invoiced.FileEndpoint
	resp := new(Client)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateAndUploadFile(filePath, fileType string) (*Client, error) {
	endpoint := invoiced.FileEndpoint
	resp := new(Client)

	err := c.Api.Upload(endpoint, filePath, "file", nil, fileType, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id int64) (*Client, error) {
	endpoint := invoiced.FileEndpoint + "/" + strconv.FormatInt(id, 10)

	file := &Client{c.Client, new(invoiced.File)}

	_, err := c.Api.Get(endpoint, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.FileEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}
