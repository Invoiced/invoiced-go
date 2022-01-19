package subscription

import (
	"errors"
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Count() (int64, error) {
	endpoint := invoiced.SubscriptionEndpoint

	count, err := c.Api.Count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *Client) Create(request *invoiced.SubscriptionRequest) (*invoiced.Subscription, error) {
	endpoint := invoiced.SubscriptionEndpoint
	resp := c.NewSubscription()

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id int64) (*invoiced.Subscription, error) {
	endpoint := invoiced.SubscriptionEndpoint + "/" + strconv.FormatInt(id, 10)

	subscription := &Client{c.Client, new(invoiced.Subscription)}

	_, err := c.Api.Get(endpoint, subscription)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (c *Client) Update(request *invoiced.SubscriptionRequest) error {
	endpoint := invoiced.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := c.NewSubscription()

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Subscription = resp.Subscription

	return nil
}

func (c *Client) Cancel() error {
	endpoint := invoiced.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAllQueryParameters(parameters map[string]string) (invoiced.Subscriptions, error) {
	endpoint := invoiced.SubscriptionEndpoint

	if len(parameters) > 0 {
		for key, value := range parameters {
			endpoint = addQueryParameter(endpoint, key, value)
		}
	}

	subscriptions := make(invoiced.Subscriptions, 0)
	subscriptionsToReturn := make(invoiced.Subscriptions, 0)

NEXT:
	tmpSubscriptions := make(invoiced.Subscriptions, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpSubscriptions)

	if err != nil {
		return nil, err
	}

	subscriptions = append(subscriptions, tmpSubscriptions...)

	if endpoint != "" {
		goto NEXT
	}

	for _, subscription := range subscriptions {
		sub := c.Client.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn, sub)
	}

	return subscriptionsToReturn, nil

}

func (c *Client) ListAllCanceled(canceled bool) (invoiced.Subscriptions, error) {
	parameters := make(map[string]string)

	if canceled {
		parameters["canceled"] = "1"
	}

	return c.ListAllQueryParameters(parameters)
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Subscriptions, error) {
	endpoint := invoiced.SubscriptionEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invoiced.Subscriptions, 0)
	subscriptionsToReturn := make(invoiced.Subscriptions, 0)

NEXT:
	tmpSubscriptions := make(invoiced.Subscriptions, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpSubscriptions)

	if err != nil {
		return nil, err
	}

	subscriptions = append(subscriptions, tmpSubscriptions...)

	if endpoint != "" {
		goto NEXT
	}

	for _, subscription := range subscriptions {
		sub := c.Client.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn, sub)
	}

	return subscriptionsToReturn, nil
}

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Subscriptions, string, error) {
	endpoint := invoiced.SubscriptionEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invoiced.Subscriptions, 0)
	subscriptionsToReturn := make(invoiced.Subscriptions, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &subscriptions)

	if err != nil {
		return nil, "", err
	}

	for _, subscription := range subscriptions {
		sub := c.Client.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn, sub)
	}

	return subscriptionsToReturn, nextEndpoint, nil
}

func (c *Client) Preview(request *invoiced.SubscriptionPreviewRequest) (*invoiced.SubscriptionPreview, error) {
	endpoint := invoiced.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/preview"

	if request == nil {
		return nil, errors.New("Client is nil")
	}

	subPreviewResp := new(invoiced.SubscriptionPreview)

	err := c.Api.Create(endpoint, request, subPreviewResp)

	if err != nil {
		return nil, err
	}

	return subPreviewResp, nil
}
