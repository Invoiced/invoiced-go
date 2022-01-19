package invoiced

type ItemClient struct {
	*Client
	*Item
}

type Items []*ItemClient

func (c *Client) NewItem() *ItemClient {
	item := new(Item)
	return &ItemClient{c, item}
}

func (c *ItemClient) Create(request *ItemRequest) (*ItemClient, error) {
	endpoint := ItemEndpoint
	resp := new(ItemClient)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ItemClient) Retrieve(id string) (*ItemClient, error) {
	endpoint := ItemEndpoint + "/" + id

	item := &ItemClient{c.Client, new(Item)}

	_, err := c.Api.Get(endpoint, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ItemClient) Update(request *ItemRequest) error {
	endpoint := ItemEndpoint + "/" + c.Id
	resp := new(ItemClient)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.Item = resp.Item

	return nil
}

func (c *ItemClient) Delete() error {
	endpoint := ItemEndpoint + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *ItemClient) ListAll(filter *Filter, sort *Sort) (Items, error) {
	endpoint := ItemEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

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
