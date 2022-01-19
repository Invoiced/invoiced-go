package event

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) ListAllByDatesAndUser(filter *invoiced.Filter, sort *invoiced.Sort, startDate int64, endDate int64, user string, objectType string, objectID int64) (invoiced.Events, error) {
	endpoint := invoiced.AddFilterAndSort("/events", filter, sort)
	endpoint = invoiced.AddQueryParameter(endpoint, "start_date", strconv.FormatInt(startDate, 10))
	endpoint = invoiced.AddQueryParameter(endpoint, "end_date", strconv.FormatInt(endDate, 10))
	endpoint = invoiced.AddQueryParameter(endpoint, "from", user)
	if len(objectType) > 0 {
		relatesTo := objectType + "," + strconv.FormatInt(objectID, 10)
		endpoint = invoiced.AddQueryParameter(endpoint, "related_to", relatesTo)
	}

	events := make(invoiced.Events, 0)

NEXT:
	tmpEvents := make(invoiced.Events, 0)

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

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Events, error) {
	endpoint := "/events"
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	events := make(invoiced.Events, 0)

NEXT:
	tmpEvents := make(invoiced.Events, 0)

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

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Events, string, error) {
	endpoint := "/events"
	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	events := make(invoiced.Events, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &events)

	if err != nil {
		return nil, "", err
	}

	return events, nextEndpoint, nil
}

func (c *Client) Retrieve(id int64) (*invoiced.Event, error) {
	resp := new(invoiced.Event)
	_, err := c.Api.Get("/events/"+strconv.FormatInt(id, 10)+"?include=user", resp)
	return resp, err
}

func (c *Client) RetrieveWithUser(id int64) (*invoiced.Event, error) {
	resp := new(invoiced.Event)
	_, err := c.Api.Get("/events/"+strconv.FormatInt(id, 10)+"?include=user", resp)
	return resp, err
}
