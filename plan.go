
package invdapi

import (
"errors"
"github.com/Invoiced/invoiced-go/invdendpoint"
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

func (c *Plan) Create(plan *Plan) (*Plan, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.PlansEndPoint)

	planResp := new(Plan)


	if plan == nil {
		return nil, errors.New("plan is nil")
	}


	//safe prune file data for creation
	invdPlanDataToCreate, err := SafePlanForCreation(plan.Plan)

	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdPlanDataToCreate, planResp)

	if apiErr != nil {
		return nil, apiErr
	}

	planResp.Connection = c.Connection

	return planResp, nil

}

func (c *Plan) Save() error {
	endPoint := c.MakeEndPointURL(invdendpoint.PlansEndPoint) + "/" + c.Id

	planResp := new(Plan)

	planDataToUpdate, err := SafePlanForUpdating(c.Plan)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, planDataToUpdate, planResp)

	if apiErr != nil {
		return apiErr
	}

	c.Plan = planResp.Plan

	return nil

}

func (c *Plan) Delete() error {
	endPoint := c.MakeEndPointURL(invdendpoint.PlansEndPoint) + "/" + c.Id

	err := c.delete(endPoint)

	if err != nil {
		return err
	}

	return nil

}

func (c *Plan) Retrieve(id string) (*Plan, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.PlansEndPoint) + "/" + id

	planEndPoint := new(invdendpoint.Plan)

	plan := &Plan{c.Connection, planEndPoint}

	_, err := c.retrieveDataFromAPI(endPoint, plan)

	if err != nil {
		return nil, err
	}

	return plan, nil

}

func (c *Plan) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Plans, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.PlansEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	plans := make(Plans, 0)

NEXT:
	tmpPlans := make(Plans, 0)

	endPointTmp, apiErr := c.retrieveDataFromAPI(endPoint, &tmpPlans)

	if apiErr != nil {
		return nil, apiErr
	}

	plans = append(plans, tmpPlans...)

	if endPointTmp != "" {
		goto NEXT
	}

	for _, plan := range plans {
		plan.Connection = c.Connection

	}

	return plans, nil

}

//SafeCustomerForCreation prunes plan data for just fields that can be used for creation of a plan
func SafePlanForCreation(plan *invdendpoint.Plan) (*invdendpoint.Plan, error) {

	if plan == nil {
		return nil, errors.New("plan is nil")
	}

	planData := new(invdendpoint.Plan)
	planData.Id = plan.Id
	planData.Object = plan.Object
	planData.CatalogItem = plan.CatalogItem
	planData.Name = plan.Name
	planData.Currency = plan.Currency
	planData.Amount = plan.Amount
	planData.PricingMode = plan.PricingMode
	planData.QuantityType = plan.QuantityType
	planData.Interval = plan.Interval
	planData.IntervalCount = plan.IntervalCount
	planData.Tiers = plan.Tiers
	planData.CreatedAt = plan.CreatedAt
	planData.Metadata = plan.Metadata

	return planData, nil
}

//SafeTaskForUpdating prunes plan data for just fields that can be used for updating of a plan
func SafePlanForUpdating(plan *invdendpoint.Plan) (*invdendpoint.Plan, error) {

	if plan == nil {
		return nil, errors.New("plan is nil")
	}

	planData := new(invdendpoint.Plan)

	planData.Name = plan.Name
	planData.Metadata = plan.Metadata

	return planData, nil
}
