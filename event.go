package invdapi

import (
	"strconv"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Event struct {
	*Connection
	*invdendpoint.Event
}

type Events []*Event

func (c *Connection) NewEvent() *Event {
	event := new(invdendpoint.Event)
	return &Event{c, event}
}

func (c *Event) ListAllByDatesAndUser(filter *invdendpoint.Filter, sort *invdendpoint.Sort, startDate int64, endDate int64, user string, objectType string, objectID int64) (Events, error) {
	endpoint := invdendpoint.EventEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)
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

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpEvents)

	if apiErr != nil {
		return nil, apiErr
	}

	events = append(events, tmpEvents...)

	if endpoint != "" {
		goto NEXT
	}

	for _, event := range events {
		event.Connection = c.Connection
	}

	return events, nil
}

func (c *Event) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Events, error) {
	endpoint := invdendpoint.EventEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	events := make(Events, 0)

NEXT:
	tmpEvents := make(Events, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpEvents)

	if apiErr != nil {
		return nil, apiErr
	}

	events = append(events, tmpEvents...)

	if endpoint != "" {
		goto NEXT
	}

	for _, event := range events {
		event.Connection = c.Connection
	}

	return events, nil
}

func (c *Event) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Events, string, error) {
	endpoint := invdendpoint.EventEndpoint
	endpoint = addFilterAndSort(endpoint, filter, sort)

	events := make(Events, 0)

	nextEndpoint, apiErr := c.retrieveDataFromAPI(endpoint, &events)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, event := range events {
		event.Connection = c.Connection
	}

	return events, nextEndpoint, nil
}

func (c *Event) Retrieve(id int64) (*Event, error) {
	endpoint := invdendpoint.EventEndpoint + "/" + strconv.FormatInt(id, 10)

	eventEndpoint := new(invdendpoint.Event)

	event := &Event{c.Connection, eventEndpoint}

	_, err := c.retrieveDataFromAPI(endpoint, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
