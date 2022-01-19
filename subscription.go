package invoiced

import (
	"errors"
	"strconv"
)

type SubscriptionClient struct {
	*Client
	*Subscription
}

func (c *SubscriptionClient) Count() (int64, error) {
	endpoint := SubscriptionEndpoint

	count, err := c.Api.Count(endpoint)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (c *SubscriptionClient) Create(request *SubscriptionRequest) (*Subscription, error) {
	endpoint := SubscriptionEndpoint
	resp := c.NewSubscription()

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *SubscriptionClient) Retrieve(id int64) (*Subscription, error) {
	endpoint := SubscriptionEndpoint + "/" + strconv.FormatInt(id, 10)

	subscription := &SubscriptionClient{c.Client, new(Subscription)}

	_, err := c.Api.Get(endpoint, subscription)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (c *SubscriptionClient) Update(request *SubscriptionRequest) error {
	endpoint := SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	resp := c.NewSubscription()

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Subscription = resp.Subscription

	return nil
}

func (c *SubscriptionClient) Cancel() error {
	endpoint := SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *SubscriptionClient) ListAllQueryParameters(parameters map[string]string) (Subscriptions, error) {
	endpoint := SubscriptionEndpoint

	if len(parameters) > 0 {
		for key, value := range parameters {
			endpoint = addQueryParameter(endpoint, key, value)
		}
	}

	subscriptions := make(Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions, 0)

NEXT:
	tmpSubscriptions := make(Subscriptions, 0)

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

func (c *SubscriptionClient) ListAllCanceled(canceled bool) (Subscriptions, error) {
	parameters := make(map[string]string)

	if canceled {
		parameters["canceled"] = "1"
	}

	return c.ListAllQueryParameters(parameters)
}

func (c *SubscriptionClient) ListAll(filter *Filter, sort *Sort) (Subscriptions, error) {
	endpoint := SubscriptionEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)

	subscriptions := make(Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions, 0)

NEXT:
	tmpSubscriptions := make(Subscriptions, 0)

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

func (c *SubscriptionClient) List(filter *Filter, sort *Sort) (Subscriptions, string, error) {
	endpoint := SubscriptionEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)

	subscriptions := make(Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions, 0)

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

func (c *SubscriptionClient) Preview(request *SubscriptionPreviewRequest) (*SubscriptionPreview, error) {
	endpoint := SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/preview"

	if request == nil {
		return nil, errors.New("SubscriptionClient is nil")
	}

	subPreviewResp := new(SubscriptionPreview)

	err := c.Api.Create(endpoint, request, subPreviewResp)

	if err != nil {
		return nil, err
	}

	return subPreviewResp, nil
}
