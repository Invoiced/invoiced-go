package file

import (
	"github.com/Invoiced/invoiced-go/v2"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.FileRequest) (*invoiced.File, error) {
	resp := new(invoiced.File)
	err := c.Api.Create("/files", request, resp)
	return resp, err
}

func (c *Client) CreateAndUploadFile(filePath, fileType string) (*invoiced.File, error) {
	resp := new(invoiced.File)
	err := c.Api.Upload("/files", filePath, "file", nil, fileType, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.File, error) {
	resp := new(invoiced.File)
	_, err := c.Api.Get("/files/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/files/" + strconv.FormatInt(id, 10))
}
