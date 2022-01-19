package invoiced

import (
	"strings"
)

type PlanClient struct {
	*Client
	*PlanClient
}

type Plans []*PlanClient

func (c *Client) NewPlan() *PlanClient {
	plan := new(PlanClient)
	return &PlanClient{c, plan}
}

func (c *PlanClient) Create(request *PlanRequest) (*PlanClient, error) {
	endpoint := PlanEndpoint
	resp := new(PlanClient)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *PlanClient) Retrieve(id string) (*PlanClient, error) {
	endpoint := PlanEndpoint + "/" + id
	plan := &PlanClient{c.Client, new(PlanClient)}

	_, err := c.Api.Get(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *PlanClient) RetrieveWithSubNumber(id string) (*PlanClient, error) {
	endpoint := PlanEndpoint + "/" + id + "?include=num_subscriptions"
	plan := &PlanClient{c.Client, new(PlanClient)}

	_, err := c.Api.Get(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *PlanClient) Update(request *PlanRequest) error {
	endpoint := PlanEndpoint + "/" + c.Id
	resp := new(PlanClient)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.PlanClient = resp.PlanClient

	return nil
}

func (c *PlanClient) Delete() error {
	endpoint := PlanEndpoint + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *PlanClient) ListAllSubNumber(filter *Filter, sort *Sort) (Plans, error) {
	endpoint := PlanEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

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

func (c *PlanClient) ListAll(filter *Filter, sort *Sort) (Plans, error) {
	endpoint := PlanEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

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
