package invoiced

import (
	"strconv"
)

type FileClient struct {
	*Client
	*File
}

type Files []*FileClient

func (c *Client) NewFile() *FileClient {
	file := new(File)
	return &FileClient{c, file}
}

func (c *FileClient) Create(request *FileRequest) (*FileClient, error) {
	endpoint := FileEndpoint
	resp := new(FileClient)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *FileClient) CreateAndUploadFile(filePath, fileType string) (*FileClient, error) {
	endpoint := FileEndpoint
	resp := new(FileClient)

	err := c.Api.Upload(endpoint, filePath, "file", nil, fileType, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *FileClient) Retrieve(id int64) (*FileClient, error) {
	endpoint := FileEndpoint + "/" + strconv.FormatInt(id, 10)

	file := &FileClient{c.Client, new(File)}

	_, err := c.Api.Get(endpoint, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *FileClient) Delete() error {
	endpoint := FileEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}
