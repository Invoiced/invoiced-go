package invdapi

import (
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type File struct {
	*Connection
	*invdendpoint.File
}

type Files []*File

func (c *Connection) NewFile() *File {
	file := new(invdendpoint.File)
	return &File{c, file}
}

func (c *File) Create(request *invdendpoint.FileRequest) (*File, error) {
	endpoint := invdendpoint.FileEndpoint
	resp := new(File)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *File) CreateAndUploadFile(filePath, fileType string) (*File, error) {
	endpoint := invdendpoint.FileEndpoint
	resp := new(File)

	err := c.upload(endpoint, filePath, "file", nil, fileType, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *File) Retrieve(id int64) (*File, error) {
	endpoint := invdendpoint.FileEndpoint + "/" + strconv.FormatInt(id, 10)

	file := &File{c.Connection, new(invdendpoint.File)}

	_, err := c.retrieveDataFromAPI(endpoint, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *File) Delete() error {
	endpoint := invdendpoint.FileEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}
