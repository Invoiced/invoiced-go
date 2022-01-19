package item

import "github.com/Invoiced/invoiced-go"

type Client struct {
	*invoiced.Api
}

type Items []*Client

func (c *Client) Create(request *invoiced.ItemRequest) (*Client, error) {
	endpoint := invoiced.ItemEndpoint
	resp := new(Client)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id string) (*Client, error) {
	endpoint := invoiced.ItemEndpoint + "/" + id

	item := &Client{c.Client, new(invoiced.Item)}

	_, err := c.Api.Get(endpoint, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *Client) Update(request *invoiced.ItemRequest) error {
	endpoint := invoiced.ItemEndpoint + "/" + c.Id
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Item = resp.Item

	return nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.ItemEndpoint + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (Items, error) {
	endpoint := invoiced.ItemEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	items := make(Items, 0)

NEXT:
	tmpItems := make(Items, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpItems)

	if err != nil {
		return nil, err
	}

	items = append(items, tmpItems...)

	if endpointTmp != "" {
		goto NEXT
	}

	return items, nil
}
