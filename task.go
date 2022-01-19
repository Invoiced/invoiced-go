package invoiced

import (
	"strconv"
)

type TaskClient struct {
	*Client
	*Task
}

type Tasks []*TaskClient

func (c *Client) NewTask() *TaskClient {
	task := new(Task)
	return &TaskClient{c, task}
}

func (c *TaskClient) Create(request *TaskRequest) (*TaskClient, error) {
	endpoint := TaskEndpoint
	resp := new(TaskClient)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TaskClient) Retrieve(id int64) (*TaskClient, error) {
	endpoint := TaskEndpoint + "/" + strconv.FormatInt(id, 10)

	taskEndpoint := new(Task)

	task := &TaskClient{c.Client, taskEndpoint}

	_, err := c.Api.Get(endpoint, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (c *TaskClient) Update(request *TaskRequest) error {
	endpoint := TaskEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(TaskClient)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Task = resp.Task

	return nil
}

func (c *TaskClient) Delete() error {
	endpoint := TaskEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *TaskClient) ListAll(filter *Filter, sort *Sort) (Tasks, error) {
	endpoint := TaskEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

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
