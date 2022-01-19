package task

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

type Tasks []*Client

func (c *Client) Create(request *invoiced.TaskRequest) (*Client, error) {
	endpoint := invoiced.TaskEndpoint
	resp := new(Client)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id int64) (*Client, error) {
	endpoint := invoiced.TaskEndpoint + "/" + strconv.FormatInt(id, 10)

	taskEndpoint := new(invoiced.Task)

	task := &Client{c.Client, taskEndpoint}

	_, err := c.Api.Get(endpoint, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (c *Client) Update(request *invoiced.TaskRequest) error {
	endpoint := invoiced.TaskEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Task = resp.Task

	return nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.TaskEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (Tasks, error) {
	endpoint := invoiced.TaskEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	tasks := make(Tasks, 0)

NEXT:
	tmpTasks := make(Tasks, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpTasks)

	if err != nil {
		return nil, err
	}

	tasks = append(tasks, tmpTasks...)

	if endpointTmp != "" {
		goto NEXT
	}

	return tasks, nil
}
