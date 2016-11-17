package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

const defaultExpandSubscription = "addons.catalog_item"

type Subscription struct {
	*Connection
	*invdendpoint.Subscription
}

type Subscriptions []*Subscription

func (c *Connection) NewSubscription() *Subscription {
	subscription := new(invdendpoint.Subscription)
	return &Subscription{c, subscription}

}

func (c *Subscription) Count() (int64, error) {
	endPoint := c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Subscription) Create(subscription *Subscription) (*Subscription, error) {
	endPoint := c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint)
	subResp := new(Subscription)

	apiErr := c.create(endPoint, subscription, subResp)

	if apiErr != nil {
		return nil, apiErr
	}

	subResp.Connection = c.Connection

	return subResp, nil

}

func (c *Subscription) Cancel() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Subscription) Save() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint), c.Id)
	subResp := new(Subscription)
	apiErr := c.update(endPoint, c, subResp)

	if apiErr != nil {
		return apiErr
	}

	c.Subscription = subResp.Subscription

	return nil

}

func (c *Subscription) Retrieve(id int64) (*Subscription, error) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint), id)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandSubscription)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	custEndPoint := new(invdendpoint.Subscription)

	subscription := &Subscription{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, subscription)

	if apiErr != nil {
		return nil, apiErr
	}

	return subscription, nil

}

func (c *Subscription) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Subscriptions, error) {
	endPoint := c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandSubscription)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	subscriptions := make(Subscriptions, 0)

NEXT:
	tmpSubscriptions := make(Subscriptions, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpSubscriptions)

	if apiErr != nil {
		return nil, apiErr
	}

	subscriptions = append(subscriptions, tmpSubscriptions...)

	if endPoint != "" {
		goto NEXT
	}

	for _, subscription := range subscriptions {
		subscription.Connection = c.Connection

	}

	return subscriptions, nil

}

func (c *Subscription) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Subscriptions, string, error) {
	endPoint := c.makeEndPointURL(invdendpoint.SubscriptionsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandSubscription)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	subscriptions := make(Subscriptions, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &subscriptions)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, subscription := range subscriptions {
		subscription.Connection = c.Connection

	}

	return subscriptions, nextEndPoint, nil

}
