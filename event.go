package invdapi

import (
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

func (c *Event) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Events, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.EventsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	events := make(Events, 0)

NEXT:
	tmpEvents := make(Events, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpEvents)

	if apiErr != nil {
		return nil, apiErr
	}

	events = append(events, tmpEvents...)

	if endPoint != "" {
		goto NEXT
	}

	for _, event := range events {
		event.Connection = c.Connection

	}

	return events, nil

}

func (c *Event) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Events, string, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.EventsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	events := make(Events, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &events)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, event := range events {
		event.Connection = c.Connection

	}

	return events, nextEndPoint, nil

}


func (c *Event) Retrieve(id int64) (*Event, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.EventsEndPoint), id)

	eventEndPoint := new(invdendpoint.Event)

	event := &Event{c.Connection, eventEndPoint}

	_, err := c.retrieveDataFromAPI(endPoint, event)

	if err != nil {
		return nil, err
	}

	return event, nil

}
