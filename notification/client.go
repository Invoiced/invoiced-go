package notification

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.NotificationRequest) (*invoiced.Notification, error) {
	resp := new(invoiced.Notification)
	err := c.Api.Create("/notifications", request, resp)
	return resp, err
}

func (c *Client) Update(request *invoiced.NotificationRequest, id int64) (*invoiced.Notification, error) {
	resp := new(invoiced.Notification)
	err := c.Api.Update("/notifications/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/notifications/" + strconv.FormatInt(id, 10))
}

func (c *Client) Retrieve(id int64) (*invoiced.Notification, error) {
	resp := new(invoiced.Notification)
	_, err := c.Api.Get("/notifications/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Notifications, error) {
	endpoint := invoiced.AddFilterAndSort("/notifications", filter, sort)

	notifications := make(invoiced.Notifications, 0)

NEXT:
	tmpNotifications := make(invoiced.Notifications, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpNotifications)

	if err != nil {
		return nil, err
	}

	notifications = append(notifications, tmpNotifications...)

	if endpoint != "" {
		goto NEXT
	}

	return notifications, nil
}
