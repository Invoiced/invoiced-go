package invdapi

import (
	"errors"

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
	endPoint := c.MakeEndPointURL(invdendpoint.TaskEndPoint)

	taskResp := new(Task)

	// safe prune file data for creation
	invdTaskDataToCreate, err := SafeTaskForCreation(task.Task)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdTaskDataToCreate, taskResp)

	if apiErr != nil {
		return nil, apiErr
	}

	taskResp.Connection = c.Connection

	return taskResp, nil
}

func (c *Task) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.TaskEndPoint), c.Id)

	taskResp := new(Task)

	taskDataToUpdate, err := SafeTaskForUpdating(c.Task)
	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, taskDataToUpdate, taskResp)

	if apiErr != nil {
		return apiErr
	}

	c.Task = taskResp.Task

	return nil
}

func (c *Task) Delete() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.TaskEndPoint), c.Id)

	err := c.delete(endPoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Task) Retrieve(id int64) (*Task, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.TaskEndPoint), id)

	taskEndPoint := new(invdendpoint.Task)

	task := &Task{c.Connection, taskEndPoint}

	_, err := c.retrieveDataFromAPI(endPoint, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (c *Task) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Tasks, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.TaskEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	tasks := make(Tasks, 0)

NEXT:
	tmpTasks := make(Tasks, 0)

	endPointTmp, apiErr := c.retrieveDataFromAPI(endPoint, &tmpTasks)

	if apiErr != nil {
		return nil, apiErr
	}

	tasks = append(tasks, tmpTasks...)

	if endPointTmp != "" {
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
