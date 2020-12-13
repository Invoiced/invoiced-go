package invdapi

import (
	"errors"
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

func (c *Task) Create(task *Task) (*Task, error) {
	endpoint := invdendpoint.TaskEndpoint

	taskResp := new(Task)

	// safe prune file data for creation
	invdTaskDataToCreate, err := SafeTaskForCreation(task.Task)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endpoint, invdTaskDataToCreate, taskResp)

	if apiErr != nil {
		return nil, apiErr
	}

	taskResp.Connection = c.Connection

	return taskResp, nil
}

func (c *Task) Save() error {
	endpoint := invdendpoint.TaskEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	taskResp := new(Task)

	taskDataToUpdate, err := SafeTaskForUpdating(c.Task)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, taskDataToUpdate, taskResp)

	if apiErr != nil {
		return apiErr
	}

	c.Task = taskResp.Task

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

func (c *Task) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Tasks, error) {
	endpoint := invdendpoint.TaskEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	tasks := make(Tasks, 0)

NEXT:
	tmpTasks := make(Tasks, 0)

	endpointTmp, apiErr := c.retrieveDataFromAPI(endpoint, &tmpTasks)

	if apiErr != nil {
		return nil, apiErr
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

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafeTaskForCreation(task *invdendpoint.Task) (*invdendpoint.Task, error) {
	if task == nil {
		return nil, errors.New("task is nil")
	}

	taskData := new(invdendpoint.Task)
	taskData.Name = task.Name
	taskData.Action = task.Action
	taskData.CustomerId = task.CustomerId
	taskData.UserId = task.UserId
	taskData.DueDate = task.DueDate

	return taskData, nil
}

// SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafeTaskForUpdating(task *invdendpoint.Task) (*invdendpoint.Task, error) {
	if task == nil {
		return nil, errors.New("task is nil")
	}

	taskData := new(invdendpoint.Task)
	taskData.Name = task.Name
	taskData.Action = task.Action
	taskData.CustomerId = task.CustomerId
	taskData.UserId = task.UserId
	taskData.DueDate = task.DueDate

	return taskData, nil
}
