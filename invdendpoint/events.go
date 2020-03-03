package invdendpoint

import (
	"encoding/json"
	"errors"
)

const EventsEndPoint = "/events"

type Events []Event

type Event struct {
	Id        int64           `json:"id,omitempty"`   //The event’s unique ID
	Type      string          `json:"type,omitempty"` //Event type
	Timestamp int64           `json:"timestamp,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"` //Contains an object property with the object that was subject of the event and an optional previous property for object.updated events that is a hash of the old values that changed during the event
}

type EventObject struct {
	Object *json.RawMessage  `json:"object,omitempty"`
}

func (e *Event) ParseEventObject() (*json.RawMessage,error) {
	data := e.Data

	eo := new(EventObject)

	b, err := data.MarshalJSON()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b,eo)

	if err != nil {
		return nil, err
	}

	if eo.Object == nil {
		return nil, errors.New("Could not parse event object")
	}

	return eo.Object,nil

}


func (e *Event) ParseInvoiceEvent() (*Invoice,error) {
	eoData, err := e.ParseEventObject()

	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()


	if err != nil {
		return nil, err
	}

	ie := new(Invoice)

	err = json.Unmarshal(b,ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}