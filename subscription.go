package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"errors"
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

func (c *Connection) NewPreviewRequest() *invdendpoint.SubscriptionPreviewRequest {
	return &invdendpoint.SubscriptionPreviewRequest{}
}

func (c *Subscription) Count() (int64, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Subscription) Create(subscription *Subscription) (*Subscription, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint)

	if subscription == nil {
		return nil, errors.New("Subscription is nil")
	}

	subDataToCreate, err := SafeSubscriptionForCreation(subscription.Subscription)

	if err != nil {
		return nil, err
	}

	subResp := new(Subscription)

	apiErr := c.create(endPoint, subDataToCreate, subResp)

	if apiErr != nil {
		return nil, apiErr
	}

	subResp.Connection = c.Connection

	return subResp, nil

}

func (c *Subscription) Cancel() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Subscription) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint), c.Id)
	subResp := new(Subscription)

	subDataToUpdate, err := SafeSubscriptionsForUpdate(c.Subscription)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, subDataToUpdate, subResp)

	if apiErr != nil {
		return apiErr
	}

	c.Subscription = subResp.Subscription

	return nil

}

func (c *Subscription) Retrieve(id int64) (*Subscription, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint), id)

	custEndPoint := new(invdendpoint.Subscription)

	subscription := &Subscription{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, subscription)

	if apiErr != nil {
		return nil, apiErr
	}

	return subscription, nil

}

func (c *Subscription) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Subscriptions, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

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
	endPoint := c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

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


func (c *Subscription) Preview(subPreviewRequest *invdendpoint.SubscriptionPreviewRequest) (*invdendpoint.SubscriptionPreview, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.SubscriptionsEndPoint), c.Id) + "/preview"

	if subPreviewRequest == nil {
		return nil, errors.New("Subscription is nil")
	}

	subPreviewResp := new(invdendpoint.SubscriptionPreview)

	apiErr := c.create(endPoint, subPreviewRequest, subPreviewResp)

	if apiErr != nil {
		return nil, apiErr
	}


	return subPreviewResp, nil

}

//SafeSubscriptionForCreation prunes subscription data for just fields that can be used for creation of a subscription
func SafeSubscriptionForCreation(sub *invdendpoint.Subscription) (*invdendpoint.Subscription, error) {
	if sub == nil  {
		return nil, errors.New("Subscription is nil")
	}

	subData :=new(invdendpoint.Subscription)
	subData.Customer = sub.Customer
	subData.Plan = sub.Plan
	subData.StartDate = sub.StartDate
	subData.BillIn = sub.BillIn
	subData.Quantity = sub.Quantity
	subData.Addons = sub.Addons
	subData.Discounts = sub.Discounts
	subData.Cycles = sub.Cycles
	subData.SnapToNthDay = sub.SnapToNthDay
	subData.Paused = sub.Paused
	subData.ContractRenewalCycles = sub.ContractRenewalCycles
	subData.ContractRenewalMode = sub.ContractRenewalMode
	subData.Taxes = sub.Taxes
	subData.CancelAtPeriodEnd = sub.CancelAtPeriodEnd
	subData.Metadata = sub.Metadata



	return subData,nil
}

//SafeSubscriptionsForUpdate prunes subscription data for just fields that can be used for updating of a subscription
func SafeSubscriptionsForUpdate(sub *invdendpoint.Subscription) (*invdendpoint.Subscription, error) {
	if sub == nil  {
		return nil, errors.New("Subscription is nil")
	}

	subData :=new(invdendpoint.Subscription)

	subData.Plan = sub.Plan
	subData.StartDate = sub.StartDate
	subData.BillIn = sub.BillIn
	subData.Quantity = sub.Quantity
	subData.Addons = sub.Addons
	subData.Paused = sub.Paused
	subData.Discounts = sub.Discounts
	subData.ContractRenewalCycles = sub.ContractRenewalCycles
	subData.ContractRenewalMode = sub.ContractRenewalMode
	subData.CancelAtPeriodEnd = sub.CancelAtPeriodEnd
	subData.Prorate = sub.Prorate
	subData.ProrationDate = sub.ProrationDate


	return subData,nil
}