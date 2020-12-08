package invdapi

import (
	"errors"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Item struct {
	*Connection
	*invdendpoint.Item
}

type Items []*Item

func (c *Connection) NewItem() *Item {
	item := new(invdendpoint.Item)
	return &Item{c, item}
}

func (c *Item) Create(item *Item) (*Item, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.ItemEndPoint)

	itemResp := new(Item)

	if item == nil {
		return nil, errors.New("item is nil")
	}

	// safe prune file data for creation
	invdItemDataToCreate, err := SafeItemForCreation(item.Item)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdItemDataToCreate, itemResp)

	if apiErr != nil {
		return nil, apiErr
	}

	itemResp.Connection = c.Connection

	return itemResp, nil
}

func (c *Item) Save() error {
	endPoint := c.MakeEndPointURL(invdendpoint.ItemEndPoint) + "/" + c.Id

	itemResp := new(Item)

	itemDataToUpdate, err := SafeItemForUpdating(c.Item)
	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, itemDataToUpdate, itemResp)

	if apiErr != nil {
		return apiErr
	}

	c.Item = itemResp.Item

	return nil
}

func (c *Item) Delete() error {
	endPoint := c.MakeEndPointURL(invdendpoint.ItemEndPoint) + "/" + c.Id

	err := c.delete(endPoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Item) Retrieve(id string) (*Item, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.ItemEndPoint) + "/" + id

	itemEndPoint := new(invdendpoint.Item)

	item := &Item{c.Connection, itemEndPoint}

	_, err := c.retrieveDataFromAPI(endPoint, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *Item) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Items, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.ItemEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	items := make(Items, 0)

NEXT:
	tmpItems := make(Items, 0)

	endPointTmp, apiErr := c.retrieveDataFromAPI(endPoint, &tmpItems)

	if apiErr != nil {
		return nil, apiErr
	}

	items = append(items, tmpItems...)

	if endPointTmp != "" {
		goto NEXT
	}

	for _, item := range items {
		item.Connection = c.Connection
	}

	return items, nil
}

// SafeForCreation prunes item data for just fields that can be used for creation of an item
func SafeItemForCreation(item *invdendpoint.Item) (*invdendpoint.Item, error) {
	if item == nil {
		return nil, errors.New("task is nil")
	}

	itemData := new(invdendpoint.Item)
	itemData.Id = item.Id
	itemData.Name = item.Name
	itemData.Currency = item.Currency
	itemData.UnitCost = item.UnitCost
	itemData.Description = item.Description
	itemData.Type = item.Type
	itemData.Taxable = item.Taxable
	itemData.AvalaraTaxCode = item.AvalaraTaxCode
	itemData.GlAccount = item.GlAccount
	itemData.Discountable = item.Discountable
	itemData.Metadata = item.Metadata

	return itemData, nil
}

// SafeForUpdating prunes item data for just fields that can be used for updating of an item
func SafeItemForUpdating(item *invdendpoint.Item) (*invdendpoint.Item, error) {
	if item == nil {
		return nil, errors.New("task is nil")
	}

	itemData := new(invdendpoint.Item)

	itemData.Name = item.Name
	itemData.Description = item.Description
	itemData.Type = item.Type
	itemData.Metadata = item.Metadata

	return itemData, nil
}
