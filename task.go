package invdapi

import (
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Task struct {
	*Connection
	*invdendpoint.Task
}

type Tasks []*Task

func (c *Connection) NewTask() *Task {
	task := new(invdendpoint.Task)
	return &Task{c, task}
}

func (c *Task) Create(request *invdendpoint.TaskRequest) (*Task, error) {
	endpoint := invdendpoint.TaskEndpoint
	resp := new(Task)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Task) Retrieve(id int64) (*Task, error) {
	endpoint := invdendpoint.TaskEndpoint + "/" + strconv.FormatInt(id, 10)

	taskEndpoint := new(invdendpoint.Task)

	task := &Task{c.Connection, taskEndpoint}

	_, err := c.retrieveDataFromAPI(endpoint, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (c *Task) Update(request *invdendpoint.TaskRequest) error {
	endpoint := invdendpoint.TaskEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := new(Task)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Task = resp.Task

	return nil
}

func (c *Task) Delete() error {
	endpoint := invdendpoint.TaskEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Task) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Tasks, error) {
	endpoint := invdendpoint.TaskEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	tasks := make(Tasks, 0)

NEXT:
	tmpTasks := make(Tasks, 0)

	endpointTmp, err := c.retrieveDataFromAPI(endpoint, &tmpTasks)

	if err != nil {
		return nil, err
	}

	tasks = append(tasks, tmpTasks...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, task := range tasks {
		task.Connection = c.Connection
	}

	return tasks, nil
}
