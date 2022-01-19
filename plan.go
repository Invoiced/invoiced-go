package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"strings"
)

type Plan struct {
	*Connection
	*invdendpoint.Plan
}

type Plans []*Plan

func (c *Connection) NewPlan() *Plan {
	plan := new(invdendpoint.Plan)
	return &Plan{c, plan}
}

func (c *Plan) Create(request *invdendpoint.PlanRequest) (*Plan, error) {
	endpoint := invdendpoint.PlanEndpoint
	resp := new(Plan)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Plan) Retrieve(id string) (*Plan, error) {
	endpoint := invdendpoint.PlanEndpoint + "/" + id
	plan := &Plan{c.Connection, new(invdendpoint.Plan)}

	_, err := c.retrieveDataFromAPI(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *Plan) RetrieveWithSubNumber(id string) (*Plan, error) {
	endpoint := invdendpoint.PlanEndpoint + "/" + id + "?include=num_subscriptions"
	plan := &Plan{c.Connection, new(invdendpoint.Plan)}

	_, err := c.retrieveDataFromAPI(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *Plan) Update(request *invdendpoint.PlanRequest) error {
	endpoint := invdendpoint.PlanEndpoint + "/" + c.Id
	resp := new(Plan)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Plan = resp.Plan

	return nil
}

func (c *Plan) Delete() error {
	endpoint := invdendpoint.PlanEndpoint + "/" + c.Id

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Plan) ListAllSubNumber(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Plans, error) {
	endpoint := invdendpoint.PlanEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	if strings.Contains(endpoint, "?") {
		endpoint = endpoint + "&include=num_subscriptions"
	} else {
		endpoint = endpoint + "?include=num_subscriptions"
	}

	plans := make(Plans, 0)

NEXT:
	tmpPlans := make(Plans, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpPlans)

	if err != nil {
		return nil, err
	}

	plans = append(plans, tmpPlans...)

	if endpoint != "" {
		goto NEXT
	}

	for _, plan := range plans {
		plan.Connection = c.Connection
	}

	return plans, nil
}

func (c *Plan) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Plans, error) {
	endpoint := invdendpoint.PlanEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	plans := make(Plans, 0)

NEXT:
	tmpPlans := make(Plans, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpPlans)

	if err != nil {
		return nil, err
	}

	plans = append(plans, tmpPlans...)

	if endpoint != "" {
		goto NEXT
	}

	for _, plan := range plans {
		plan.Connection = c.Connection
	}

	return plans, nil
}
