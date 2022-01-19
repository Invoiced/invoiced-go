package invoiced

import (
	"strconv"
)

type NotificationClient struct {
	*Client
	*NotificationClient
}

func (c *Client) NewNotification() *NotificationClient {
	notification := new(NotificationClient)
	return &NotificationClient{c, notification}
}

func (c *NotificationClient) Create(notificationRequest *NotificationRequest) (*Notification, error) {
	endpoint := NotificationEndpoint

	notificationResp := new(Notification)

	err := c.Api.Create(endpoint, notificationRequest, notificationResp)

	if err != nil {
		return nil, err
	}

	return notificationResp, nil
}

func (c *NotificationClient) Save(notificationRequest *NotificationRequest, id int64) error {
	endpoint := NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	notifResp := new(Notification)

	err := c.Api.Update(endpoint, notificationRequest, notifResp)

	if err != nil {
		return err
	}

	return nil
}

func (c *NotificationClient) Delete(id int64) error {
	endpoint := NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	err := c.Api.Delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *NotificationClient) Retrieve(id int64) (*Notification, error) {
	endpoint := NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	notifResp := new(Notification)

	_, err := c.Api.Get(endpoint, notifResp)
	if err != nil {
		return nil, err
	}

	return notifResp, nil
}

func (c *NotificationClient) ListAll(filter *Filter, sort *Sort) (Notifications, error) {
	endpoint := NotificationEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

	notifications := make(Notifications, 0)

NEXT:
	tmpNotifications := make(Notifications, 0)

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
