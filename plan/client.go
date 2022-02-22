package plan

import (
	"strings"

	"github.com/strongdm/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.PlanRequest) (*invoiced.Plan, error) {
	resp := new(invoiced.Plan)
	err := c.Api.Create("/plans", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id string) (*invoiced.Plan, error) {
	resp := new(invoiced.Plan)
	_, err := c.Api.Get("/plans/"+id, resp)
	return resp, err
}

func (c *Client) RetrieveWithSubNumber(id string) (*invoiced.Plan, error) {
	resp := new(invoiced.Plan)
	_, err := c.Api.Get("/plans/"+id+"?include=num_subscriptions", resp)
	return resp, err
}

func (c *Client) Update(id string, request *invoiced.PlanRequest) (*invoiced.Plan, error) {
	resp := new(invoiced.Plan)
	err := c.Api.Update("/plans/"+id, request, resp)
	return resp, err
}

func (c *Client) Delete(id string) error {
	return c.Api.Delete("/plans/" + id)
}

func (c *Client) ListAllSubNumber(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Plans, error) {
	endpoint := invoiced.AddFilterAndSort("/plans", filter, sort)

	if strings.Contains(endpoint, "?") {
		endpoint = endpoint + "&include=num_subscriptions"
	} else {
		endpoint = endpoint + "?include=num_subscriptions"
	}

	plans := make(invoiced.Plans, 0)

NEXT:
	tmpPlans := make(invoiced.Plans, 0)

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

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Plans, error) {
	endpoint := invoiced.AddFilterAndSort("/plans", filter, sort)

	plans := make(invoiced.Plans, 0)

NEXT:
	tmpPlans := make(invoiced.Plans, 0)

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
