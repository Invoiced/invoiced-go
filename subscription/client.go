package subscription

import (
	"github.com/Invoiced/invoiced-go/v2"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.SubscriptionRequest) (*invoiced.Subscription, error) {
	resp := new(invoiced.Subscription)
	err := c.Api.Create("/subscriptions", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.Subscription, error) {
	resp := new(invoiced.Subscription)
	_, err := c.Api.Get("/subscriptions/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) RetrievePlanCustomerExpanded(id int64) (*invoiced.Subscription, error) {
	resp := new(invoiced.Subscription)
	_, err := c.Api.Get("/subscriptions/"+strconv.FormatInt(id, 10) + "?expand=plan,customer,addons.catalog_item,addons.plan", resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.SubscriptionRequest) (*invoiced.Subscription, error) {
	endpoint := "/subscriptions/" + strconv.FormatInt(id, 10)
	resp := new(invoiced.Subscription)
	err := c.Api.Update(endpoint, request, resp)
	return resp, err
}

func (c *Client) Cancel(id int64) error {
	endpoint := "/subscriptions/" + strconv.FormatInt(id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Count() (int64, error) {
	return c.Api.Count("/subscriptions")
}

func (c *Client) ListAllQueryParameters(parameters map[string]string) (invoiced.Subscriptions, error) {
	endpoint := "/subscriptions"

	if len(parameters) > 0 {
		for key, value := range parameters {
			endpoint = invoiced.AddQueryParameter(endpoint, key, value)
		}
	}

	subscriptions := make(invoiced.Subscriptions, 0)

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

	return subscriptions, nil

}

func (c *Client) ListAllCanceled(canceled bool) (invoiced.Subscriptions, error) {
	parameters := make(map[string]string)

	if canceled {
		parameters["canceled"] = "1"
	}

	return c.ListAllQueryParameters(parameters)
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Subscriptions, error) {
	endpoint := "/subscriptions"
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invoiced.Subscriptions, 0)

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

	return subscriptions, nil
}

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Subscriptions, string, error) {
	endpoint := "/subscriptions"
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invoiced.Subscriptions, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &subscriptions)

	if err != nil {
		return nil, "", err
	}

	return subscriptions, nextEndpoint, nil
}

func (c *Client) Preview(request *invoiced.SubscriptionPreviewRequest) (*invoiced.SubscriptionPreview, error) {
	resp := new(invoiced.SubscriptionPreview)
	err := c.Api.Create("/subscriptions/preview", request, resp)
	return resp, err
}
