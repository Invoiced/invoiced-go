package invdapi

import (
	"errors"
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

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
	endpoint := invdendpoint.SubscriptionEndpoint

	count, err := c.count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Subscription) Create(request *invdendpoint.SubscriptionRequest) (*Subscription, error) {
	endpoint := invdendpoint.SubscriptionEndpoint
	resp := c.NewSubscription()

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Subscription) Retrieve(id int64) (*Subscription, error) {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(id, 10)

	subscription := &Subscription{c.Connection, new(invdendpoint.Subscription)}

	_, err := c.retrieveDataFromAPI(endpoint, subscription)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (c *Subscription) Update(request *invdendpoint.SubscriptionRequest) error {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := c.NewSubscription()

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Subscription = resp.Subscription

	return nil
}

func (c *Subscription) Cancel() error {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Subscription) ListAllQueryParameters(parameters map[string]string) (Subscriptions, error) {
	endpoint := invdendpoint.SubscriptionEndpoint

	if len(parameters) > 0 {
		for key, value := range parameters {
			endpoint = addQueryParameter(endpoint, key, value)
		}
	}

	subscriptions := make(invdendpoint.Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions, 0)

NEXT:
	tmpSubscriptions := make(invdendpoint.Subscriptions, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpSubscriptions)

	if err != nil {
		return nil, err
	}

	subscriptions = append(subscriptions, tmpSubscriptions...)

	if endpoint != "" {
		goto NEXT
	}

	for _, subscription := range subscriptions {
		sub := c.Connection.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn, sub)
	}

	return subscriptionsToReturn, nil

}

func (c *Subscription) ListAllCanceled(canceled bool) (Subscriptions, error) {
	parameters := make(map[string]string)

	if canceled {
		parameters["canceled"] = "1"
	}

	return c.ListAllQueryParameters(parameters)
}

func (c *Subscription) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Subscriptions, error) {
	endpoint := invdendpoint.SubscriptionEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invdendpoint.Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions, 0)

NEXT:
	tmpSubscriptions := make(invdendpoint.Subscriptions, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpSubscriptions)

	if err != nil {
		return nil, err
	}

	subscriptions = append(subscriptions, tmpSubscriptions...)

	if endpoint != "" {
		goto NEXT
	}

	for _, subscription := range subscriptions {
		sub := c.Connection.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn, sub)
	}

	return subscriptionsToReturn, nil
}

func (c *Subscription) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Subscriptions, string, error) {
	endpoint := invdendpoint.SubscriptionEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invdendpoint.Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions, 0)

	nextEndpoint, err := c.retrieveDataFromAPI(endpoint, &subscriptions)

	if err != nil {
		return nil, "", err
	}

	for _, subscription := range subscriptions {
		sub := c.Connection.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn, sub)
	}

	return subscriptionsToReturn, nextEndpoint, nil
}

func (c *Subscription) Preview(request *invdendpoint.SubscriptionPreviewRequest) (*invdendpoint.SubscriptionPreview, error) {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/preview"

	if request == nil {
		return nil, errors.New("Subscription is nil")
	}

	subPreviewResp := new(invdendpoint.SubscriptionPreview)

	err := c.create(endpoint, request, subPreviewResp)

	if err != nil {
		return nil, err
	}

	return subPreviewResp, nil
}
