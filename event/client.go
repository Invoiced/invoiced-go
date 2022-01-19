package event

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

type Events []*Client

func (c *Client) ListAllByDatesAndUser(filter *invoiced.Filter, sort *invoiced.Sort, startDate int64, endDate int64, user string, objectType string, objectID int64) (Events, error) {
	endpoint := invoiced.EventEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)
	endpoint = addQueryParameter(endpoint, "start_date", strconv.FormatInt(startDate, 10))
	endpoint = addQueryParameter(endpoint, "end_date", strconv.FormatInt(endDate, 10))
	endpoint = addQueryParameter(endpoint, "from", user)
	if len(objectType) > 0 {
		relatesTo := objectType + "," + strconv.FormatInt(objectID, 10)
		endpoint = addQueryParameter(endpoint, "related_to", relatesTo)
	}

	events := make(Events, 0)

NEXT:
	tmpEvents := make(Events, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpEvents)

	if err != nil {
		return nil, err
	}

	events = append(events, tmpEvents...)

	if endpoint != "" {
		goto NEXT
	}

	return events, nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (Events, error) {
	endpoint := invoiced.EventEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	events := make(Events, 0)

NEXT:
	tmpEvents := make(Events, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpEvents)

	if err != nil {
		return nil, err
	}

	events = append(events, tmpEvents...)

	if endpoint != "" {
		goto NEXT
	}

	return events, nil
}

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (Events, string, error) {
	endpoint := invoiced.EventEndpoint
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	events := make(Events, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &events)

	if err != nil {
		return nil, "", err
	}

	return events, nextEndpoint, nil
}

func (c *Client) Retrieve(id int64) (*Client, error) {
	endpoint := invoiced.EventEndpoint + "/" + strconv.FormatInt(id, 10) + "?include=user"

	eventEndpoint := new(invoiced.Event)

	event := &Client{c.Client, eventEndpoint}

	_, err := c.Api.Get(endpoint, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (c *Client) RetrieveWithUser(id int64) (*Client, error) {
	endpoint := invoiced.EventEndpoint + "/" + strconv.FormatInt(id, 10) + "?include=user"

	eventEndpoint := new(invoiced.Event)

	event := &Client{c.Client, eventEndpoint}

	_, err := c.Api.Get(endpoint, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
