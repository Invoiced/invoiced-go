package plan

import (
	"github.com/Invoiced/invoiced-go"
	"strings"
)

type Client struct {
	*invoiced.Api
}

type Plans []*Client

func (c *Client) Create(request *invoiced.PlanRequest) (*Client, error) {
	endpoint := invoiced.PlanEndpoint
	resp := new(Client)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id string) (*Client, error) {
	endpoint := invoiced.PlanEndpoint + "/" + id
	plan := &Client{c.Client, new(Client)}

	_, err := c.Api.Get(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *Client) RetrieveWithSubNumber(id string) (*Client, error) {
	endpoint := invoiced.PlanEndpoint + "/" + id + "?include=num_subscriptions"
	plan := &Client{c.Client, new(Client)}

	_, err := c.Api.Get(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *Client) Update(request *invoiced.PlanRequest) error {
	endpoint := invoiced.PlanEndpoint + "/" + c.Id
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Client = resp.Client

	return nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.PlanEndpoint + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAllSubNumber(filter *invoiced.Filter, sort *invoiced.Sort) (Plans, error) {
	endpoint := invoiced.PlanEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	if strings.Contains(endpoint, "?") {
		endpoint = endpoint + "&include=num_subscriptions"
	} else {
		endpoint = endpoint + "?include=num_subscriptions"
	}

	plans := make(Plans, 0)

NEXT:
	tmpPlans := make(Plans, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpPlans)

	if err != nil {
		return nil, err
	}

	plans = append(plans, tmpPlans...)

	if endpoint != "" {
		goto NEXT
	}

	return plans, nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (Plans, error) {
	endpoint := invoiced.PlanEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	plans := make(Plans, 0)

NEXT:
	tmpPlans := make(Plans, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpPlans)

	if err != nil {
		return nil, err
	}

	plans = append(plans, tmpPlans...)

	if endpoint != "" {
		goto NEXT
	}

	return plans, nil
}
