package invdapi

import (
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

func (c *Item) Create(request *invdendpoint.ItemRequest) (*Item, error) {
	endpoint := invdendpoint.ItemEndpoint
	resp := new(Item)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Item) Retrieve(id string) (*Item, error) {
	endpoint := invdendpoint.ItemEndpoint + "/" + id

	item := &Item{c.Connection, new(invdendpoint.Item)}

	_, err := c.retrieveDataFromAPI(endpoint, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *Item) Update(request *invdendpoint.ItemRequest) error {
	endpoint := invdendpoint.ItemEndpoint + "/" + c.Id
	resp := new(Item)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Item = resp.Item

	return nil
}

func (c *Item) Delete() error {
	endpoint := invdendpoint.ItemEndpoint + "/" + c.Id

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Item) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Items, error) {
	endpoint := invdendpoint.ItemEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	items := make(Items, 0)

NEXT:
	tmpItems := make(Items, 0)

	endpointTmp, err := c.retrieveDataFromAPI(endpoint, &tmpItems)

	if err != nil {
		return nil, err
	}

	items = append(items, tmpItems...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, item := range items {
		item.Connection = c.Connection
	}

	return items, nil
}
