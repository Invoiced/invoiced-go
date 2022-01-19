package invoiced

import (
	"strconv"
)

type EventClient struct {
	*Client
	*Event
}

type Events []*EventClient

func (c *Client) NewEvent() *EventClient {
	event := new(Event)
	return &EventClient{c, event}
}

func (c *EventClient) ListAllByDatesAndUser(filter *Filter, sort *Sort, startDate int64, endDate int64, user string, objectType string, objectID int64) (Events, error) {
	endpoint := EventEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)
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

func (c *EventClient) ListAll(filter *Filter, sort *Sort) (Events, error) {
	endpoint := EventEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)

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

func (c *EventClient) List(filter *Filter, sort *Sort) (Events, string, error) {
	endpoint := EventEndpoint
	endpoint = AddFilterAndSort(endpoint, filter, sort)

	events := make(Events, 0)

	nextEndpoint, err := c.Api.Get(endpoint, &events)

	if err != nil {
		return nil, "", err
	}

	return events, nextEndpoint, nil
}

func (c *EventClient) Retrieve(id int64) (*EventClient, error) {
	endpoint := EventEndpoint + "/" + strconv.FormatInt(id, 10) + "?include=user"

	eventEndpoint := new(Event)

	event := &EventClient{c.Client, eventEndpoint}

	_, err := c.Api.Get(endpoint, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (c *EventClient) RetrieveWithUser(id int64) (*EventClient, error) {
	endpoint := EventEndpoint + "/" + strconv.FormatInt(id, 10) + "?include=user"

	eventEndpoint := new(Event)

	event := &EventClient{c.Client, eventEndpoint}

	_, err := c.Api.Get(endpoint, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
