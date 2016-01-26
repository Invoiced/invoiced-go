package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

func (c *Connection) ListSubscription(id int64) (*invdendpoint.Subscription, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint), id)

	subscription := new(invdendpoint.Subscription)

	_, apiErr := c.retrieveDataFromAPI(endPoint, subscription)

	if apiErr != nil {
		return nil, apiErr
	}

	return subscription, apiErr

}

func (c *Connection) DeleteSubscription(id int64) *APIError {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint), id)

	apiErr := c.delete(endPoint)

	return apiErr

}

func (c *Connection) UpdateSubscription(id int64, subscriptionToUpdate *invdendpoint.Subscription) (*invdendpoint.Subscription, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint), id)

	subscriptionCreated := new(invdendpoint.Subscription)

	apiErr := c.update(endPoint, subscriptionToUpdate, subscriptionCreated)

	return subscriptionCreated, apiErr

}

func (c *Connection) CreateSubscription(subscriptionToCreate *invdendpoint.Subscription) (*invdendpoint.Subscription, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint)

	subscriptionCreated := new(invdendpoint.Subscription)

	apiErr := c.create(endPoint, subscriptionToCreate, subscriptionCreated)

	return subscriptionCreated, apiErr

}
