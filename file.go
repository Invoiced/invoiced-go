package invdapi

import (
	"errors"
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

func (c *File) Create(file *File) (*File, error) {
	endpoint := invdendpoint.FileEndpoint
	fileResp := new(File)

	// safe prune file data for creation
	invdFileDataToCreate, err := SafeFileForCreation(file.File)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endpoint, invdFileDataToCreate, fileResp)

	if apiErr != nil {
		return nil, apiErr
	}

	fileResp.Connection = c.Connection

	return fileResp, nil
}

func (c *File) CreateAndUploadFile(filePath,fileType string) (*File, error) {
	endpoint := invdendpoint.FileEndpoint
	fileResp := new(File)


	apiErr := c.upload(endpoint,filePath,"file",nil,fileType,fileResp)

	if apiErr != nil {
		return nil, apiErr
	}

	fileResp.Connection = c.Connection

	return fileResp, nil
}

func (c *File) Delete() error {
	endpoint := invdendpoint.FileEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *File) Retrieve(id int64) (*File, error) {
	endpoint := invdendpoint.FileEndpoint + "/" + strconv.FormatInt(id, 10)

	custEndpoint := new(invdendpoint.File)

	file := &File{c.Connection, custEndpoint}

	_, err := c.retrieveDataFromAPI(endpoint, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafeFileForCreation(file *invdendpoint.File) (*invdendpoint.File, error) {
	if file == nil {
		return nil, errors.New("file is nil")
	}

	fileData := new(invdendpoint.File)
	fileData.Name = file.Name
	fileData.Size = file.Size
	fileData.Type = file.Type
	fileData.Url = file.Url

	return fileData, nil
}
