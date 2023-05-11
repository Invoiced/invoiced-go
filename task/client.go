package task

import (
	"github.com/Invoiced/invoiced-go/v2"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.TaskRequest) (*invoiced.Task, error) {
	resp := new(invoiced.Task)
	err := c.Api.Create("/tasks", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.Task, error) {
	resp := new(invoiced.Task)
	_, err := c.Api.Get("/tasks/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.TaskRequest) (*invoiced.Task, error) {
	resp := new(invoiced.Task)
	err := c.Api.Update("/tasks/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/tasks/" + strconv.FormatInt(id, 10))
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Tasks, error) {
	endpoint := invoiced.AddFilterAndSort("/tasks", filter, sort)

	tasks := make(invoiced.Tasks, 0)

NEXT:
	tmpTasks := make(invoiced.Tasks, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpTasks)

	if err != nil {
		return nil, err
	}

	tasks = append(tasks, tmpTasks...)

	if endpoint != "" {
		goto NEXT
	}

	return tasks, nil
}
