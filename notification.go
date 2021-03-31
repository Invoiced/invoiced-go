package invdapi

import (
	"strconv"
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Notification struct {
	*Connection
	*invdendpoint.Notification
}

type Notifications []*Notification

func (c *Connection) NewNotification() *Notification {
	notification := new(invdendpoint.Notification)
	return &Notification{c, notification}
}

func (c *Notification) Create(notificationRequest *invdendpoint.NotificationRequest) (*Notification, error) {
	endpoint := invdendpoint.NotificationEndpoint

	notificationResp := new(Notification)

	apiErr := c.create(endpoint, notificationRequest, notificationResp)

	if apiErr != nil {
		return nil, apiErr
	}

	notificationResp.Connection = c.Connection

	return notificationResp, nil
}

func (c *Notification) Save(notificationRequest *invdendpoint.NotificationRequest, id int64) error {
	endpoint := invdendpoint.NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	notifResp := new(Notification)

	apiErr := c.update(endpoint, notificationRequest, notifResp)

	if apiErr != nil {
		return apiErr
	}

	notifResp.Connection = c.Connection

	return nil
}

func (c *Notification) Delete(id int64) error {
	endpoint := invdendpoint.NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	err := c.delete(endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (c *Notification) Retrieve(id int64) (*Notification, error) {
	endpoint := invdendpoint.NotificationEndpoint + "/" + strconv.FormatInt(id, 10)

	notifResp := new(Notification)

	_, err := c.retrieveDataFromAPI(endpoint, notifResp)
	if err != nil {
		return nil, err
	}

	notifResp.Connection = c.Connection

	return notifResp, nil
}

func (c *Notification) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Notifications, error) {
	endpoint := invdendpoint.NotificationEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	notifications := make(Notifications, 0)

NEXT:
	tmpNotifications := make(Notifications, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpNotifications)

	if apiErr != nil {
		return nil, apiErr
	}

	notifications = append(notifications, tmpNotifications...)

	if endpoint != "" {
		goto NEXT
	}

	for _, notification := range notifications {
		notification.Connection = c.Connection
	}

	return notifications, nil
}


