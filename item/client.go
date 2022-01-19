package item

import "github.com/Invoiced/invoiced-go"

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.ItemRequest) (*invoiced.Item, error) {
	resp := new(invoiced.Item)
	err := c.Api.Create("/items", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id string) (*invoiced.Item, error) {
	resp := new(invoiced.Item)
	_, err := c.Api.Get("/items/"+id, resp)
	return resp, err
}

func (c *Client) Update(id string, request *invoiced.ItemRequest) (*invoiced.Item, error) {
	resp := new(invoiced.Item)
	err := c.Api.Update("/items/"+id, request, resp)
	return resp, err
}

func (c *Client) Delete(id string) error {
	return c.Api.Delete("/items/" + id)
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Items, error) {
	endpoint := invoiced.AddFilterAndSort("/items", filter, sort)

	items := make(invoiced.Items, 0)

NEXT:
	tmpItems := make(invoiced.Items, 0)

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
