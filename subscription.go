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

func (c *Connection) NewPreviewRequest() *invdendpoint.SubscriptionPreviewRequest {
	return &invdendpoint.SubscriptionPreviewRequest{}
}

func (c *Subscription) Count() (int64, error) {
	endpoint := invdendpoint.SubscriptionEndpoint

	count, apiErr := c.count(endpoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil
}

func (c *Subscription) Create(subscription *Subscription) (*Subscription, error) {
	endpoint := invdendpoint.SubscriptionEndpoint

	if subscription == nil {
		return nil, errors.New("Subscription is nil")
	}

	subDataToCreate, err := SafeSubscriptionForCreation(subscription.Subscription)
	if err != nil {
		return nil, err
	}

	subResp := c.NewSubscription()

	apiErr := c.create(endpoint, subDataToCreate, subResp)

	if apiErr != nil {
		return nil, apiErr
	}

	subResp.Connection = c.Connection

	return subResp, nil
}

func (c *Subscription) Cancel() error {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)

	apiErr := c.delete(endpoint)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (c *Subscription) Save() error {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10)
	subResp := c.NewSubscription()

	subDataToUpdate, err := SafeSubscriptionsForUpdate(c.Subscription)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, subDataToUpdate, subResp)

	if apiErr != nil {
		return apiErr
	}

	c.Subscription = subResp.Subscription

	return nil
}

func (c *Subscription) Retrieve(id int64) (*Subscription, error) {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(id, 10)

	custEndpoint := new(invdendpoint.Subscription)

	subscription := &Subscription{c.Connection, custEndpoint}

	_, apiErr := c.retrieveDataFromAPI(endpoint, subscription)

	if apiErr != nil {
		return nil, apiErr
	}

	return subscription, nil
}

func (c *Subscription) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Subscriptions, error) {
	endpoint := invdendpoint.SubscriptionEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invdendpoint.Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions,0)

NEXT:
	tmpSubscriptions := make(invdendpoint.Subscriptions, 0)


	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpSubscriptions)

	if apiErr != nil {
		return nil, apiErr
	}

	subscriptions = append(subscriptions, tmpSubscriptions...)

	if endpoint != "" {
		goto NEXT
	}

	for _, subscription := range subscriptions {
		sub := c.Connection.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn,sub)
	}

	return subscriptionsToReturn, nil
}

func (c *Subscription) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Subscriptions, string, error) {
	endpoint := invdendpoint.SubscriptionEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	subscriptions := make(invdendpoint.Subscriptions, 0)
	subscriptionsToReturn := make(Subscriptions,0)

	nextEndpoint, apiErr := c.retrieveDataFromAPI(endpoint, &subscriptions)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, subscription := range subscriptions {
		sub := c.Connection.NewSubscription()
		subData := subscription
		sub.Subscription = &subData
		subscriptionsToReturn = append(subscriptionsToReturn,sub)
	}

	return subscriptionsToReturn, nextEndpoint, nil
}

func (c *Subscription) Preview(subPreviewRequest *invdendpoint.SubscriptionPreviewRequest) (*invdendpoint.SubscriptionPreview, error) {
	endpoint := invdendpoint.SubscriptionEndpoint + "/" + strconv.FormatInt(c.Id, 10) + "/preview"

	if subPreviewRequest == nil {
		return nil, errors.New("Subscription is nil")
	}

	subPreviewResp := new(invdendpoint.SubscriptionPreview)

	apiErr := c.create(endpoint, subPreviewRequest, subPreviewResp)

	if apiErr != nil {
		return nil, apiErr
	}

	return subPreviewResp, nil
}

// SafeSubscriptionForCreation prunes subscription data for just fields that can be used for creation of a subscription
func SafeSubscriptionForCreation(sub *invdendpoint.Subscription) (*invdendpoint.Subscription, error) {
	if sub == nil {
		return nil, errors.New("Subscription is nil")
	}

	subData := new(invdendpoint.Subscription)
	subData.Customer = sub.Customer
	subData.Plan = sub.Plan
	subData.StartDate = sub.StartDate
	subData.BillIn = sub.BillIn
	subData.BillInAdvanceDays = sub.BillInAdvanceDays
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

	return subData, nil
}

// SafeSubscriptionsForUpdate prunes subscription data for just fields that can be used for updating of a subscription
func SafeSubscriptionsForUpdate(sub *invdendpoint.Subscription) (*invdendpoint.Subscription, error) {
	if sub == nil {
		return nil, errors.New("Subscription is nil")
	}

	subData := new(invdendpoint.Subscription)

	subData.Plan = sub.Plan
	subData.StartDate = sub.StartDate
	subData.BillIn = sub.BillIn
	subData.BillInAdvanceDays = sub.BillInAdvanceDays
	subData.Quantity = sub.Quantity
	subData.Addons = sub.Addons
	subData.Paused = sub.Paused
	subData.Discounts = sub.Discounts
	subData.ContractRenewalCycles = sub.ContractRenewalCycles
	subData.ContractRenewalMode = sub.ContractRenewalMode
	subData.CancelAtPeriodEnd = sub.CancelAtPeriodEnd
	subData.Prorate = sub.Prorate
	subData.ProrationDate = sub.ProrationDate

	return subData, nil
}
