package invdendpoint

import (
	"encoding/json"
)

const EventsEndPoint = "/events"

type Event struct {
	Id        int64           `json:"id,omitempty"`   //The event’s unique ID
	Type      string          `json:"type,omitempty"` //Event type
	Timestamp int64           `json:"timestamp,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"` //Contains an object property with the object that was subject of the event and an optional previous property for object.updated events that is a hash of the old values that changed during the event
}
