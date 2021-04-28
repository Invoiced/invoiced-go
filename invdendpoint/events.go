package invdendpoint

import (
	"encoding/json"
	"errors"
	"strings"
)

const EventEndpoint = "/events"

type Events []Event

type Event struct {
	Id        int64           `json:"id,omitempty"` // The event’s unique ID
	Object    string          `json:"object,omitempty"`
	Type      string          `json:"type,omitempty"` // Event type
	Timestamp int64           `json:"timestamp,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"` // Contains an object property with the object that was subject of the event and an optional previous property for object.updated events that is a hash of the old values that changed during the event
}

type EventObject struct {
	Object         *json.RawMessage `json:"object,omitempty"`
	PreviousObject *json.RawMessage `json:"previous,omitempty"`
}

func (e *Event) ParseEventObject() (*json.RawMessage, error) {
	data := e.Data

	eo := new(EventObject)

	b, err := data.MarshalJSON()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, eo)

	if err != nil {
		return nil, err
	}

	if eo.Object == nil {
		return nil, errors.New("Could not parse event object")
	}

	return eo.Object, nil
}

func (e *Event) ParseEventPreviousObject() (*json.RawMessage, error) {
	data := e.Data

	eo := new(EventObject)

	b, err := data.MarshalJSON()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, eo)

	if err != nil {
		return nil, err
	}

	if eo.Object == nil {
		return nil, errors.New("Could not parse event object")
	}

	return eo.PreviousObject, nil
}

func (e *Event) ParseInvoiceEvent() (*Invoice, error) {
	eoData, err := e.ParseEventObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Invoice)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParseInvoicePreviousEvent() (*Invoice, error) {
	eoData, err := e.ParseEventPreviousObject()
	if err != nil {
		return nil, err
	}

	if eoData == nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Invoice)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParsePaymentEvent() (*Payment, error) {
	eoData, err := e.ParseEventObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Payment)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func CleanMetaDataArray(b []byte) []byte {
	s := string(b)
	s1 := strings.Replace(s, `"metadata": []`, ` "metadata": null`, -1)
	s1 = strings.Replace(s1, `"metadata":[]`, ` "metadata": null`, -1)
	return []byte(s1)
}
