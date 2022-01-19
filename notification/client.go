package notification

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(notificationRequest *invoiced.NotificationRequest) (*invoiced.Notification, error) {
	endpoint := invoiced.NotificationEndpoint

	notificationResp := new(invoiced.Notification)

	err := c.Api.Create(endpoint, notificationRequest, notificationResp)

	if err != nil {
		return nil, err
	}

	return notificationResp, nil
}

func (c *Client) Save(notificationRequest *invoiced.NotificationRequest, id int64) error {
	endpoint := invoiced.NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	notifResp := new(invoiced.Notification)

	err := c.Api.Update(endpoint, notificationRequest, notifResp)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(id int64) error {
	endpoint := invoiced.NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	err := c.Api.Delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Retrieve(id int64) (*invoiced.Notification, error) {
	endpoint := invoiced.NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	notifResp := new(invoiced.Notification)

	_, err := c.Api.Get(endpoint, notifResp)
	if err != nil {
		return nil, err
	}

	return notifResp, nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Notifications, error) {
	endpoint := invoiced.NotificationEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

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
