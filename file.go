package invdapi

import (
	"github.com/ActiveState/invoiced-go/invdendpoint"
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
	endPoint := c.makeEndPointURL(invdendpoint.FilesEndPoint)
	fileResp := new(File)

	apiErr := c.create(endPoint, file, fileResp)

	if apiErr != nil {
		return nil, apiErr
	}

	fileResp.Connection = c.Connection

	return fileResp, nil

}

func (c *File) Delete() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.FilesEndPoint), c.Id)

	err := c.delete(endPoint)

	if err != nil {
		return err
	}

	return nil

}

func (c *File) Retrieve(id int64) (*File, error) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.FilesEndPoint), id)

	custEndPoint := new(invdendpoint.File)

	file := &File{c.Connection, custEndPoint}

	_, err := c.retrieveDataFromAPI(endPoint, file)

	if err != nil {
		return nil, err
	}

	return file, nil

}
